package biz

import (
	"encoding/json"
	"net/http"
)

type QueryES struct{}

func (qes *QueryES) Query(resp http.ResponseWriter, req *http.Request) {
	data := qes.getData()
	jsonRet, _ := json.Marshal(data)
	_, _ = resp.Write(jsonRet)
}

func (qes *QueryES) getData() *ESData {
	//TODO 从ElasticSearch获取线上数据
	//Mock
	return &ESData{
		Dau:        10333,
		Qps:        34556,
		QueryDelay: 23,
	}
}

type ESData struct {
	Qps        int32
	Dau        int32
	QueryDelay int32
}
