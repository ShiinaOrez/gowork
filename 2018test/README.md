## 木犀团队2019年秋季招新 For 2018 测试任务

### 1. 完成对于Go语言基本语法的学习

参考网站：

+ [菜鸟教程](https://www.runoob.com/go/go-tutorial.html)
+ [Go语言中文网](https://studygolang.com/)
+ [Go语言官网](https://golang.org/)
+ [在线Go语言编译平台](https://play.golang.org/)

完成以下几道题目（使用Go语言）：

+ [Leetcode - 最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/)
+ [Leetcode - 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)
+ [Leetcode - 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

完成后以截图的形式验收（包括时间和空间效率的截图）

### 2. 学习gin框架，完成以下接口：

注意，这个环节应该在你电脑的本地实现，因此需要**安装Go语言**

```
[POST] localhost:8080/register

Request Raw Data(JSON):
{
    "username": string,
    "password": string
}

Response:
Status Code: 200 | 注册成功
Status Code: 401 | 用户名已经使用
```

```
[POST] localhost:8080/login

Request Raw Data(JSON):
{
    "username": string,
    "password": string
}

Response:
Status Code: 200 | 登录成功
Status Code: 401 | 密码错误
Status Code: 404 | 没有该用户
```

```
[GET]  localhost:8080/getPassword

Request Parameters:

:username | string

Response:
Status Code: 200 | (JSON): { "password": string }
Status Code: 404 | 没有该用户
```

成果以代码压缩包的形式验收（注意不要把编译生成的二进制文件也打包了），会在以后问你们代码中的具体实现细节，所以请**不要**抄代码。一旦发现直接out.

你的接口应该保证数据的持久化，最起码应该保证**在同一次程序运行的条件下，可以完成数据的存储和访问**(不然你也没法实现登录和获取密码的接口)

##### 先行知识：

+ [请求和响应](https://zhuanlan.zhihu.com/p/68328899)
+ [什么是JSON](https://cloud.tencent.com/developer/article/1454409)
+ [我写的gin教程](https://github.com/ShiinaOrez/Tutor-Go/blob/master/ginny/README.md)
+ [gin的Github仓库](https://github.com/gin-gonic/gin)