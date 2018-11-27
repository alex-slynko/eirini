// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeV3UnsetEnvActor struct {
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
	UnsetEnvironmentVariableByApplicationNameAndSpaceStub        func(string, string, string) (v3action.Warnings, error)
	unsetEnvironmentVariableByApplicationNameAndSpaceMutex       sync.RWMutex
	unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	unsetEnvironmentVariableByApplicationNameAndSpaceReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	unsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3UnsetEnvActor) CloudControllerAPIVersion() string {
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

func (fake *FakeV3UnsetEnvActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3UnsetEnvActor) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeV3UnsetEnvActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3UnsetEnvActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
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

func (fake *FakeV3UnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpace(arg1 string, arg2 string, arg3 string) (v3action.Warnings, error) {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall[len(fake.unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall)]
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall = append(fake.unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("UnsetEnvironmentVariableByApplicationNameAndSpace", []interface{}{arg1, arg2, arg3})
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Unlock()
	if fake.UnsetEnvironmentVariableByApplicationNameAndSpaceStub != nil {
		return fake.UnsetEnvironmentVariableByApplicationNameAndSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV3UnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpaceCallCount() int {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RLock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RUnlock()
	return len(fake.unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall)
}

func (fake *FakeV3UnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpaceCalls(stub func(string, string, string) (v3action.Warnings, error)) {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Lock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Unlock()
	fake.UnsetEnvironmentVariableByApplicationNameAndSpaceStub = stub
}

func (fake *FakeV3UnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall(i int) (string, string, string) {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RLock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RUnlock()
	argsForCall := fake.unsetEnvironmentVariableByApplicationNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV3UnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpaceReturns(result1 v3action.Warnings, result2 error) {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Lock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Unlock()
	fake.UnsetEnvironmentVariableByApplicationNameAndSpaceStub = nil
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3UnsetEnvActor) UnsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Lock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.Unlock()
	fake.UnsetEnvironmentVariableByApplicationNameAndSpaceStub = nil
	if fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall == nil {
		fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3UnsetEnvActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RLock()
	defer fake.unsetEnvironmentVariableByApplicationNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3UnsetEnvActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.V3UnsetEnvActor = new(FakeV3UnsetEnvActor)
