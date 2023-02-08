// Code generated by "./generator ./system/com.deepin.system.network"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package network

import "fmt"
import "github.com/godbus/dbus"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockNetwork struct {
	MockInterfaceNetwork // interface com.deepin.system.Network
	proxy.MockObject
}

type MockInterfaceNetwork struct {
	mock.Mock
}

// method EnableDevice

func (v *MockInterfaceNetwork) GoEnableDevice(flags dbus.Flags, ch chan *dbus.Call, pathOrIface string, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, pathOrIface, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) EnableDevice(flags dbus.Flags, pathOrIface string, enabled bool) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, pathOrIface, enabled)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method IsDeviceEnabled

func (v *MockInterfaceNetwork) GoIsDeviceEnabled(flags dbus.Flags, ch chan *dbus.Call, pathOrIface string) *dbus.Call {
	mockArgs := v.Called(flags, ch, pathOrIface)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) IsDeviceEnabled(flags dbus.Flags, pathOrIface string) (bool, error) {
	mockArgs := v.Called(flags, pathOrIface)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method Ping

func (v *MockInterfaceNetwork) GoPing(flags dbus.Flags, ch chan *dbus.Call, host string) *dbus.Call {
	mockArgs := v.Called(flags, ch, host)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) Ping(flags dbus.Flags, host string) error {
	mockArgs := v.Called(flags, host)

	return mockArgs.Error(0)
}

// method ToggleWirelessEnabled

func (v *MockInterfaceNetwork) GoToggleWirelessEnabled(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) ToggleWirelessEnabled(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// signal DeviceEnabled

func (v *MockInterfaceNetwork) ConnectDeviceEnabled(cb func(devPath dbus.ObjectPath, enabled bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property VpnEnabled b

func (v *MockInterfaceNetwork) VpnEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
