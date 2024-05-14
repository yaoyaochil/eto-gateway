# ETO·网关服务

[![License: Apache2.0](https://img.shields.io/badge/License-Apache2.0-red.svg)](https://opensource.org/licenses/Apache)
[![Go](https://img.shields.io/badge/Go-1.22.2-blue)](https://golang.dev/)


**注意：** 本项目需要搭配 [eto-gateway-admin](https://github.com/yaoyaochil/eto-gateway-admin) 前端项目使用

**如果：** 使用docker-compose无需单独运行

## 项目介绍
ETO·网关服务是一个基于gin框架的网关服务，主要用于对接游戏服务，提供统一的接口，方便对接多个游戏服务。

## 项目特点
- 支持邮件单独发送 无需配置物品占用
- 邮件直接发送金币
- Frida动态代码注入
- 快速对接游戏服务
- 超快速的响应速度,对比php的网关服务，响应速度提升了10倍
- 支持多平台快速部署
- 占用资源少，运行稳定,支持高并发

## 项目截图
![img.png](https://github.com/yaoyaochil/eto-gateway/assets/49603204/9f024abf-2e11-496e-a2f8-ba59d42ff8aa)
<img width="1919" style="border-radius:8px;margin-bottom:10px" alt="image" src="https://github.com/yaoyaochil/eto-gateway-admin/assets/49603204/26875c19-aa15-42bf-82aa-4e98c6b4f83b">
<img width="1919" style="border-radius:8px;margin-bottom:10px" alt="image" src="https://github.com/yaoyaochil/eto-gateway-admin/assets/49603204/8cd61bad-12fe-4698-a18e-5e0c11a337de">


```shell
# 配置文件 将eto.config.yaml 重命名为config.yaml 并修改配置文件为你的配置
mv eto.config.yaml config.yaml


# mod 包管理
go mod tidy

# 编译
go build -o eto-gateway

# 运行
./eto-gateway

# 如果你的服务正常运行，你可以通过curl访问登陆接口测试 post 请求 body 为json格式 {"accountname":"","password":""}
curl -X POST http://localhost:8080/login -d '{"accountname":"admin","password":"admin"}' -H "Content-Type: application/json"

# 获取到的数据中 data 为游戏token base64编码后的数据
```

##### 成功返回结果
```json
{
    "code": 0,
    "data": "MMAjO5ubD7LawXw/....",
    "msg": "登录成功"
}
```

```shell
#注意 新版本登陆游戏和注册游戏的前置路由增加/base/ 例如登陆游戏的路由为 /d_taiwan/dnfLogin  新版本为 /base/d_taiwan/dnfLogin

# 完整的登陆游戏的路由为 /base/d_taiwan/dnfLogin
# 完整的注册游戏的路由为 /base/d_taiwan/Register

# 传递参数不变
```

##### ![img.png](https://github.com/yaoyaochil/eto-gateway/assets/49603204/b97408e8-8b63-4db7-b47b-343ce29b2ece)



## 交流群
```shell
QQ群： 575053531
```


## 如果你觉得这个项目对你有帮助，你可以请我喝杯咖啡：
<p float="left">
    <img src="resource/wechat_pay.JPG" width="200" style="margin-right: 50px;" />
    <img src="resource/aili_pay.JPG" width="200" />
</p>