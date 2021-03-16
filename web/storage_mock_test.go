// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package web

import (
	"sync"
)

// Ensure, that HumanStorageMock does implement HumanStorage.
// If this is not the case, regenerate this file with moq.
var _ HumanStorage = &HumanStorageMock{}

// HumanStorageMock is a mock implementation of HumanStorage.
//
// 	func TestSomethingThatUsesHumanStorage(t *testing.T) {
//
// 		// make and configure a mocked HumanStorage
// 		mockedHumanStorage := &HumanStorageMock{
// 			AddFunc: func(man Man) error {
// 				panic("mock out the Add method")
// 			},
// 			DelFunc: func(id string) error {
// 				panic("mock out the Del method")
// 			},
// 			EditFunc: func(man Man) error {
// 				panic("mock out the Edit method")
// 			},
// 			GetFunc: func(id string) (Man, error) {
// 				panic("mock out the Get method")
// 			},
// 			GetAllFunc: func() ([]Man, error) {
// 				panic("mock out the GetAll method")
// 			},
// 		}
//
// 		// use mockedHumanStorage in code that requires HumanStorage
// 		// and then make assertions.
//
// 	}
type HumanStorageMock struct {
	// AddFunc mocks the Add method.
	AddFunc func(man Man) error

	// DelFunc mocks the Del method.
	DelFunc func(id string) error

	// EditFunc mocks the Edit method.
	EditFunc func(man Man) error

	// GetFunc mocks the Get method.
	GetFunc func(id string) (Man, error)

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func() ([]Man, error)

	// calls tracks calls to the methods.
	calls struct {
		// Add holds details about calls to the Add method.
		Add []struct {
			// Man is the man argument value.
			Man Man
		}
		// Del holds details about calls to the Del method.
		Del []struct {
			// ID is the id argument value.
			ID string
		}
		// Edit holds details about calls to the Edit method.
		Edit []struct {
			// Man is the man argument value.
			Man Man
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// ID is the id argument value.
			ID string
		}
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
		}
	}
	lockAdd    sync.RWMutex
	lockDel    sync.RWMutex
	lockEdit   sync.RWMutex
	lockGet    sync.RWMutex
	lockGetAll sync.RWMutex
}

// Add calls AddFunc.
func (mock *HumanStorageMock) Add(man Man) error {
	if mock.AddFunc == nil {
		panic("HumanStorageMock.AddFunc: method is nil but HumanStorage.Add was just called")
	}
	callInfo := struct {
		Man Man
	}{
		Man: man,
	}
	mock.lockAdd.Lock()
	mock.calls.Add = append(mock.calls.Add, callInfo)
	mock.lockAdd.Unlock()
	return mock.AddFunc(man)
}

// AddCalls gets all the calls that were made to Add.
// Check the length with:
//     len(mockedHumanStorage.AddCalls())
func (mock *HumanStorageMock) AddCalls() []struct {
	Man Man
} {
	var calls []struct {
		Man Man
	}
	mock.lockAdd.RLock()
	calls = mock.calls.Add
	mock.lockAdd.RUnlock()
	return calls
}

// Del calls DelFunc.
func (mock *HumanStorageMock) Del(id string) error {
	if mock.DelFunc == nil {
		panic("HumanStorageMock.DelFunc: method is nil but HumanStorage.Del was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockDel.Lock()
	mock.calls.Del = append(mock.calls.Del, callInfo)
	mock.lockDel.Unlock()
	return mock.DelFunc(id)
}

// DelCalls gets all the calls that were made to Del.
// Check the length with:
//     len(mockedHumanStorage.DelCalls())
func (mock *HumanStorageMock) DelCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockDel.RLock()
	calls = mock.calls.Del
	mock.lockDel.RUnlock()
	return calls
}

// Edit calls EditFunc.
func (mock *HumanStorageMock) Edit(man Man) error {
	if mock.EditFunc == nil {
		panic("HumanStorageMock.EditFunc: method is nil but HumanStorage.Edit was just called")
	}
	callInfo := struct {
		Man Man
	}{
		Man: man,
	}
	mock.lockEdit.Lock()
	mock.calls.Edit = append(mock.calls.Edit, callInfo)
	mock.lockEdit.Unlock()
	return mock.EditFunc(man)
}

// EditCalls gets all the calls that were made to Edit.
// Check the length with:
//     len(mockedHumanStorage.EditCalls())
func (mock *HumanStorageMock) EditCalls() []struct {
	Man Man
} {
	var calls []struct {
		Man Man
	}
	mock.lockEdit.RLock()
	calls = mock.calls.Edit
	mock.lockEdit.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *HumanStorageMock) Get(id string) (Man, error) {
	if mock.GetFunc == nil {
		panic("HumanStorageMock.GetFunc: method is nil but HumanStorage.Get was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedHumanStorage.GetCalls())
func (mock *HumanStorageMock) GetCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// GetAll calls GetAllFunc.
func (mock *HumanStorageMock) GetAll() ([]Man, error) {
	if mock.GetAllFunc == nil {
		panic("HumanStorageMock.GetAllFunc: method is nil but HumanStorage.GetAll was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc()
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//     len(mockedHumanStorage.GetAllCalls())
func (mock *HumanStorageMock) GetAllCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}
