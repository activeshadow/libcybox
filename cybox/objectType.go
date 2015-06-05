// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package cybox

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/freestix/libcybox/common"
	"github.com/freestix/libcybox/defs"
)

type ObjectType struct {
	Id         string                `json:"id,omitempty"`
	IdRef      string                `json:"idref,omitempty"`
	HasChanged bool                  `json:"hasChanged,omitempty"`
	VocabName  string                `json:"vocabName,omitempty"`
	VocabType  string                `json:"vocabType,omitempty"`
	UriObjects *common.UriObjectType `json:"uriObj,omitempty"`
}

// ----------------------------------------------------------------------
// Methods
// ----------------------------------------------------------------------

func (o *ObjectType) CreateId() {
	o.Id = defs.COMPANY + ":cyboxObject-" + uuid.New()
}

func (o *ObjectType) AddUriObject(u common.UriObjectType) {
	o.UriObjects = &u
}
