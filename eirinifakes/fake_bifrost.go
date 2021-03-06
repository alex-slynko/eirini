// Code generated by counterfeiter. DO NOT EDIT.
package eirinifakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/eirini"
	"code.cloudfoundry.org/eirini/models/cf"
	"code.cloudfoundry.org/eirini/opi"
)

type FakeBifrost struct {
	TransferStub        func(ctx context.Context, request cf.DesireLRPRequest) error
	transferMutex       sync.RWMutex
	transferArgsForCall []struct {
		ctx     context.Context
		request cf.DesireLRPRequest
	}
	transferReturns struct {
		result1 error
	}
	transferReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(ctx context.Context) ([]*models.DesiredLRPSchedulingInfo, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		ctx context.Context
	}
	listReturns struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}
	UpdateStub        func(ctx context.Context, update cf.UpdateDesiredLRPRequest) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		ctx    context.Context
		update cf.UpdateDesiredLRPRequest
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func(ctx context.Context, identifier opi.LRPIdentifier) error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		ctx        context.Context
		identifier opi.LRPIdentifier
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	GetAppStub        func(ctx context.Context, identifier opi.LRPIdentifier) *models.DesiredLRP
	getAppMutex       sync.RWMutex
	getAppArgsForCall []struct {
		ctx        context.Context
		identifier opi.LRPIdentifier
	}
	getAppReturns struct {
		result1 *models.DesiredLRP
	}
	getAppReturnsOnCall map[int]struct {
		result1 *models.DesiredLRP
	}
	GetInstancesStub        func(ctx context.Context, identifier opi.LRPIdentifier) ([]*cf.Instance, error)
	getInstancesMutex       sync.RWMutex
	getInstancesArgsForCall []struct {
		ctx        context.Context
		identifier opi.LRPIdentifier
	}
	getInstancesReturns struct {
		result1 []*cf.Instance
		result2 error
	}
	getInstancesReturnsOnCall map[int]struct {
		result1 []*cf.Instance
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBifrost) Transfer(ctx context.Context, request cf.DesireLRPRequest) error {
	fake.transferMutex.Lock()
	ret, specificReturn := fake.transferReturnsOnCall[len(fake.transferArgsForCall)]
	fake.transferArgsForCall = append(fake.transferArgsForCall, struct {
		ctx     context.Context
		request cf.DesireLRPRequest
	}{ctx, request})
	fake.recordInvocation("Transfer", []interface{}{ctx, request})
	fake.transferMutex.Unlock()
	if fake.TransferStub != nil {
		return fake.TransferStub(ctx, request)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.transferReturns.result1
}

func (fake *FakeBifrost) TransferCallCount() int {
	fake.transferMutex.RLock()
	defer fake.transferMutex.RUnlock()
	return len(fake.transferArgsForCall)
}

func (fake *FakeBifrost) TransferArgsForCall(i int) (context.Context, cf.DesireLRPRequest) {
	fake.transferMutex.RLock()
	defer fake.transferMutex.RUnlock()
	return fake.transferArgsForCall[i].ctx, fake.transferArgsForCall[i].request
}

func (fake *FakeBifrost) TransferReturns(result1 error) {
	fake.TransferStub = nil
	fake.transferReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBifrost) TransferReturnsOnCall(i int, result1 error) {
	fake.TransferStub = nil
	if fake.transferReturnsOnCall == nil {
		fake.transferReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.transferReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBifrost) List(ctx context.Context) ([]*models.DesiredLRPSchedulingInfo, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		ctx context.Context
	}{ctx})
	fake.recordInvocation("List", []interface{}{ctx})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(ctx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listReturns.result1, fake.listReturns.result2
}

func (fake *FakeBifrost) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeBifrost) ListArgsForCall(i int) context.Context {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].ctx
}

func (fake *FakeBifrost) ListReturns(result1 []*models.DesiredLRPSchedulingInfo, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeBifrost) ListReturnsOnCall(i int, result1 []*models.DesiredLRPSchedulingInfo, result2 error) {
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []*models.DesiredLRPSchedulingInfo
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeBifrost) Update(ctx context.Context, update cf.UpdateDesiredLRPRequest) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		ctx    context.Context
		update cf.UpdateDesiredLRPRequest
	}{ctx, update})
	fake.recordInvocation("Update", []interface{}{ctx, update})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(ctx, update)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateReturns.result1
}

func (fake *FakeBifrost) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeBifrost) UpdateArgsForCall(i int) (context.Context, cf.UpdateDesiredLRPRequest) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].ctx, fake.updateArgsForCall[i].update
}

