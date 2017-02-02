// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package object

import (
	"github.com/pborman/uuid"
	"github.com/activeshadow/libcybox/defs"
)

// ----------------------------------------------------------------------
// Defines Type
// ----------------------------------------------------------------------

type PropertiesType struct {
	Id         string          `json:"id,omitempty"`
	IdRef      string          `json:"idref,omitempty"`
	Type       string          `json:"type,omitempty"`
	UriObjects []UriObjectType `json:"uri_objects,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func NewProperties() PropertiesType {
	obj := CreateCyboxPropertiesObject()
	obj.CreateId()
	return obj
}

func CreateCyboxPropertiesObject() PropertiesType {
	var obj PropertiesType
	return obj
}

// ----------------------------------------------------------------------
// Methods
// ----------------------------------------------------------------------

func (this *PropertiesType) CreateId() {
	this.Id = defs.COMPANY + ":object-" + uuid.New()
}

func (this *PropertiesType) AddType(t string) {
	this.Type = t
}

func (this *PropertiesType) AddEqualsUriValue(uri string) {
	uriobj := NewUriObject()
	uriobj.SetConditionEquals()
	uriobj.AddValue(uri)
	this.AddUriObject(uriobj)
}

func (this *PropertiesType) AddUriObject(obj UriObjectType) {
	if this.UriObjects == nil {
		a := make([]UriObjectType, 0)
		this.UriObjects = a
	}
	this.UriObjects = append(this.UriObjects, obj)
}
