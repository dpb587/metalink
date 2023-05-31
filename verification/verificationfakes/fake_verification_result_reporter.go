// Code generated by counterfeiter. DO NOT EDIT.
package verificationfakes

import (
	"sync"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/verification"
)

type FakeVerificationResultReporter struct {
	ReportVerificationResultStub        func(metalink.File, verification.VerificationResult) error
	reportVerificationResultMutex       sync.RWMutex
	reportVerificationResultArgsForCall []struct {
		arg1 metalink.File
		arg2 verification.VerificationResult
	}
	reportVerificationResultReturns struct {
		result1 error
	}
	reportVerificationResultReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVerificationResultReporter) ReportVerificationResult(arg1 metalink.File, arg2 verification.VerificationResult) error {
	fake.reportVerificationResultMutex.Lock()
	ret, specificReturn := fake.reportVerificationResultReturnsOnCall[len(fake.reportVerificationResultArgsForCall)]
	fake.reportVerificationResultArgsForCall = append(fake.reportVerificationResultArgsForCall, struct {
		arg1 metalink.File
		arg2 verification.VerificationResult
	}{arg1, arg2})
	stub := fake.ReportVerificationResultStub
	fakeReturns := fake.reportVerificationResultReturns
	fake.recordInvocation("ReportVerificationResult", []interface{}{arg1, arg2})
	fake.reportVerificationResultMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVerificationResultReporter) ReportVerificationResultCallCount() int {
	fake.reportVerificationResultMutex.RLock()
	defer fake.reportVerificationResultMutex.RUnlock()
	return len(fake.reportVerificationResultArgsForCall)
}

func (fake *FakeVerificationResultReporter) ReportVerificationResultCalls(stub func(metalink.File, verification.VerificationResult) error) {
	fake.reportVerificationResultMutex.Lock()
	defer fake.reportVerificationResultMutex.Unlock()
	fake.ReportVerificationResultStub = stub
}

func (fake *FakeVerificationResultReporter) ReportVerificationResultArgsForCall(i int) (metalink.File, verification.VerificationResult) {
	fake.reportVerificationResultMutex.RLock()
	defer fake.reportVerificationResultMutex.RUnlock()
	argsForCall := fake.reportVerificationResultArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVerificationResultReporter) ReportVerificationResultReturns(result1 error) {
	fake.reportVerificationResultMutex.Lock()
	defer fake.reportVerificationResultMutex.Unlock()
	fake.ReportVerificationResultStub = nil
	fake.reportVerificationResultReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVerificationResultReporter) ReportVerificationResultReturnsOnCall(i int, result1 error) {
	fake.reportVerificationResultMutex.Lock()
	defer fake.reportVerificationResultMutex.Unlock()
	fake.ReportVerificationResultStub = nil
	if fake.reportVerificationResultReturnsOnCall == nil {
		fake.reportVerificationResultReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.reportVerificationResultReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVerificationResultReporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.reportVerificationResultMutex.RLock()
	defer fake.reportVerificationResultMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVerificationResultReporter) recordInvocation(key string, args []interface{}) {
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

var _ verification.VerificationResultReporter = new(FakeVerificationResultReporter)
