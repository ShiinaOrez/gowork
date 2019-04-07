# 文档自动生成

### 1.环境准备

首先获取bee 和 beego两个包
```
go get github.com/astaxie/beego
go get github.com/beego/bee
```

beego的文档自动生成是通过扫描工程下面的注释来实现的。

### 2.生成文档
进入autoAPI 文件夹之后，运行
```
bee run -downdoc=true -gendoc=true
```
即可自动生成文档。
此时访问 `http://127.0.0.1:8080/swagger/index.html` 即可获取文档。

在运行 `bee run -downdoc=true -gendoc=true` 状态时，如果修改了文件，beego会自动重新生成文档并重新启动。

### 3.修改文档

所有文档的注释源文件都在 `/autoAPI/controllers/` 文件夹下的对应文件中。
每个API都为一个函数，相应的API文档中的字段都写在了各个函数的注释中。
在注释中可以指定 POST过来的相应的Object 和 返回的Object等，为了方便，所有注释中用到的Object都放在了`/autoAPI/models/user.go`中。

修改文档后，注意重新生成查看无误后再推。

### 4.格式

**GET方法**

```
// @Title CompanyInformationByID
// @Description Get company information by id. (information less)
// @Param   Token       header    string      true        "The token to conform"
// @Param   Cid         path      int         true        "The company ID for get"
// @Success 200 {object} models.CompanyInformation
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | company not existed
// @router /:Cid/info [get]
func (u *CompanyController) CompanyInfo() {
}
```
这个是一个需要 `Token` Header 和 `Cid` Path参数的Get方法。
在成功时返回 `models.CompanyInformation` 
在失败时返回401 和 404等

**PUT方法**

```
// @Title EditCompanyInformation
// @Description Edit company information
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.CompanyInfoModify   true   "edit company info body"
// @Success 200 {string} edit ok
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @Failure 406 {string} image error
// @router /info [put]
func (u *CompanyController) EditCompanyInfo() {
}
```
这是一个Header中有Token, 请求体为 `models.CompanyInfoModify` 的方法，注意 `@Param   Payload ` 这一行的行尾的true 和 描述字符串不可掉。



### 5. 格式修改
因为前期的一些疏漏，导致API格式有很多不对的地方。 被分配到各个部分的同学要负责相应API文档的修改。

格式要求：

* POST PUT 等携带 json requests body 的API，都要变成 json的格式（类似上文PUT方法中的`models.CompanyInfoModify` ），现有文档很多是 string 类型。  
* 所有函数的返回体中携带response body的，都要改成json格式。因为有些返回的body类型误写成了{string},正确应为 {object}。
* 注意状态码的修改。现有API很多状态码不正确，如很多不存在使用的为402状态码。

