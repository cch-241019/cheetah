##### 复制表

下面语句复制表结构。  
方法一：

```mysql
CREATE TABLE 新表名 LIKE 源数据库.源表名;
```

方法二：

```mysql
CREATE TABLE 新表名 AS
SELECT *
FROM 源数据库.源表名;
```

复制表结构，同时复制数据。

```mysql
CREATE TABLE 新表名 AS
SELECT *
FROM 源数据库.源表名
WHERE 1 = 0;
```

加`WHERE 1=0`这个永久假的条件可以防止`SELECT`查询出数据，从而防止复制数据。