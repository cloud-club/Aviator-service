// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockServerInterface is an autogenerated mock type for the ServerInterface type
type MockServerInterface struct {
	mock.Mock
}

type MockServerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServerInterface) EXPECT() *MockServerInterface_Expecter {
	return &MockServerInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: url, payload
func (_m *MockServerInterface) Create(url string, payload interface{}) error {
	ret := _m.Called(url, payload)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(url, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServerInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockServerInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - url string
//   - payload interface{}
func (_e *MockServerInterface_Expecter) Create(url interface{}, payload interface{}) *MockServerInterface_Create_Call {
	return &MockServerInterface_Create_Call{Call: _e.mock.On("Create", url, payload)}
}

func (_c *MockServerInterface_Create_Call) Run(run func(url string, payload interface{})) *MockServerInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *MockServerInterface_Create_Call) Return(_a0 error) *MockServerInterface_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerInterface_Create_Call) RunAndReturn(run func(string, interface{}) error) *MockServerInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: url
func (_m *MockServerInterface) Delete(url string) error {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServerInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockServerInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - url string
func (_e *MockServerInterface_Expecter) Delete(url interface{}) *MockServerInterface_Delete_Call {
	return &MockServerInterface_Delete_Call{Call: _e.mock.On("Delete", url)}
}

func (_c *MockServerInterface_Delete_Call) Run(run func(url string)) *MockServerInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockServerInterface_Delete_Call) Return(_a0 error) *MockServerInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerInterface_Delete_Call) RunAndReturn(run func(string) error) *MockServerInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: url
func (_m *MockServerInterface) Get(url string) error {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServerInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockServerInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - url string
func (_e *MockServerInterface_Expecter) Get(url interface{}) *MockServerInterface_Get_Call {
	return &MockServerInterface_Get_Call{Call: _e.mock.On("Get", url)}
}

func (_c *MockServerInterface_Get_Call) Run(run func(url string)) *MockServerInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockServerInterface_Get_Call) Return(_a0 error) *MockServerInterface_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerInterface_Get_Call) RunAndReturn(run func(string) error) *MockServerInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: url
func (_m *MockServerInterface) List(url string) error {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServerInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockServerInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - url string
func (_e *MockServerInterface_Expecter) List(url interface{}) *MockServerInterface_List_Call {
	return &MockServerInterface_List_Call{Call: _e.mock.On("List", url)}
}

func (_c *MockServerInterface_List_Call) Run(run func(url string)) *MockServerInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockServerInterface_List_Call) Return(_a0 error) *MockServerInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerInterface_List_Call) RunAndReturn(run func(string) error) *MockServerInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: url
func (_m *MockServerInterface) Update(url string) error {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServerInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockServerInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - url string
func (_e *MockServerInterface_Expecter) Update(url interface{}) *MockServerInterface_Update_Call {
	return &MockServerInterface_Update_Call{Call: _e.mock.On("Update", url)}
}

func (_c *MockServerInterface_Update_Call) Run(run func(url string)) *MockServerInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockServerInterface_Update_Call) Return(_a0 error) *MockServerInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerInterface_Update_Call) RunAndReturn(run func(string) error) *MockServerInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockServerInterface creates a new instance of MockServerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServerInterface {
	mock := &MockServerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}