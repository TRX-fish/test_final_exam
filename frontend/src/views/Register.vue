<template>
  <el-row justify="center" style="margin-top: 100px;">
    <el-col :span="8">
      <el-card>
        <h2 style="text-align:center;">用户注册</h2>
        <el-form :model="form" :rules="rules" ref="registerForm" label-width="80px">
          <el-form-item label="用户名" prop="username">
            <el-input v-model="form.username" autocomplete="off" />
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input v-model="form.password" type="password" autocomplete="off" />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="form.email" autocomplete="off" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onRegister" :loading="loading">注册</el-button>
            <el-button type="text" @click="$router.push('/login')">返回登录</el-button>
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
const form = ref({ username: '', password: '', email: '' })
const loading = ref(false)
const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  email: [{ type: 'email', message: '请输入正确的邮箱', trigger: 'blur' }]
}

const registerForm = ref(null)

const onRegister = () => {
  registerForm.value.validate(async valid => {
    if (!valid) return
    loading.value = true
    try {
      console.log('发送注册请求:', form.value)
      const res = await api.post('/api/register', form.value)
      console.log('注册响应:', res.data)
      if (res.data.code === '2000') {
        ElMessage.success('注册成功，请登录')
        router.push('/login')
      } else {
        ElMessage.error(res.data.msg || '注册失败')
      }
    } catch (e) {
      console.error('注册错误:', e)
      console.error('错误详情:', e.response?.data)
      ElMessage.error(e.response?.data?.msg || '注册失败')
    } finally {
      loading.value = false
    }
  })
}
</script> 