// Code generated by "./generator ./system/org.deepin.dde.procs1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package procs1

import "fmt"
import "github.com/godbus/dbus"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockProcs struct {
	MockInterfaceProcs // interface org.deepin.dde.Procs1
	proxy.MockObject
}

type MockInterfaceProcs struct {
	mock.Mock
}

// signal ExecProc

func (v *MockInterfaceProcs) ConnectExecProc(cb func(ExecPath string, CGroupPath string, Pid string, PPid string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ExitProc

func (v *MockInterfaceProcs) ConnectExitProc(cb func(ExecPath string, CGroupPath string, Pid string, PPid string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Procs a{s(ssss)}

func (v *MockInterfaceProcs) Procs() PropProcsMap {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropProcsMap)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropProcsMap struct {
	mock.Mock
}

func (p MockPropProcsMap) Get(flags dbus.Flags) (value map[string]ProcMessage, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(map[string]ProcMessage)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropProcsMap) Set(flags dbus.Flags, value map[string]ProcMessage) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropProcsMap) ConnectChanged(cb func(hasValue bool, value map[string]ProcMessage)) error {
	args := p.Called(cb)

	return args.Error(0)
}
