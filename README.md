# 大学生期末考试收录平台

一个基于 **Go + Vue 3** 的期末考试题库管理系统，用于收录往年期末考试题目与解析，帮助学生高效复习！

## 🚀 技术栈

### 后端（Go 重构版）
- **框架**: Gin Web Framework
- **ORM**: GORM
- **数据库**: MySQL/MariaDB
- **认证**: JWT (golang-jwt/jwt)
- **密码加密**: bcrypt

### 前端
- **框架**: Vue 3
- **构建工具**: Vite
- **UI 框架**: Element Plus
- **路由**: Vue Router
- **HTTP**: Axios

## 📁 项目结构

```
test_final_exam/
├── backend-go/          # Go 后端
│   ├── main.go         # 主程序入口
│   ├── config/         # 配置管理
│   ├── models/         # 数据模型
│   ├── handlers/       # 路由处理器
│   ├── services/       # 业务逻辑层
│   ├── middleware/     # 中间件（JWT认证等）
│   ├── utils/          # 工具函数
│   ├── database/       # 数据库连接
│   └── static/         # 静态文件
├── frontend/           # Vue 3 前端
│   ├── src/           # 源代码
│   ├── public/        # 公共资源
│   └── package.json   # 依赖配置
└── config.ini         # 数据库配置
```

## 🔧 环境要求

- Go 1.20+
- Node.js 16+
- MySQL 5.7+ / MariaDB

## 📦 安装部署

### 1. 克隆项目

```bash
git clone https://github.com/TRX-fish/test_final_exam.git
cd test_final_exam
```

### 2. 配置数据库

编辑 `config.ini` 文件：

```ini
[mysql]
user = root
password = your_password
host = localhost
port = 3306
db = final_exam
```

创建数据库：

```sql
CREATE DATABASE final_exam CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

### 3. 启动后端

```bash
cd backend-go
go mod tidy
go run main.go
```

后端将运行在 `http://localhost:5000`

### 4. 启动前端

```bash
cd frontend
npm install
npm run dev
```

前端将运行在 `http://localhost:5173`

## 🌐 API 接口

### 用户相关
- `POST /api/register` - 用户注册
- `POST /api/login` - 用户登录
- `GET /api/user/me` - 获取当前用户信息

### 试卷相关
- `POST /api/papers` - 添加试卷
- `GET /api/papers` - 获取试卷列表
- `PUT /api/papers/:id/image` - 更新试卷图片
- `GET /api/stats` - 获取统计数据

### 题目相关
- `POST /api/questions` - 添加题目
- `GET /api/questions` - 获取题目列表
- `GET /api/questions/search` - 搜索题目
- `PUT /api/questions/:id/image` - 更新题目图片

### 上传相关
- `POST /api/upload/paper-image` - 上传试卷图片
- `POST /api/upload/question-image` - 上传题目图片
- `DELETE /api/upload/delete-image` - 删除图片

## 💾 数据库表结构

### user (用户表)
- id - 用户ID
- username - 用户名
- password_hash - 密码哈希
- email - 邮箱
- role - 角色（user/admin）
- is_active - 是否激活
- created_at - 注册时间
- updated_at - 更新时间

### paper (试卷表)
- id - 试卷ID
- course - 课程名称
- year - 年份
- term - 学期
- college - 学院
- image_path - 图片路径
- created_at - 创建时间

### question (题目表)
- id - 题目ID
- paper_id - 试卷ID
- content - 题目内容
- type - 题目类型
- options - 选项
- answer - 答案
- explanation - 解析
- image_path - 图片路径
- created_at - 创建时间

## 📝 开发说明

### 后端开发

Go 后端使用 Gin 框架，采用分层架构：
- `handlers`: 处理 HTTP 请求
- `services`: 业务逻辑
- `models`: 数据模型
- `middleware`: 中间件

### 前端开发

Vue 3 使用 Composition API，Element Plus 作为 UI 框架。

## 🔒 安全特性

- JWT Token 认证
- Bcrypt 密码加密
- CORS 跨域配置
- 文件类型验证

## 📄 License

MIT License

## 👨‍💻 作者

TRX-fish

---

⭐ 如果这个项目对你有帮助，请给个 Star！ 