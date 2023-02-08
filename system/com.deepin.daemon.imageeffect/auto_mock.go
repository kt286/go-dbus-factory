// Code generated by "./generator ./system/com.deepin.daemon.imageeffect"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package imageeffect

import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockImageEffect struct {
	MockInterfaceImageEffect // interface com.deepin.daemon.ImageEffect
	proxy.MockObject
}

type MockInterfaceImageEffect struct {
	mock.Mock
}

// method Delete

func (v *MockInterfaceImageEffect) GoDelete(flags dbus.Flags, ch chan *dbus.Call, effect string, filename string) *dbus.Call {
	mockArgs := v.Called(flags, ch, effect, filename)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceImageEffect) Delete(flags dbus.Flags, effect string, filename string) error {
	mockArgs := v.Called(flags, effect, filename)

	return mockArgs.Error(0)
}

// method Get

func (v *MockInterfaceImageEffect) GoGet(flags dbus.Flags, ch chan *dbus.Call, effect string, filename string) *dbus.Call {
	mockArgs := v.Called(flags, ch, effect, filename)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceImageEffect) Get(flags dbus.Flags, effect string, filename string) (string, error) {
	mockArgs := v.Called(flags, effect, filename)

	return mockArgs.String(0), mockArgs.Error(1)
}