func (fake *FakeBifrost) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBifrost) UpdateReturnsOnCall(i int, result1 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBifrost) Stop(ctx context.Context, identifier opi.LRPIdentifier) error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		ctx        context.Context
		identifier opi.LRPIdentifier
	}{ctx, identifier})
	fake.recordInvocation("Stop", []interface{}{ctx, identifier})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub(ctx, identifier)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stopReturns.result1
}

func (fake *FakeBifrost) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeBifrost) StopArgsForCall(i int) (context.Context, opi.LRPIdentifier) {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.stopArgsForCall[i].ctx, fake.stopArgsForCall[i].identifier
}

func (fake *FakeBifrost) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBifrost) StopReturnsOnCall(i int, result1 error) {
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

func (fake *FakeBifrost) GetApp(ctx context.Context, identifier opi.LRPIdentifier) *models.DesiredLRP {
	fake.getAppMutex.Lock()
	ret, specificReturn := fake.getAppReturnsOnCall[len(fake.getAppArgsForCall)]
	fake.getAppArgsForCall = append(fake.getAppArgsForCall, struct {
		ctx        context.Context
		identifier opi.LRPIdentifier
	}{ctx, identifier})
	fake.recordInvocation("GetApp", []interface{}{ctx, identifier})
	fake.getAppMutex.Unlock()
	if fake.GetAppStub != nil {
		return fake.GetAppStub(ctx, identifier)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getAppReturns.result1
}

func (fake *FakeBifrost) GetAppCallCount() int {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return len(fake.getAppArgsForCall)
}

func (fake *FakeBifrost) GetAppArgsForCall(i int) (context.Context, opi.LRPIdentifier) {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return fake.getAppArgsForCall[i].ctx, fake.getAppArgsForCall[i].identifier
}

func (fake *FakeBifrost) GetAppReturns(result1 *models.DesiredLRP) {
	fake.GetAppStub = nil
	fake.getAppReturns = struct {
		result1 *models.DesiredLRP
	}{result1}
}

func (fake *FakeBifrost) GetAppReturnsOnCall(i int, result1 *models.DesiredLRP) {
	fake.GetAppStub = nil
	if fake.getAppReturnsOnCall == nil {
		fake.getAppReturnsOnCall = make(map[int]struct {
			result1 *models.DesiredLRP
		})
	}
	fake.getAppReturnsOnCall[i] = struct {
		result1 *models.DesiredLRP
	}{result1}
}

func (fake *FakeBifrost) GetInstances(ctx context.Context, identifier opi.LRPIdentifier) ([]*cf.Instance, error) {
	fake.getInstancesMutex.Lock()
	ret, specificReturn := fake.getInstancesReturnsOnCall[len(fake.getInstancesArgsForCall)]
	fake.getInstancesArgsForCall = append(fake.getInstancesArgsForCall, struct {
		ctx        context.Context
		identifier opi.LRPIdentifier
	}{ctx, identifier})
	fake.recordInvocation("GetInstances", []interface{}{ctx, identifier})
	fake.getInstancesMutex.Unlock()
	if fake.GetInstancesStub != nil {
		return fake.GetInstancesStub(ctx, identifier)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getInstancesReturns.result1, fake.getInstancesReturns.result2
}

func (fake *FakeBifrost) GetInstancesCallCount() int {
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	return len(fake.getInstancesArgsForCall)
}

func (fake *FakeBifrost) GetInstancesArgsForCall(i int) (context.Context, opi.LRPIdentifier) {
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	return fake.getInstancesArgsForCall[i].ctx, fake.getInstancesArgsForCall[i].identifier
}

func (fake *FakeBifrost) GetInstancesReturns(result1 []*cf.Instance, result2 error) {
	fake.GetInstancesStub = nil
	fake.getInstancesReturns = struct {
		result1 []*cf.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeBifrost) GetInstancesReturnsOnCall(i int, result1 []*cf.Instance, result2 error) {
	fake.GetInstancesStub = nil
	if fake.getInstancesReturnsOnCall == nil {
		fake.getInstancesReturnsOnCall = make(map[int]struct {
			result1 []*cf.Instance
			result2 error
		})
	}
	fake.getInstancesReturnsOnCall[i] = struct {
		result1 []*cf.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeBifrost) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.transferMutex.RLock()
	defer fake.transferMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBifrost) recordInvocation(key string, args []interface{}) {
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

var _ eirini.Bifrost = new(FakeBifrost)
