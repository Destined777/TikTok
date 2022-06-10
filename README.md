# 抖声项目介绍文档
## 配置文件
 config.json文件
 - DbSettings MySQL配置 包括Username, Password, Hostname, Dbname
## 域名
抖声项目的域名，存在consts.go里，需要根据运行环境进行修改。
## 大体框架
- 从上至下分为controller层，service层，dao层。
- 数据库的表模型存在model层，而dao层是对model层的操作。
- service层封装了单个接口的业务实现，调用dao层的函数。
- controller层是接收处理移动端表单输入，调用service层的函数，并处理错误，返回响应。
- public文件夹存静态资源，在本项目中即为用户上传的视频。
- util层存其它各层由于业务需求调用的工具类函数。
- global与config的作用是处理配置文件，初始化环境配置，方便其它层对数据库的操作。
- http_param文件下存接收表单输入的结构体
## 功能
- 用户注册
- 登录 
- jwt身份验证（只是通过jwt生成token并没有通过请求头来验证）
- 视频发布与获取
- 评论
- 点赞
- 关注
## 说明
- go mod tidy 引入依赖
- go run main.go 运行程序