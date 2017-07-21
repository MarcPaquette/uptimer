// Code generated by counterfeiter. DO NOT EDIT.
package measurementfakes

import (
	"sync"
	"time"

	"github.com/cloudfoundry/uptimer/measurement"
)

type FakeResultSet struct {
	RecordSuccessStub                    func()
	recordSuccessMutex                   sync.RWMutex
	recordSuccessArgsForCall             []struct{}
	RecordFailureStub                    func()
	recordFailureMutex                   sync.RWMutex
	recordFailureArgsForCall             []struct{}
	SuccessesSinceLastFailureStub        func() (int, time.Time)
	successesSinceLastFailureMutex       sync.RWMutex
	successesSinceLastFailureArgsForCall []struct{}
	successesSinceLastFailureReturns     struct {
		result1 int
		result2 time.Time
	}
	successesSinceLastFailureReturnsOnCall map[int]struct {
		result1 int
		result2 time.Time
	}
	SuccessfulStub        func() int
	successfulMutex       sync.RWMutex
	successfulArgsForCall []struct{}
	successfulReturns     struct {
		result1 int
	}
	successfulReturnsOnCall map[int]struct {
		result1 int
	}
	FailedStub        func() int
	failedMutex       sync.RWMutex
	failedArgsForCall []struct{}
	failedReturns     struct {
		result1 int
	}
	failedReturnsOnCall map[int]struct {
		result1 int
	}
	TotalStub        func() int
	totalMutex       sync.RWMutex
	totalArgsForCall []struct{}
	totalReturns     struct {
		result1 int
	}
	totalReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResultSet) RecordSuccess() {
	fake.recordSuccessMutex.Lock()
	fake.recordSuccessArgsForCall = append(fake.recordSuccessArgsForCall, struct{}{})
	fake.recordInvocation("RecordSuccess", []interface{}{})
	fake.recordSuccessMutex.Unlock()
	if fake.RecordSuccessStub != nil {
		fake.RecordSuccessStub()
	}
}

func (fake *FakeResultSet) RecordSuccessCallCount() int {
	fake.recordSuccessMutex.RLock()
	defer fake.recordSuccessMutex.RUnlock()
	return len(fake.recordSuccessArgsForCall)
}

func (fake *FakeResultSet) RecordFailure() {
	fake.recordFailureMutex.Lock()
	fake.recordFailureArgsForCall = append(fake.recordFailureArgsForCall, struct{}{})
	fake.recordInvocation("RecordFailure", []interface{}{})
	fake.recordFailureMutex.Unlock()
	if fake.RecordFailureStub != nil {
		fake.RecordFailureStub()
	}
}

func (fake *FakeResultSet) RecordFailureCallCount() int {
	fake.recordFailureMutex.RLock()
	defer fake.recordFailureMutex.RUnlock()
	return len(fake.recordFailureArgsForCall)
}

