// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/xiazeyin/fabric-gm/core/container"
	"github.com/xiazeyin/fabric-gm/core/container/ccintf"
)

type Instance struct {
	ChaincodeServerInfoStub        func() (*ccintf.ChaincodeServerInfo, error)
	chaincodeServerInfoMutex       sync.RWMutex
	chaincodeServerInfoArgsForCall []struct {
	}
	chaincodeServerInfoReturns struct {
		result1 *ccintf.ChaincodeServerInfo
		result2 error
	}
	chaincodeServerInfoReturnsOnCall map[int]struct {
		result1 *ccintf.ChaincodeServerInfo
		result2 error
	}
	StartStub        func(*ccintf.PeerConnection) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 *ccintf.PeerConnection
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	WaitStub        func() (int, error)
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
	}
	waitReturns struct {
		result1 int
		result2 error
	}
	waitReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Instance) ChaincodeServerInfo() (*ccintf.ChaincodeServerInfo, error) {
	fake.chaincodeServerInfoMutex.Lock()
	ret, specificReturn := fake.chaincodeServerInfoReturnsOnCall[len(fake.chaincodeServerInfoArgsForCall)]
	fake.chaincodeServerInfoArgsForCall = append(fake.chaincodeServerInfoArgsForCall, struct {
	}{})
	fake.recordInvocation("ChaincodeServerInfo", []interface{}{})
	fake.chaincodeServerInfoMutex.Unlock()
	if fake.ChaincodeServerInfoStub != nil {
		return fake.ChaincodeServerInfoStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.chaincodeServerInfoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Instance) ChaincodeServerInfoCallCount() int {
	fake.chaincodeServerInfoMutex.RLock()
	defer fake.chaincodeServerInfoMutex.RUnlock()
	return len(fake.chaincodeServerInfoArgsForCall)
}

func (fake *Instance) ChaincodeServerInfoCalls(stub func() (*ccintf.ChaincodeServerInfo, error)) {
	fake.chaincodeServerInfoMutex.Lock()
	defer fake.chaincodeServerInfoMutex.Unlock()
	fake.ChaincodeServerInfoStub = stub
}

func (fake *Instance) ChaincodeServerInfoReturns(result1 *ccintf.ChaincodeServerInfo, result2 error) {
	fake.chaincodeServerInfoMutex.Lock()
	defer fake.chaincodeServerInfoMutex.Unlock()
	fake.ChaincodeServerInfoStub = nil
	fake.chaincodeServerInfoReturns = struct {
		result1 *ccintf.ChaincodeServerInfo
		result2 error
	}{result1, result2}
}

func (fake *Instance) ChaincodeServerInfoReturnsOnCall(i int, result1 *ccintf.ChaincodeServerInfo, result2 error) {
	fake.chaincodeServerInfoMutex.Lock()
	defer fake.chaincodeServerInfoMutex.Unlock()
	fake.ChaincodeServerInfoStub = nil
	if fake.chaincodeServerInfoReturnsOnCall == nil {
		fake.chaincodeServerInfoReturnsOnCall = make(map[int]struct {
			result1 *ccintf.ChaincodeServerInfo
			result2 error
		})
	}
	fake.chaincodeServerInfoReturnsOnCall[i] = struct {
		result1 *ccintf.ChaincodeServerInfo
		result2 error
	}{result1, result2}
}

func (fake *Instance) Start(arg1 *ccintf.PeerConnection) error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 *ccintf.PeerConnection
	}{arg1})
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.startReturns
	return fakeReturns.result1
}

func (fake *Instance) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *Instance) StartCalls(stub func(*ccintf.PeerConnection) error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *Instance) StartArgsForCall(i int) *ccintf.PeerConnection {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Instance) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *Instance) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Instance) Stop() error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stopReturns
	return fakeReturns.result1
}

func (fake *Instance) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *Instance) StopCalls(stub func() error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *Instance) StopReturns(result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *Instance) StopReturnsOnCall(i int, result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Instance) Wait() (int, error) {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
	}{})
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.waitReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Instance) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *Instance) WaitCalls(stub func() (int, error)) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *Instance) WaitReturns(result1 int, result2 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *Instance) WaitReturnsOnCall(i int, result1 int, result2 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *Instance) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chaincodeServerInfoMutex.RLock()
	defer fake.chaincodeServerInfoMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Instance) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ container.Instance = new(Instance)
