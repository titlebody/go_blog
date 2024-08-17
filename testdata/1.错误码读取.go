package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go_blog/model/res"
	"os"
)

const file = "model/res/err_code.json"

type ErrMap map[res.ErrCode]string

func main() {
	byteData, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap ErrMap
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap)
	fmt.Println(errMap[res.SettingsError])
}
