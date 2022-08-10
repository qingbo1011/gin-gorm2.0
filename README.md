# gin-gorm2.0
gin+gorm v2的实践。

## Getting Start

在`conf/`下创建`config.ini`：

```ini
[service]
HttpPort = :8080

[mysql]
MysqlHost = 127.0.0.1
MysqlPort = 3308
MysqlUser = root
MysqlPassword = 123456
MysqlDataBase = gorm2
MysqlCharset = utf8mb4
MysqlLogMode = 4
MysqlDefaultStringSize = 256
MysqlDisableDatetimePrecision = true
MysqlDontSupportRenameIndex = true
MysqlDontSupportRenameColumn = true
MysqlSkipInitializeWithVersion = false
MysqlSingularTable = true
MysqlMaxIdleConns = 20
MysqlMaxOpenConns = 100
MysqlConnMaxLifetime = 30
```

> HttpPort的含义就不用多说了，指定gin的端口为8080。
>
> 主要说一下Msyql相关配置：
>
> - `MysqlHost = 127.0.0.1`：MySQL主机Host
> - `MysqlPort = 3308`：MySQL端口
> - `MysqlUser = root`：MySQL用户名
> - `MysqlPassword = 123456`：MySQL密码
> - `MysqlDataBase = gorm2`：MySQL数据库名
> - `MysqlCharset = utf8mb4`：MySQL数据库编码
> - `MysqlLogMode = 4`：Gorm2.0的日志级别（1,2,3,4分别代表`Silent`、`Error`、`Warn`、`Info`）
> - `MysqlDefaultStringSize = 256`：string类型字段的默认长度
> - `MysqlDisableDatetimePrecision = true`：禁用datetime精度，MySQL5.6之前的数据库不支持
> - `MysqlDontSupportRenameIndex = true`： 重命名索引时采用删除并新建的方式，MySQL5.7之前的数据库和MariaDB不支持重命名索引
> - `MysqlDontSupportRenameColumn = true`：用change重命名列，MySQL8之前的数据库和MariaDB不支持重命名列
> - `MysqlSkipInitializeWithVersion = false`：根据当前MySQL版本自动配置
> - `MysqlSingularTable = true`：表名不加s
> - `MysqlMaxIdleConns = 20`： 设置空闲连接池中的最大连接数
> - `MysqlMaxOpenConns = 100`：设置数据库连接最大打开数
> - `MysqlConnMaxLifetime = 30`：设置可重用连接的最长时间

接下来运行`main.go`即可。