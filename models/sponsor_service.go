package models
import (
	"net/http"
	"bytes"
	_"net/url"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"
)


func Pay(money float32) interface{} {
	data := make(map[string]interface{})
	data["order_price"] = money
	data["product_name"] = "充值"
	data["body"] = "赞助beegostudy" + strconv.FormatFloat(float64(money), 'f', -1, 32)
	var bmobApplicationId, bmobRestApiKey interface{}

	bmobApplicationId, _ = ParameterValue("x-bmob-application-id")
	bmobRestApiKey, _ = ParameterValue("x-bmob-rest-api-key")

	b, err := json.Marshal(data)

	if err != nil {
		fmt.Println("json err:", err)
	}
	body := bytes.NewBuffer([]byte(b))
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.bmob.cn/1/webpay", body)
	req.Header.Add("X-Bmob-Application-Id", bmobApplicationId.(string))
	req.Header.Add("X-Bmob-REST-API-Key", bmobRestApiKey.(string))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	dataBody, _ := ioutil.ReadAll(resp.Body)
	var jsonString = string(dataBody)
	data = make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &data)
	return data["html"]
}
