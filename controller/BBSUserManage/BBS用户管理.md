## 说明

包括：

1. 登录：``BBSUserLogin.go``
2. 检测登录：``IsOnline``，默认是以 cookie 形式存储（httponly），如果需要单纯以 token 为判断，重写 ``InOnline.go`` 里面的逻辑即可；
3. 