package db

import (
	"context"
	"fmt"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/olivere/elastic/v7"
)

var EDB *elastic.Client

func init() {
	var host = fmt.Sprintf("%s:%d", config.SettingsConfig.Es.Host, config.SettingsConfig.Es.Port)
	// 创建一个 Elasticsearch 客户端
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(config.SettingsConfig.Es.Username, config.SettingsConfig.Es.Password))

	if err != nil {
		klog.Error("Failed to connect to Elasticsearch: %v", err)
	}
	EDB = client
	// Ping Elasticsearch 服务器是否正常
	info, code, err := EDB.Ping(host).Do(context.Background())
	if err != nil {
		klog.Error("Failed to ping Elasticsearch: %v", err)
	}
	klog.Info("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

}
