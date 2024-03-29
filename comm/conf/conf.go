package conf

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Conf .
type Conf struct {
	App       App       `json:"app" yaml:"app"`
	Keys      Keys      `json:"keys" yaml:"keys"`
	RedisConf RedisConf `json:"redis" yaml:"redis"`
	ChatGPT   ChatGPT   `json:"chatgpt" yaml:"chatgpt"`
}

type RedisConf struct {
	IP     string `json:"ip" yaml:"ip"`
	Port   string `json:"port" yaml:"port"`
	Passwd string `json:"passwd" yaml:"passwd"`
}

// App .
type App struct {
	Env string `json:"env" yaml:"env"`
}

type Keys struct {
	ChristmasHatURL    string `json:"christmas_hat_url" yaml:"christmas_hat_url"`
	BotName            string `json:"bot_name" yaml:"bot_name"`
	WeatherKey         string `json:"weather_key" yaml:"weather_key"`
	TianapiKey         string `json:"tianapi_key" yaml:"tianapi_key"`
	TianapiKey1        string `json:"tianapi_key1" yaml:"tianapi_key1"`
	HoneyLove          string `json:"honey_love" yaml:"honey_love"`
	LoverChName        string `json:"lover_ch_name" yaml:"lover_ch_name"`
	MasterAccount      string `json:"master_account" yaml:"master_account"`
	HouchangcunFans    string `json:"houchangcun_fans" yaml:"houchangcun_fans"`
	BanzhuanGroup      string `json:"banzhuan_group" yaml:"banzhuan_group"`
	BubeiGroup         string `json:"bubei_group" yaml:"bubei_group"`
	QweatherKey        string `json:"qweather_key" yaml:"qweather_key"`
	BubeiStartDate     string `json:"bubei_start_date" yaml:"bubei_start_date"`
	WuZhuangShiMembers string `json:"wu_zhuang_shi_members" yaml:"wu_zhuang_shi_members"`
	RemindMsg          string `json:"remind_msg" yaml:"remind_msg"`
}

type ChatGPT struct {
	ApiKey string `json:"api_key" yaml:"api_key"`
}

// GetConf .
func GetConf(cfg string) (conf *Conf, err error) {
	var (
		yamlFile = make([]byte, 0)
	)

	filepath := fmt.Sprintf("%s", cfg)
	logrus.Infof("filepath: %s", filepath)
	yamlFile, err = ioutil.ReadFile(filepath)
	if err != nil {
		err = errors.Wrapf(err, "ReadFile error")
		logrus.Errorf(err.Error())
		return conf, err
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		err = errors.Wrapf(err, "yaml.Unmarshal error")
		logrus.Errorf(err.Error())
		return conf, err
	}

	return conf, nil
}
