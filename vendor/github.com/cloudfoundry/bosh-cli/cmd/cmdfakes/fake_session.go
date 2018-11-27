// Code generated by counterfeiter. DO NOT EDIT.
package cmdfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/cmd"
	cmdconf "github.com/cloudfoundry/bosh-cli/cmd/config"
	boshdir "github.com/cloudfoundry/bosh-cli/director"
	boshuaa "github.com/cloudfoundry/bosh-cli/uaa"
)

type FakeSession struct {
	EnvironmentStub        func() string
	environmentMutex       sync.RWMutex
	environmentArgsForCall []struct{}
	environmentReturns     struct {
		result1 string
	}
	environmentReturnsOnCall map[int]struct {
		result1 string
	}
	CredentialsStub        func() cmdconf.Creds
	credentialsMutex       sync.RWMutex
	credentialsArgsForCall []struct{}
	credentialsReturns     struct {
		result1 cmdconf.Creds
	}
	credentialsReturnsOnCall map[int]struct {
		result1 cmdconf.Creds
	}
	UAAStub        func() (boshuaa.UAA, error)
	uAAMutex       sync.RWMutex
	uAAArgsForCall []struct{}
	uAAReturns     struct {
		result1 boshuaa.UAA
		result2 error
	}
	uAAReturnsOnCall map[int]struct {
		result1 boshuaa.UAA
		result2 error
	}
	DirectorStub        func() (boshdir.Director, error)
	directorMutex       sync.RWMutex
	directorArgsForCall []struct{}
	directorReturns     struct {
		result1 boshdir.Director
		result2 error
	}
	directorReturnsOnCall map[int]struct {
		result1 boshdir.Director
		result2 error
	}
	AnonymousDirectorStub        func() (boshdir.Director, error)
	anonymousDirectorMutex       sync.RWMutex
	anonymousDirectorArgsForCall []struct{}
	anonymousDirectorReturns     struct {
		result1 boshdir.Director
		result2 error
	}
	anonymousDirectorReturnsOnCall map[int]struct {
		result1 boshdir.Director
		result2 error
	}
	DeploymentStub        func() (boshdir.Deployment, error)
	deploymentMutex       sync.RWMutex
	deploymentArgsForCall []struct{}
	deploymentReturns     struct {
		result1 boshdir.Deployment
		result2 error
	}
	deploymentReturnsOnCall map[int]struct {
		result1 boshdir.Deployment
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSession) Environment() string {
	fake.environmentMutex.Lock()
	ret, specificReturn := fake.environmentReturnsOnCall[len(fake.environmentArgsForCall)]
	fake.environmentArgsForCall = append(fake.environmentArgsForCall, struct{}{})
	fake.recordInvocation("Environment", []interface{}{})
	fake.environmentMutex.Unlock()
	if fake.EnvironmentStub != nil {
		return fake.EnvironmentStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.environmentReturns.result1
}

func (fake *FakeSession) EnvironmentCallCount() int {
	fake.environmentMutex.RLock()
	defer fake.environmentMutex.RUnlock()
	return len(fake.environmentArgsForCall)
}

func (fake *FakeSession) EnvironmentReturns(result1 string) {
	fake.EnvironmentStub = nil
	fake.environmentReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSession) EnvironmentReturnsOnCall(i int, result1 string) {
	fake.EnvironmentStub = nil
	if fake.environmentReturnsOnCall == nil {
		fake.environmentReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.environmentReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSession) Credentials() cmdconf.Creds {
	fake.credentialsMutex.Lock()
	ret, specificReturn := fake.credentialsReturnsOnCall[len(fake.credentialsArgsForCall)]
	fake.credentialsArgsForCall = append(fake.credentialsArgsForCall, struct{}{})
	fake.recordInvocation("Credentials", []interface{}{})
	fake.credentialsMutex.Unlock()
	if fake.CredentialsStub != nil {
		return fake.CredentialsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.credentialsReturns.result1
}

func (fake *FakeSession) CredentialsCallCount() int {
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	return len(fake.credentialsArgsForCall)
}

func (fake *FakeSession) CredentialsReturns(result1 cmdconf.Creds) {
	fake.CredentialsStub = nil
	fake.credentialsReturns = struct {
		result1 cmdconf.Creds
	}{result1}
}

func (fake *FakeSession) CredentialsReturnsOnCall(i int, result1 cmdconf.Creds) {
	fake.CredentialsStub = nil
	if fake.credentialsReturnsOnCall == nil {
		fake.credentialsReturnsOnCall = make(map[int]struct {
			result1 cmdconf.Creds
		})
	}
	fake.credentialsReturnsOnCall[i] = struct {
		result1 cmdconf.Creds
	}{result1}
}

func (fake *FakeSession) UAA() (boshuaa.UAA, error) {
	fake.uAAMutex.Lock()
	ret, specificReturn := fake.uAAReturnsOnCall[len(fake.uAAArgsForCall)]
	fake.uAAArgsForCall = append(fake.uAAArgsForCall, struct{}{})
	fake.recordInvocation("UAA", []interface{}{})
	fake.uAAMutex.Unlock()
	if fake.UAAStub != nil {
		return fake.UAAStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.uAAReturns.result1, fake.uAAReturns.result2
}

func (fake *FakeSession) UAACallCount() int {
	fake.uAAMutex.RLock()
	defer fake.uAAMutex.RUnlock()
	return len(fake.uAAArgsForCall)
}

func (fake *FakeSession) UAAReturns(result1 boshuaa.UAA, result2 error) {
	fake.UAAStub = nil
	fake.uAAReturns = struct {
		result1 boshuaa.UAA
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) UAAReturnsOnCall(i int, result1 boshuaa.UAA, result2 error) {
	fake.UAAStub = nil
	if fake.uAAReturnsOnCall == nil {
		fake.uAAReturnsOnCall = make(map[int]struct {
			result1 boshuaa.UAA
			result2 error
		})
	}
	fake.uAAReturnsOnCall[i] = struct {
		result1 boshuaa.UAA
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) Director() (boshdir.Director, error) {
	fake.directorMutex.Lock()
	ret, specificReturn := fake.directorReturnsOnCall[len(fake.directorArgsForCall)]
	fake.directorArgsForCall = append(fake.directorArgsForCall, struct{}{})
	fake.recordInvocation("Director", []interface{}{})
	fake.directorMutex.Unlock()
	if fake.DirectorStub != nil {
		return fake.DirectorStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.directorReturns.result1, fake.directorReturns.result2
}

func (fake *FakeSession) DirectorCallCount() int {
	fake.directorMutex.RLock()
	defer fake.directorMutex.RUnlock()
	return len(fake.directorArgsForCall)
}

func (fake *FakeSession) DirectorReturns(result1 boshdir.Director, result2 error) {
	fake.DirectorStub = nil
	fake.directorReturns = struct {
		result1 boshdir.Director
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) DirectorReturnsOnCall(i int, result1 boshdir.Director, result2 error) {
	fake.DirectorStub = nil
	if fake.directorReturnsOnCall == nil {
		fake.directorReturnsOnCall = make(map[int]struct {
			result1 boshdir.Director
			result2 error
		})
	}
	fake.directorReturnsOnCall[i] = struct {
		result1 boshdir.Director
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) AnonymousDirector() (boshdir.Director, error) {
	fake.anonymousDirectorMutex.Lock()
	ret, specificReturn := fake.anonymousDirectorReturnsOnCall[len(fake.anonymousDirectorArgsForCall)]
	fake.anonymousDirectorArgsForCall = append(fake.anonymousDirectorArgsForCall, struct{}{})
	fake.recordInvocation("AnonymousDirector", []interface{}{})
	fake.anonymousDirectorMutex.Unlock()
	if fake.AnonymousDirectorStub != nil {
		return fake.AnonymousDirectorStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.anonymousDirectorReturns.result1, fake.anonymousDirectorReturns.result2
}

func (fake *FakeSession) AnonymousDirectorCallCount() int {
	fake.anonymousDirectorMutex.RLock()
	defer fake.anonymousDirectorMutex.RUnlock()
	return len(fake.anonymousDirectorArgsForCall)
}

func (fake *FakeSession) AnonymousDirectorReturns(result1 boshdir.Director, result2 error) {
	fake.AnonymousDirectorStub = nil
	fake.anonymousDirectorReturns = struct {
		result1 boshdir.Director
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) AnonymousDirectorReturnsOnCall(i int, result1 boshdir.Director, result2 error) {
	fake.AnonymousDirectorStub = nil
	if fake.anonymousDirectorReturnsOnCall == nil {
		fake.anonymousDirectorReturnsOnCall = make(map[int]struct {
			result1 boshdir.Director
			result2 error
		})
	}
	fake.anonymousDirectorReturnsOnCall[i] = struct {
		result1 boshdir.Director
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) Deployment() (boshdir.Deployment, error) {
	fake.deploymentMutex.Lock()
	ret, specificReturn := fake.deploymentReturnsOnCall[len(fake.deploymentArgsForCall)]
	fake.deploymentArgsForCall = append(fake.deploymentArgsForCall, struct{}{})
	fake.recordInvocation("Deployment", []interface{}{})
	fake.deploymentMutex.Unlock()
	if fake.DeploymentStub != nil {
		return fake.DeploymentStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deploymentReturns.result1, fake.deploymentReturns.result2
}

func (fake *FakeSession) DeploymentCallCount() int {
	fake.deploymentMutex.RLock()
	defer fake.deploymentMutex.RUnlock()
	return len(fake.deploymentArgsForCall)
}

func (fake *FakeSession) DeploymentReturns(result1 boshdir.Deployment, result2 error) {
	fake.DeploymentStub = nil
	fake.deploymentReturns = struct {
		result1 boshdir.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) DeploymentReturnsOnCall(i int, result1 boshdir.Deployment, result2 error) {
	fake.DeploymentStub = nil
	if fake.deploymentReturnsOnCall == nil {
		fake.deploymentReturnsOnCall = make(map[int]struct {
			result1 boshdir.Deployment
			result2 error
		})
	}
	fake.deploymentReturnsOnCall[i] = struct {
		result1 boshdir.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.environmentMutex.RLock()
	defer fake.environmentMutex.RUnlock()
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	fake.uAAMutex.RLock()
	defer fake.uAAMutex.RUnlock()
	fake.directorMutex.RLock()
	defer fake.directorMutex.RUnlock()
	fake.anonymousDirectorMutex.RLock()
	defer fake.anonymousDirectorMutex.RUnlock()
	fake.deploymentMutex.RLock()
	defer fake.deploymentMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSession) recordInvocation(key string, args []interface{}) {
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

var _ cmd.Session = new(FakeSession)
