// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/xiazeyin/fabric-gm/common/ledger"
)

type HistoryQueryExecutor struct {
	GetHistoryForKeyStub        func(string, string) (ledger.ResultsIterator, error)
	getHistoryForKeyMutex       sync.RWMutex
	getHistoryForKeyArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getHistoryForKeyReturns struct {
		result1 ledger.ResultsIterator
		result2 error
	}
	getHistoryForKeyReturnsOnCall map[int]struct {
		result1 ledger.ResultsIterator
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *HistoryQueryExecutor) GetHistoryForKey(arg1 string, arg2 string) (ledger.ResultsIterator, error) {
	fake.getHistoryForKeyMutex.Lock()
	ret, specificReturn := fake.getHistoryForKeyReturnsOnCall[len(fake.getHistoryForKeyArgsForCall)]
	fake.getHistoryForKeyArgsForCall = append(fake.getHistoryForKeyArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetHistoryForKey", []interface{}{arg1, arg2})
	fake.getHistoryForKeyMutex.Unlock()
	if fake.GetHistoryForKeyStub != nil {
		return fake.GetHistoryForKeyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getHistoryForKeyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *HistoryQueryExecutor) GetHistoryForKeyCallCount() int {
	fake.getHistoryForKeyMutex.RLock()
	defer fake.getHistoryForKeyMutex.RUnlock()
	return len(fake.getHistoryForKeyArgsForCall)
}

func (fake *HistoryQueryExecutor) GetHistoryForKeyCalls(stub func(string, string) (ledger.ResultsIterator, error)) {
	fake.getHistoryForKeyMutex.Lock()
	defer fake.getHistoryForKeyMutex.Unlock()
	fake.GetHistoryForKeyStub = stub
}

func (fake *HistoryQueryExecutor) GetHistoryForKeyArgsForCall(i int) (string, string) {
	fake.getHistoryForKeyMutex.RLock()
	defer fake.getHistoryForKeyMutex.RUnlock()
	argsForCall := fake.getHistoryForKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *HistoryQueryExecutor) GetHistoryForKeyReturns(result1 ledger.ResultsIterator, result2 error) {
	fake.getHistoryForKeyMutex.Lock()
	defer fake.getHistoryForKeyMutex.Unlock()
	fake.GetHistoryForKeyStub = nil
	fake.getHistoryForKeyReturns = struct {
		result1 ledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *HistoryQueryExecutor) GetHistoryForKeyReturnsOnCall(i int, result1 ledger.ResultsIterator, result2 error) {
	fake.getHistoryForKeyMutex.Lock()
	defer fake.getHistoryForKeyMutex.Unlock()
	fake.GetHistoryForKeyStub = nil
	if fake.getHistoryForKeyReturnsOnCall == nil {
		fake.getHistoryForKeyReturnsOnCall = make(map[int]struct {
			result1 ledger.ResultsIterator
			result2 error
		})
	}
	fake.getHistoryForKeyReturnsOnCall[i] = struct {
		result1 ledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *HistoryQueryExecutor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getHistoryForKeyMutex.RLock()
	defer fake.getHistoryForKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *HistoryQueryExecutor) recordInvocation(key string, args []interface{}) {
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
