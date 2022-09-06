// Code generated by "./generator ./com.deepin.api.device"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package device

import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Device interface {
	device // interface com.deepin.api.Device
	proxy.Object
}

type objectDevice struct {
	interfaceDevice // interface com.deepin.api.Device
	proxy.ImplObject
}

func NewDevice(conn *dbus.Conn) Device {
	obj := new(objectDevice)
	obj.ImplObject.Init_(conn, "com.deepin.api.Device", "/com/deepin/api/Device")
	return obj
}

type device interface {
	GoHasBluetoothDeviceBlocked(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	HasBluetoothDeviceBlocked(flags dbus.Flags) (bool, error)
	GoUnblockBluetoothDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	UnblockBluetoothDevices(flags dbus.Flags) error
}

type interfaceDevice struct{}

func (v *interfaceDevice) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceDevice) GetInterfaceName_() string {
	return "com.deepin.api.Device"
}

// method HasBluetoothDeviceBlocked

func (v *interfaceDevice) GoHasBluetoothDeviceBlocked(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".HasBluetoothDeviceBlocked", flags, ch)
}

func (*interfaceDevice) StoreHasBluetoothDeviceBlocked(call *dbus.Call) (has bool, err error) {
	err = call.Store(&has)
	return
}

func (v *interfaceDevice) HasBluetoothDeviceBlocked(flags dbus.Flags) (bool, error) {
	return v.StoreHasBluetoothDeviceBlocked(
		<-v.GoHasBluetoothDeviceBlocked(flags, make(chan *dbus.Call, 1)).Done)
}

// method UnblockBluetoothDevices

func (v *interfaceDevice) GoUnblockBluetoothDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnblockBluetoothDevices", flags, ch)
}

func (v *interfaceDevice) UnblockBluetoothDevices(flags dbus.Flags) error {
	return (<-v.GoUnblockBluetoothDevices(flags, make(chan *dbus.Call, 1)).Done).Err
}
