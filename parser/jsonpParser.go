package parser

import (
	"regexp"

	"github.com/ErosZy/labour/common"
	"github.com/ErosZy/labour/models"
)

func ParseJSONP(body string, jsonItems []models.JSONItem) []models.KeyValuePair {
	reg, err := regexp.Compile(`(?:[\s\S]+?)\(([\s\S]*)\)`)

	if err != nil {
		common.Logger(common.LOG_WARNING, err.Error())
		return []models.KeyValuePair{}
	}

	matches := reg.FindAllStringSubmatch(body, -1)

	if matches == nil {
		common.Logger(common.LOG_WARNING, "can't find the submatches in jsonp!")
		return []models.KeyValuePair{}
	}

	return ParseJSON(matches[0][1], jsonItems)
}
