package process

import (
	"encoding/json"
	_common "labour/common"
	"labour/models"
	"labour/parser"
	"labour/router"
	"net/url"

	"github.com/ErosZy/singoriensis/common"
)

type Process struct {
	router *router.Router
}

func NewProcess(taskInfo *models.TaskInfoModel) *Process {
	router := router.NewRouter()
	schedulers := taskInfo.Schedulers
	pages := taskInfo.Pages

	for index, schedulerItem := range schedulers {
		router.Add(
			schedulerItem.Route,
			getProcessRouterHandler(&pages[index], &schedulerItem),
		)
	}

	return &Process{router}
}

func getProcessRouterHandler(pageItem *models.PageItemModel, schedulerItem *models.SchedulerItemModel) router.RouterHandlerFunc {
	return func(urlStr string, params ...interface{}) {
		var pageItems [][]models.KeyValuePair
		var schedulers []models.KeyValuePair

		page := params[0].(*common.Page)
		bodyStr := _common.CharsetIconv(page)

		if pageItem.Text != nil {
			pageItems = make([][]models.KeyValuePair, 0)
			pXpath := pageItem.Text.Xpath
			if len(pXpath) > 0 {
				tmp := parser.ParsePageItemXpath(bodyStr, pXpath)

				for _, v := range tmp {
					pageItems = append(pageItems, v)
				}
			}

			pRegex := pageItem.Text.Regex
			if len(pRegex) > 0 {
				tmp := parser.ParseRegex(bodyStr, pRegex)
				for _, v := range tmp {
					pageItems = append(pageItems, v)
				}
			}

			pJson := pageItem.Json
			if pJson.JSONs != nil {
				switch pJson.Type {
				case 0:
					pageItems = append(pageItems, parser.ParseJSON(bodyStr, pJson.JSONs))
				case 1:
					pageItems = append(pageItems, parser.ParseJSONP(bodyStr, pJson.JSONs))
				}
			}

			addPageItems(page, pageItems, pageItem.MainKey)
		}

		schedulers = make([]models.KeyValuePair, 0)

		sXpath := schedulerItem.Text.Xpath
		if len(sXpath) > 0 {
			schedulers = concat(schedulers, parser.ParseSchedulerItemXpath(bodyStr, sXpath))
		}

		sRegex := schedulerItem.Text.Regex
		if len(sRegex) > 0 {
			tmp := parser.ParseRegex(bodyStr, sRegex)
			for _, v := range tmp {
				schedulers = concat(schedulers, v)
			}
		}

		sJson := schedulerItem.Json
		if sJson.JSONs != nil {
			switch sJson.Type {
			case 0:
				schedulers = concat(schedulers, parser.ParseJSON(bodyStr, sJson.JSONs))
			case 1:
				schedulers = concat(schedulers, parser.ParseJSONP(bodyStr, sJson.JSONs))
			}
		}

		addSchedulers(page, schedulers)
	}
}

func addSchedulers(page *common.Page, schedulers []models.KeyValuePair) {
	for _, v := range schedulers {
		u, err := url.Parse(v.Value)
		if err != nil {
			_common.Logger(_common.LOG_FATAL, err.Error())
		}

		var rUrl string
		if !u.IsAbs() {
			rUrl = page.Req.URL.ResolveReference(u).String()
		} else {
			rUrl = v.Value
		}

		page.AddElem(common.NewElementItem(rUrl))
	}
}

func addPageItems(page *common.Page, pageItems [][]models.KeyValuePair, mainKey string) {
	tmp := make([]models.KeyValuePair, 0)
	pipelinerItem := _common.PipelinerItem{}

	for _, v := range pageItems {
		tmp = append(tmp, v...)
	}

	data, err := json.Marshal(tmp)

	if err != nil {
		_common.Logger(_common.LOG_FATAL, err.Error())
	}

	pipelinerItem.Body = string(data)
	pipelinerItem.BaseUrl = page.Req.URL.String()
	pipelinerItem.MainKey = mainKey
	pipelinerItem.MainKeyValue = getMainKeyValue(tmp, mainKey)
	pipelinerItem.BodyType = "text/html"

	page.AddItem(pipelinerItem)
}

func getMainKeyValue(items []models.KeyValuePair, mainKey string) string {
	for _, v := range items {
		if v.Key == mainKey {
			return v.Value
		}
	}

	return "-1"
}

func concat(o1 []models.KeyValuePair, o2 []models.KeyValuePair) []models.KeyValuePair {
	for _, v := range o2 {
		o1 = append(o1, v)
	}

	return o1
}
