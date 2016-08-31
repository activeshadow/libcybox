// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package object

import (
	"github.com/freestix/libcybox"
	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type ObjectType struct {
	Id         string          `json:"id,omitempty"`
	IdRef      string          `json:"idref,omitempty"`
	HasChanged bool            `json:"has_changed,omitempty"`
	Properties *PropertiesType `json:"properties,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func NewObject() ObjectType {
	obj := CreateObject()
	obj.CreateId()
	return obj
}

func CreateObject() ObjectType {
	var obj ObjectType
	return obj
}

// ----------------------------------------------------------------------
// Methods
// ----------------------------------------------------------------------

func (this *ObjectType) CreateId() {
	this.Id = libcybox.Company + ":cyboxObject-" + uuid.New()
}

func (this *ObjectType) AddProperties(u PropertiesType) {
	this.Properties = &u
}
