package utils

import (
	"bytes"
	"github.com/astaxie/beego"
	"net/url"
	"io/ioutil"
	"io"
)

func BuildQueryString(params map[string]string) string {
	reqStr := bytes.NewBufferString("?")

	for key, val := range params {
		reqStr.WriteString(key)
		reqStr.WriteString("=")
		reqStr.WriteString(val)
		reqStr.WriteString("&")
	}

	return reqStr.String()

}

func ExtractResponse(body io.ReadCloser) (map[string]string, error) {
	result, err := ioutil.ReadAll(body)

	if nil != err {
		beego.Error(err)
		return nil, err;
	}
	if nil != err {
		beego.Error(err)
		return nil, err;
	}

	content := string(result)
	values, err := url.ParseQuery(content)
	if nil != err {
		beego.Error(err)
		return nil, err;
	}

	paramMap := make(map[string]string)
	for i, v := range values {
		len := len(v)
		if len > 0 {
			paramMap[i] = v[0]
		}
	}

	return paramMap, nil
}