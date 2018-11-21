// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"net/url"
	"sync"

	"github.com/alphagov/paas-prometheus-exporter/exporter"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type FakeCFClient struct {
	ListAppsByQueryStub        func(url.Values) ([]cfclient.App, error)
	listAppsByQueryMutex       sync.RWMutex
	listAppsByQueryArgsForCall []struct {
		arg1 url.Values
	}
	listAppsByQueryReturns struct {
		result1 []cfclient.App
		result2 error
	}
	listAppsByQueryReturnsOnCall map[int]struct {
		result1 []cfclient.App
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFClient) ListAppsByQuery(arg1 url.Values) ([]cfclient.App, error) {
	fake.listAppsByQueryMutex.Lock()
	ret, specificReturn := fake.listAppsByQueryReturnsOnCall[len(fake.listAppsByQueryArgsForCall)]
	fake.listAppsByQueryArgsForCall = append(fake.listAppsByQueryArgsForCall, struct {
		arg1 url.Values
	}{arg1})
	fake.recordInvocation("ListAppsByQuery", []interface{}{arg1})
	fake.listAppsByQueryMutex.Unlock()
	if fake.ListAppsByQueryStub != nil {
		return fake.ListAppsByQueryStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listAppsByQueryReturns.result1, fake.listAppsByQueryReturns.result2
}

func (fake *FakeCFClient) ListAppsByQueryCallCount() int {
	fake.listAppsByQueryMutex.RLock()
	defer fake.listAppsByQueryMutex.RUnlock()
	return len(fake.listAppsByQueryArgsForCall)
}

func (fake *FakeCFClient) ListAppsByQueryArgsForCall(i int) url.Values {
	fake.listAppsByQueryMutex.RLock()
	defer fake.listAppsByQueryMutex.RUnlock()
	return fake.listAppsByQueryArgsForCall[i].arg1
}

func (fake *FakeCFClient) ListAppsByQueryReturns(result1 []cfclient.App, result2 error) {
	fake.ListAppsByQueryStub = nil
	fake.listAppsByQueryReturns = struct {
		result1 []cfclient.App
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListAppsByQueryReturnsOnCall(i int, result1 []cfclient.App, result2 error) {
	fake.ListAppsByQueryStub = nil
	if fake.listAppsByQueryReturnsOnCall == nil {
		fake.listAppsByQueryReturnsOnCall = make(map[int]struct {
			result1 []cfclient.App
			result2 error
		})
	}
	fake.listAppsByQueryReturnsOnCall[i] = struct {
		result1 []cfclient.App
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listAppsByQueryMutex.RLock()
	defer fake.listAppsByQueryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFClient) recordInvocation(key string, args []interface{}) {
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

var _ exporter.CFClient = new(FakeCFClient)
