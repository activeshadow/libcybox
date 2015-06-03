// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package common

// TODO: change the value to use time.Time and import time
type DateTimeWithPrecisionType struct {
	Precision string `json:"precision,omitempty"`
	Value     string `json:"value,omitempty"`
}

// Issues: Need to support ISO 8601, specifically 2015-01-15T14:16:00-07:00
type TimeType struct {
	StartTime    *DateTimeWithPrecisionType `json:"startTime,omitempty"`
	EndTime      *DateTimeWithPrecisionType `json:"endTime,omitempty"`
	ProducedTime *DateTimeWithPrecisionType `json:"producedTime,omitempty"`
	ReceivedTime *DateTimeWithPrecisionType `json:"receivedTime,omitempty"`
}

// TODO need to finish this type
type ToolInformationType struct {
	Id    string `json:"id,omitempty"`
	IdRef string `json:"idref,omitempty"`
}
