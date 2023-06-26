# Zinx-V0.3

新增Request模块和Router模块

## Request 模块
Request模块是对请求连接进行封装，包括连接对象，传输的数据等及其以后需要的参数都可以统一封装到Request中
继承BaseRouter 即可实现Router对象（maybe BaseRouter 叫 BaseHandler 更好）

## Router 模块
> 现在本质就是Handler

Router 模块接收Request对象，主要为服务端处理用户连接时的一些前置和后置操作。

