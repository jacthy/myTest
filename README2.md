## 系统介绍
go backend test demo from liaojuntao，一个简易的用户管理系统demo
### 系统简析
![image](./系统架构图.png)
- 系统主要分三层：接口层，业务层，基础服务层
    - 接口层主要是注册路由，规范restful服务接口，以及拦截器（后期可添加限流，熔断，鉴权等高级特性）
    - 业务层主要是处理业务逻辑，然后调用基础服务层（其实就是dto层）实现持久化存储
    - 基础服务层，这里抽象了一层repo，为的是将数据库解耦，便于插拔式剪裁组件。这里用作demo，简易为主，用了sqlite做存储，后期可以改用mysql，redis等实现存储
### tips for developer
1. 接口层是一个基于go的基础包http实现的一个简易web框架，只添加了两个组件router和interceptor,分别实现路由配置，以及拦截器功能
2. 