# bank
银行


### api文档

- `xxxxx`
- [登录API](#登录API)
- [注册API](#注册API)
- [存钱API](#存钱API)
- [取钱API](#取钱API)
- [转账API](#转账API)


#### 登录API

接口  POST `v1/loginBank`

REQUEST：
``` 
{
    "ID": "" ,  //字符串 账号ID
    "Password":"" //字符串 账户密码
}
```

RESPONSE:
``` 
{
    {
    "code": 0,  //状态码
    "message": "", 
    "data": {   //数据
        "name": "李四", //用户名
        "id": "9", //账号ID
        "password": "111", //账户密码
        "balance": 439 // 账户余额
    }
}
}
```

备注:

xxxx

#### 注册API

接口 POST `v1/createUser`

REQUEST:
```
{
    "username":"", //字符串 用户名
    "password": "8888"  //字符串 密码
}
```

RESPONSE
```
{
    "code": 0, //状态码
    "message": "",
    "data": 29 //账号ID
}
```

备注:

xxx

#### 存钱API

接口 PUT `v1/saveMoney`

REQUEST:
```
{
    "user_id":"29", //字符串 账户id
    "money": 500 //浮点型 存入金额
}
```

RESPONSE:
```
{
    "code": 0, //状态码
    "message": "",
    "data": "ok" //操作成功
}
```

备注:

xxx

#### 取钱API

接口 PUT `v1/withdrawMoney`
REQUEST:
```
{
    "user_id":"29", //字符串 账户id
    "money": 500 //浮点型 取款金额
}
```

RESPONSE:
```
{
    "code": 0, //状态码
    "message": "",
    "data": "ok" //操作成功
}
```



#### 转账API

接口 PUT `v1/transfer`

REQUEST:
```
{
    "out_user_id":"29", //字符串 转出账户ID
    "enter_user_id":"27", //字符串 转入账户ID
    "money": 500    // 浮点型 转账金额
}
```

RESPONSE
```
{
    "code": 0,  //状态吗
    "message": "", 
    "data": "ok" //操作成功
}
```

备注:

xxxx