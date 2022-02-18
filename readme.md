# Gf-Vben-Admin
## 前后端分离后台管理系统
### 本仓库为后端部分

### 前端部分
>https://github.com/vbenjs/gf-vben-admin

#### 后端语言：golang
#### 后端框架：[GoFrame](https://github.com/gogf/gf)
#### 前端语言：Vue3.0
#### 前端框架：[Vben Admin](https://github.com/anncwb/vue-vben-admin)

### 基本组件

1. 鉴权： jwt 
>  https://github.com/jinmao88/gf-jwt 
2. 权限控制： casbin  
>  https://github.com/casbin/casbin
3. 雪花ID： 雪花漂移算法
>  https://github.com/yitter/IdGenerator
4. 后端路由

## Pgsql数据库相关

* 只提供了全局的curd接口 作为demo
* 数据库自己创建


### user表sql语句
```sql
create table user
(
    id        int auto_increment comment 'primary id',
    username  varchar(120)         not null comment 'username',
    password  varchar(64)          null comment 'password',
    note      varchar(255)         null,
    nick_name varchar(120)         null comment 'nickName',
    status    tinyint(1) default 1 null comment '1:enable 2:disable',
    create_at timestamp            null,
    update_at timestamp            null,
    delete_at timestamp            null,
    primary key (id, username)
)
    charset = utf8mb4;


```

### casbin表sql语句

```sql
create table casbin_rule
(
    ptype varchar(10)  null,
    v0    varchar(256) null,
    v1    varchar(256) null,
    v2    varchar(256) null,
    v3    varchar(256) null,
    v4    varchar(256) null,
    v5    varchar(256) null
);


```

### 路由表sql语句
```sql
create table router
(
    id        int auto_increment
        primary key,
    path      varchar(20)  null,
    name      varchar(20)  null,
    redirect  varchar(50)  null,
    title     varchar(30)  null,
    icon      varchar(100) null,
    component varchar(100) null,
    parent    int          null,
    orderNo   int          null,
    status    tinyint(1)   null,
    create_at timestamp    null,
    update_at timestamp    null,
    delete_at timestamp    null
);


```
