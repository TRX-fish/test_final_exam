<template>
  <div class="login-container">
    <div class="login-card">
      <!-- 左侧装饰区 -->
      <div class="left-section">
        <div class="logo-area">
          <div class="logo-icon">
            <el-icon :size="48"><Reading /></el-icon>
          </div>
          <h1 class="brand-title">期末考试题库</h1>
          <p class="brand-subtitle">收录往年考试真题，助力高效复习备考</p>
        </div>
        <div class="illustration">
          <el-icon :size="200" color="#E8F1FF"><Document /></el-icon>
        </div>
      </div>

      <!-- 右侧登录表单 -->
      <div class="right-section">
        <div class="form-container">
          <h2 class="form-title">欢迎登录</h2>
          <p class="form-subtitle">使用您的账号登录系统</p>

          <el-form 
            :model="form" 
            :rules="rules" 
            ref="loginFormRef" 
            class="login-form"
            @submit.prevent="onLogin"
          >
            <el-form-item prop="username">
              <el-input 
                v-model="form.username" 
                placeholder="请输入用户名"
                size="large"
                :prefix-icon="User"
                clearable
              />
            </el-form-item>

            <el-form-item prop="password">
              <el-input 
                v-model="form.password" 
                type="password" 
                placeholder="请输入密码"
                size="large"
                :prefix-icon="Lock"
                show-password
                @keyup.enter="onLogin"
              />
            </el-form-item>

            <el-form-item>
              <el-button 
                type="primary" 
                size="large"
                class="submit-btn"
                @click="onLogin" 
                :loading="loading"
              >
                登录
              </el-button>
            </el-form-item>
          </el-form>

          <div class="form-footer">
            <span class="footer-text">还没有账号？</span>
            <el-button type="primary" link @click="$router.push('/register')">
              立即注册
            </el-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '../api/request'
import { ElMessage } from 'element-plus'
import { User, Lock, Reading, Document } from '@element-plus/icons-vue'

const router = useRouter()
const form = ref({ username: '', password: '' })
const loading = ref(false)
const loginFormRef = ref(null)

const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 50, message: '用户名长度在 3 到 50 个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于 6 个字符', trigger: 'blur' }
  ]
}

const onLogin = () => {
  loginFormRef.value.validate(async valid => {
    if (!valid) return
    
    loading.value = true
    try {
      const res = await api.post('/api/login', form.value)
      if (res.data.code === '2000') {
        localStorage.setItem('token', res.data.data.token)
        localStorage.setItem('userRole', res.data.data.role)
        localStorage.setItem('username', form.value.username)
        ElMessage.success({
          message: '登录成功，欢迎回来！',
          duration: 2000
        })
        setTimeout(() => {
          router.push('/')
        }, 500)
      } else {
        ElMessage.error(res.data.msg || '登录失败')
      }
    } catch (e) {
      ElMessage.error('网络连接失败，请稍后重试')
    } finally {
      loading.value = false
    }
  })
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #E8F1FF 0%, #F5F9FF 100%);
  padding: 24px;
}

.login-card {
  display: flex;
  width: 100%;
  max-width: 1000px;
  min-height: 600px;
  background: #ffffff;
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(58, 122, 254, 0.08);
  overflow: hidden;
}

/* 左侧装饰区 */
.left-section {
  flex: 1;
  background: linear-gradient(135deg, #3A7AFE 0%, #5B8DFF 100%);
  padding: 48px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  color: white;
}

.logo-area {
  margin-bottom: 40px;
}

.logo-icon {
  width: 64px;
  height: 64px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 24px;
  backdrop-filter: blur(10px);
}

.brand-title {
  font-size: 32px;
  font-weight: 600;
  margin: 0 0 16px 0;
  line-height: 1.2;
}

.brand-subtitle {
  font-size: 16px;
  opacity: 0.9;
  line-height: 1.5;
  margin: 0;
}

.illustration {
  text-align: center;
  opacity: 0.3;
}

/* 右侧表单区 */
.right-section {
  flex: 1;
  padding: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.form-container {
  width: 100%;
  max-width: 400px;
}

.form-title {
  font-size: 28px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 8px 0;
  line-height: 1.2;
}

.form-subtitle {
  font-size: 14px;
  color: #666;
  margin: 0 0 40px 0;
  line-height: 1.5;
}

.login-form {
  margin-top: 32px;
}

.login-form :deep(.el-form-item) {
  margin-bottom: 24px;
}

.login-form :deep(.el-input__wrapper) {
  border-radius: 8px;
  box-shadow: 0 0 0 1px #e5e5e5 inset;
  transition: all 0.3s;
}

.login-form :deep(.el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px #3A7AFE inset;
}

.login-form :deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 2px #3A7AFE inset;
}

.submit-btn {
  width: 100%;
  height: 48px;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 500;
  background: #3A7AFE;
  border: none;
  margin-top: 8px;
  transition: all 0.3s;
}

.submit-btn:hover {
  background: #2969ed;
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(58, 122, 254, 0.3);
}

.submit-btn:active {
  transform: translateY(0);
}

.form-footer {
  text-align: center;
  margin-top: 32px;
  padding-top: 24px;
  border-top: 1px solid #f0f0f0;
}

.footer-text {
  color: #666;
  font-size: 14px;
  margin-right: 8px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .left-section {
    display: none;
  }
  
  .login-card {
    max-width: 480px;
  }
  
  .right-section {
    padding: 32px 24px;
  }
}
</style> 