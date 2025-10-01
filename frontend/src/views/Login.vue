<template>
  <el-row justify="center" style="margin-top: 100px;">
    <el-col :span="8">
      <el-card>
        <h2 style="text-align:center;">用户登录</h2>
        <el-form :model="form" :rules="rules" ref="loginForm" label-width="80px">
          <el-form-item label="用户名" prop="username">
            <el-input v-model="form.username" autocomplete="off" />
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input v-model="form.password" type="password" autocomplete="off" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onLogin" :loading="loading">登录</el-button>
            <el-button type="text" @click="$router.push('/register')">注册</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '../api/request'
import { ElMessage } from 'element-plus'

const router = useRouter()
const form = ref({ username: '', password: '' })
const loading = ref(false)
const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

const loginForm = ref(null)

const onLogin = () => {
  loginForm.value.validate(async valid => {
    if (!valid) return
    loading.value = true
    try {
      const res = await api.post('/api/login', form.value)
      if (res.data.code === '2000') {
        localStorage.setItem('token', res.data.data.token)
        ElMessage.success('登录成功')
        // 跳转到首页或其它页面
        router.push('/')
      } else {
        ElMessage.error(res.data.msg)
      }
    } catch (e) {
      ElMessage.error('登录失败')
    } finally {
      loading.value = false
    }
  })
}
</script> 