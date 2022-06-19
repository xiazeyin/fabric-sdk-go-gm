/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package channelconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
	cb "github.com/xiazeyin/fabric-protos-go-gm/common"
)

func TestConsortiums(t *testing.T) {
	_, err := NewConsortiumsConfig(&cb.ConfigGroup{}, nil)
	assert.NoError(t, err)
}
