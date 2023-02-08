// Code generated by "./generator ./system/org.deepin.dde.authenticate1.face"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package face

import "errors"
import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Face interface {
	face // interface org.deepin.dde.Authenticate1.Face.Device
	proxy.Object
}

type objectFace struct {
	interfaceFace // interface org.deepin.dde.Authenticate1.Face.Device
	proxy.ImplObject
}

func NewFace(conn *dbus.Conn, serviceName string, path dbus.ObjectPath) (Face, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectFace)
	obj.ImplObject.Init_(conn, serviceName, path)
	return obj, nil
}

type face interface {
	SetInterfaceName_(name string)
	GoClaim(flags dbus.Flags, ch chan *dbus.Call, claimed bool) *dbus.Call
	Claim(flags dbus.Flags, claimed bool) error
	GoDeleteFace(flags dbus.Flags, ch chan *dbus.Call, uuid string, faceName string) *dbus.Call
	DeleteFace(flags dbus.Flags, uuid string, faceName string) error
	GoDeleteFaces(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call
	DeleteFaces(flags dbus.Flags, uuid string) error
	GoGetShareMemInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetShareMemInfo(flags dbus.Flags) (string, string, int32, error)
	GoListFaces(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call
	ListFaces(flags dbus.Flags, uuid string) ([]string, error)
	GoRenameFace(flags dbus.Flags, ch chan *dbus.Call, uuid string, oldFaceName string, newFaceName string) *dbus.Call
	RenameFace(flags dbus.Flags, uuid string, oldFaceName string, newFaceName string) error
	GoSetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call, device string) *dbus.Call
	SetDefaultDevice(flags dbus.Flags, device string) error
	GoStartEnroll(flags dbus.Flags, ch chan *dbus.Call, uuid string, faceName string) *dbus.Call
	StartEnroll(flags dbus.Flags, uuid string, faceName string) error
	GoStartVerify(flags dbus.Flags, ch chan *dbus.Call, uuid string, timeout int32) *dbus.Call
	StartVerify(flags dbus.Flags, uuid string, timeout int32) error
	GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	StopEnroll(flags dbus.Flags) error
	GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	StopVerify(flags dbus.Flags) error
	ConnectEnrollStatus(cb func(uuid string, code int32, description string)) (dbusutil.SignalHandlerId, error)
	ConnectVerifyStatus(cb func(uuid string, code int32, description string)) (dbusutil.SignalHandlerId, error)
	SupportedDevices() proxy.PropStringArray
	DefaultDevice() proxy.PropString
	Name() proxy.PropString
	Claimed() proxy.PropBool
	Capability() proxy.PropInt32
	Status() proxy.PropInt32
}

type interfaceFace struct{}

func (v *interfaceFace) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (v *interfaceFace) SetInterfaceName_(name string) {
	v.GetObject_().SetExtra("customIfc", name)
}

func (v *interfaceFace) GetInterfaceName_() string {
	ifcName, _ := v.GetObject_().GetExtra("customIfc")
	ifcNameStr, _ := ifcName.(string)
	return ifcNameStr
}

// method Claim

func (v *interfaceFace) GoClaim(flags dbus.Flags, ch chan *dbus.Call, claimed bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Claim", flags, ch, claimed)
}

func (v *interfaceFace) Claim(flags dbus.Flags, claimed bool) error {
	return (<-v.GoClaim(flags, make(chan *dbus.Call, 1), claimed).Done).Err
}

// method DeleteFace

func (v *interfaceFace) GoDeleteFace(flags dbus.Flags, ch chan *dbus.Call, uuid string, faceName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteFace", flags, ch, uuid, faceName)
}

func (v *interfaceFace) DeleteFace(flags dbus.Flags, uuid string, faceName string) error {
	return (<-v.GoDeleteFace(flags, make(chan *dbus.Call, 1), uuid, faceName).Done).Err
}

// method DeleteFaces

func (v *interfaceFace) GoDeleteFaces(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteFaces", flags, ch, uuid)
}

func (v *interfaceFace) DeleteFaces(flags dbus.Flags, uuid string) error {
	return (<-v.GoDeleteFaces(flags, make(chan *dbus.Call, 1), uuid).Done).Err
}

// method GetShareMemInfo

func (v *interfaceFace) GoGetShareMemInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetShareMemInfo", flags, ch)
}

