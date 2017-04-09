package common

import (
	"github.com/ErosZy/singoriensis/common"

	"github.com/axgle/mahonia"
	"golang.org/x/net/html/charset"
)

func CharsetIconv(page *common.Page) string {
	bodyStr := page.ResBody
	_, charset, _ := charset.DetermineEncoding([]byte(bodyStr), "text/html")

	if charset == "gbk" {
		decoder := mahonia.NewDecoder(charset)
		bodyStr = decoder.ConvertString(bodyStr)
	}

	return bodyStr
}
