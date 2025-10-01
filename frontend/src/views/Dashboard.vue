<template>
  <el-row justify="center" style="margin-top: 40px;">
    <el-col :span="16">
      <el-card>
        <h2 style="text-align:center;">大学生期末考试收录平台</h2>
        <p style="text-align:center; margin-bottom: 20px;">收录往年期末考试题目与解析，助力高效复习！</p>
        <el-row :gutter="20" justify="center">
          <el-col :span="6">
            <el-statistic title="试卷总数" :value="stats.papers" />
          </el-col>
          <el-col :span="6">
            <el-statistic title="题目总数" :value="stats.questions" />
          </el-col>
          <el-col :span="6">
            <el-statistic title="注册用户" :value="stats.users" />
          </el-col>
        </el-row>
        <el-divider />
        <el-row justify="center" style="margin-top: 20px;">
          <el-button type="primary" @click="$router.push('/papers')">进入试卷列表</el-button>
          <el-button @click="$router.push('/user-center')">个人中心</el-button>
        </el-row>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../api/request'

const stats = ref({ papers: 0, questions: 0, users: 0 })

onMounted(async () => {
  try {
    // 假设后端有统计接口 /api/stats
    const res = await api.get('/api/stats')
    if (res.data.code === '2000') {
      stats.value = res.data.data
    }
  } catch (e) {
    // 可忽略错误，默认显示0
  }
})
</script> 