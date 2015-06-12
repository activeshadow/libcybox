// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package object

// ----------------------------------------------------------------------
// Defines Type
// ----------------------------------------------------------------------

type UriObjectType struct {
	Condition string `json:"condition,omitempty"`
	Value     string `json:"value,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func NewUriObject() UriObjectType {
	obj := CreateUriObject()
	return obj
}

func CreateUriObject() UriObjectType {
	var obj UriObjectType
	return obj
}

// ----------------------------------------------------------------------
// Methods
// ----------------------------------------------------------------------

func (this *UriObjectType) SetConditionEquals() {
	this.Condition = "Equals"
}

func (this *UriObjectType) AddValue(v string) {
	this.Value = v
}
