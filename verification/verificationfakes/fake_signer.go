// Code generated by counterfeiter. DO NOT EDIT.
package verificationfakes

import (
	"sync"

	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/verification"
)

type FakeSigner struct {
	SignStub        func(file.Reference) (verification.Verification, error)
	signMutex       sync.RWMutex
	signArgsForCall []struct {
		arg1 file.Reference
	}
	signReturns struct {
		result1 verification.Verification
		result2 error
	}
	signReturnsOnCall map[int]struct {
		result1 verification.Verification
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSigner) Sign(arg1 file.Reference) (verification.Verification, error) {
	fake.signMutex.Lock()
	ret, specificReturn := fake.signReturnsOnCall[len(fake.signArgsForCall)]
	fake.signArgsForCall = append(fake.signArgsForCall, struct {
		arg1 file.Reference
	}{arg1})
	stub := fake.SignStub
	fakeReturns := fake.signReturns
	fake.recordInvocation("Sign", []interface{}{arg1})
	fake.signMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSigner) SignCallCount() int {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return len(fake.signArgsForCall)
}

func (fake *FakeSigner) SignCalls(stub func(file.Reference) (verification.Verification, error)) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = stub
}

func (fake *FakeSigner) SignArgsForCall(i int) file.Reference {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	argsForCall := fake.signArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSigner) SignReturns(result1 verification.Verification, result2 error) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = nil
	fake.signReturns = struct {
		result1 verification.Verification
		result2 error
	}{result1, result2}
}

func (fake *FakeSigner) SignReturnsOnCall(i int, result1 verification.Verification, result2 error) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = nil
	if fake.signReturnsOnCall == nil {
		fake.signReturnsOnCall = make(map[int]struct {
			result1 verification.Verification
			result2 error
		})
	}
	fake.signReturnsOnCall[i] = struct {
		result1 verification.Verification
		result2 error
	}{result1, result2}
}

func (fake *FakeSigner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSigner) recordInvocation(key string, args []interface{}) {
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

var _ verification.Signer = new(FakeSigner)
