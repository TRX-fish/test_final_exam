<template>
  <div class="register-container">
    <div class="register-card">
      <!-- 左侧装饰区 -->
      <div class="left-section">
        <div class="logo-area">
          <div class="logo-icon">
            <el-icon :size="48"><Reading /></el-icon>
          </div>
          <h1 class="brand-title">加入我们</h1>
          <p class="brand-subtitle">开始构建您的专属学习题库</p>
        </div>
        <div class="features">
          <div class="feature-item">
            <el-icon :size="24" color="#fff"><Check /></el-icon>
            <span>海量真题资源</span>
          </div>
          <div class="feature-item">
            <el-icon :size="24" color="#fff"><Check /></el-icon>
            <span>智能题目搜索</span>
          </div>
          <div class="feature-item">
            <el-icon :size="24" color="#fff"><Check /></el-icon>
            <span>详细答案解析</span>
          </div>
        </div>
      </div>

      <!-- 右侧注册表单 -->
      <div class="right-section">
        <div class="form-container">
          <h2 class="form-title">创建账号</h2>
          <p class="form-subtitle">填写信息完成注册</p>

          <el-form 
            :model="form" 
            :rules="rules" 
            ref="registerFormRef" 
            class="register-form"
            @submit.prevent="onRegister"
          >
            <el-form-item prop="username">
              <el-input 
                v-model="form.username" 
                placeholder="用户名（3-50个字符）"
                size="large"
                :prefix-icon="User"
                clearable
              />
            </el-form-item>

            <el-form-item prop="email">
              <el-input 
                v-model="form.email" 
                placeholder="邮箱地址（选填）"
                size="large"
                :prefix-icon="Message"
                clearable
              />
            </el-form-item>

            <el-form-item prop="password">
              <el-input 
                v-model="form.password" 
                type="password" 
                placeholder="密码（至少6个字符）"
                size="large"
                :prefix-icon="Lock"
                show-password
              />
            </el-form-item>

            <el-form-item prop="confirmPassword">
              <el-input 
                v-model="form.confirmPassword" 
                type="password" 
                placeholder="确认密码"
                size="large"
                :prefix-icon="Lock"
                show-password
                @keyup.enter="onRegister"
              />
            </el-form-item>

            <el-form-item>
              <el-button 
                type="primary" 
                size="large"
                class="submit-btn"
                @click="onRegister" 
                :loading="loading"
              >
                注册
              </el-button>
            </el-form-item>
          </el-form>

          <div class="form-footer">
            <span class="footer-text">已有账号？</span>
            <el-button type="primary" link @click="$router.push('/login')">
              立即登录
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
import { User, Lock, Message, Reading, Check } from '@element-plus/icons-vue'

const router = useRouter()
const form = ref({ 
  username: '', 
  email: '', 
  password: '', 
  confirmPassword: '' 
})
const loading = ref(false)
const registerFormRef = ref(null)

const validatePass = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请输入密码'))
  } else if (value.length < 6) {
    callback(new Error('密码长度不能少于 6 个字符'))
  } else {
    if (form.value.confirmPassword !== '') {
      registerFormRef.value.validateField('confirmPassword')
    }
    callback()
  }
}

const validatePass2 = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请再次输入密码'))
  } else if (value !== form.value.password) {
    callback(new Error('两次输入密码不一致'))
  } else {
    callback()
  }
}

const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 50, message: '用户名长度在 3 到 50 个字符', trigger: 'blur' },
    { pattern: /^[a-zA-Z0-9_]+$/, message: '用户名只能包含字母、数字和下划线', trigger: 'blur' }
  ],
  email: [
    { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' }
  ],
  password: [
    { validator: validatePass, trigger: 'blur' }
  ],
  confirmPassword: [
    { validator: validatePass2, trigger: 'blur' }
  ]
}

const onRegister = () => {
  registerFormRef.value.validate(async valid => {
    if (!valid) return
    
    loading.value = true
    try {
      const { confirmPassword, ...data } = form.value
      const res = await api.post('/api/register', data)
      if (res.data.code === '2000') {
        ElMessage.success({
          message: '注册成功！即将跳转到登录页...',
          duration: 2000
        })
        setTimeout(() => {
          router.push('/login')
        }, 1500)
      } else {
        ElMessage.error(res.data.msg || '注册失败')
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
.register-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #E8F1FF 0%, #F5F9FF 100%);
  padding: 24px;
}

.register-card {
  display: flex;
  width: 100%;
  max-width: 1000px;
  min-height: 640px;
  background: #ffffff;
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(58, 122, 254, 0.08);
  overflow: hidden;
}

/* 左侧装饰区 */
.left-section {
  flex: 1;
  background: linear-gradient(135deg, #00C2A8 0%, #00D4B5 100%);
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

.features {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.feature-item {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 16px;
  opacity: 0.95;
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

.register-form {
  margin-top: 32px;
}

.register-form :deep(.el-form-item) {
  margin-bottom: 20px;
}

.register-form :deep(.el-input__wrapper) {
  border-radius: 8px;
  box-shadow: 0 0 0 1px #e5e5e5 inset;
  transition: all 0.3s;
}

.register-form :deep(.el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px #00C2A8 inset;
}

.register-form :deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 2px #00C2A8 inset;
}

.submit-btn {
  width: 100%;
  height: 48px;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 500;
  background: #00C2A8;
  border: none;
  margin-top: 8px;
  transition: all 0.3s;
}

.submit-btn:hover {
  background: #00B199;
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(0, 194, 168, 0.3);
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
  
  .register-card {
    max-width: 480px;
  }
  
  .right-section {
    padding: 32px 24px;
  }
}
</style> 