func (*interfaceFace) StoreGetShareMemInfo(call *dbus.Call) (sockPath string, key string, size int32, err error) {
	err = call.Store(&sockPath, &key, &size)
	return
}

func (v *interfaceFace) GetShareMemInfo(flags dbus.Flags) (string, string, int32, error) {
	return v.StoreGetShareMemInfo(
		<-v.GoGetShareMemInfo(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListFaces

func (v *interfaceFace) GoListFaces(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListFaces", flags, ch, uuid)
}

func (*interfaceFace) StoreListFaces(call *dbus.Call) (faces []string, err error) {
	err = call.Store(&faces)
	return
}

func (v *interfaceFace) ListFaces(flags dbus.Flags, uuid string) ([]string, error) {
	return v.StoreListFaces(
		<-v.GoListFaces(flags, make(chan *dbus.Call, 1), uuid).Done)
}

// method RenameFace

func (v *interfaceFace) GoRenameFace(flags dbus.Flags, ch chan *dbus.Call, uuid string, oldFaceName string, newFaceName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RenameFace", flags, ch, uuid, oldFaceName, newFaceName)
}

func (v *interfaceFace) RenameFace(flags dbus.Flags, uuid string, oldFaceName string, newFaceName string) error {
	return (<-v.GoRenameFace(flags, make(chan *dbus.Call, 1), uuid, oldFaceName, newFaceName).Done).Err
}

// method SetDefaultDevice

func (v *interfaceFace) GoSetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call, device string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDefaultDevice", flags, ch, device)
}

func (v *interfaceFace) SetDefaultDevice(flags dbus.Flags, device string) error {
	return (<-v.GoSetDefaultDevice(flags, make(chan *dbus.Call, 1), device).Done).Err
}

// method StartEnroll

func (v *interfaceFace) GoStartEnroll(flags dbus.Flags, ch chan *dbus.Call, uuid string, faceName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartEnroll", flags, ch, uuid, faceName)
}

func (v *interfaceFace) StartEnroll(flags dbus.Flags, uuid string, faceName string) error {
	return (<-v.GoStartEnroll(flags, make(chan *dbus.Call, 1), uuid, faceName).Done).Err
}

// method StartVerify

func (v *interfaceFace) GoStartVerify(flags dbus.Flags, ch chan *dbus.Call, uuid string, timeout int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartVerify", flags, ch, uuid, timeout)
}

func (v *interfaceFace) StartVerify(flags dbus.Flags, uuid string, timeout int32) error {
	return (<-v.GoStartVerify(flags, make(chan *dbus.Call, 1), uuid, timeout).Done).Err
}

// method StopEnroll

func (v *interfaceFace) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopEnroll", flags, ch)
}

func (v *interfaceFace) StopEnroll(flags dbus.Flags) error {
	return (<-v.GoStopEnroll(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method StopVerify

func (v *interfaceFace) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopVerify", flags, ch)
}

func (v *interfaceFace) StopVerify(flags dbus.Flags) error {
	return (<-v.GoStopVerify(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal EnrollStatus

func (v *interfaceFace) ConnectEnrollStatus(cb func(uuid string, code int32, description string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EnrollStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EnrollStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var uuid string
		var code int32
		var description string
		err := dbus.Store(sig.Body, &uuid, &code, &description)
		if err == nil {
			cb(uuid, code, description)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *interfaceFace) ConnectVerifyStatus(cb func(uuid string, code int32, description string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var uuid string
		var code int32
		var description string
		err := dbus.Store(sig.Body, &uuid, &code, &description)
		if err == nil {
			cb(uuid, code, description)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property SupportedDevices as

func (v *interfaceFace) SupportedDevices() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "SupportedDevices",
	}
}

// property DefaultDevice s

func (v *interfaceFace) DefaultDevice() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "DefaultDevice",
	}
}

// property Name s

func (v *interfaceFace) Name() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Name",
	}
}

// property Claimed b

func (v *interfaceFace) Claimed() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Claimed",
	}
}

// property Capability i

func (v *interfaceFace) Capability() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "Capability",
	}
}

// property Status i

func (v *interfaceFace) Status() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "Status",
	}
}
