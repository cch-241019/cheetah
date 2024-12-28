##### 文件权限

```shell
-rw-r--r--. 1 root root 1969 Dec 26 10:48 /etc/profile
```

```text
文件类型-文件所有者权限-用户组权限-其他用户权限
```

**修改文件权限**

```shell
chmod u+rwx,g+rw,o=r file.txt
```

##### 文件所属

文件“所属”包含:

1. 所有者(Owner)
2. 所属组(Group)

**修改文件所属**

```shell
sudo chown user2:group2 data.txt
```

##### 文件夹下的内容查询

```shell
find . -type f -exec grep "example" {} +
```

##### GNU findutils命令工具

[GNU Findutils](https://www.gnu.org/software/findutils/manual/html_mono/find.html)

