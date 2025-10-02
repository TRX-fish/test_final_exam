<template>
  <div class="dashboard-container">
    <!-- 顶部欢迎区 -->
    <div class="welcome-section">
      <div>
        <h1 class="welcome-title">欢迎回来！</h1>
        <p class="welcome-subtitle">开始您的高效复习之旅</p>
      </div>
    </div>

    <!-- 快捷操作区 -->
    <div class="section-card">
      <div class="section-header">
        <h3 class="section-title">快捷操作</h3>
      </div>
      <el-row :gutter="16" class="quick-actions">
        <el-col :span="6" :xs="12" :sm="8" :md="6">
          <div class="action-item" @click="$router.push('/papers')">
            <div class="action-icon">
              <el-icon :size="24"><List /></el-icon>
            </div>
            <div class="action-text">浏览试卷</div>
          </div>
        </el-col>
        <el-col :span="6" :xs="12" :sm="8" :md="6">
          <div class="action-item" @click="$router.push('/questions')">
            <div class="action-icon">
              <el-icon :size="24"><Search /></el-icon>
            </div>
            <div class="action-text">搜索题目</div>
          </div>
        </el-col>
        <el-col :span="6" :xs="12" :sm="8" :md="6">
          <div class="action-item" @click="$router.push('/favorites')">
            <div class="action-icon">
              <el-icon :size="24"><Star /></el-icon>
            </div>
            <div class="action-text">我的收藏</div>
          </div>
        </el-col>
        <el-col :span="6" :xs="12" :sm="8" :md="6">
          <div class="action-item" @click="$router.push('/mistakes')">
            <div class="action-icon">
              <el-icon :size="24"><Edit /></el-icon>
            </div>
            <div class="action-text">错题本</div>
          </div>
        </el-col>
      </el-row>
    </div>

    <!-- 热门试卷 -->
    <div class="section-card">
      <div class="section-header">
        <h3 class="section-title">热门试卷</h3>
        <el-button type="primary" link @click="$router.push('/papers')">
          查看全部 <el-icon><ArrowRight /></el-icon>
        </el-button>
      </div>
      
      <el-empty v-if="hotPapers.length === 0" description="暂无热门试卷" />
      
      <div v-else class="paper-grid">
        <div 
          v-for="paper in hotPapers.slice(0, 4)" 
          :key="paper.id" 
          class="paper-card hot-paper"
          @click="$router.push(`/papers/${paper.id}`)"
        >
          <div class="hot-badge">
            <el-icon><TrendCharts /></el-icon>
            <span>{{ paper.views }}浏览</span>
          </div>
          <div class="paper-header">
            <h4 class="paper-title">{{ paper.course }}</h4>
            <el-tag size="small" type="danger">{{ paper.year }}年</el-tag>
          </div>
          <div class="paper-meta">
            <span class="meta-item">
              <el-icon :size="14"><School /></el-icon>
              {{ paper.college }}
            </span>
            <span class="meta-item">
              <el-icon :size="14"><Calendar /></el-icon>
              {{ paper.term }}
            </span>
          </div>
          <div class="paper-footer">
            <span class="footer-time">{{ formatDate(paper.created_at) }}</span>
            <el-icon class="arrow-icon"><ArrowRight /></el-icon>
          </div>
        </div>
      </div>
    </div>

    <!-- 最近更新试卷 -->
    <div class="section-card">
      <div class="section-header">
        <h3 class="section-title">最近更新</h3>
        <el-button type="primary" link @click="$router.push('/papers')">
          查看全部 <el-icon><ArrowRight /></el-icon>
        </el-button>
      </div>
      
      <el-empty v-if="recentPapers.length === 0" description="暂无试卷数据" />
      
      <div v-else class="paper-grid">
        <div 
          v-for="paper in recentPapers.slice(0, 4)" 
          :key="paper.id" 
          class="paper-card"
          @click="$router.push(`/papers/${paper.id}`)"
        >
          <div class="paper-header">
            <h4 class="paper-title">{{ paper.course }}</h4>
            <el-tag size="small" type="info">{{ paper.year }}年</el-tag>
          </div>
          <div class="paper-meta">
            <span class="meta-item">
              <el-icon :size="14"><School /></el-icon>
              {{ paper.college }}
            </span>
            <span class="meta-item">
              <el-icon :size="14"><Calendar /></el-icon>
              {{ paper.term }}
            </span>
          </div>
          <div class="paper-footer">
            <span class="footer-time">{{ formatDate(paper.created_at) }}</span>
            <el-icon class="arrow-icon"><ArrowRight /></el-icon>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '../api/request'
