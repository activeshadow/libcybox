// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

// TODO: change the value to use time.Time and import time
type DateTimeWithPrecisionType struct {
	Precision string `json:"precision,omitempty"`
	Value     string `json:"value,omitempty"`
}

// Issues: Need to support ISO 8601, specifically 2015-01-15T14:16:00-07:00
type TimeType struct {
	StartTime    *DateTimeWithPrecisionType `json:"start_time,omitempty"`
	EndTime      *DateTimeWithPrecisionType `json:"end_time,omitempty"`
	ProducedTime *DateTimeWithPrecisionType `json:"produced_time,omitempty"`
	ReceivedTime *DateTimeWithPrecisionType `json:"received_time,omitempty"`
}

// TODO need to finish this type
type ToolInformationType struct {
	Id    string `json:"id,omitempty"`
	IdRef string `json:"idref,omitempty"`
}
