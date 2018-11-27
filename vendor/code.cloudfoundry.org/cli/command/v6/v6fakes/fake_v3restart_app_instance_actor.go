// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeV3RestartAppInstanceActor struct {
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
	DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub        func(string, string, string, int) (v3action.Warnings, error)
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex       sync.RWMutex
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 int
	}
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3RestartAppInstanceActor) CloudControllerAPIVersion() string {
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

func (fake *FakeV3RestartAppInstanceActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3RestartAppInstanceActor) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeV3RestartAppInstanceActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3RestartAppInstanceActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
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

func (fake *FakeV3RestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndex(arg1 string, arg2 string, arg3 string, arg4 int) (v3action.Warnings, error) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	ret, specificReturn := fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall[len(fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall)]
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall = append(fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 int
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("DeleteInstanceByApplicationNameSpaceProcessTypeAndIndex", []interface{}{arg1, arg2, arg3, arg4})
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	if fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub != nil {
		return fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV3RestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexCallCount() int {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	return len(fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall)
}

func (fake *FakeV3RestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexCalls(stub func(string, string, string, int) (v3action.Warnings, error)) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub = stub
}

func (fake *FakeV3RestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall(i int) (string, string, string, int) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	argsForCall := fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeV3RestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns(result1 v3action.Warnings, result2 error) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub = nil
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3RestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub = nil
	if fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall == nil {
		fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3RestartAppInstanceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3RestartAppInstanceActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.V3RestartAppInstanceActor = new(FakeV3RestartAppInstanceActor)
