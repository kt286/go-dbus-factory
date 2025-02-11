// Code generated by "./generator ./session/org.deepin.dde.application1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package application1

import "errors"
import "fmt"
import "github.com/godbus/dbus/v5"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Application interface {
	application // interface org.deepin.dde.Application1
	proxy.Object
}

type objectApplication struct {
	interfaceApplication // interface org.deepin.dde.Application1
	proxy.ImplObject
}

func NewApplication(conn *dbus.Conn) Application {
	obj := new(objectApplication)
	obj.ImplObject.Init_(conn, "org.deepin.dde.Application1", "/org/deepin/dde/Application1")
	return obj
}

type application interface {
	GoName(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call
	Name(flags dbus.Flags, locale string) (string, error)
	GoComment(flags dbus.Flags, ch chan *dbus.Call, comment string) *dbus.Call
	Comment(flags dbus.Flags, comment string) (string, error)
	Categories() proxy.PropStringArray
	Mimetypes() proxy.PropStringArray
	Id() proxy.PropString
	Icon() proxy.PropString
	Instances() proxy.PropObjectPathArray
}

type interfaceApplication struct{}

func (v *interfaceApplication) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceApplication) GetInterfaceName_() string {
	return "org.deepin.dde.Application1"
}

// method Name

func (v *interfaceApplication) GoName(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Name", flags, ch, locale)
}

func (*interfaceApplication) StoreName(call *dbus.Call) (value string, err error) {
	err = call.Store(&value)
	return
}

func (v *interfaceApplication) Name(flags dbus.Flags, locale string) (string, error) {
	return v.StoreName(
		<-v.GoName(flags, make(chan *dbus.Call, 1), locale).Done)
}

// method Comment

func (v *interfaceApplication) GoComment(flags dbus.Flags, ch chan *dbus.Call, comment string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Comment", flags, ch, comment)
}

func (*interfaceApplication) StoreComment(call *dbus.Call) (value string, err error) {
	err = call.Store(&value)
	return
}

func (v *interfaceApplication) Comment(flags dbus.Flags, comment string) (string, error) {
	return v.StoreComment(
		<-v.GoComment(flags, make(chan *dbus.Call, 1), comment).Done)
}

// property categories as

func (v *interfaceApplication) Categories() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "categories",
	}
}

// property mimetypes as

func (v *interfaceApplication) Mimetypes() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "mimetypes",
	}
}

// property id s

func (v *interfaceApplication) Id() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "id",
	}
}

// property icon s

func (v *interfaceApplication) Icon() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "icon",
	}
}

// property instances ao

func (v *interfaceApplication) Instances() proxy.PropObjectPathArray {
	return &proxy.ImplPropObjectPathArray{
		Impl: v,
		Name: "instances",
	}
}

type Manager interface {
	manager // interface org.deepin.dde.Application1.Manager
	proxy.Object
}

type objectManager struct {
	interfaceManager // interface org.deepin.dde.Application1.Manager
	proxy.ImplObject
}

func NewManager(conn *dbus.Conn) Manager {
	obj := new(objectManager)
	obj.ImplObject.Init_(conn, "org.deepin.dde.Application1", "/org/deepin/dde/Application1/Manager")
	return obj
}

type manager interface {
	GoAddAutostart(flags dbus.Flags, ch chan *dbus.Call, fileNamae string) *dbus.Call
	AddAutostart(flags dbus.Flags, fileNamae string) (bool, error)
	GoAutostartList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	AutostartList(flags dbus.Flags) ([]string, error)
	GoIsAutostart(flags dbus.Flags, ch chan *dbus.Call, fileName string) *dbus.Call
	IsAutostart(flags dbus.Flags, fileName string) (bool, error)
	GoRemoveAutostart(flags dbus.Flags, ch chan *dbus.Call, fileNamae string) *dbus.Call
	RemoveAutostart(flags dbus.Flags, fileNamae string) (bool, error)
	GoLaunch(flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call
	Launch(flags dbus.Flags, desktopFile string) error
	GoLaunchApp(flags dbus.Flags, ch chan *dbus.Call, desktopFile string, timestamp uint32, files []string) *dbus.Call
	LaunchApp(flags dbus.Flags, desktopFile string, timestamp uint32, files []string) error
	GoLaunchAppAction(flags dbus.Flags, ch chan *dbus.Call, desktopFile string, action string, timestamp uint32) *dbus.Call
	LaunchAppAction(flags dbus.Flags, desktopFile string, action string, timestamp uint32) error
	ConnectAutostartChanged(cb func(status string, filePath string)) (dbusutil.SignalHandlerId, error)
}

type interfaceManager struct{}

func (v *interfaceManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceManager) GetInterfaceName_() string {
	return "org.deepin.dde.Application1.Manager"
}

// method AddAutostart

func (v *interfaceManager) GoAddAutostart(flags dbus.Flags, ch chan *dbus.Call, fileNamae string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddAutostart", flags, ch, fileNamae)
}

func (*interfaceManager) StoreAddAutostart(call *dbus.Call) (outArg0 bool, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceManager) AddAutostart(flags dbus.Flags, fileNamae string) (bool, error) {
	return v.StoreAddAutostart(
		<-v.GoAddAutostart(flags, make(chan *dbus.Call, 1), fileNamae).Done)
}

// method AutostartList

func (v *interfaceManager) GoAutostartList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AutostartList", flags, ch)
}

