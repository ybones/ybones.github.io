## 常用SHOW
```SQL
-- 查看全局变量信息
SHOW GLOBAL VARIABLES;

-- 查看会话变量信息
show SESSION VARIABLES;

-- 修改变量信息
# 全局变量修改
SET GLOBAL sort_buffer_size = 10 * 1024 * 1024; 
# 会话变量修改
SET SESSION sort_buffer_size = 10 * 1024 * 1024;
# 会话变量修改
SET sort_buffer_size = 10 * 1024 * 1024;

-- 列出server状态信息
SHOW STATUS;

-- 返回所有数据库列表
SHOW DATABASES;

-- 使用某一数据库
USE mysql;

-- 获得一个数据库内表的列表
SHOW TABLES;

-- 获得表的结构信息
SHOW COLUMNS FROM tbl_name;

-- 获得数据库信息
SHOW TABLES [FROM db_name];

-- 获取表索引信息
SHOW INDEX FROM tbl_name [FROM db_name];

-- 显示错误或警告信息
SHOW ERRORS;
SHOW WARNINGS;
```

```SQL
SELECT
{* | <字段列名>}
FROM <表 1>, <表 2>…
[WHERE <表达式>]
[GROUP BY <group by definition>]
[HAVING <expression> [{<operator> <expression>}…]]
[ORDER BY <order by definition> [DESC]]
[LIMIT[<offset>,] <row count>]

-- 去重
SELECT DISTINCT 列名称 FROM 表名称

-- 匹配
SELECT 列名称 FROM 表名称 WHERE 列名称 REGEXP "definition"
SELECT 列名称 FROM 表名称 WHERE 列名称 LIKE "definition"
```

[MySQL 数据类型](https://www.runoob.com/mysql/mysql-data-types.html)

```SQL
INSERT INTO table_name ( field1, field2,...fieldN )
                       VALUES
                       ( value1, value2,...valueN );
```

```SQL
UPDATE table_name 
SET field1=new-value1, field2=new-value2
[WHERE Clause]
```

```SQL
DELETE FROM table_name 
[WHERE Clause]
```

```SQL
CREATE TABLE IF NOT EXISTS `runoob_tbl`(
   `runoob_id` INT UNSIGNED AUTO_INCREMENT,
   `runoob_title` VARCHAR(100) NOT NULL,
   `runoob_author` VARCHAR(40) NOT NULL,
   `submission_date` DATE,
   PRIMARY KEY ( `runoob_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```