// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package object

// ----------------------------------------------------------------------
// Defines Type
// ----------------------------------------------------------------------

type AddressObjectType struct {
	Condition string `json:"condition,omitempty"`
	Value     string `json:"address_value,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func NewAddressObject() AddressObjectType {
	obj := CreateAddressObject()
	return obj
}

func CreateAddressObject() AddressObjectType {
	var obj AddressObjectType
	return obj
}

// ----------------------------------------------------------------------
// Methods
// ----------------------------------------------------------------------

func (this *AddressObjectType) SetConditionEquals() {
	this.Condition = "Equals"
}

func (this *AddressObjectType) AddValue(v string) {
	this.Value = v
}
