// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package cybox

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/jordan2175/freestix/libcybox/defs"
)

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
// Methods
// ----------------------------------------------------------------------

func (o *ObservableType) CreateId() {
	o.Id = defs.COMPANY + ":observable-" + uuid.New()
}

func (o *ObservableType) AddIdRef(idref string) {
	o.IdRef = idref
}

func (o *ObservableType) SetNegate(b bool) {
	o.Negate = b
}

func (o *ObservableType) IncreaseSightingCount(u uint) {
	o.SightingCount = +u
}

func (o *ObservableType) AddObject(obj ObjectType) {
	o.Object = &obj
}
