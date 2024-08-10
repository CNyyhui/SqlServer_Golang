# SqlServer_Golang

使用golang对数据库进行增删查改方法包
该包是基于
github.com/denisenkom/go-mssqldb
的进一步开发

前提：
配置本地go环境加入这位大佬写的包
github.com/denisenkom/go-mssqldb
作为踩过很多坑的新手，我建议把git上引入的包全丢在你的golang安装包下的src
新建一个github.com文件夹,然后放入github上下载的文件

比如你golang的下载路径是 盘C 用户user 文件夹go
那么你就应该在go文件夹内的src下新建一个名为github.com的包，并且把下载包放到github.com包下才能在本地索引

你应该放的本地文件路径（在该例子中）：C:\\User\go\src\github.com\下载包名
在golang编译器中的索引："github.com/denisenkom/go-mssqldb"



方法设置：
目前没有通用的方法，该包目前只针对三个变量，后续会更新变量索引使得更加有弹性（duang~~~）

Domain       ：所有所需变量的定义区，但你如果熟知go语言，请教教我 *\^-^/*
AddDaTabase  ：创建表，后续会扩容至ADD——新建库，表一体
Connection   ：测试是否连接到对应的数据库
DeleteData   ：删除数据、表、库
InsertData   ：插入数据，新建列    目前还没有实现传值和弹性变量 只有三个变量
QueryData    ：查询数据
UpdateData   ：更新数据
test         ：想干嘛干嘛  do anything u want ~~~
main         ：运行方法，实现操作
