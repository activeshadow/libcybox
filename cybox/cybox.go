// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package cybox

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/freestix/libcybox/defs"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------
// TODO: Finish filling out this data structure

type ObservableType struct {
	Id            string      `json:"id,omitempty"`
	IdRef         string      `json:"idref,omitempty"`
	Negate        bool        `json:"negate,omitempty"`
	SightingCount uint        `json:"sightingCount,omitempty"`
	Title         string      `json:"title,omitempty"`
	Object        *ObjectType `json:"object,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func New() ObservableType {
	obj := CreateObservable()
	obj.CreateId()
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

func (this *ObservableType) AddObject(obj ObjectType) {
	this.Object = &obj
}
