package parser

import (
	"encoding/json"
	"github.com/ErosZy/labour/common"
	"github.com/ErosZy/labour/models"
	"regexp"
	"strings"
)

func ParseRegex(body string, regex []*models.RegexItem) [][]models.KeyValuePair {
	items := make([][]models.KeyValuePair, 0)
	body = common.UnicodeConvert(body)

	for _, v := range regex {
		r, err := regexp.Compile(v.RegexStr)
		if err != nil {
			common.Logger(common.LOG_WARNING, err.Error())
			return items
		}

		pair := make([]string, 0)
		for _, v1 := range v.Match {
			pair = append(pair, `{"key":"`+v1.Key+`","value":"`+v1.Value+`"}`)
		}

		jsonStr := `[` + strings.Join(pair, ",") + `]`
		dst := []byte("")
		match := r.FindAllStringSubmatchIndex(body, -1)

		for _, v2 := range match {
			var kv []models.KeyValuePair
			tmp := r.ExpandString(dst, jsonStr, body, v2)

			err := json.Unmarshal(tmp, &kv)
			if err != nil {
				common.Logger(common.LOG_WARNING, err.Error())
				return items
			}

			items = append(items, kv)
		}
	}

	return items
}
