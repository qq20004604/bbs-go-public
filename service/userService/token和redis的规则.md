## token和redis的规则

### 1. redis

#### 1.1 数据

用户登录后，其token和个人数据即存于redis里

redis里的键值对有两组

1. 第一组的key是token，其值是用户数据
2. 第二组的key是用户ID，其值是token

token的格式是：``bbs-随机生成的唯一KEY``

#### 1.2 唯一性验证

在生成token后，需要去redis里查询是否有相同token，如果有，则需要重新查询，直到没有为止