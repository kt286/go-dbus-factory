// Code generated by "./generator ./session/com.deepin.dde.osd"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package osd

import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockOSD struct {
	MockInterfaceOsd // interface com.deepin.dde.osd
	proxy.MockObject
}

type MockInterfaceOsd struct {
	mock.Mock
}

// method ShowOSD

func (v *MockInterfaceOsd) GoShowOSD(flags dbus.Flags, ch chan *dbus.Call, osd string) *dbus.Call {
	mockArgs := v.Called(flags, ch, osd)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceOsd) ShowOSD(flags dbus.Flags, osd string) error {
	mockArgs := v.Called(flags, osd)

	return mockArgs.Error(0)
}
