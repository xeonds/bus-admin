# Bus-Admin 公交调度管理系统

借助`Gorm`和`MySQL`，实现的一个简单的公交调度管理系统。

## 项目结构

```bash
.
├── server/     # 后端
├── web/        # 前端
├── README.md
├── Makefile
└── .gitignore
```

## 技术栈

- 后端：`Golang` + `Gin` + `Gorm` + `MySQL`
- 前端：`Vue` + `ElementUI`

## 项目部署

项目在Linux环境下构建，使用Ubuntu 22.04 LTS测试通过。

1. 安装`Golang`和`Node.js`，并配置好环境变量
2. 在config.yaml中配置好数据库地址等信息
3. 运行如下指令：

```bash
make && make run
```
