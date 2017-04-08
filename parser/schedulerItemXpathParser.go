package parser

import (
	"labour/common"
	"labour/models"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseSchedulerItemXpath(body string, xpath []*models.XpathItem) []models.KeyValuePair {
	reader := strings.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)
	r := make([]models.KeyValuePair, 0)
	filter := regexp.MustCompile("\\s+")

	if err != nil {
		common.Logger(common.LOG_WARNING, err.Error())
		return r
	}

	for _, v := range xpath {
		doc.Find(v.DomStr).Each(func(i int, s *goquery.Selection) {
			if v.Type == 0 {
				r = append(r, models.KeyValuePair{v.Key, filter.ReplaceAllString(s.Text(), "")})
			} else if v.Type == 1 {
				attrValue, exists := s.Attr(v.AttrKey)

				if exists {
					r = append(r, models.KeyValuePair{v.Key, filter.ReplaceAllString(attrValue, "")})
				}
			} else {
				ret, _ := s.Html()
				r = append(r, models.KeyValuePair{v.Key, filter.ReplaceAllString(ret, "")})
			}
		})
	}

	return r
}
