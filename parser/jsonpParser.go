package parser

import (
	"labour/common"
	"labour/models"
	"regexp"
)

/*
	body := `
	test({
			"name":"zyEros",
			"age" : 24,
			"female": false,
			"project":{
				"name":[
					"hello","zyEros"
				]
			}
		});
	`

	pairs := make([]models.JSONItem,0)

	pairs = append(pairs,models.JSONItem{models.KeyValuePair{"project","project.name[1]"},"Object.Array"},
		models.JSONItem{models.KeyValuePair{"name","name"},"String"},
		models.JSONItem{models.KeyValuePair{"age","age"},"Int"},
		models.JSONItem{models.KeyValuePair{"female","female"},"Bool"})

	fmt.Println(parser.ParseJSONP(body,pairs))
	return
*/

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
