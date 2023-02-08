// Code generated by "./generator ./system/com.deepin.daemon.authenticate.face"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package face

import "fmt"
import "github.com/godbus/dbus"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockFace struct {
	MockInterfaceFace // interface com.deepin.daemon.Authenticate.Face.Device
	proxy.MockObject
}

type MockInterfaceFace struct {
	mock.Mock
}

func (v *MockInterfaceFace) SetInterfaceName_(string) {
}

// method Claim

func (v *MockInterfaceFace) GoClaim(flags dbus.Flags, ch chan *dbus.Call, claimed bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, claimed)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) Claim(flags dbus.Flags, claimed bool) error {
	mockArgs := v.Called(flags, claimed)

	return mockArgs.Error(0)
}

// method DeleteFace

func (v *MockInterfaceFace) GoDeleteFace(flags dbus.Flags, ch chan *dbus.Call, uuid string, faceName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, faceName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) DeleteFace(flags dbus.Flags, uuid string, faceName string) error {
	mockArgs := v.Called(flags, uuid, faceName)

	return mockArgs.Error(0)
}

// method DeleteFaces

func (v *MockInterfaceFace) GoDeleteFaces(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) DeleteFaces(flags dbus.Flags, uuid string) error {
	mockArgs := v.Called(flags, uuid)

	return mockArgs.Error(0)
}

// method GetShareMemInfo

func (v *MockInterfaceFace) GoGetShareMemInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) GetShareMemInfo(flags dbus.Flags) (string, string, int32, error) {
	mockArgs := v.Called(flags)

	ret2, ok := mockArgs.Get(2).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 2, mockArgs.Get(2)))
	}

	return mockArgs.String(0), mockArgs.String(1), ret2, mockArgs.Error(3)
}

// method ListFaces

func (v *MockInterfaceFace) GoListFaces(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) ListFaces(flags dbus.Flags, uuid string) ([]string, error) {
	mockArgs := v.Called(flags, uuid)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RenameFace

func (v *MockInterfaceFace) GoRenameFace(flags dbus.Flags, ch chan *dbus.Call, uuid string, oldFaceName string, newFaceName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, oldFaceName, newFaceName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) RenameFace(flags dbus.Flags, uuid string, oldFaceName string, newFaceName string) error {
	mockArgs := v.Called(flags, uuid, oldFaceName, newFaceName)

	return mockArgs.Error(0)
}

// method SetDefaultDevice

func (v *MockInterfaceFace) GoSetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call, device string) *dbus.Call {
	mockArgs := v.Called(flags, ch, device)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) SetDefaultDevice(flags dbus.Flags, device string) error {
	mockArgs := v.Called(flags, device)

	return mockArgs.Error(0)
}

// method StartEnroll

func (v *MockInterfaceFace) GoStartEnroll(flags dbus.Flags, ch chan *dbus.Call, uuid string, faceName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, faceName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) StartEnroll(flags dbus.Flags, uuid string, faceName string) error {
	mockArgs := v.Called(flags, uuid, faceName)

	return mockArgs.Error(0)
}

// method StartVerify

func (v *MockInterfaceFace) GoStartVerify(flags dbus.Flags, ch chan *dbus.Call, uuid string, timeout int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, timeout)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) StartVerify(flags dbus.Flags, uuid string, timeout int32) error {
	mockArgs := v.Called(flags, uuid, timeout)

	return mockArgs.Error(0)
}

// method StopEnroll

func (v *MockInterfaceFace) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) StopEnroll(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method StopVerify

func (v *MockInterfaceFace) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) StopVerify(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// signal EnrollStatus

func (v *MockInterfaceFace) ConnectEnrollStatus(cb func(uuid string, code int32, description string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyStatus

func (v *MockInterfaceFace) ConnectVerifyStatus(cb func(uuid string, code int32, description string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property SupportedDevices as

func (v *MockInterfaceFace) SupportedDevices() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DefaultDevice s

func (v *MockInterfaceFace) DefaultDevice() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Name s

func (v *MockInterfaceFace) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Claimed b

func (v *MockInterfaceFace) Claimed() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Capability i

func (v *MockInterfaceFace) Capability() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Status i

func (v *MockInterfaceFace) Status() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
