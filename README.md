# Go 微信机器人

⚠️本项目只是作者的一个玩具，大部分功能只为个人定制，并不通用。

## 功能介绍
微信+百度AI+ChatGPT 实现你的虚拟情人

## 部署说明

clone 项目到本地，然后进入项目目录，将 `config/dev.yaml` 文件改成 `config/prod.yaml`， yaml 配置文件需要配置下，可以去对应的网站获取 apiKey。

执行如下命令：

```shell
go mod tidy # 下载依赖

go build -v -o wxbot  # 编译

nohup ./wxbot > core.log & # 后台运行, 可以查看日志 core.log
```

`less core.log` 可以查看日志，日志里有二维码，可以扫码登录。

## 功能列表

基于 [openwechat](https://github.com/eatmoreapple/openwechat)，[百度AI](https://cloud.baidu.com/doc/SPEECH/s/mlciskuqn)，[ChatGPT](https://github.com/kkdai/chatgpt) 开发，感谢作者。


