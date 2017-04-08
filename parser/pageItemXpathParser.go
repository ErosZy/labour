package parser

import (
	"labour/common"
	"labour/models"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParsePageItemXpath(body string, xpath []*models.PrefixXpathItem) [][]models.KeyValuePair {
	reader := strings.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)
	r := make([][]models.KeyValuePair, 0)
	filter := regexp.MustCompile("\\s+")

	if err != nil {
		common.Logger(common.LOG_WARNING, err.Error())
		return r
	}

	for _, v := range xpath {
		doc.Find(v.Prefix).Each(func(i int, s *goquery.Selection) {
			tmp := make([]models.KeyValuePair, 0)
			items := v.Arr

			for _, item := range items {
				sel := s.Find(item.DomStr)
				if sel.Length() != 0 {
					sel.Each(func(j int, s1 *goquery.Selection) {
						if item.Type == 0 {
							tmp = append(tmp, models.KeyValuePair{item.Key, filter.ReplaceAllString(s1.Text(), "")})
						} else if item.Type == 1 {
							attrValue, exists := s1.Attr(item.AttrKey)

							if exists {
								tmp = append(tmp, models.KeyValuePair{item.Key, filter.ReplaceAllString(attrValue, "")})
							}
						} else {
							ret, _ := s1.Html()
							tmp = append(tmp, models.KeyValuePair{item.Key, filter.ReplaceAllString(ret, "")})
						}
					})
				} else {
					tmp = append(tmp, models.KeyValuePair{item.Key, ""})
				}
			}

			r = append(r, tmp)
		})
	}

	return r
}
