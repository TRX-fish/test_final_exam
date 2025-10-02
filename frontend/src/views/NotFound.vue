<template>
  <div class="notfound-container">
    <div class="error-content">
      <div class="error-icon">
        <el-icon :size="120" color="#E6E8EB"><Warning /></el-icon>
      </div>
      <h1 class="error-code">404</h1>
      <h2 class="error-title">页面未找到</h2>
      <p class="error-desc">抱歉，您访问的页面不存在或已被移除</p>
      <div class="error-actions">
        <el-button type="primary" size="large" @click="goHome">
          <el-icon><HomeFilled /></el-icon>
          返回首页
        </el-button>
        <el-button size="large" @click="goBack">
          <el-icon><Back /></el-icon>
          返回上页
        </el-button>
      </div>
      
      <!-- 快捷链接 -->
      <div class="quick-links">
        <p class="links-title">您可能在找：</p>
        <div class="links-list">
          <el-button link @click="$router.push('/papers')">试卷列表</el-button>
          <span>|</span>
          <el-button link @click="$router.push('/questions')">题目搜索</el-button>
          <span>|</span>
          <el-button link @click="$router.push('/help')">帮助中心</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { Warning, HomeFilled, Back } from '@element-plus/icons-vue'

const router = useRouter()

const goHome = () => {
  router.push('/')
}

const goBack = () => {
  router.back()
}
</script>

<style scoped>
.notfound-container {
  min-height: calc(100vh - 60px);
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f8f9fa;
  padding: 32px;
}

.error-content {
  text-align: center;
  max-width: 600px;
}

.error-icon {
  margin-bottom: 24px;
}

.error-code {
  font-size: 96px;
  font-weight: 700;
  color: #E6E8EB;
  margin: 0 0 16px 0;
  line-height: 1;
}

.error-title {
  font-size: 28px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 12px 0;
}

.error-desc {
  font-size: 16px;
  color: #666;
  margin: 0 0 40px 0;
  line-height: 1.6;
}

.error-actions {
  display: flex;
  gap: 16px;
  justify-content: center;
  margin-bottom: 48px;
}

.quick-links {
  padding-top: 32px;
  border-top: 1px solid #e5e5e5;
}

.links-title {
  font-size: 14px;
  color: #999;
  margin: 0 0 16px 0;
}

.links-list {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  font-size: 14px;
  color: #999;
}

@media (max-width: 768px) {
  .error-code {
    font-size: 72px;
  }
  
  .error-title {
    font-size: 24px;
  }
  
  .error-actions {
    flex-direction: column;
  }
  
  .error-actions .el-button {
    width: 100%;
  }
}
</style> 