// Code generated by "./generator ./com.deepin.system.bluetooth"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package bluetooth

import "fmt"
import "github.com/godbus/dbus"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockBluetooth struct {
	MockInterfaceBluetooth // interface com.deepin.system.Bluetooth
	proxy.MockObject
}

type MockInterfaceBluetooth struct {
	mock.Mock
}

// method ClearUnpairedDevice

func (v *MockInterfaceBluetooth) GoClearUnpairedDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) ClearUnpairedDevice(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ConnectDevice

func (v *MockInterfaceBluetooth) GoConnectDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath, adapterPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath, adapterPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) ConnectDevice(flags dbus.Flags, devPath dbus.ObjectPath, adapterPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, devPath, adapterPath)

	return mockArgs.Error(0)
}

// method DebugInfo

func (v *MockInterfaceBluetooth) GoDebugInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) DebugInfo(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method DisconnectAudioDevices

func (v *MockInterfaceBluetooth) GoDisconnectAudioDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) DisconnectAudioDevices(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method DisconnectDevice

func (v *MockInterfaceBluetooth) GoDisconnectDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) DisconnectDevice(flags dbus.Flags, devPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, devPath)

	return mockArgs.Error(0)
}

// method GetAdapters

func (v *MockInterfaceBluetooth) GoGetAdapters(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) GetAdapters(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetDevices

func (v *MockInterfaceBluetooth) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, adapterPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) GetDevices(flags dbus.Flags, adapterPath dbus.ObjectPath) (string, error) {
	mockArgs := v.Called(flags, adapterPath)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method RegisterAgent

func (v *MockInterfaceBluetooth) GoRegisterAgent(flags dbus.Flags, ch chan *dbus.Call, agentPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, agentPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) RegisterAgent(flags dbus.Flags, agentPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, agentPath)

	return mockArgs.Error(0)
}

// method RemoveDevice

func (v *MockInterfaceBluetooth) GoRemoveDevice(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, adapterPath, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) RemoveDevice(flags dbus.Flags, adapterPath dbus.ObjectPath, devPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, adapterPath, devPath)

	return mockArgs.Error(0)
}

// method RequestDiscovery

func (v *MockInterfaceBluetooth) GoRequestDiscovery(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, adapterPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) RequestDiscovery(flags dbus.Flags, adapterPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, adapterPath)

	return mockArgs.Error(0)
}

// method SetAdapterAlias

func (v *MockInterfaceBluetooth) GoSetAdapterAlias(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, alias string) *dbus.Call {
	mockArgs := v.Called(flags, ch, adapterPath, alias)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) SetAdapterAlias(flags dbus.Flags, adapterPath dbus.ObjectPath, alias string) error {
	mockArgs := v.Called(flags, adapterPath, alias)

	return mockArgs.Error(0)
}

// method SetAdapterDiscoverable

func (v *MockInterfaceBluetooth) GoSetAdapterDiscoverable(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, discoverable bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, adapterPath, discoverable)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) SetAdapterDiscoverable(flags dbus.Flags, adapterPath dbus.ObjectPath, discoverable bool) error {
	mockArgs := v.Called(flags, adapterPath, discoverable)

	return mockArgs.Error(0)
}

// method SetAdapterDiscoverableTimeout

func (v *MockInterfaceBluetooth) GoSetAdapterDiscoverableTimeout(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, discoverableTimeout uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, adapterPath, discoverableTimeout)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) SetAdapterDiscoverableTimeout(flags dbus.Flags, adapterPath dbus.ObjectPath, discoverableTimeout uint32) error {
	mockArgs := v.Called(flags, adapterPath, discoverableTimeout)

	return mockArgs.Error(0)
}

// method SetAdapterDiscovering

func (v *MockInterfaceBluetooth) GoSetAdapterDiscovering(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, discovering bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, adapterPath, discovering)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) SetAdapterDiscovering(flags dbus.Flags, adapterPath dbus.ObjectPath, discovering bool) error {
	mockArgs := v.Called(flags, adapterPath, discovering)

	return mockArgs.Error(0)
}

// method SetAdapterPowered

func (v *MockInterfaceBluetooth) GoSetAdapterPowered(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, powered bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, adapterPath, powered)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) SetAdapterPowered(flags dbus.Flags, adapterPath dbus.ObjectPath, powered bool) error {
	mockArgs := v.Called(flags, adapterPath, powered)

	return mockArgs.Error(0)
}

// method SetDeviceAlias

func (v *MockInterfaceBluetooth) GoSetDeviceAlias(flags dbus.Flags, ch chan *dbus.Call, device dbus.ObjectPath, alias string) *dbus.Call {
	mockArgs := v.Called(flags, ch, device, alias)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) SetDeviceAlias(flags dbus.Flags, device dbus.ObjectPath, alias string) error {
	mockArgs := v.Called(flags, device, alias)

	return mockArgs.Error(0)
}

// method SetDeviceTrusted

func (v *MockInterfaceBluetooth) GoSetDeviceTrusted(flags dbus.Flags, ch chan *dbus.Call, device dbus.ObjectPath, trusted bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, device, trusted)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) SetDeviceTrusted(flags dbus.Flags, device dbus.ObjectPath, trusted bool) error {
	mockArgs := v.Called(flags, device, trusted)

	return mockArgs.Error(0)
}

// method UnregisterAgent

func (v *MockInterfaceBluetooth) GoUnregisterAgent(flags dbus.Flags, ch chan *dbus.Call, agentPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, agentPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) UnregisterAgent(flags dbus.Flags, agentPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, agentPath)

	return mockArgs.Error(0)
}

// signal AdapterAdded

func (v *MockInterfaceBluetooth) ConnectAdapterAdded(cb func(adapterJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal AdapterRemoved

func (v *MockInterfaceBluetooth) ConnectAdapterRemoved(cb func(adapterJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal AdapterPropertiesChanged

func (v *MockInterfaceBluetooth) ConnectAdapterPropertiesChanged(cb func(adapterJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal DeviceAdded

func (v *MockInterfaceBluetooth) ConnectDeviceAdded(cb func(devJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal DeviceRemoved

func (v *MockInterfaceBluetooth) ConnectDeviceRemoved(cb func(devJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal DevicePropertiesChanged

func (v *MockInterfaceBluetooth) ConnectDevicePropertiesChanged(cb func(devJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property CanSendFile b

func (v *MockInterfaceBluetooth) CanSendFile() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property State u

func (v *MockInterfaceBluetooth) State() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
