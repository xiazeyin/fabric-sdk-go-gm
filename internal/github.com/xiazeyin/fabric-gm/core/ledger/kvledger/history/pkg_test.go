/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package history

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xiazeyin/fabric-gm/common/ledger/blkstorage"
	"github.com/xiazeyin/fabric-gm/common/metrics/disabled"
	"github.com/xiazeyin/fabric-gm/core/ledger/kvledger/bookkeeping"
	"github.com/xiazeyin/fabric-gm/core/ledger/kvledger/txmgmt/privacyenabledstate"
	"github.com/xiazeyin/fabric-gm/core/ledger/kvledger/txmgmt/txmgr"
	"github.com/xiazeyin/fabric-gm/core/ledger/mock"
	"github.com/xiazeyin/gmgo/sm3"
)

var (
	testHashFunc = func(data []byte) ([]byte, error) {
		h := sm3.New()
		if _, err := h.Write(data); err != nil {
			return nil, err
		}
		return h.Sum(nil), nil
	}
)

type levelDBLockBasedHistoryEnv struct {
	t                     testing.TB
	testBlockStorageEnv   *testBlockStoreEnv
	testDBEnv             privacyenabledstate.TestEnv
	testBookkeepingEnv    *bookkeeping.TestEnv
	txmgr                 *txmgr.LockBasedTxMgr
	testHistoryDBProvider *DBProvider
	testHistoryDB         *DB
	testHistoryDBPath     string
}

func newTestHistoryEnv(t *testing.T) *levelDBLockBasedHistoryEnv {
	testLedgerID := "TestLedger"

	blockStorageTestEnv := newBlockStorageTestEnv(t)

	testDBEnv := &privacyenabledstate.LevelDBTestEnv{}
	testDBEnv.Init(t)
	testDB := testDBEnv.GetDBHandle(testLedgerID)
	testBookkeepingEnv := bookkeeping.NewTestEnv(t)

	testHistoryDBPath, err := ioutil.TempDir("", "historyldb")
	if err != nil {
		t.Fatalf("Failed to create history database directory: %s", err)
	}

	txmgrInitializer := &txmgr.Initializer{
		LedgerID:            testLedgerID,
		DB:                  testDB,
		StateListeners:      nil,
		BtlPolicy:           nil,
		BookkeepingProvider: testBookkeepingEnv.TestProvider,
		CCInfoProvider:      &mock.DeployedChaincodeInfoProvider{},
		CustomTxProcessors:  nil,
		HashFunc:            testHashFunc,
	}
	txMgr, err := txmgr.NewLockBasedTxMgr(txmgrInitializer)

	assert.NoError(t, err)
	testHistoryDBProvider, err := NewDBProvider(testHistoryDBPath)
	assert.NoError(t, err)
	testHistoryDB, err := testHistoryDBProvider.GetDBHandle("TestHistoryDB")
	assert.NoError(t, err)

	return &levelDBLockBasedHistoryEnv{
		t,
		blockStorageTestEnv,
		testDBEnv,
		testBookkeepingEnv,
		txMgr,
		testHistoryDBProvider,
		testHistoryDB,
		testHistoryDBPath,
	}
}

func (env *levelDBLockBasedHistoryEnv) cleanup() {
	env.txmgr.Shutdown()
	env.testDBEnv.Cleanup()
	env.testBlockStorageEnv.cleanup()
	env.testBookkeepingEnv.Cleanup()
	// clean up history
	env.testHistoryDBProvider.Close()
	os.RemoveAll(env.testHistoryDBPath)
}

/////// testBlockStoreEnv//////

type testBlockStoreEnv struct {
	t               testing.TB
	provider        *blkstorage.BlockStoreProvider
	blockStorageDir string
}

func newBlockStorageTestEnv(t testing.TB) *testBlockStoreEnv {

	testPath, err := ioutil.TempDir("", "historyleveldb-")
	if err != nil {
		panic(err)
	}
	conf := blkstorage.NewConf(testPath, 0)

	attrsToIndex := []blkstorage.IndexableAttr{
		blkstorage.IndexableAttrBlockHash,
		blkstorage.IndexableAttrBlockNum,
		blkstorage.IndexableAttrTxID,
		blkstorage.IndexableAttrBlockNumTranNum,
	}
	indexConfig := &blkstorage.IndexConfig{AttrsToIndex: attrsToIndex}

	p, err := blkstorage.NewProvider(conf, indexConfig, &disabled.Provider{})
	assert.NoError(t, err)
	return &testBlockStoreEnv{t, p, testPath}
}

func (env *testBlockStoreEnv) cleanup() {
	env.provider.Close()
	env.removeFSPath()
}

func (env *testBlockStoreEnv) removeFSPath() {
	fsPath := env.blockStorageDir
	os.RemoveAll(fsPath)
}
