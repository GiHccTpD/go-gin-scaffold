[TOC]



## 介绍

使用Go Module作为依赖管理，基于Gin搭建Go的Web服务器，使用Endless来使服务器平滑重启，使用Swagger来自动生成Api文档。



## 安装

```bash
export GO111MODULE=on
export GOPROXY=https://goproxy.io,direct
go get
go run main.go
```



## 注意

- go的版本需要高于1.11
- 使用goland时，需要确保Go module是Enable状态



## 环境区分

复制 `config` 目录下的 `.env.example` 文件，修改为 `.env.development`， 并将内部的配置项改为开发的配置参数



## Swagger 配置



### 使用

- 查看[Swagger API文档](http://localhost:10080/swagger/index.html)

- 在 `controllers`文件夹中每个 `endpoint` 前面添加上述注释用于生产 `Swagger API` , 如添加在 `controllers/login/login.go` 中 `Login` 前的注释

  ```go
  
  // @Summary 登录
  // @Description 登录
  // @Accept  json
  // @Produce  json
  // @Param auth body Auth true "校验"
  // @Success 200 {string} json "{"code":200,"data":{"token": "token"},"msg":"ok"}"
  // @Failure 500 {string} json "{"code":500,"data":null,"msg":"TOKEN校验失败"}"
  // @Failure 500 {string} json "{"code":500,"data":null,"msg":"TOKEN校验超时"}"
  // @Failure 500 {string} json "{"code":500,"data":null,"msg":"生成TOKEN失败"}"
  // @Failure 500 {string} json "{"code":500,"data":null,"msg":"错误的TOKEN"}"
  // @Router /login [post]
  func Login(c *gin.Context) {
  	appG := util.Gin{C: c}
  }
  
  ```

  

> ### 示意
>
> - @Summary 是对该接口的一个描述
>
> - @Id 是一个全局标识符，所有的接口文档中 Id 不能标注
>
> - @Tags 是对接口的标注，同一个 tag 为一组，这样方便我们整理接口
>
> - @Version 表明该接口的版本
>
> - @Accept 表示该该请求的请求类型
>
> - @Param 表示参数 分别有以下参数 参数名词 参数类型 数据类型 是否必须 注释 属性(可选参数),参数之间用空格隔开。
>
> - @Success 表示请求成功后返回，它有以下参数 请求返回状态码，参数类型，数据类型，注释
>
> - @Failure 请求失败后返回，参数同上
>
> - @Router 该函数定义了请求路由并且包含路由的请求方式。

- 在 `endpoint` 添加上面的备注，在执行 `go run main.go` 之前执行 `swag init`

* 查看 [http://localhost:10080/swagger/index.html](http://localhost:10080/swagger/index.html) 新的API文档已经存在

### 参考

1.[Swagger API Operation](https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html#api-operation)

2.[Examples](https://github.com/swaggo/swag/blob/master/example/celler/controller/accounts.go)

