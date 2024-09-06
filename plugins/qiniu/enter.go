package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go_blog/config"
	"go_blog/global"
	"time"
)

// 获取token
func getToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// 获取上传配置
func getCfg(q config.QiNiu) storage.Config {
	cfg := storage.Config{}
	//  空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 是否使用CDN上传加速
	cfg.UseCdnDomains = false
	return cfg
}

// 上传图片 文件数组 前缀
func UploadImages(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.Config.QiNiu.IsEnable {
		return "", errors.New("七牛未开启")
	}
	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("七牛未配置AccessKey SecretKey")
	}
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("图片大小超过限制")
	}
	upToken := getToken(q)
	cfg := getCfg(q)
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))
	// 获取当前时间
	now := time.Now().Format("20060102150406")
	key := fmt.Sprintf("%s/%s__%s", prefix, now, imageName)

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil
}
