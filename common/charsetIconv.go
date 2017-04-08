package common

import (
	"regexp"
	"strings"

	"github.com/ErosZy/singoriensis/common"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

func CharsetIconv(page *common.Page) string {
	contentType := page.Res.Header.Get("Content-Type")
	bodyStr := page.ResBody

	reg := regexp.MustCompile("text\\s*/\\s*html\\s*;\\s*charset\\s*=\\s*(.+)")
	isGBK := false

	matches := reg.FindStringSubmatch(contentType)

	if contentType != "" && (len(matches) > 1 && "gbk" == strings.ToLower(matches[1])) {
		isGBK = true
	} else if contentType == "" {
		reader := strings.NewReader(bodyStr)
		doc, err := goquery.NewDocumentFromReader(reader)

		if err != nil {
			Logger(LOG_WARNING, err.Error())
			return bodyStr
		}

		doc.Find("meta").Each(func(i int, s1 *goquery.Selection) {
			if !isGBK {
				val, exists := s1.Attr("charset")
				if exists && "gbk" == strings.ToLower(val) {
					isGBK = true
				}

				val, exists = s1.Attr("content")
				if exists {
					matches := reg.FindStringSubmatch(val)
					if "gbk" == strings.ToLower(matches[1]) {
						isGBK = true
					}
				}
			}
		})
	}

	if isGBK {
		decoder := mahonia.NewDecoder("gbk")
		bodyStr = decoder.ConvertString(bodyStr)
	}

	return bodyStr
}
