# Unite Server [Unite 服务端]

> 采用全新的微服务架构

```
-- user
---- ironic
------- pb
---- controller
---- subscriber
```

## 关于新的服务架构，最优开发实践的构想

* 自动化生成飞业务代码，Controller, Entity, Service
* Controller 内建 Restful

### 最终的项目结构类似与这样

- SimpleSRV
-- handler
-- service
-- 