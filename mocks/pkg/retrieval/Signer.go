// Code generated by mockery v2.0.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Signer is an autogenerated mock type for the Signer type
type Signer struct {
	mock.Mock
}

// Sign provides a mock function with given fields: _a0
func (_m *Signer) Sign(_a0 []byte) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
