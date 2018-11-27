// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v2action "code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeCreateRouteActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct {
	}
	cloudControllerAPIVersionReturns struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	CreateRouteWithExistenceCheckStub        func(string, string, v2action.Route, bool) (v2action.Route, v2action.Warnings, error)
	createRouteWithExistenceCheckMutex       sync.RWMutex
	createRouteWithExistenceCheckArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 v2action.Route
		arg4 bool
	}
	createRouteWithExistenceCheckReturns struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	createRouteWithExistenceCheckReturnsOnCall map[int]struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreateRouteActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerAPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeCreateRouteActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeCreateRouteActor) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeCreateRouteActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreateRouteActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreateRouteActor) CreateRouteWithExistenceCheck(arg1 string, arg2 string, arg3 v2action.Route, arg4 bool) (v2action.Route, v2action.Warnings, error) {
	fake.createRouteWithExistenceCheckMutex.Lock()
	ret, specificReturn := fake.createRouteWithExistenceCheckReturnsOnCall[len(fake.createRouteWithExistenceCheckArgsForCall)]
	fake.createRouteWithExistenceCheckArgsForCall = append(fake.createRouteWithExistenceCheckArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 v2action.Route
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CreateRouteWithExistenceCheck", []interface{}{arg1, arg2, arg3, arg4})
	fake.createRouteWithExistenceCheckMutex.Unlock()
	if fake.CreateRouteWithExistenceCheckStub != nil {
		return fake.CreateRouteWithExistenceCheckStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createRouteWithExistenceCheckReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCreateRouteActor) CreateRouteWithExistenceCheckCallCount() int {
	fake.createRouteWithExistenceCheckMutex.RLock()
	defer fake.createRouteWithExistenceCheckMutex.RUnlock()
	return len(fake.createRouteWithExistenceCheckArgsForCall)
}

func (fake *FakeCreateRouteActor) CreateRouteWithExistenceCheckCalls(stub func(string, string, v2action.Route, bool) (v2action.Route, v2action.Warnings, error)) {
	fake.createRouteWithExistenceCheckMutex.Lock()
	defer fake.createRouteWithExistenceCheckMutex.Unlock()
	fake.CreateRouteWithExistenceCheckStub = stub
}

func (fake *FakeCreateRouteActor) CreateRouteWithExistenceCheckArgsForCall(i int) (string, string, v2action.Route, bool) {
	fake.createRouteWithExistenceCheckMutex.RLock()
	defer fake.createRouteWithExistenceCheckMutex.RUnlock()
	argsForCall := fake.createRouteWithExistenceCheckArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeCreateRouteActor) CreateRouteWithExistenceCheckReturns(result1 v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.createRouteWithExistenceCheckMutex.Lock()
	defer fake.createRouteWithExistenceCheckMutex.Unlock()
	fake.CreateRouteWithExistenceCheckStub = nil
	fake.createRouteWithExistenceCheckReturns = struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateRouteActor) CreateRouteWithExistenceCheckReturnsOnCall(i int, result1 v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.createRouteWithExistenceCheckMutex.Lock()
	defer fake.createRouteWithExistenceCheckMutex.Unlock()
	fake.CreateRouteWithExistenceCheckStub = nil
	if fake.createRouteWithExistenceCheckReturnsOnCall == nil {
		fake.createRouteWithExistenceCheckReturnsOnCall = make(map[int]struct {
			result1 v2action.Route
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.createRouteWithExistenceCheckReturnsOnCall[i] = struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateRouteActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.createRouteWithExistenceCheckMutex.RLock()
	defer fake.createRouteWithExistenceCheckMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreateRouteActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.CreateRouteActor = new(FakeCreateRouteActor)
