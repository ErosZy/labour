package models

import (
	"encoding/json"

	"github.com/ErosZy/labour/common"
)

type RequestHeaderItem struct {
	KeyValuePair
}

type RequestHeaderModel struct {
	RequestHeaders []*RequestHeaderItem
}

func NewRequestHeaderModel(headerStr string) *RequestHeaderModel {
	instance := &RequestHeaderModel{}

	if headerStr != "" {
		err := json.Unmarshal([]byte(headerStr), &instance.RequestHeaders)

		if err != nil {
			common.Logger(common.LOG_WARNING, err.Error())
		}
	}

	return instance
}
