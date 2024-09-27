# GVB Server 晨曦博客后端

晨曦一代博客

项目技术栈：
`go` `gin` `gorm` `elastic` 

# 项目运行



# 项目部署
````shell
// 交叉编译
set GOARCH=amd64
set COOS=linux
go build -o main
// 记得编译后改回
set COOS=windows
// 导出数据库
mysqldump -uroot -p123456 go_blog >gvb.sql
//编码为UTF-8,否则会报错
//ssh 使用rz上传到服务器
main gvb.sql config.yaml uploads docs
// main转可执行文件
chmod +x main
// 后台执行文件 防止关闭ssh 进程结束
nohup ./main &

