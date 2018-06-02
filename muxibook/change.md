### Muxibook Version 2.0 Is Coming!

#### Repair

+ url:"/book" 将bookname & booknum 一起作为重复添加的标准
+ url:"/book" 恢复token的作用
+ url:"/book" 重置了StatusCode的标准
+ url:"/lendbook" 添加了对book.available的检查
+ url:"/lendbook" 重置了StatusCode的标准
+ url:"/returnbook" 重置了StatusCode的标准
+ url:"/returnbook" 修复了无限还书bug
+ url:"/renewbook" 改变了对还书时间的赋值方式
+ url:"/renewbook" 重置了StatusCode的标准


### 关于本次后台的命名标准

-----

#### URL:

+ 命名形式如

	/prefix/(小写)verb+book

-----

#### Models:(驼峰命名法)

+ User: ID,Username,Password,Confirmed,BookCount,Books[]
+ Kind: ID,Books[]
+ Book: gorm.Models,Bookname,Booknum,Available,Realname,LendTime,ReturnTime,KindID,UserID

-----

#### ReadJSON 结构体:

+ 命名形式如：

	type (驼峰命名法)verb+Book struct {...}

-----

#### api函数:

+ 命名形式如：

	func (驼峰命名法)Book+verb (ctx iris.Context){...}

-----
