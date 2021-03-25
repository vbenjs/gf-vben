# Gf-Vben-Admin
## 前后端分离后台管理系统
### 本仓库为后端部分

#### 后端语言：golang，后端框架：GoFrame
#### 前端框架：Vben

### 基本组件

1. 鉴权： jwt
2. 权限控制：casbin //todo


> 只提供了全局的curd接口 作为demo
> 数据库自己创建

### user表sql语句
```sql
create table app_user
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