func (*interfaceManager) StoreAutostartList(call *dbus.Call) (outArg0 []string, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceManager) AutostartList(flags dbus.Flags) ([]string, error) {
	return v.StoreAutostartList(
		<-v.GoAutostartList(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsAutostart

func (v *interfaceManager) GoIsAutostart(flags dbus.Flags, ch chan *dbus.Call, fileName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsAutostart", flags, ch, fileName)
}

func (*interfaceManager) StoreIsAutostart(call *dbus.Call) (outArg0 bool, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceManager) IsAutostart(flags dbus.Flags, fileName string) (bool, error) {
	return v.StoreIsAutostart(
		<-v.GoIsAutostart(flags, make(chan *dbus.Call, 1), fileName).Done)
}

// method RemoveAutostart

func (v *interfaceManager) GoRemoveAutostart(flags dbus.Flags, ch chan *dbus.Call, fileNamae string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveAutostart", flags, ch, fileNamae)
}

func (*interfaceManager) StoreRemoveAutostart(call *dbus.Call) (outArg0 bool, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceManager) RemoveAutostart(flags dbus.Flags, fileNamae string) (bool, error) {
	return v.StoreRemoveAutostart(
		<-v.GoRemoveAutostart(flags, make(chan *dbus.Call, 1), fileNamae).Done)
}

// method Launch

func (v *interfaceManager) GoLaunch(flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Launch", flags, ch, desktopFile)
}

func (v *interfaceManager) Launch(flags dbus.Flags, desktopFile string) error {
	return (<-v.GoLaunch(flags, make(chan *dbus.Call, 1), desktopFile).Done).Err
}

// method LaunchApp

func (v *interfaceManager) GoLaunchApp(flags dbus.Flags, ch chan *dbus.Call, desktopFile string, timestamp uint32, files []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchApp", flags, ch, desktopFile, timestamp, files)
}

func (v *interfaceManager) LaunchApp(flags dbus.Flags, desktopFile string, timestamp uint32, files []string) error {
	return (<-v.GoLaunchApp(flags, make(chan *dbus.Call, 1), desktopFile, timestamp, files).Done).Err
}

// method LaunchAppAction

func (v *interfaceManager) GoLaunchAppAction(flags dbus.Flags, ch chan *dbus.Call, desktopFile string, action string, timestamp uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchAppAction", flags, ch, desktopFile, action, timestamp)
}

func (v *interfaceManager) LaunchAppAction(flags dbus.Flags, desktopFile string, action string, timestamp uint32) error {
	return (<-v.GoLaunchAppAction(flags, make(chan *dbus.Call, 1), desktopFile, action, timestamp).Done).Err
}

// signal AutostartChanged

func (v *interfaceManager) ConnectAutostartChanged(cb func(status string, filePath string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AutostartChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AutostartChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var status string
		var filePath string
		err := dbus.Store(sig.Body, &status, &filePath)
		if err == nil {
			cb(status, filePath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

type Instance interface {
	instance // interface org.deepin.dde.Application1.Instance
	proxy.Object
}

type objectInstance struct {
	interfaceInstance // interface org.deepin.dde.Application1.Instance
	proxy.ImplObject
}

func NewInstance(conn *dbus.Conn) Instance {
	obj := new(objectInstance)
	obj.ImplObject.Init_(conn, "org.deepin.dde.Application1", "/org/deepin/dde/Application1/Instance")
	return obj
}

type instance interface {
	GoExit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Exit(flags dbus.Flags) error
	GoKill(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Kill(flags dbus.Flags) error
	Id() proxy.PropObjectPath
}

type interfaceInstance struct{}

func (v *interfaceInstance) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceInstance) GetInterfaceName_() string {
	return "org.deepin.dde.Application1.Instance"
}

// method Exit

func (v *interfaceInstance) GoExit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Exit", flags, ch)
}

func (v *interfaceInstance) Exit(flags dbus.Flags) error {
	return (<-v.GoExit(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Kill

func (v *interfaceInstance) GoKill(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Kill", flags, ch)
}

func (v *interfaceInstance) Kill(flags dbus.Flags) error {
	return (<-v.GoKill(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property id o

func (v *interfaceInstance) Id() proxy.PropObjectPath {
	return &proxy.ImplPropObjectPath{
		Impl: v,
		Name: "id",
	}
}
