package common

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

func Extend(o1 interface{}, o2 interface{}) interface{} {
	ov1 := reflect.ValueOf(o1)
	ov2 := reflect.ValueOf(o2)

	if ov1.Elem().NumField() != ov2.Elem().NumField() {
		panic("o1 and o2 need to be the same struct!")
	}

	fieldNum := ov1.Elem().NumField()

	for i := 0; i < fieldNum; i++ {
		of1 := ov1.Elem().Field(i)
		of2 := ov2.Elem().Field(i)
		kind := of2.Kind()

		if of1.CanSet() {
			switch kind {
			case reflect.Int:
				of1.SetInt(of2.Int())
			case reflect.String:
				of1.SetString(of2.String())
			case reflect.Bool:
				of1.SetBool(of2.Bool())
			}
		}
	}

	return ov1.Elem().Interface()
}

func Object2JsonStr(obj interface{}) string {
	data, err := json.Marshal(obj)

	if err != nil {
		data = nil
		Logger(LOG_WARNING, err.Error())
	}

	return string(data)
}

func ToReaderCloser(reader io.Reader) io.ReadCloser {
	rc, ok := reader.(io.ReadCloser)
	if !ok && reader != nil {
		rc = ioutil.NopCloser(reader)
	}

	return rc
}

func UrlEncoded(str string) (string, error) {
	filterRegex := regexp.MustCompile("%00|\n|\r\n")

	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return filterRegex.ReplaceAllString(u.String(), ""), nil
}

func UnicodeConvert(str string) string {
	buf := bytes.NewBuffer(nil)
	i, j := 0, len(str)
	for i < j {
		x := i + 6
		if x > j {
			buf.WriteString(str[i:])
			break
		}
		if str[i] == '\\' && str[i+1] == 'u' {
			hex := str[i+2 : x]
			r, err := strconv.ParseUint(hex, 16, 64)
			if err == nil {
				buf.WriteRune(rune(r))
			} else {
				buf.WriteString(str[i:x])
			}
			i = x
		} else {
			buf.WriteByte(str[i])
			i++
		}
	}

	return buf.String()
}

func RegStrQuote(regStr string) string {
	r := regexp.MustCompile(`\\"|\\.`)
	return r.ReplaceAllStringFunc(regStr, func(str string) string {
		if str == `\"` {
			return `\\` + str
		}
		return `\` + str
	})
}

func RegStrUnquote(regStr string) string {
	r := regexp.MustCompile(`\\`)
	return r.ReplaceAllString(regStr, `\`)
}
