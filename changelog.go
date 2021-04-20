package api

import (
	"fmt"
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/models"
	"github.com/spf13/viper"
)

type Upgrade struct {
	Date  string
	Version string
	Tags []string
	Message []string
}

type config struct {
	Upgrade []Upgrade
}

func GetSignedInChangelog(c *models.ReqContext) response.Response {

	v := viper.New()
	v.SetConfigFile("changelog.yml")
	v.ReadInConfig()
	cfg := &config{}
	v.Unmarshal(cfg)
	fmt.Println(cfg.Upgrade)
	return response.JSON(200, cfg.Upgrade)
}
