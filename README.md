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
## 功能（已实现要求的全部功能）
- 用户注册
- 登录 
- jwt身份验证（只是通过jwt生成token并没有通过请求头来验证）
- 视频发布与获取
- 根据视频内容抽帧自动生成封面（ffmpeg）
- 评论
- 点赞
- 关注
## 说明
- go mod tidy 引入依赖
- go run main.go 运行程序

## 演示

- 接口文档 

  [链接](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-1834514)

  

  ![](https://github.com/Destined777/TikTok/blob/main/public/image-20220613105520766.png)

  注册测试

  ![](https://github.com/Destined777/TikTok/blob/main/public/image-20220613105726392.png)

  

  移动端登录成功

  ![](https://github.com/Destined777/TikTok/blob/main/public/image-20220613115917111.png)

  发布视频成功
  
  ![](https://github.com/Destined777/TikTok/blob/main/public/image-20220613115946828.png)

​				可以在public文件夹下看到刚生成的封面和图片

<video src="https://github.com/Destined777/TikTok/blob/main/public/84d7b315ee8709321dac4f128976db71.mp4"></video>

其它功能演示视频
