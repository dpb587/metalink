// Code generated by counterfeiter. DO NOT EDIT.
package metaurlfakes

import (
	"sync"

	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/file"
	"github.com/dpb587/metalink/file/metaurl"
)

type FakeLoader struct {
	LoadMetaURLStub        func(metalink.MetaURL) (file.Reference, error)
	loadMetaURLMutex       sync.RWMutex
	loadMetaURLArgsForCall []struct {
		arg1 metalink.MetaURL
	}
	loadMetaURLReturns struct {
		result1 file.Reference
		result2 error
	}
	loadMetaURLReturnsOnCall map[int]struct {
		result1 file.Reference
		result2 error
	}
	SupportsMetaURLStub        func(metalink.MetaURL) bool
	supportsMetaURLMutex       sync.RWMutex
	supportsMetaURLArgsForCall []struct {
		arg1 metalink.MetaURL
	}
	supportsMetaURLReturns struct {
		result1 bool
	}
	supportsMetaURLReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLoader) LoadMetaURL(arg1 metalink.MetaURL) (file.Reference, error) {
	fake.loadMetaURLMutex.Lock()
	ret, specificReturn := fake.loadMetaURLReturnsOnCall[len(fake.loadMetaURLArgsForCall)]
	fake.loadMetaURLArgsForCall = append(fake.loadMetaURLArgsForCall, struct {
		arg1 metalink.MetaURL
	}{arg1})
	stub := fake.LoadMetaURLStub
	fakeReturns := fake.loadMetaURLReturns
	fake.recordInvocation("LoadMetaURL", []interface{}{arg1})
	fake.loadMetaURLMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLoader) LoadMetaURLCallCount() int {
	fake.loadMetaURLMutex.RLock()
	defer fake.loadMetaURLMutex.RUnlock()
	return len(fake.loadMetaURLArgsForCall)
}

func (fake *FakeLoader) LoadMetaURLCalls(stub func(metalink.MetaURL) (file.Reference, error)) {
	fake.loadMetaURLMutex.Lock()
	defer fake.loadMetaURLMutex.Unlock()
	fake.LoadMetaURLStub = stub
}

func (fake *FakeLoader) LoadMetaURLArgsForCall(i int) metalink.MetaURL {
	fake.loadMetaURLMutex.RLock()
	defer fake.loadMetaURLMutex.RUnlock()
	argsForCall := fake.loadMetaURLArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLoader) LoadMetaURLReturns(result1 file.Reference, result2 error) {
	fake.loadMetaURLMutex.Lock()
	defer fake.loadMetaURLMutex.Unlock()
	fake.LoadMetaURLStub = nil
	fake.loadMetaURLReturns = struct {
		result1 file.Reference
		result2 error
	}{result1, result2}
}

func (fake *FakeLoader) LoadMetaURLReturnsOnCall(i int, result1 file.Reference, result2 error) {
	fake.loadMetaURLMutex.Lock()
	defer fake.loadMetaURLMutex.Unlock()
	fake.LoadMetaURLStub = nil
	if fake.loadMetaURLReturnsOnCall == nil {
		fake.loadMetaURLReturnsOnCall = make(map[int]struct {
			result1 file.Reference
			result2 error
		})
	}
	fake.loadMetaURLReturnsOnCall[i] = struct {
		result1 file.Reference
		result2 error
	}{result1, result2}
}

func (fake *FakeLoader) SupportsMetaURL(arg1 metalink.MetaURL) bool {
	fake.supportsMetaURLMutex.Lock()
	ret, specificReturn := fake.supportsMetaURLReturnsOnCall[len(fake.supportsMetaURLArgsForCall)]
	fake.supportsMetaURLArgsForCall = append(fake.supportsMetaURLArgsForCall, struct {
		arg1 metalink.MetaURL
	}{arg1})
	stub := fake.SupportsMetaURLStub
	fakeReturns := fake.supportsMetaURLReturns
	fake.recordInvocation("SupportsMetaURL", []interface{}{arg1})
	fake.supportsMetaURLMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLoader) SupportsMetaURLCallCount() int {
	fake.supportsMetaURLMutex.RLock()
	defer fake.supportsMetaURLMutex.RUnlock()
	return len(fake.supportsMetaURLArgsForCall)
}

func (fake *FakeLoader) SupportsMetaURLCalls(stub func(metalink.MetaURL) bool) {
	fake.supportsMetaURLMutex.Lock()
	defer fake.supportsMetaURLMutex.Unlock()
	fake.SupportsMetaURLStub = stub
}

func (fake *FakeLoader) SupportsMetaURLArgsForCall(i int) metalink.MetaURL {
	fake.supportsMetaURLMutex.RLock()
	defer fake.supportsMetaURLMutex.RUnlock()
	argsForCall := fake.supportsMetaURLArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLoader) SupportsMetaURLReturns(result1 bool) {
	fake.supportsMetaURLMutex.Lock()
	defer fake.supportsMetaURLMutex.Unlock()
	fake.SupportsMetaURLStub = nil
	fake.supportsMetaURLReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeLoader) SupportsMetaURLReturnsOnCall(i int, result1 bool) {
	fake.supportsMetaURLMutex.Lock()
	defer fake.supportsMetaURLMutex.Unlock()
	fake.SupportsMetaURLStub = nil
	if fake.supportsMetaURLReturnsOnCall == nil {
		fake.supportsMetaURLReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.supportsMetaURLReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeLoader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loadMetaURLMutex.RLock()
	defer fake.loadMetaURLMutex.RUnlock()
	fake.supportsMetaURLMutex.RLock()
	defer fake.supportsMetaURLMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLoader) recordInvocation(key string, args []interface{}) {
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

var _ metaurl.Loader = new(FakeLoader)
