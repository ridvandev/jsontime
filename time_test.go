package gojsontime_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	rc "github.com/ridvandev/jsontime"
)

type TestDate struct {
	Date rc.DateTime `json:"date"`
}

func Test_DateTime(t *testing.T) {
	now := TestDate{Date: rc.DateTime{Time: time.Date(2022, 12, 28, 23, 12, 5, 0, time.UTC)}}
	nowStr, err := json.Marshal(now)
	assert.Nil(t, err)
	assert.Equal(t, "{\"date\":\"2022-12-28T23:12:05.000Z\"}", string(nowStr))

	var now2 TestDate
	err = json.Unmarshal(nowStr, &now2)
	assert.Nil(t, err)
	assert.Equal(t, now, now2)

	var empty TestDate
	emptyTimeJson := "{\"date\":\"\"}"
	err = json.Unmarshal([]byte(emptyTimeJson), &empty)
	assert.Nil(t, err)
	assert.Equal(t, TestDate{}, empty)
	assert.Equal(t, rc.DateTime{}, empty.Date)
	assert.Equal(t, time.Time{}, empty.Date.Time)
	assert.Equal(t, true, empty.Date.Time.IsZero())

	emptyTimeJsonTest, err := json.Marshal(empty)
	assert.Nil(t, err)
	assert.Equal(t, emptyTimeJson, string(emptyTimeJsonTest))
}
