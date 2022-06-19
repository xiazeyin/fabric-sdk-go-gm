/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blockledger

import (
	"github.com/golang/protobuf/proto"
	"github.com/xiazeyin/fabric-gm/common/flogging"
	"github.com/xiazeyin/fabric-gm/protoutil"
	cb "github.com/xiazeyin/fabric-protos-go-gm/common"
	ab "github.com/xiazeyin/fabric-protos-go-gm/orderer"
)

var logger = flogging.MustGetLogger("common.ledger.blockledger.util")

var closedChan chan struct{}

func init() {
	closedChan = make(chan struct{})
	close(closedChan)
}

// NotFoundErrorIterator simply always returns an error of cb.Status_NOT_FOUND,
// and is generally useful for implementations of the Reader interface
type NotFoundErrorIterator struct{}

// Next returns nil, cb.Status_NOT_FOUND
func (nfei *NotFoundErrorIterator) Next() (*cb.Block, cb.Status) {
	return nil, cb.Status_NOT_FOUND
}

// ReadyChan returns a closed channel
func (nfei *NotFoundErrorIterator) ReadyChan() <-chan struct{} {
	return closedChan
}

// Close does nothing
func (nfei *NotFoundErrorIterator) Close() {}

// CreateNextBlock provides a utility way to construct the next block from
// contents and metadata for a given ledger
// XXX This will need to be modified to accept marshaled envelopes
//     to accommodate non-deterministic marshaling
func CreateNextBlock(rl Reader, messages []*cb.Envelope) *cb.Block {
	var nextBlockNumber uint64
	var previousBlockHash []byte
	var err error

	if rl.Height() > 0 {
		it, _ := rl.Iterator(&ab.SeekPosition{
			Type: &ab.SeekPosition_Newest{
				Newest: &ab.SeekNewest{},
			},
		})
		block, status := it.Next()
		if status != cb.Status_SUCCESS {
			panic("Error seeking to newest block for chain with non-zero height")
		}
		nextBlockNumber = block.Header.Number + 1
		previousBlockHash = protoutil.BlockHeaderHash(block.Header)
	}

	data := &cb.BlockData{
		Data: make([][]byte, len(messages)),
	}

	for i, msg := range messages {
		data.Data[i], err = proto.Marshal(msg)
		if err != nil {
			panic(err)
		}
	}

	block := protoutil.NewBlock(nextBlockNumber, previousBlockHash)
	block.Header.DataHash = protoutil.BlockDataHash(data)
	block.Data = data

	return block
}

// GetBlock is a utility method for retrieving a single block
func GetBlock(rl Reader, index uint64) *cb.Block {
	iterator, _ := rl.Iterator(&ab.SeekPosition{
		Type: &ab.SeekPosition_Specified{
			Specified: &ab.SeekSpecified{Number: index},
		},
	})
	if iterator == nil {
		return nil
	}
	defer iterator.Close()
	block, status := iterator.Next()
	if status != cb.Status_SUCCESS {
		return nil
	}
	return block
}

func GetBlockByNumber(rl Reader, blockNum uint64) (*cb.Block, error) {
	logger.Debugw("Retrieving block", "blockNum", blockNum)
	return rl.RetrieveBlockByNumber(blockNum)
}
