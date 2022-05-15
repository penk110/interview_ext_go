

##### 参考
```text
https://casbin.org/docs/zh-CN/overview
```

##### 模型定义
request_definition
```text
[request_definition]
r = sub, obj, act

分别是：
    访问实体 (Subject)
    访问资源 (Object)
    访问方法 (Action)
```

policy_effect
```text
[policy_effect]
e = some(where (p.eft == allow))


```

1.举例
```text
GET /users

添加GET权限：
    tester,/users,GET
```