import { Document, Reading, User, List, Search, Star, Edit, ArrowRight, School, Calendar, TrendCharts } from '@element-plus/icons-vue'

const router = useRouter()
const stats = ref({ papers: 0, questions: 0, users: 0 })
const recentPapers = ref([])
const hotPapers = ref([])

const fetchStats = async () => {
  try {
    const res = await api.get('/api/stats')
    if (res.data.code === '2000') {
      stats.value = res.data.data
    }
  } catch (e) {
    console.error('获取统计数据失败', e)
  }
}

const fetchRecentPapers = async () => {
  try {
    const res = await api.get('/api/papers', { params: { page: 1, per_page: 4 } })
    if (res.data.code === '2000') {
      recentPapers.value = res.data.data.items || []
    }
  } catch (e) {
    console.error('获取最近试卷失败', e)
  }
}

const fetchHotPapers = async () => {
  try {
    const res = await api.get('/api/papers', { params: { page: 1, per_page: 4 } })
    if (res.data.code === '2000') {
      // 模拟添加浏览量数据
      const papers = res.data.data.items || []
      hotPapers.value = papers.map((paper, index) => ({
        ...paper,
        views: 500 - index * 50 // 模拟浏览量，实际应从后端获取
      }))
    }
  } catch (e) {
    console.error('获取热门试卷失败', e)
  }
}

const formatDate = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  const now = new Date()
  const diff = now - date
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))
  
  if (days === 0) return '今天'
  if (days === 1) return '昨天'
  if (days < 7) return `${days}天前`
  return dateStr.split(' ')[0]
}

onMounted(() => {
  fetchStats()
  fetchHotPapers()
  fetchRecentPapers()
})
</script>

<style scoped>
.dashboard-container {
  padding: 32px;
  max-width: 1400px;
  margin: 0 auto;
}

/* 欢迎区 */
.welcome-section {
  margin-bottom: 32px;
}

.welcome-title {
  font-size: 32px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 8px 0;
  line-height: 1.2;
}

.welcome-subtitle {
  font-size: 16px;
  color: #666;
  margin: 0;
  line-height: 1.5;
}

/* 区块卡片 */
.section-card {
  background: #fff;
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0;
  line-height: 1.2;
}

/* 快捷操作 */
.quick-actions {
  margin-top: 16px;
}

.action-item {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 24px 16px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s;
  border: 1px solid transparent;
}

.action-item:hover {
  background: #fff;
  border-color: #3A7AFE;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(58, 122, 254, 0.1);
}

.action-icon {
  width: 48px;
  height: 48px;
  background: #3A7AFE;
  color: #fff;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 12px;
  transition: all 0.3s;
}

.action-item:hover .action-icon {
  background: #2969ed;
}

.action-text {
  font-size: 14px;
  color: #333;
  line-height: 1.5;
}

/* 试卷网格 */
.paper-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
}

.paper-card {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 20px;
  cursor: pointer;
  transition: all 0.3s;
  border: 1px solid transparent;
  position: relative;
}

/* 热门试卷特殊样式 */
.hot-paper {
  background: linear-gradient(135deg, #FFF5F5 0%, #FFF9F0 100%);
  border: 1px solid #FFE5E5;
}

.hot-badge {
  position: absolute;
  top: 12px;
  right: 12px;
  background: linear-gradient(135deg, #FF6B6B 0%, #FF8E53 100%);
  color: white;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  display: flex;
  align-items: center;
  gap: 4px;
  font-weight: 500;
  box-shadow: 0 2px 8px rgba(255, 107, 107, 0.3);
}

.paper-card:hover {
  background: #fff;
  border-color: #3A7AFE;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(58, 122, 254, 0.1);
}

.hot-paper:hover {
  background: #fff;
  border-color: #FF6B6B;
  box-shadow: 0 4px 16px rgba(255, 107, 107, 0.2);
}

.paper-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 16px;
}

.paper-title {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0;
  line-height: 1.5;
  flex: 1;
  margin-right: 8px;
}

.paper-meta {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 16px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #666;
}

.paper-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 12px;
  border-top: 1px solid #f0f0f0;
}

.footer-time {
  font-size: 12px;
  color: #999;
}

.arrow-icon {
  color: #3A7AFE;
  transition: transform 0.3s;
}

.paper-card:hover .arrow-icon {
  transform: translateX(4px);
}

/* 响应式 */
@media (max-width: 768px) {
  .dashboard-container {
    padding: 16px;
  }
  
  .paper-grid {
    grid-template-columns: 1fr;
  }
}
</style> 