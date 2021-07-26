// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admi_v02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, Event2{}.Validate())
	assert.NotNil(t, SystemEventNotificationV02{}.Validate())
	assert.NotNil(t, RequestDetails3{}.Validate())
	assert.NotNil(t, StaticDataRequestV02{}.Validate())
	assert.NotNil(t, ReportParameter1{}.Validate())
	assert.NotNil(t, RequestDetails4{}.Validate())
	assert.NotNil(t, RequestDetails5{}.Validate())
	assert.NotNil(t, StaticDataReportV02{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
}
