// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1
// Last Updated: 2015-01-26
// Status: Level 1

package cybox

import (
	"github.com/jordan2175/freestix/libcybox/common"
)

// ----------------------------------------------------------------------
// Top Level Cybox Objects
// ----------------------------------------------------------------------

func CreateObservable() ObservableType {
	var o ObservableType
	return o
}

func CreateObject(vocabName, value string) ObjectType {
	var o ObjectType
	o.VocabName = vocabName
	o.VocabType = value
	return o
}

// ----------------------------------------------------------------------
// Common Cybox Objects
// ----------------------------------------------------------------------

func CreateUriObject() common.UriObjectType {
	var u common.UriObjectType
	return u
}
