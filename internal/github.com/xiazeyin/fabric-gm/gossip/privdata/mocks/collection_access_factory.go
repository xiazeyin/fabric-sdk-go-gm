// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import peer "github.com/xiazeyin/fabric-protos-go-gm/peer"
import privdata "github.com/xiazeyin/fabric-gm/core/common/privdata"

// CollectionAccessFactory is an autogenerated mock type for the CollectionAccessFactory type
type CollectionAccessFactory struct {
	mock.Mock
}

// AccessPolicy provides a mock function with given fields: config, chainID
func (_m *CollectionAccessFactory) AccessPolicy(config *peer.CollectionConfig, chainID string) (privdata.CollectionAccessPolicy, error) {
	ret := _m.Called(config, chainID)

	var r0 privdata.CollectionAccessPolicy
	if rf, ok := ret.Get(0).(func(*peer.CollectionConfig, string) privdata.CollectionAccessPolicy); ok {
		r0 = rf(config, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(privdata.CollectionAccessPolicy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*peer.CollectionConfig, string) error); ok {
		r1 = rf(config, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
