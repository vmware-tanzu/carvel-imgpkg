// Code generated by counterfeiter. DO NOT EDIT.
package imagefakes

import (
	"sync"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/k14s/imgpkg/pkg/imgpkg/image"
)

type FakeImagesMetadata struct {
	GenericStub        func(name.Reference) (v1.Descriptor, error)
	genericMutex       sync.RWMutex
	genericArgsForCall []struct {
		arg1 name.Reference
	}
	genericReturns struct {
		result1 v1.Descriptor
		result2 error
	}
	genericReturnsOnCall map[int]struct {
		result1 v1.Descriptor
		result2 error
	}
	ImageStub        func(name.Reference) (v1.Image, error)
	imageMutex       sync.RWMutex
	imageArgsForCall []struct {
		arg1 name.Reference
	}
	imageReturns struct {
		result1 v1.Image
		result2 error
	}
	imageReturnsOnCall map[int]struct {
		result1 v1.Image
		result2 error
	}
	IndexStub        func(name.Reference) (v1.ImageIndex, error)
	indexMutex       sync.RWMutex
	indexArgsForCall []struct {
		arg1 name.Reference
	}
	indexReturns struct {
		result1 v1.ImageIndex
		result2 error
	}
	indexReturnsOnCall map[int]struct {
		result1 v1.ImageIndex
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImagesMetadata) Generic(arg1 name.Reference) (v1.Descriptor, error) {
	fake.genericMutex.Lock()
	ret, specificReturn := fake.genericReturnsOnCall[len(fake.genericArgsForCall)]
	fake.genericArgsForCall = append(fake.genericArgsForCall, struct {
		arg1 name.Reference
	}{arg1})
	fake.recordInvocation("Generic", []interface{}{arg1})
	fake.genericMutex.Unlock()
	if fake.GenericStub != nil {
		return fake.GenericStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.genericReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImagesMetadata) GenericCallCount() int {
	fake.genericMutex.RLock()
	defer fake.genericMutex.RUnlock()
	return len(fake.genericArgsForCall)
}

func (fake *FakeImagesMetadata) GenericCalls(stub func(name.Reference) (v1.Descriptor, error)) {
	fake.genericMutex.Lock()
	defer fake.genericMutex.Unlock()
	fake.GenericStub = stub
}

func (fake *FakeImagesMetadata) GenericArgsForCall(i int) name.Reference {
	fake.genericMutex.RLock()
	defer fake.genericMutex.RUnlock()
	argsForCall := fake.genericArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImagesMetadata) GenericReturns(result1 v1.Descriptor, result2 error) {
	fake.genericMutex.Lock()
	defer fake.genericMutex.Unlock()
	fake.GenericStub = nil
	fake.genericReturns = struct {
		result1 v1.Descriptor
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesMetadata) GenericReturnsOnCall(i int, result1 v1.Descriptor, result2 error) {
	fake.genericMutex.Lock()
	defer fake.genericMutex.Unlock()
	fake.GenericStub = nil
	if fake.genericReturnsOnCall == nil {
		fake.genericReturnsOnCall = make(map[int]struct {
			result1 v1.Descriptor
			result2 error
		})
	}
	fake.genericReturnsOnCall[i] = struct {
		result1 v1.Descriptor
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesMetadata) Image(arg1 name.Reference) (v1.Image, error) {
	fake.imageMutex.Lock()
	ret, specificReturn := fake.imageReturnsOnCall[len(fake.imageArgsForCall)]
	fake.imageArgsForCall = append(fake.imageArgsForCall, struct {
		arg1 name.Reference
	}{arg1})
	fake.recordInvocation("Image", []interface{}{arg1})
	fake.imageMutex.Unlock()
	if fake.ImageStub != nil {
		return fake.ImageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.imageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImagesMetadata) ImageCallCount() int {
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	return len(fake.imageArgsForCall)
}

func (fake *FakeImagesMetadata) ImageCalls(stub func(name.Reference) (v1.Image, error)) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = stub
}

func (fake *FakeImagesMetadata) ImageArgsForCall(i int) name.Reference {
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	argsForCall := fake.imageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImagesMetadata) ImageReturns(result1 v1.Image, result2 error) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = nil
	fake.imageReturns = struct {
		result1 v1.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesMetadata) ImageReturnsOnCall(i int, result1 v1.Image, result2 error) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = nil
	if fake.imageReturnsOnCall == nil {
		fake.imageReturnsOnCall = make(map[int]struct {
			result1 v1.Image
			result2 error
		})
	}
	fake.imageReturnsOnCall[i] = struct {
		result1 v1.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesMetadata) Index(arg1 name.Reference) (v1.ImageIndex, error) {
	fake.indexMutex.Lock()
	ret, specificReturn := fake.indexReturnsOnCall[len(fake.indexArgsForCall)]
	fake.indexArgsForCall = append(fake.indexArgsForCall, struct {
		arg1 name.Reference
	}{arg1})
	fake.recordInvocation("Index", []interface{}{arg1})
	fake.indexMutex.Unlock()
	if fake.IndexStub != nil {
		return fake.IndexStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.indexReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImagesMetadata) IndexCallCount() int {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	return len(fake.indexArgsForCall)
}

func (fake *FakeImagesMetadata) IndexCalls(stub func(name.Reference) (v1.ImageIndex, error)) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = stub
}

func (fake *FakeImagesMetadata) IndexArgsForCall(i int) name.Reference {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	argsForCall := fake.indexArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImagesMetadata) IndexReturns(result1 v1.ImageIndex, result2 error) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = nil
	fake.indexReturns = struct {
		result1 v1.ImageIndex
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesMetadata) IndexReturnsOnCall(i int, result1 v1.ImageIndex, result2 error) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = nil
	if fake.indexReturnsOnCall == nil {
		fake.indexReturnsOnCall = make(map[int]struct {
			result1 v1.ImageIndex
			result2 error
		})
	}
	fake.indexReturnsOnCall[i] = struct {
		result1 v1.ImageIndex
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesMetadata) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.genericMutex.RLock()
	defer fake.genericMutex.RUnlock()
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImagesMetadata) recordInvocation(key string, args []interface{}) {
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

var _ image.ImagesMetadata = new(FakeImagesMetadata)