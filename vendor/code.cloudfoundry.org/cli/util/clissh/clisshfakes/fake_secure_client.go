// Code generated by counterfeiter. DO NOT EDIT.
package clisshfakes

import (
	net "net"
	sync "sync"

	clissh "code.cloudfoundry.org/cli/util/clissh"
	ssh "golang.org/x/crypto/ssh"
)

type FakeSecureClient struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	ConnStub        func() ssh.Conn
	connMutex       sync.RWMutex
	connArgsForCall []struct {
	}
	connReturns struct {
		result1 ssh.Conn
	}
	connReturnsOnCall map[int]struct {
		result1 ssh.Conn
	}
	DialStub        func(string, string) (net.Conn, error)
	dialMutex       sync.RWMutex
	dialArgsForCall []struct {
		arg1 string
		arg2 string
	}
	dialReturns struct {
		result1 net.Conn
		result2 error
	}
	dialReturnsOnCall map[int]struct {
		result1 net.Conn
		result2 error
	}
	NewSessionStub        func() (clissh.SecureSession, error)
	newSessionMutex       sync.RWMutex
	newSessionArgsForCall []struct {
	}
	newSessionReturns struct {
		result1 clissh.SecureSession
		result2 error
	}
	newSessionReturnsOnCall map[int]struct {
		result1 clissh.SecureSession
		result2 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
	}
	waitReturns struct {
		result1 error
	}
	waitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecureClient) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeSecureClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSecureClient) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeSecureClient) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureClient) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureClient) Conn() ssh.Conn {
	fake.connMutex.Lock()
	ret, specificReturn := fake.connReturnsOnCall[len(fake.connArgsForCall)]
	fake.connArgsForCall = append(fake.connArgsForCall, struct {
	}{})
	fake.recordInvocation("Conn", []interface{}{})
	fake.connMutex.Unlock()
	if fake.ConnStub != nil {
		return fake.ConnStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.connReturns
	return fakeReturns.result1
}

func (fake *FakeSecureClient) ConnCallCount() int {
	fake.connMutex.RLock()
	defer fake.connMutex.RUnlock()
	return len(fake.connArgsForCall)
}

func (fake *FakeSecureClient) ConnCalls(stub func() ssh.Conn) {
	fake.connMutex.Lock()
	defer fake.connMutex.Unlock()
	fake.ConnStub = stub
}

func (fake *FakeSecureClient) ConnReturns(result1 ssh.Conn) {
	fake.connMutex.Lock()
	defer fake.connMutex.Unlock()
	fake.ConnStub = nil
	fake.connReturns = struct {
		result1 ssh.Conn
	}{result1}
}

func (fake *FakeSecureClient) ConnReturnsOnCall(i int, result1 ssh.Conn) {
	fake.connMutex.Lock()
	defer fake.connMutex.Unlock()
	fake.ConnStub = nil
	if fake.connReturnsOnCall == nil {
		fake.connReturnsOnCall = make(map[int]struct {
			result1 ssh.Conn
		})
	}
	fake.connReturnsOnCall[i] = struct {
		result1 ssh.Conn
	}{result1}
}

func (fake *FakeSecureClient) Dial(arg1 string, arg2 string) (net.Conn, error) {
	fake.dialMutex.Lock()
	ret, specificReturn := fake.dialReturnsOnCall[len(fake.dialArgsForCall)]
	fake.dialArgsForCall = append(fake.dialArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Dial", []interface{}{arg1, arg2})
	fake.dialMutex.Unlock()
	if fake.DialStub != nil {
		return fake.DialStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.dialReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecureClient) DialCallCount() int {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return len(fake.dialArgsForCall)
}

func (fake *FakeSecureClient) DialCalls(stub func(string, string) (net.Conn, error)) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = stub
}

func (fake *FakeSecureClient) DialArgsForCall(i int) (string, string) {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	argsForCall := fake.dialArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecureClient) DialReturns(result1 net.Conn, result2 error) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = nil
	fake.dialReturns = struct {
		result1 net.Conn
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureClient) DialReturnsOnCall(i int, result1 net.Conn, result2 error) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = nil
	if fake.dialReturnsOnCall == nil {
		fake.dialReturnsOnCall = make(map[int]struct {
			result1 net.Conn
			result2 error
		})
	}
	fake.dialReturnsOnCall[i] = struct {
		result1 net.Conn
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureClient) NewSession() (clissh.SecureSession, error) {
	fake.newSessionMutex.Lock()
	ret, specificReturn := fake.newSessionReturnsOnCall[len(fake.newSessionArgsForCall)]
	fake.newSessionArgsForCall = append(fake.newSessionArgsForCall, struct {
	}{})
	fake.recordInvocation("NewSession", []interface{}{})
	fake.newSessionMutex.Unlock()
	if fake.NewSessionStub != nil {
		return fake.NewSessionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newSessionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecureClient) NewSessionCallCount() int {
	fake.newSessionMutex.RLock()
	defer fake.newSessionMutex.RUnlock()
	return len(fake.newSessionArgsForCall)
}

func (fake *FakeSecureClient) NewSessionCalls(stub func() (clissh.SecureSession, error)) {
	fake.newSessionMutex.Lock()
	defer fake.newSessionMutex.Unlock()
	fake.NewSessionStub = stub
}

func (fake *FakeSecureClient) NewSessionReturns(result1 clissh.SecureSession, result2 error) {
	fake.newSessionMutex.Lock()
	defer fake.newSessionMutex.Unlock()
	fake.NewSessionStub = nil
	fake.newSessionReturns = struct {
		result1 clissh.SecureSession
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureClient) NewSessionReturnsOnCall(i int, result1 clissh.SecureSession, result2 error) {
	fake.newSessionMutex.Lock()
	defer fake.newSessionMutex.Unlock()
	fake.NewSessionStub = nil
	if fake.newSessionReturnsOnCall == nil {
		fake.newSessionReturnsOnCall = make(map[int]struct {
			result1 clissh.SecureSession
			result2 error
		})
	}
	fake.newSessionReturnsOnCall[i] = struct {
		result1 clissh.SecureSession
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureClient) Wait() error {
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
		return ret.result1
	}
	fakeReturns := fake.waitReturns
	return fakeReturns.result1
}

func (fake *FakeSecureClient) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeSecureClient) WaitCalls(stub func() error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeSecureClient) WaitReturns(result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureClient) WaitReturnsOnCall(i int, result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.connMutex.RLock()
	defer fake.connMutex.RUnlock()
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	fake.newSessionMutex.RLock()
	defer fake.newSessionMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecureClient) recordInvocation(key string, args []interface{}) {
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

var _ clissh.SecureClient = new(FakeSecureClient)
