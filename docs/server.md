# 组
## 创建组
    groupadd nt
## 删除组
    groupdel nt
## 查看所有组
    cat /etc/group
# 用户
## 增加用户
    增加用户yu；使用目录/home/work；增加nt组
    useradd -d /home/work -g nt yu
## 删除用户
    userdel yu
## 设置密码
    passwd yu
## 切换用户
    su yu
# 文件权限
## 修改权限
    chmod 755 file
## 常用数字权限
    常用的linux文件权限：
    444 r--r--r--
    600 rw-------
    644 rw-r--r--
    666 rw-rw-rw-
    700 rwx------
    744 rwxr--r--
    755 rwxr-xr-x
    777 rwxrwxrwx
    若要 rwx 属性则 4+2+1=7；
    若要 rw- 属性则 4+2=6；
    若要 r-x 属性则 4+1=5。
### 例子755
    7:  4+2+1   rwx     所有者具有读取、写入、执行权限；
    5:  4+0+1   r-x     同组用户具有读取、执行权限但没有写入权限；
    5:  4+0+1   r-x     其他用户具有读取、执行权限但没有写入权限；
# 安装软件
## 安装python3
## 安装golang
## 安装mysql
## 安装redis
## 安装mongo
## 安装nginx
## 安装docker
