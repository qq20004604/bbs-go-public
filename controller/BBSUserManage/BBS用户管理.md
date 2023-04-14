### 1. 接口说明

接口：

1. 具体查看 ``router/index.go``
2. 登录；
3. 在线状态检测；
4. 注册新用户；
5. 登出；
6. 获取用户信息；

### 2. 登录状态

包括：

1. 登录状态：同时提供 Cookie 和 token 两种方式，默认是以 cookie（httponly） 方式存储和判断；
2. 检测登录：``IsOnline``，默认是以 cookie 形式存储，如果需要单纯以 token 为判断，重写 ``InOnline.go`` 里面的逻辑即可；
3. 登录状态在 redis 里进行缓存，默认是 168 个小时（参照yml里的LoginExpireTime）

### 3. 注册账号

1. 默认新账号注册后即是正常账号；
2. 通过 yml 配置的 RegistrationRequiresAdminApproval 项，设置为 true 后，新用户注册后为待审核状态，只有管理员审批后才能通过。

### 4. 账号管理

1. 管理员可以看到所有账号状态；
2. 可以批量改变账号的状态；