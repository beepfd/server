# BeepFD Server

BeepFD Server 是一个基于 eBPF 技术的故障检测服务器端实现。该项目提供了一套完整的故障检测、监控和管理功能。

## 功能特性

- 基于 eBPF 的系统级监控
- Prometheus 指标集成
- RESTful API 接口
- 性能分析支持
- MySQL 数据持久化
- 完整的日志记录系统

## 技术栈

- Go 1.24+
- eBPF
- Gin Web 框架
- GORM
- Prometheus
- MySQL

## 系统要求

- Linux 内核版本 >= 4.18
- Go 1.24 或更高版本
- MySQL 5.7 或更高版本

## 快速开始

### 安装

1. 克隆仓库：

```bash
git clone https://github.com/beepfd/server.git
cd server
```

2. 安装依赖：

```bash
go mod download
```

3. 编译项目：

```bash
make build
```

### 配置

在运行服务之前，请确保正确配置 `config.json` 文件。主要配置项包括：

- 数据库连接信息
- 服务监听端口
- 日志级别
- 监控配置

### 运行

```bash
./cmd/server
```

## 项目结构

```
.
├── cmd/            # 主程序入口
├── conf/           # 配置文件
├── internal/       # 内部包
├── models/         # 数据模型
├── pkg/            # 公共包
├── router/         # 路由定义
├── sql/           # SQL 脚本
└── testdata/      # 测试数据
```

## 许可证

本项目采用 [LICENSE](LICENSE) 许可证。

## 贡献指南

欢迎提交 Pull Request 或创建 Issue。

## 联系方式

如有问题，请通过 Issue 系统反馈。
