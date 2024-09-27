package flag

import (
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"go_blog/global"
	"golang.org/x/net/context"
	"os"
)

type EsRawMessage struct {
	Row json.RawMessage `json:"row"`
	ID  string          `json:"id"`
}

type EsIndexResponse struct {
	Data    []EsRawMessage `json:"data"`
	Mapping interface{}    `json:"mapping"`
}

// DumpIndex 获取索引信息
// go run main.go -es -dump article_index
// go run main.go -es -dump full_text_index
func DumpIndex(index string) {
	result, err := global.ESClient.
		Search(index).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	mapping, err := global.ESClient.GetMapping().Index(index).Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	var jsonList []EsRawMessage

	for _, hit := range result.Hits.Hits {
		var jsonData EsRawMessage
		jsonData.Row = hit.Source
		jsonData.ID = hit.Id
		jsonList = append(jsonList, jsonData)
	}
	if len(jsonList) == 0 {
		logrus.Info("没有数据")
		return
	}
	indexMapping, ok := mapping[index]
	if !ok {
		logrus.Error("没有索引")
		return
	}
	esIndexREsponse := EsIndexResponse{
		Data:    jsonList,
		Mapping: indexMapping,
	}

	file, err := os.Create("./data/" + index + ".json")
	defer file.Close()
	if err != nil {
		logrus.Error(err)
		return
	}
	byteData, _ := json.Marshal(esIndexREsponse)
	file.Write(byteData)
	logrus.Info("导出成功")
}