func (fake *FakeResultSet) SuccessesSinceLastFailure() (int, time.Time) {
	fake.successesSinceLastFailureMutex.Lock()
	ret, specificReturn := fake.successesSinceLastFailureReturnsOnCall[len(fake.successesSinceLastFailureArgsForCall)]
	fake.successesSinceLastFailureArgsForCall = append(fake.successesSinceLastFailureArgsForCall, struct{}{})
	fake.recordInvocation("SuccessesSinceLastFailure", []interface{}{})
	fake.successesSinceLastFailureMutex.Unlock()
	if fake.SuccessesSinceLastFailureStub != nil {
		return fake.SuccessesSinceLastFailureStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.successesSinceLastFailureReturns.result1, fake.successesSinceLastFailureReturns.result2
}

func (fake *FakeResultSet) SuccessesSinceLastFailureCallCount() int {
	fake.successesSinceLastFailureMutex.RLock()
	defer fake.successesSinceLastFailureMutex.RUnlock()
	return len(fake.successesSinceLastFailureArgsForCall)
}

func (fake *FakeResultSet) SuccessesSinceLastFailureReturns(result1 int, result2 time.Time) {
	fake.SuccessesSinceLastFailureStub = nil
	fake.successesSinceLastFailureReturns = struct {
		result1 int
		result2 time.Time
	}{result1, result2}
}

func (fake *FakeResultSet) SuccessesSinceLastFailureReturnsOnCall(i int, result1 int, result2 time.Time) {
	fake.SuccessesSinceLastFailureStub = nil
	if fake.successesSinceLastFailureReturnsOnCall == nil {
		fake.successesSinceLastFailureReturnsOnCall = make(map[int]struct {
			result1 int
			result2 time.Time
		})
	}
	fake.successesSinceLastFailureReturnsOnCall[i] = struct {
		result1 int
		result2 time.Time
	}{result1, result2}
}

func (fake *FakeResultSet) Successful() int {
	fake.successfulMutex.Lock()
	ret, specificReturn := fake.successfulReturnsOnCall[len(fake.successfulArgsForCall)]
	fake.successfulArgsForCall = append(fake.successfulArgsForCall, struct{}{})
	fake.recordInvocation("Successful", []interface{}{})
	fake.successfulMutex.Unlock()
	if fake.SuccessfulStub != nil {
		return fake.SuccessfulStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.successfulReturns.result1
}

func (fake *FakeResultSet) SuccessfulCallCount() int {
	fake.successfulMutex.RLock()
	defer fake.successfulMutex.RUnlock()
	return len(fake.successfulArgsForCall)
}

func (fake *FakeResultSet) SuccessfulReturns(result1 int) {
	fake.SuccessfulStub = nil
	fake.successfulReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) SuccessfulReturnsOnCall(i int, result1 int) {
	fake.SuccessfulStub = nil
	if fake.successfulReturnsOnCall == nil {
		fake.successfulReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.successfulReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) Failed() int {
	fake.failedMutex.Lock()
	ret, specificReturn := fake.failedReturnsOnCall[len(fake.failedArgsForCall)]
	fake.failedArgsForCall = append(fake.failedArgsForCall, struct{}{})
	fake.recordInvocation("Failed", []interface{}{})
	fake.failedMutex.Unlock()
	if fake.FailedStub != nil {
		return fake.FailedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.failedReturns.result1
}

func (fake *FakeResultSet) FailedCallCount() int {
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	return len(fake.failedArgsForCall)
}

func (fake *FakeResultSet) FailedReturns(result1 int) {
	fake.FailedStub = nil
	fake.failedReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) FailedReturnsOnCall(i int, result1 int) {
	fake.FailedStub = nil
	if fake.failedReturnsOnCall == nil {
		fake.failedReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.failedReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) Total() int {
	fake.totalMutex.Lock()
	ret, specificReturn := fake.totalReturnsOnCall[len(fake.totalArgsForCall)]
	fake.totalArgsForCall = append(fake.totalArgsForCall, struct{}{})
	fake.recordInvocation("Total", []interface{}{})
	fake.totalMutex.Unlock()
	if fake.TotalStub != nil {
		return fake.TotalStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.totalReturns.result1
}

func (fake *FakeResultSet) TotalCallCount() int {
	fake.totalMutex.RLock()
	defer fake.totalMutex.RUnlock()
	return len(fake.totalArgsForCall)
}

func (fake *FakeResultSet) TotalReturns(result1 int) {
	fake.TotalStub = nil
	fake.totalReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) TotalReturnsOnCall(i int, result1 int) {
	fake.TotalStub = nil
	if fake.totalReturnsOnCall == nil {
		fake.totalReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.totalReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResultSet) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.recordSuccessMutex.RLock()
	defer fake.recordSuccessMutex.RUnlock()
	fake.recordFailureMutex.RLock()
	defer fake.recordFailureMutex.RUnlock()
	fake.successesSinceLastFailureMutex.RLock()
	defer fake.successesSinceLastFailureMutex.RUnlock()
	fake.successfulMutex.RLock()
	defer fake.successfulMutex.RUnlock()
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	fake.totalMutex.RLock()
	defer fake.totalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResultSet) recordInvocation(key string, args []interface{}) {
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

var _ measurement.ResultSet = new(FakeResultSet)
