package flag

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"go_blog/global"
	"os"
	"strings"
)

func LoadIndex(jsonName string) {
	// 读取文件·
	byteData, err := os.ReadFile(jsonName)
	if err != nil {
		global.Log.Error(err)
		return
	}
	var indexResp EsIndexResponse
	err = json.Unmarshal(byteData, &indexResp)
	if err != nil {
		global.Log.Error(err)
		return
	}
	_list := strings.Split(jsonName, ".")
	if len(_list) != 2 {
		logrus.Error("文件名错误")
		return
	}
	index := _list[0]
	bd, _ := json.Marshal(indexResp.Mapping)
	// 创建索引
	CreateIndex(index, string(bd))
	for _, row := range indexResp.Data {
		_, err := global.ESClient.Index().Index(index).Id(row.ID).BodyJson(row.Row).Do(context.Background())
		if err != nil {
			global.Log.Error(err)
			continue
		}
		logrus.Infof("导入数据 %s 成功", row.ID)
	}
	logrus.Info("导入完成")
}

func CreateIndex(index string, s string) {
	logrus.Info("索引不存在,创建索引")
	_, err := global.ESClient.CreateIndex(index).BodyString(s).Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		global.Log.Error(err)
		return
	}

}

func RemoveIndex(index string) {
	logrus.Info("索引已存在,删除索引")
	// 删除索引
	_, err := global.ESClient.DeleteIndex(index).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		global.Log.Error(err)
		return
	}
}
