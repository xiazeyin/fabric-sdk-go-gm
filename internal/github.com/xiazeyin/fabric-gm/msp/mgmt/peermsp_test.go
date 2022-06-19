/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mgmt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xiazeyin/fabric-gm/bccsp/sw"
	"github.com/xiazeyin/fabric-gm/core/config/configtest"
)

func TestLocalMSP(t *testing.T) {
	mspDir := configtest.GetDevMspDir()
	err := LoadLocalMsp(mspDir, nil, "SampleOrg")
	if err != nil {
		t.Fatalf("LoadLocalMsp failed, err %s", err)
	}

	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	assert.NoError(t, err)
	_, err = GetLocalMSP(cryptoProvider).GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetDefaultSigningIdentity failed, err %s", err)
	}
}
