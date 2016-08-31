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
	Id             string              `json:"id,omitempty"`
	IdRef          string              `json:"idref,omitempty"`
	Type           string              `json:"type,omitempty"`
	Category       string              `json:"category,omitempty"`
	UriObjects     []UriObjectType     `json:"uri_objects,omitempty"`
	AddressObjects []AddressObjectType `json:"address_objects,omitempty"`
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

func (this *PropertiesType) AddCategory(c string) {
	this.Category = c
}

func (this *PropertiesType) AddEqualsUriValue(uri string) {
	uriobj := NewUriObject()
	uriobj.SetConditionEquals()
	uriobj.AddValue(uri)
	this.AddUriObject(uriobj)
}

func (this *PropertiesType) AddEqualsAddressValue(addr string) {
	addrobj := NewAddressObject()
	addrobj.SetConditionEquals()
	addrobj.AddValue(addr)
	this.AddAddressObject(addrobj)
}

func (this *PropertiesType) AddUriObject(obj UriObjectType) {
	if this.UriObjects == nil {
		a := make([]UriObjectType, 0)
		this.UriObjects = a
	}
	this.UriObjects = append(this.UriObjects, obj)
}

func (this *PropertiesType) AddAddressObject(obj AddressObjectType) {
	// Don't really need this... append will make the slice for you if it
	// doesn't already exist.
	if this.AddressObjects == nil {
		a := make([]AddressObjectType, 0)
		this.AddressObjects = a
	}
	this.AddressObjects = append(this.AddressObjects, obj)
}
