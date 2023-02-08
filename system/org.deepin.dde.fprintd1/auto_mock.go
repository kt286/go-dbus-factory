// Code generated by "./generator ./system/org.deepin.dde.fprintd1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package fprintd1

import "fmt"
import "github.com/godbus/dbus"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockFprintd struct {
	MockInterfaceFprintd // interface org.deepin.dde.Fprintd1
	proxy.MockObject
}

type MockInterfaceFprintd struct {
	mock.Mock
}

// method GetDefaultDevice

func (v *MockInterfaceFprintd) GoGetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFprintd) GetDefaultDevice(flags dbus.Flags) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetDevices

func (v *MockInterfaceFprintd) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFprintd) GetDevices(flags dbus.Flags) ([]dbus.ObjectPath, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method TriggerUDevEvent

func (v *MockInterfaceFprintd) GoTriggerUDevEvent(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFprintd) TriggerUDevEvent(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// property Devices ao

func (v *MockInterfaceFprintd) Devices() proxy.PropObjectPathArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPathArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockDevice struct {
	MockInterfaceDevice // interface org.deepin.dde.Fprintd1.Device
	proxy.MockObject
}

type MockInterfaceDevice struct {
	mock.Mock
}

// method Claim

func (v *MockInterfaceDevice) GoClaim(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) Claim(flags dbus.Flags, username string) error {
	mockArgs := v.Called(flags, username)

	return mockArgs.Error(0)
}

// method ClaimForce

func (v *MockInterfaceDevice) GoClaimForce(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) ClaimForce(flags dbus.Flags, username string) error {
	mockArgs := v.Called(flags, username)

	return mockArgs.Error(0)
}

// method DeleteEnrolledFinger

func (v *MockInterfaceDevice) GoDeleteEnrolledFinger(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) DeleteEnrolledFinger(flags dbus.Flags, username string, finger string) error {
	mockArgs := v.Called(flags, username, finger)

	return mockArgs.Error(0)
}

// method DeleteEnrolledFingers

func (v *MockInterfaceDevice) GoDeleteEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) DeleteEnrolledFingers(flags dbus.Flags, username string) error {
	mockArgs := v.Called(flags, username)

	return mockArgs.Error(0)
}

// method EnrollStart

func (v *MockInterfaceDevice) GoEnrollStart(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) EnrollStart(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// method EnrollStop

func (v *MockInterfaceDevice) GoEnrollStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) EnrollStop(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method GetCapabilities

func (v *MockInterfaceDevice) GoGetCapabilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) GetCapabilities(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ListEnrolledFingers

func (v *MockInterfaceDevice) GoListEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) ListEnrolledFingers(flags dbus.Flags, username string) ([]string, error) {
	mockArgs := v.Called(flags, username)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method Release

func (v *MockInterfaceDevice) GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) Release(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method VerifyStart

func (v *MockInterfaceDevice) GoVerifyStart(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) VerifyStart(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// method VerifyStop

func (v *MockInterfaceDevice) GoVerifyStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) VerifyStop(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// signal EnrollStatus

func (v *MockInterfaceDevice) ConnectEnrollStatus(cb func(status string, done bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyStatus

func (v *MockInterfaceDevice) ConnectVerifyStatus(cb func(status string, done bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyFingerSelected

func (v *MockInterfaceDevice) ConnectVerifyFingerSelected(cb func(finger string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property ScanType s

func (v *MockInterfaceDevice) ScanType() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
