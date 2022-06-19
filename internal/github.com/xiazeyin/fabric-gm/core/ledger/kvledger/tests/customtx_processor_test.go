/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xiazeyin/fabric-gm/bccsp/sw"
	"github.com/xiazeyin/fabric-gm/core/ledger"
	"github.com/xiazeyin/fabric-gm/core/ledger/ledgermgmt"
	"github.com/xiazeyin/fabric-gm/core/ledger/mock"
	"github.com/xiazeyin/fabric-protos-go-gm/common"
	protopeer "github.com/xiazeyin/fabric-protos-go-gm/peer"
)

func TestReadWriteCustomTxProcessor(t *testing.T) {
	fakeTxProcessor := &mock.CustomTxProcessor{}

	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	assert.NoError(t, err)

	env := newEnvWithInitializer(
		t,
		&ledgermgmt.Initializer{
			CustomTxProcessors: map[common.HeaderType]ledger.CustomTxProcessor{
				100: fakeTxProcessor,
			},
			HashProvider: cryptoProvider,
		},
	)
	defer env.cleanup()
	env.initLedgerMgmt()

	h := env.newTestHelperCreateLgr("ledger1", t)
	h.simulateDataTx("tx0", func(s *simulator) {
		s.setState("ns", "key1", "value1")
		s.setState("ns", "key2", "value2")
		s.setState("ns", "key3", "value3")
	})
	h.cutBlockAndCommitLegacy() // commit block-1 to populate initial state

	valueCounter := 0
	fakeTxProcessor.GenerateSimulationResultsStub =
		// tx processor reads and modifies key1
		func(txEnvelop *common.Envelope, s ledger.TxSimulator, initializingLedger bool) error {
			valKey1, err := s.GetState("ns", "key1")
			assert.NoError(t, err)
			assert.Equal(t, []byte("value1"), valKey1)
			valueCounter++
			return s.SetState("ns", "key1", []byte(fmt.Sprintf("value1_%d", valueCounter)))
		}

	// block-2 with two post order transactions
	h.addPostOrderTx("tx1", 100)
	h.addPostOrderTx("tx2", 100)
	h.cutBlockAndCommitLegacy()

	// Tx1 should be valid and tx2 should be marked as invalid because, tx1 has already modified key1
	// in the same block
	h.verifyTxValidationCode("tx1", protopeer.TxValidationCode_VALID)
	h.verifyTxValidationCode("tx2", protopeer.TxValidationCode_MVCC_READ_CONFLICT)
	h.verifyPubState("ns", "key1", "value1_1")
}

func TestRangeReadAndWriteCustomTxProcessor(t *testing.T) {
	fakeTxProcessor1 := &mock.CustomTxProcessor{}
	fakeTxProcessor2 := &mock.CustomTxProcessor{}
	fakeTxProcessor3 := &mock.CustomTxProcessor{}

	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	assert.NoError(t, err)

	env := newEnvWithInitializer(
		t,
		&ledgermgmt.Initializer{
			CustomTxProcessors: map[common.HeaderType]ledger.CustomTxProcessor{
				101: fakeTxProcessor1,
				102: fakeTxProcessor2,
				103: fakeTxProcessor3,
			},
			HashProvider: cryptoProvider,
		},
	)
	defer env.cleanup()
	env.initLedgerMgmt()

	h := env.newTestHelperCreateLgr("ledger1", t)
	h.simulateDataTx("tx0", func(s *simulator) {
		s.setState("ns", "key1", "value1")
		s.setState("ns", "key2", "value2")
		s.setState("ns", "key3", "value3")
	})
	h.cutBlockAndCommitLegacy() // commit block-1 to populate initial state

	fakeTxProcessor1.GenerateSimulationResultsStub =
		// tx processor for txtype 101 sets key1
		func(txEnvelop *common.Envelope, s ledger.TxSimulator, initializingLedger bool) error {
			return s.SetState("ns", "key1", []byte("value1_new"))
		}

	fakeTxProcessor2.GenerateSimulationResultsStub =
		// tx processor for txtype 102 reads a range (that covers key1) and sets key2
		func(txEnvelop *common.Envelope, s ledger.TxSimulator, initializingLedger bool) error {
			itr, err := s.GetStateRangeScanIterator("ns", "key1", "key2")
			assert.NoError(t, err)
			for {
				res, err := itr.Next()
				assert.NoError(t, err)
				if res == nil {
					break
				}
			}
			return s.SetState("ns", "key2", []byte("value2_new"))
		}

	fakeTxProcessor3.GenerateSimulationResultsStub =
		// tx processor for txtype 103 reads a range (that does not include key1) and sets key2
		func(txEnvelop *common.Envelope, s ledger.TxSimulator, initializingLedger bool) error {
			itr, err := s.GetStateRangeScanIterator("ns", "key2", "key3")
			assert.NoError(t, err)
			for {
				res, err := itr.Next()
				assert.NoError(t, err)
				if res == nil {
					break
				}
			}
			return s.SetState("ns", "key3", []byte("value3_new"))
		}

	// block-2 with three post order transactions
	h.addPostOrderTx("tx1", 101)
	h.addPostOrderTx("tx2", 102)
	h.addPostOrderTx("tx3", 103)
	h.cutBlockAndCommitLegacy()

	// Tx1 should be valid and tx2 should be marked as invalid because, tx1 has already modified key1
	// in the same block (and tx2 does a range iteration that includes key1)
	// However, tx3 should be fine as this performs a range itertaes that does not include key1
	h.verifyTxValidationCode("tx1", protopeer.TxValidationCode_VALID)
	h.verifyTxValidationCode("tx2", protopeer.TxValidationCode_PHANTOM_READ_CONFLICT)
	h.verifyTxValidationCode("tx3", protopeer.TxValidationCode_VALID)
	h.verifyPubState("ns", "key1", "value1_new")
	h.verifyPubState("ns", "key2", "value2")
	h.verifyPubState("ns", "key3", "value3_new")
}
