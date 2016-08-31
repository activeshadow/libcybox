// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package observable

import (
	"github.com/freestix/libcybox/defs"
	"github.com/freestix/libcybox/object"
	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
// Defines Type
// ----------------------------------------------------------------------
// TODO Finish filling out this data structure

type ObservableType struct {
	Id            string             `json:"id,omitempty"`
	IdRef         string             `json:"idref,omitempty"`
	Negate        bool               `json:"negate,omitempty"`
	SightingCount uint               `json:"sighting_count,omitempty"`
	Title         string             `json:"title,omitempty"`
	Object        *object.ObjectType `json:"object,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func New() ObservableType {
	obj := CreateObservable()
	obj.CreateId()
	cyboxobj := object.NewObject()
	prop := object.NewProperties()
	cyboxobj.AddProperties(prop)
	obj.AddObject(cyboxobj)
	return obj
}

func CreateObservable() ObservableType {
	var obj ObservableType
	return obj
}

// ----------------------------------------------------------------------
// Methods
// ----------------------------------------------------------------------

func (this *ObservableType) CreateId() {
	this.Id = defs.COMPANY + ":observable-" + uuid.New()
}

func (this *ObservableType) AddIdRef(idref string) {
	this.IdRef = idref
}

func (this *ObservableType) SetNegate(b bool) {
	this.Negate = b
}

func (this *ObservableType) IncreaseSightingCount(u uint) {
	this.SightingCount = +u
}

func (this *ObservableType) AddObject(obj object.ObjectType) {
	this.Object = &obj
}

func (this *ObservableType) GetObject() *object.ObjectType {
	return this.Object
}

func (this *ObservableType) GetObjectProperties() *object.PropertiesType {
	return this.Object.Properties
}
