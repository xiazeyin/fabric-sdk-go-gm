/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package statecouchdb

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xiazeyin/fabric-gm/common/metrics/disabled"
	"github.com/xiazeyin/fabric-gm/core/ledger"
	"github.com/xiazeyin/fabric-gm/integration/runner"
)

// StartCouchDB starts the CouchDB if it is not running already
func StartCouchDB(t *testing.T, binds []string) (addr string, stopCouchDBFunc func()) {
	couchDB := &runner.CouchDB{Binds: binds}
	require.NoError(t, couchDB.Start())
	return couchDB.Address(), func() { couchDB.Stop() }
}

// IsEmpty returns whether or not the couchdb is empty
func IsEmpty(t testing.TB, config *ledger.CouchDBConfig) bool {
	couchInstance, err := createCouchInstance(config, &disabled.Provider{})
	require.NoError(t, err)
	dbEmpty, err := couchInstance.isEmpty(nil)
	require.NoError(t, err)
	return dbEmpty
}
