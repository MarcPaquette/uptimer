// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry/uptimer/cfCmdGenerator"
	"github.com/cloudfoundry/uptimer/cmdStartWaiter"
)

type FakeCfCmdGenerator struct {
	ApiStub        func(url string) cmdStartWaiter.CmdStartWaiter
	apiMutex       sync.RWMutex
	apiArgsForCall []struct {
		url string
	}
	apiReturns struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	apiReturnsOnCall map[int]struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	AuthStub        func(username, password string) cmdStartWaiter.CmdStartWaiter
	authMutex       sync.RWMutex
	authArgsForCall []struct {
		username string
		password string
	}
	authReturns struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	authReturnsOnCall map[int]struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	CreateOrgStub        func(org string) cmdStartWaiter.CmdStartWaiter
	createOrgMutex       sync.RWMutex
	createOrgArgsForCall []struct {
		org string
	}
	createOrgReturns struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	createOrgReturnsOnCall map[int]struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	CreateSpaceStub        func(org, space string) cmdStartWaiter.CmdStartWaiter
	createSpaceMutex       sync.RWMutex
	createSpaceArgsForCall []struct {
		org   string
		space string
	}
	createSpaceReturns struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	createSpaceReturnsOnCall map[int]struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	TargetStub        func(org, space string) cmdStartWaiter.CmdStartWaiter
	targetMutex       sync.RWMutex
	targetArgsForCall []struct {
		org   string
		space string
	}
	targetReturns struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	targetReturnsOnCall map[int]struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	PushStub        func(name, path string) cmdStartWaiter.CmdStartWaiter
	pushMutex       sync.RWMutex
	pushArgsForCall []struct {
		name string
		path string
	}
	pushReturns struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	pushReturnsOnCall map[int]struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	DeleteStub        func(name string) cmdStartWaiter.CmdStartWaiter
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		name string
	}
	deleteReturns struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	deleteReturnsOnCall map[int]struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	DeleteOrgStub        func(org string) cmdStartWaiter.CmdStartWaiter
	deleteOrgMutex       sync.RWMutex
	deleteOrgArgsForCall []struct {
		org string
	}
	deleteOrgReturns struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	deleteOrgReturnsOnCall map[int]struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	LogOutStub        func() cmdStartWaiter.CmdStartWaiter
	logOutMutex       sync.RWMutex
	logOutArgsForCall []struct{}
	logOutReturns     struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	logOutReturnsOnCall map[int]struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	RecentLogsStub        func(appName string) cmdStartWaiter.CmdStartWaiter
	recentLogsMutex       sync.RWMutex
	recentLogsArgsForCall []struct {
		appName string
	}
	recentLogsReturns struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	recentLogsReturnsOnCall map[int]struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCfCmdGenerator) Api(url string) cmdStartWaiter.CmdStartWaiter {
	fake.apiMutex.Lock()
	ret, specificReturn := fake.apiReturnsOnCall[len(fake.apiArgsForCall)]
	fake.apiArgsForCall = append(fake.apiArgsForCall, struct {
		url string
	}{url})
	fake.recordInvocation("Api", []interface{}{url})
	fake.apiMutex.Unlock()
	if fake.ApiStub != nil {
		return fake.ApiStub(url)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.apiReturns.result1
}

func (fake *FakeCfCmdGenerator) ApiCallCount() int {
	fake.apiMutex.RLock()
	defer fake.apiMutex.RUnlock()
	return len(fake.apiArgsForCall)
}

func (fake *FakeCfCmdGenerator) ApiArgsForCall(i int) string {
	fake.apiMutex.RLock()
	defer fake.apiMutex.RUnlock()
	return fake.apiArgsForCall[i].url
}

func (fake *FakeCfCmdGenerator) ApiReturns(result1 cmdStartWaiter.CmdStartWaiter) {
	fake.ApiStub = nil
	fake.apiReturns = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) ApiReturnsOnCall(i int, result1 cmdStartWaiter.CmdStartWaiter) {
	fake.ApiStub = nil
	if fake.apiReturnsOnCall == nil {
		fake.apiReturnsOnCall = make(map[int]struct {
			result1 cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.apiReturnsOnCall[i] = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) Auth(username string, password string) cmdStartWaiter.CmdStartWaiter {
	fake.authMutex.Lock()
	ret, specificReturn := fake.authReturnsOnCall[len(fake.authArgsForCall)]
	fake.authArgsForCall = append(fake.authArgsForCall, struct {
		username string
		password string
	}{username, password})
	fake.recordInvocation("Auth", []interface{}{username, password})
	fake.authMutex.Unlock()
	if fake.AuthStub != nil {
		return fake.AuthStub(username, password)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.authReturns.result1
}

func (fake *FakeCfCmdGenerator) AuthCallCount() int {
	fake.authMutex.RLock()
	defer fake.authMutex.RUnlock()
	return len(fake.authArgsForCall)
}

func (fake *FakeCfCmdGenerator) AuthArgsForCall(i int) (string, string) {
	fake.authMutex.RLock()
	defer fake.authMutex.RUnlock()
	return fake.authArgsForCall[i].username, fake.authArgsForCall[i].password
}

func (fake *FakeCfCmdGenerator) AuthReturns(result1 cmdStartWaiter.CmdStartWaiter) {
	fake.AuthStub = nil
	fake.authReturns = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) AuthReturnsOnCall(i int, result1 cmdStartWaiter.CmdStartWaiter) {
	fake.AuthStub = nil
	if fake.authReturnsOnCall == nil {
		fake.authReturnsOnCall = make(map[int]struct {
			result1 cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.authReturnsOnCall[i] = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) CreateOrg(org string) cmdStartWaiter.CmdStartWaiter {
	fake.createOrgMutex.Lock()
	ret, specificReturn := fake.createOrgReturnsOnCall[len(fake.createOrgArgsForCall)]
	fake.createOrgArgsForCall = append(fake.createOrgArgsForCall, struct {
		org string
	}{org})
	fake.recordInvocation("CreateOrg", []interface{}{org})
	fake.createOrgMutex.Unlock()
	if fake.CreateOrgStub != nil {
		return fake.CreateOrgStub(org)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createOrgReturns.result1
}

func (fake *FakeCfCmdGenerator) CreateOrgCallCount() int {
	fake.createOrgMutex.RLock()
	defer fake.createOrgMutex.RUnlock()
	return len(fake.createOrgArgsForCall)
}

func (fake *FakeCfCmdGenerator) CreateOrgArgsForCall(i int) string {
	fake.createOrgMutex.RLock()
	defer fake.createOrgMutex.RUnlock()
	return fake.createOrgArgsForCall[i].org
}

func (fake *FakeCfCmdGenerator) CreateOrgReturns(result1 cmdStartWaiter.CmdStartWaiter) {
	fake.CreateOrgStub = nil
	fake.createOrgReturns = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) CreateOrgReturnsOnCall(i int, result1 cmdStartWaiter.CmdStartWaiter) {
	fake.CreateOrgStub = nil
	if fake.createOrgReturnsOnCall == nil {
		fake.createOrgReturnsOnCall = make(map[int]struct {
			result1 cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.createOrgReturnsOnCall[i] = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) CreateSpace(org string, space string) cmdStartWaiter.CmdStartWaiter {
	fake.createSpaceMutex.Lock()
	ret, specificReturn := fake.createSpaceReturnsOnCall[len(fake.createSpaceArgsForCall)]
	fake.createSpaceArgsForCall = append(fake.createSpaceArgsForCall, struct {
		org   string
		space string
	}{org, space})
	fake.recordInvocation("CreateSpace", []interface{}{org, space})
	fake.createSpaceMutex.Unlock()
	if fake.CreateSpaceStub != nil {
		return fake.CreateSpaceStub(org, space)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createSpaceReturns.result1
}

func (fake *FakeCfCmdGenerator) CreateSpaceCallCount() int {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	return len(fake.createSpaceArgsForCall)
}

func (fake *FakeCfCmdGenerator) CreateSpaceArgsForCall(i int) (string, string) {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	return fake.createSpaceArgsForCall[i].org, fake.createSpaceArgsForCall[i].space
}

func (fake *FakeCfCmdGenerator) CreateSpaceReturns(result1 cmdStartWaiter.CmdStartWaiter) {
	fake.CreateSpaceStub = nil
	fake.createSpaceReturns = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) CreateSpaceReturnsOnCall(i int, result1 cmdStartWaiter.CmdStartWaiter) {
	fake.CreateSpaceStub = nil
	if fake.createSpaceReturnsOnCall == nil {
		fake.createSpaceReturnsOnCall = make(map[int]struct {
			result1 cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.createSpaceReturnsOnCall[i] = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) Target(org string, space string) cmdStartWaiter.CmdStartWaiter {
	fake.targetMutex.Lock()
	ret, specificReturn := fake.targetReturnsOnCall[len(fake.targetArgsForCall)]
	fake.targetArgsForCall = append(fake.targetArgsForCall, struct {
		org   string
		space string
	}{org, space})
	fake.recordInvocation("Target", []interface{}{org, space})
	fake.targetMutex.Unlock()
	if fake.TargetStub != nil {
		return fake.TargetStub(org, space)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.targetReturns.result1
}

func (fake *FakeCfCmdGenerator) TargetCallCount() int {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return len(fake.targetArgsForCall)
}

func (fake *FakeCfCmdGenerator) TargetArgsForCall(i int) (string, string) {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return fake.targetArgsForCall[i].org, fake.targetArgsForCall[i].space
}

func (fake *FakeCfCmdGenerator) TargetReturns(result1 cmdStartWaiter.CmdStartWaiter) {
	fake.TargetStub = nil
	fake.targetReturns = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) TargetReturnsOnCall(i int, result1 cmdStartWaiter.CmdStartWaiter) {
	fake.TargetStub = nil
	if fake.targetReturnsOnCall == nil {
		fake.targetReturnsOnCall = make(map[int]struct {
			result1 cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.targetReturnsOnCall[i] = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) Push(name string, path string) cmdStartWaiter.CmdStartWaiter {
	fake.pushMutex.Lock()
	ret, specificReturn := fake.pushReturnsOnCall[len(fake.pushArgsForCall)]
	fake.pushArgsForCall = append(fake.pushArgsForCall, struct {
		name string
		path string
	}{name, path})
	fake.recordInvocation("Push", []interface{}{name, path})
	fake.pushMutex.Unlock()
	if fake.PushStub != nil {
		return fake.PushStub(name, path)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pushReturns.result1
}

func (fake *FakeCfCmdGenerator) PushCallCount() int {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return len(fake.pushArgsForCall)
}

func (fake *FakeCfCmdGenerator) PushArgsForCall(i int) (string, string) {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return fake.pushArgsForCall[i].name, fake.pushArgsForCall[i].path
}

func (fake *FakeCfCmdGenerator) PushReturns(result1 cmdStartWaiter.CmdStartWaiter) {
	fake.PushStub = nil
	fake.pushReturns = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) PushReturnsOnCall(i int, result1 cmdStartWaiter.CmdStartWaiter) {
	fake.PushStub = nil
	if fake.pushReturnsOnCall == nil {
		fake.pushReturnsOnCall = make(map[int]struct {
			result1 cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.pushReturnsOnCall[i] = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) Delete(name string) cmdStartWaiter.CmdStartWaiter {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Delete", []interface{}{name})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeCfCmdGenerator) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeCfCmdGenerator) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].name
}

func (fake *FakeCfCmdGenerator) DeleteReturns(result1 cmdStartWaiter.CmdStartWaiter) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) DeleteReturnsOnCall(i int, result1 cmdStartWaiter.CmdStartWaiter) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) DeleteOrg(org string) cmdStartWaiter.CmdStartWaiter {
	fake.deleteOrgMutex.Lock()
	ret, specificReturn := fake.deleteOrgReturnsOnCall[len(fake.deleteOrgArgsForCall)]
	fake.deleteOrgArgsForCall = append(fake.deleteOrgArgsForCall, struct {
		org string
	}{org})
	fake.recordInvocation("DeleteOrg", []interface{}{org})
	fake.deleteOrgMutex.Unlock()
	if fake.DeleteOrgStub != nil {
		return fake.DeleteOrgStub(org)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteOrgReturns.result1
}

func (fake *FakeCfCmdGenerator) DeleteOrgCallCount() int {
	fake.deleteOrgMutex.RLock()
	defer fake.deleteOrgMutex.RUnlock()
	return len(fake.deleteOrgArgsForCall)
}

func (fake *FakeCfCmdGenerator) DeleteOrgArgsForCall(i int) string {
	fake.deleteOrgMutex.RLock()
	defer fake.deleteOrgMutex.RUnlock()
	return fake.deleteOrgArgsForCall[i].org
}

func (fake *FakeCfCmdGenerator) DeleteOrgReturns(result1 cmdStartWaiter.CmdStartWaiter) {
	fake.DeleteOrgStub = nil
	fake.deleteOrgReturns = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) DeleteOrgReturnsOnCall(i int, result1 cmdStartWaiter.CmdStartWaiter) {
	fake.DeleteOrgStub = nil
	if fake.deleteOrgReturnsOnCall == nil {
		fake.deleteOrgReturnsOnCall = make(map[int]struct {
			result1 cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.deleteOrgReturnsOnCall[i] = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) LogOut() cmdStartWaiter.CmdStartWaiter {
	fake.logOutMutex.Lock()
	ret, specificReturn := fake.logOutReturnsOnCall[len(fake.logOutArgsForCall)]
	fake.logOutArgsForCall = append(fake.logOutArgsForCall, struct{}{})
	fake.recordInvocation("LogOut", []interface{}{})
	fake.logOutMutex.Unlock()
	if fake.LogOutStub != nil {
		return fake.LogOutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.logOutReturns.result1
}

func (fake *FakeCfCmdGenerator) LogOutCallCount() int {
	fake.logOutMutex.RLock()
	defer fake.logOutMutex.RUnlock()
	return len(fake.logOutArgsForCall)
}

func (fake *FakeCfCmdGenerator) LogOutReturns(result1 cmdStartWaiter.CmdStartWaiter) {
	fake.LogOutStub = nil
	fake.logOutReturns = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) LogOutReturnsOnCall(i int, result1 cmdStartWaiter.CmdStartWaiter) {
	fake.LogOutStub = nil
	if fake.logOutReturnsOnCall == nil {
		fake.logOutReturnsOnCall = make(map[int]struct {
			result1 cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.logOutReturnsOnCall[i] = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) RecentLogs(appName string) cmdStartWaiter.CmdStartWaiter {
	fake.recentLogsMutex.Lock()
	ret, specificReturn := fake.recentLogsReturnsOnCall[len(fake.recentLogsArgsForCall)]
	fake.recentLogsArgsForCall = append(fake.recentLogsArgsForCall, struct {
		appName string
	}{appName})
	fake.recordInvocation("RecentLogs", []interface{}{appName})
	fake.recentLogsMutex.Unlock()
	if fake.RecentLogsStub != nil {
		return fake.RecentLogsStub(appName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.recentLogsReturns.result1
}

func (fake *FakeCfCmdGenerator) RecentLogsCallCount() int {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	return len(fake.recentLogsArgsForCall)
}

func (fake *FakeCfCmdGenerator) RecentLogsArgsForCall(i int) string {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	return fake.recentLogsArgsForCall[i].appName
}

func (fake *FakeCfCmdGenerator) RecentLogsReturns(result1 cmdStartWaiter.CmdStartWaiter) {
	fake.RecentLogsStub = nil
	fake.recentLogsReturns = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) RecentLogsReturnsOnCall(i int, result1 cmdStartWaiter.CmdStartWaiter) {
	fake.RecentLogsStub = nil
	if fake.recentLogsReturnsOnCall == nil {
		fake.recentLogsReturnsOnCall = make(map[int]struct {
			result1 cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.recentLogsReturnsOnCall[i] = struct {
		result1 cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfCmdGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.apiMutex.RLock()
	defer fake.apiMutex.RUnlock()
	fake.authMutex.RLock()
	defer fake.authMutex.RUnlock()
	fake.createOrgMutex.RLock()
	defer fake.createOrgMutex.RUnlock()
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteOrgMutex.RLock()
	defer fake.deleteOrgMutex.RUnlock()
	fake.logOutMutex.RLock()
	defer fake.logOutMutex.RUnlock()
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCfCmdGenerator) recordInvocation(key string, args []interface{}) {
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

var _ cfCmdGenerator.CfCmdGenerator = new(FakeCfCmdGenerator)
