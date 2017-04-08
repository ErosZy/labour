package parser

import (
	"encoding/json"
	"labour/common"
	"labour/models"
	"regexp"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
)

func ParseJSON(body string, jsonItems []models.JSONItem) []models.KeyValuePair {
	j, err := simplejson.NewJson([]byte(body))
	tmp := make([]models.KeyValuePair, 0)

	if err != nil {
		common.Logger(common.LOG_WARNING, err.Error())
		return tmp
	}

	for _, v := range jsonItems {
		keys := strings.Split(v.Value, ".")
		typePaths := strings.Split(v.TypeStr, ".")
		val := getRecursionValue(j, keys, typePaths)

		valStr := ""
		switch val.(type) {
		case bool:
			valStr = strconv.FormatBool(val.(bool))
		case json.Number:
			valStr = string(val.(json.Number))
		case string:
			valStr = val.(string)
		}

		tmp = append(tmp, models.KeyValuePair{v.Key, valStr})
	}

	return tmp
}

func getRecursionValue(json *simplejson.Json, keys []string, typePaths []string) interface{} {
	index := 0
	filter := regexp.MustCompile("\\s+")

	for ; index < len(keys); index++ {
		key := filter.ReplaceAllString(keys[index], "")
		itemType := strings.ToLower(filter.ReplaceAllString(typePaths[index], ""))

		if itemType == "array" {
			startIndex := strings.Index(key, "[")
			endIndex := strings.LastIndex(key, "]")
			realKey := key[0:startIndex]
			realIndex, err := strconv.Atoi(key[startIndex+1 : endIndex])

			if err != nil {
				common.Logger(common.LOG_WARNING, err.Error())
				return ""
			}

			json = json.Get(realKey).GetIndex(realIndex)
		} else {
			json = json.Get(key)
		}
	}

	return json.Interface()
}
