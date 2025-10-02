<template>
  <div id="app">
    <!-- 顶部导航栏 -->
    <el-header v-if="showNav" style="background-color: #409EFF; color: white; height: 60px; padding: 0 20px;">
      <div style="display: flex; justify-content: space-between; align-items: center; height: 100%;">
        <div style="display: flex; align-items: center;">
          <h2 style="margin: 0; cursor: pointer;" @click="$router.push('/')">期末考试题库系统</h2>
          <el-menu
            mode="horizontal"
            :default-active="activeIndex"
            background-color="#409EFF"
            text-color="#fff"
            active-text-color="#ffd04b"
            style="margin-left: 40px; border: none;"
            @select="handleSelect"
          >
            <el-menu-item index="/">首页</el-menu-item>
            <el-menu-item index="/papers">试卷</el-menu-item>
            <el-menu-item index="/questions">题目</el-menu-item>
            <el-sub-menu index="study">
              <template #title>学习</template>
              <el-menu-item index="/favorites">我的收藏</el-menu-item>
              <el-menu-item index="/mistakes">错题本</el-menu-item>
              <el-menu-item index="/history">学习记录</el-menu-item>
              <el-menu-item index="/notes">我的笔记</el-menu-item>
            </el-sub-menu>
            <el-menu-item index="/discussion">讨论区</el-menu-item>
          </el-menu>
        </div>
        <div style="display: flex; align-items: center; gap: 24px;">
          <!-- 会员按钮 -->
          <el-button 
            v-if="!isMember"
            type="warning" 
            size="small"
            style="background: linear-gradient(135deg, #FFD700 0%, #FFA500 100%); border: none; font-weight: 600;"
            @click="$router.push('/membership')"
          >
            <el-icon><Wallet /></el-icon>
            开通会员
          </el-button>
          <el-tag v-else type="warning" effect="dark" style="cursor: pointer;" @click="$router.push('/membership')">
            <el-icon><Medal /></el-icon>
            VIP会员
          </el-tag>
          
          <!-- 通知 -->
          <el-badge :value="unreadCount" :hidden="unreadCount === 0">
            <el-icon :size="20" style="cursor: pointer; color: white;" @click="$router.push('/notifications')">
              <Bell />
            </el-icon>
          </el-badge>
          
          <!-- 用户菜单 -->
          <el-dropdown @command="handleCommand">
            <span style="cursor: pointer; color: white; display: flex; align-items: center; gap: 8px;">
              <el-icon><User /></el-icon>
              <span>{{ username }}</span>
              <el-icon><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="admin" v-if="isAdmin" divided>
                  <el-icon><Setting /></el-icon>
                  管理后台
                </el-dropdown-item>
                <el-dropdown-item command="help">
                  <el-icon><QuestionFilled /></el-icon>
                  帮助中心
                </el-dropdown-item>
                <el-dropdown-item command="about">
                  <el-icon><InfoFilled /></el-icon>
                  关于我们
                </el-dropdown-item>
                <el-dropdown-item command="logout" divided>
                  <el-icon><SwitchButton /></el-icon>
                  退出登录
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
    </el-header>

    <!-- 主内容区 -->
    <el-main style="padding: 0; min-height: calc(100vh - 60px);">
      <router-view />
    </el-main>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import { 
  User, ArrowDown, Bell, Setting, QuestionFilled, InfoFilled, SwitchButton, Wallet, Medal
} from '@element-plus/icons-vue'

const router = useRouter()
const route = useRoute()

// 不显示导航栏的页面
const noNavPages = ['/login', '/register']
const showNav = computed(() => !noNavPages.includes(route.path))
const activeIndex = ref(route.path)

// 用户信息
const username = ref(localStorage.getItem('username') || 'admin')
const isAdmin = computed(() => localStorage.getItem('userRole') === 'admin')
const unreadCount = ref(2)

// 会员状态
const isMember = computed(() => {
  const membership = localStorage.getItem('membership')
  if (!membership) return false
  
  try {
    const memberData = JSON.parse(membership)
    const endDate = new Date(memberData.endDate)
    const now = new Date()
    return memberData.status === 'active' && endDate > now
  } catch {
    return false
  }
})

// 监听路由变化
watch(() => route.path, (newPath) => {
  activeIndex.value = newPath
})

const handleSelect = (index) => {
  router.push(index)
}

const handleCommand = (command) => {
  if (command === 'logout') {
    localStorage.removeItem('token')
    localStorage.removeItem('userRole')
    localStorage.removeItem('username')
    ElMessage.success('退出登录成功')
    router.push('/login')
  } else if (command === 'admin') {
    router.push('/admin')
  } else if (command === 'help') {
    router.push('/help')
  } else if (command === 'about') {
    router.push('/about')
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', Arial, sans-serif;
}

body {
  margin: 0;
  background-color: #f5f5f5;
}
</style> 