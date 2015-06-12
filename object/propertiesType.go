// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package object

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/freestix/libcybox/defs"
)

// ----------------------------------------------------------------------
// Defines Type
// ----------------------------------------------------------------------

type PropertiesType struct {
	Id         string   `json:"id,omitempty"`
	IdRef      string   `json:"idref,omitempty"`
	Type       string   `json:"type,omitempty"`
	UriObjects []string `json:"uri_objects,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func NewProperties() PropertiesType {
	obj := CreateUriObject()
	obj.CreateId()
	return obj
}

func CreateUriObject() PropertiesType {
	var obj PropertiesType
	return obj
}

// ----------------------------------------------------------------------
// Methods
// ----------------------------------------------------------------------

func (this *PropertiesType) CreateId() {
	this.Id = defs.COMPANY + ":UriObject-" + uuid.New()
}

func (this *PropertiesType) AddType(t string) {
	this.Type = t
}

func (this *PropertiesType) AddUriObject(value string) {
	if this.UriObjects == nil {
		a := make([]string, 0)
		this.UriObjects = a
	}
	this.UriObjects = append(this.UriObjects, value)
}
