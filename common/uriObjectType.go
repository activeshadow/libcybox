// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package common

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/freestix/libcybox/defs"
)

type UriObjectType struct {
	Id         string   `json:"id,omitempty"`
	IdRef      string   `json:"idref,omitempty"`
	UriObjects []string `json:"uriObjects,omitempty"`
}

// ----------------------------------------------------------------------
// Methods
// ----------------------------------------------------------------------

func (u *UriObjectType) CreateId() {
	u.Id = defs.COMPANY + ":UriObject-" + uuid.New()
}

func (u *UriObjectType) AddUriObject(value string) {
	if u.UriObjects == nil {
		a := make([]string, 0)
		u.UriObjects = a
	}
	u.UriObjects = append(u.UriObjects, value)
}
