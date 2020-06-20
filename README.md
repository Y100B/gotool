## gin网站开发模板项目

### 1.修改config.env。

填写MySQL账号以及密码，数据库名字填写后程序会自动创建相应的数据库，如果数据库存在则跳过创建。


### 2.用户注册。

示例：
```
api:http://localhost:8005/v1/api/user/register
method:post
type:json

Request-data:

{
    "NickName": "Golang",
    "Email": "1144620122@qq.com",
    "PassWord": "PassWordPassWord",
    "UserAnswer": "https://api.syrme.top/v1/api/file/image/upload/headers/107.png"
}

Response-data:

{
    "code": 200,
    "data": {
        "ID": 1,
        "CreatedAt": "2020-06-20T23:10:39.0466167+08:00",
        "UpdatedAt": "2020-06-20T23:10:39.0466167+08:00",
        "DeletedAt": null,
        "NickName": "Golang",
        "Email": "1144620122@qq.com",
        "UserAnswer": "https://api.syrme.top/v1/api/file/image/upload/headers/107.png"
    },
    "msg": "注册成功！"
}

```

### 3.用户登录
示例：
```
api:http://localhost:8005/v1/api/user/register
method:post
type:json

Request-data:

{
    "Email": "1144620122@qq.com",
    "PassWord": "PassWordPassWord"
}

Response-data:

{
    "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7IklEIjoxLCJDcmVhdGVkQXQiOiIyMDIwLTA2LTIwVDIzOjEwOjM5KzA4OjAwIiwiVXBkYXRlZEF0IjoiMjAyMC0wNi0yMFQyMzoxMDozOSswODowMCIsIkRlbGV0ZWRBdCI6bnVsbCwiTmlja05hbWUiOiJHb2xhbmciLCJFbWFpbCI6IjExNDQ2MjAxMjJAcXEuY29tIiwiVXNlckFuc3dlciI6Imh0dHBzOi8vYXBpLnN5cm1lLnRvcC92MS9hcGkvZmlsZS9pbWFnZS91cGxvYWQvaGVhZGVycy8xMDcucG5nIn0sImV4cCI6MTU5Mjc1Mjk4NSwiaXNzIjoiQWxmcmVkbyBNZW5kb3phIn0.fJpBWZBb8d1Dso5WIl56cWHGBjaNFDZ-Q2cpVdYCZok",
    "code": 200,
    "msg": "登录成功！"
}
```

### 4.获取用户信息

```
api:http://localhost:8005/v1/api/user/data?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7IklEIjoxLCJDcmVhdGVkQXQiOiIyMDIwLTA2LTIwVDIzOjEwOjM5KzA4OjAwIiwiVXBkYXRlZEF0IjoiMjAyMC0wNi0yMFQyMzoxMDozOSswODowMCIsIkRlbGV0ZWRBdCI6bnVsbCwiTmlja05hbWUiOiJHb2xhbmciLCJFbWFpbCI6IjExNDQ2MjAxMjJAcXEuY29tIiwiVXNlckFuc3dlciI6Imh0dHBzOi8vYXBpLnN5cm1lLnRvcC92MS9hcGkvZmlsZS9pbWFnZS91cGxvYWQvaGVhZGVycy8xMDcucG5nIn0sImV4cCI6MTU5Mjc1Mjk4NSwiaXNzIjoiQWxmcmVkbyBNZW5kb3phIn0.fJpBWZBb8d1Dso5WIl56cWHGBjaNFDZ-Q2cpVdYCZok
method:get or post

Response-data:
{
    "code": 200,
    "data": {
        "ID": 1,
        "CreatedAt": "2020-06-20T23:10:39+08:00",
        "UpdatedAt": "2020-06-20T23:10:39+08:00",
        "DeletedAt": null,
        "NickName": "Golang",
        "Email": "1144620122@qq.com",
        "UserAnswer": "https://api.syrme.top/v1/api/file/image/upload/headers/107.png"
    },
    "msg": "获取用户信息成功！"
}
```
