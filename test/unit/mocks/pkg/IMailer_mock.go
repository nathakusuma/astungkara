// Code generated by mockery v2.51.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockIMailer is an autogenerated mock type for the IMailer type
type MockIMailer struct {
	mock.Mock
}

type MockIMailer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIMailer) EXPECT() *MockIMailer_Expecter {
	return &MockIMailer_Expecter{mock: &_m.Mock}
}

// Send provides a mock function with given fields: recipientEmail, subject, templateName, data
func (_m *MockIMailer) Send(recipientEmail string, subject string, templateName string, data map[string]any) error {
	ret := _m.Called(recipientEmail, subject, templateName, data)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, map[string]any) error); ok {
		r0 = rf(recipientEmail, subject, templateName, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIMailer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockIMailer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - recipientEmail string
//   - subject string
//   - templateName string
//   - data map[string]any
func (_e *MockIMailer_Expecter) Send(recipientEmail interface{}, subject interface{}, templateName interface{}, data interface{}) *MockIMailer_Send_Call {
	return &MockIMailer_Send_Call{Call: _e.mock.On("Send", recipientEmail, subject, templateName, data)}
}

func (_c *MockIMailer_Send_Call) Run(run func(recipientEmail string, subject string, templateName string, data map[string]any)) *MockIMailer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(map[string]any))
	})
	return _c
}

func (_c *MockIMailer_Send_Call) Return(_a0 error) *MockIMailer_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIMailer_Send_Call) RunAndReturn(run func(string, string, string, map[string]any) error) *MockIMailer_Send_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIMailer creates a new instance of MockIMailer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIMailer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIMailer {
	mock := &MockIMailer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
