// Code generated by counterfeiter. DO NOT EDIT.
package verificationfakes

import (
	"sync"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/verification"
)

type FakeVerifier struct {
	VerifyStub        func(file.Reference, metalink.File) verification.VerificationResult
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct {
		arg1 file.Reference
		arg2 metalink.File
	}
	verifyReturns struct {
		result1 verification.VerificationResult
	}
	verifyReturnsOnCall map[int]struct {
		result1 verification.VerificationResult
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVerifier) Verify(arg1 file.Reference, arg2 metalink.File) verification.VerificationResult {
	fake.verifyMutex.Lock()
	ret, specificReturn := fake.verifyReturnsOnCall[len(fake.verifyArgsForCall)]
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct {
		arg1 file.Reference
		arg2 metalink.File
	}{arg1, arg2})
	stub := fake.VerifyStub
	fakeReturns := fake.verifyReturns
	fake.recordInvocation("Verify", []interface{}{arg1, arg2})
	fake.verifyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVerifier) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *FakeVerifier) VerifyCalls(stub func(file.Reference, metalink.File) verification.VerificationResult) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = stub
}

func (fake *FakeVerifier) VerifyArgsForCall(i int) (file.Reference, metalink.File) {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	argsForCall := fake.verifyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVerifier) VerifyReturns(result1 verification.VerificationResult) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 verification.VerificationResult
	}{result1}
}

func (fake *FakeVerifier) VerifyReturnsOnCall(i int, result1 verification.VerificationResult) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = nil
	if fake.verifyReturnsOnCall == nil {
		fake.verifyReturnsOnCall = make(map[int]struct {
			result1 verification.VerificationResult
		})
	}
	fake.verifyReturnsOnCall[i] = struct {
		result1 verification.VerificationResult
	}{result1}
}

func (fake *FakeVerifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVerifier) recordInvocation(key string, args []interface{}) {
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

var _ verification.Verifier = new(FakeVerifier)
