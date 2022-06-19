/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package protoutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xiazeyin/fabric-gm/protoutil"
	"github.com/xiazeyin/fabric-protos-go-gm/common"
)

func TestNewConfigGroup(t *testing.T) {
	assert.Equal(t,
		&common.ConfigGroup{
			Groups:   make(map[string]*common.ConfigGroup),
			Values:   make(map[string]*common.ConfigValue),
			Policies: make(map[string]*common.ConfigPolicy),
		},
		protoutil.NewConfigGroup(),
	)
}
