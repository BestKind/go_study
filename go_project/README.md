## 问题
按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

## 回答
1. 大体框架
- internal 使用 biz、data、service 等目录，携带 myapp 应用名（单项目可以省去这层）
- pkg 为项目里共用
- cmd/myapp，cmd 下需要带上 app 名字
- api/service/v1，按照版本号
- configs 放配置文件
2. 对象初始化，biz、data、service，依赖的对象作为参数传入，在 main 里使用 wire 构建和消费资源
3. biz 中定义 repository 的接口，并对 DomainObject 定义，实现在 data 目录中
4. main.go 中使用 wire 进行对象后，对 lifecycle 进行服务的注册和启动
