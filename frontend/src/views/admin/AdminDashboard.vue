<template>
  <div class="admin-dashboard">
    <div class="page-header">
      <div>
        <h2 class="page-title">管理员控制台</h2>
        <p class="page-subtitle">系统概览与数据统计</p>
      </div>
    </div>

    <!-- 核心数据统计 -->
    <el-row :gutter="24" class="stats-row">
      <el-col :span="6" :xs="12" :sm="12" :md="6">
        <div class="stat-card primary">
          <div class="stat-icon">
            <el-icon :size="32"><Document /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.papers }}</div>
            <div class="stat-label">试卷总数</div>
            <div class="stat-trend">
              <el-icon color="#00C2A8"><TrendCharts /></el-icon>
              <span>+12%</span>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :span="6" :xs="12" :sm="12" :md="6">
        <div class="stat-card success">
          <div class="stat-icon">
            <el-icon :size="32"><Reading /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.questions }}</div>
            <div class="stat-label">题目总数</div>
            <div class="stat-trend">
              <el-icon color="#00C2A8"><TrendCharts /></el-icon>
              <span>+8%</span>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :span="6" :xs="12" :sm="12" :md="6">
        <div class="stat-card warning">
          <div class="stat-icon">
            <el-icon :size="32"><User /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.users }}</div>
            <div class="stat-label">用户总数</div>
            <div class="stat-trend">
              <el-icon color="#00C2A8"><TrendCharts /></el-icon>
              <span>+5%</span>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :span="6" :xs="12" :sm="12" :md="6">
        <div class="stat-card info">
          <div class="stat-icon">
            <el-icon :size="32"><View /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ todayVisits }}</div>
            <div class="stat-label">今日访问</div>
            <div class="stat-trend">
              <el-icon color="#00C2A8"><TrendCharts /></el-icon>
              <span>+15%</span>
            </div>
          </div>
        </div>
      </el-col>
    </el-row>

    <!-- 快捷管理入口 -->
    <div class="section-card">
      <div class="section-header">
        <h3 class="section-title">快捷管理</h3>
      </div>
      <el-row :gutter="16">
        <el-col :span="6" :xs="12" :sm="8" :md="6">
          <div class="manage-item" @click="$router.push('/admin/papers')">
            <div class="manage-icon" style="background: #E8F1FF;">
              <el-icon :size="24" color="#3A7AFE"><Tickets /></el-icon>
            </div>
            <div class="manage-text">试卷管理</div>
          </div>
        </el-col>
        <el-col :span="6" :xs="12" :sm="8" :md="6">
          <div class="manage-item" @click="$router.push('/admin/questions')">
            <div class="manage-icon" style="background: #E7F9F6;">
              <el-icon :size="24" color="#00C2A8"><EditPen /></el-icon>
            </div>
            <div class="manage-text">题目管理</div>
          </div>
        </el-col>
        <el-col :span="6" :xs="12" :sm="8" :md="6">
          <div class="manage-item" @click="$router.push('/admin/users')">
            <div class="manage-icon" style="background: #FFF5E6;">
              <el-icon :size="24" color="#FF9800"><UserFilled /></el-icon>
            </div>
            <div class="manage-text">用户管理</div>
          </div>
        </el-col>
        <el-col :span="6" :xs="12" :sm="8" :md="6">
          <div class="manage-item" @click="$router.push('/admin/settings')">
            <div class="manage-icon" style="background: #F5F5F5;">
              <el-icon :size="24" color="#666"><Setting /></el-icon>
            </div>
            <div class="manage-text">系统设置</div>
          </div>
        </el-col>
      </el-row>
    </div>

    <el-row :gutter="24">
      <!-- 最近活动 -->
      <el-col :span="12" :xs="24">
        <div class="section-card">
          <div class="section-header">
            <h3 class="section-title">最近活动</h3>
          </div>
          <div class="activity-list">
            <div v-for="activity in recentActivities" :key="activity.id" class="activity-item">
              <div class="activity-icon">
                <el-icon :size="16" :color="activity.color">
                  <component :is="activity.icon" />
                </el-icon>
              </div>
              <div class="activity-content">
                <div class="activity-text">{{ activity.text }}</div>
                <div class="activity-time">{{ activity.time }}</div>
              </div>
            </div>
          </div>
        </div>
      </el-col>

      <!-- 待处理事项 -->
      <el-col :span="12" :xs="24">
        <div class="section-card">
          <div class="section-header">
            <h3 class="section-title">待处理事项</h3>
            <el-badge :value="pendingCount" class="badge">
              <el-icon :size="20"><Bell /></el-icon>
            </el-badge>
          </div>
          <div class="pending-list">
            <div v-for="item in pendingItems" :key="item.id" class="pending-item">
              <div class="pending-content">
                <div class="pending-title">{{ item.title }}</div>
                <div class="pending-desc">{{ item.desc }}</div>
              </div>
              <el-button type="primary" link @click="handlePending(item.id)">
                处理
              </el-button>
            </div>
          </div>
          <el-empty v-if="pendingItems.length === 0" description="暂无待处理事项" />
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '../../api/request'
import { 
  Document, Reading, User, View, TrendCharts, Tickets, EditPen, 
  UserFilled, Setting, Bell, Plus, Delete, Edit 
} from '@element-plus/icons-vue'

const router = useRouter()
const stats = ref({ papers: 0, questions: 0, users: 0 })
const todayVisits = ref(156)

const recentActivities = ref([
  { id: 1, icon: Plus, color: '#00C2A8', text: 'user1 添加了试卷《数据结构》', time: '10分钟前' },
  { id: 2, icon: Edit, color: '#3A7AFE', text: 'admin 编辑了题目ID: 25', time: '30分钟前' },
  { id: 3, icon: UserFilled, color: '#FF9800', text: '新用户 student123 注册', time: '1小时前' },
  { id: 4, icon: Delete, color: '#F56C6C', text: 'admin 删除了试卷《旧版C语言》', time: '2小时前' },
])

const pendingItems = ref([
  { id: 1, title: '用户反馈待处理', desc: '3条新反馈等待回复' },
  { id: 2, title: '试卷审核', desc: '2份试卷等待审核' },
])

const pendingCount = computed(() => pendingItems.value.length)

const fetchStats = async () => {
  try {
    const res = await api.get('/api/stats')
    if (res.data.code === '2000') {
      stats.value = res.data.data
    }
  } catch (e) {
    console.error('获取统计失败', e)
  }
}

const handlePending = (id) => {
  pendingItems.value = pendingItems.value.filter(item => item.id !== id)
}

onMounted(fetchStats)
</script>

<style scoped>
.admin-dashboard {
  padding: 32px;
  max-width: 1400px;
  margin: 0 auto;
}

.page-header {
  margin-bottom: 32px;
}

.page-title {
  font-size: 28px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 8px 0;
}

.page-subtitle {
  font-size: 14px;
  color: #666;
  margin: 0;
}

/* 统计卡片 */
.stats-row {
  margin-bottom: 24px;
}

.stat-card {
  background: #fff;
  border-radius: 12px;
  padding: 24px;
  display: flex;
  gap: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  transition: all 0.3s;
  border-left: 4px solid transparent;
}

.stat-card.primary { border-left-color: #3A7AFE; }
.stat-card.success { border-left-color: #00C2A8; }
.stat-card.warning { border-left-color: #FF9800; }
.stat-card.info { border-left-color: #909399; }

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
}

.stat-icon {
  width: 64px;
  height: 64px;
  background: #f8f9fa;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 32px;
  font-weight: 600;
  color: #1a1a1a;
  line-height: 1;
  margin-bottom: 8px;
}

.stat-label {
  font-size: 14px;
  color: #666;
  margin-bottom: 8px;
}

.stat-trend {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 13px;
  color: #00C2A8;
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
  margin-bottom: 20px;
}

.section-title {
  font-size: 18px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0;
}

/* 管理入口 */
.manage-item {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 24px 16px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s;
  border: 1px solid transparent;
}

.manage-item:hover {
  background: #fff;
  border-color: #3A7AFE;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(58, 122, 254, 0.1);
}

.manage-icon {
  width: 56px;
  height: 56px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 12px;
}

.manage-text {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

/* 活动列表 */
.activity-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.activity-item {
  display: flex;
  gap: 12px;
  padding: 12px;
  border-radius: 6px;
  background: #fafafa;
  transition: background 0.3s;
}

.activity-item:hover {
  background: #f0f0f0;
}

.activity-icon {
  width: 32px;
  height: 32px;
  background: #fff;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.activity-content {
  flex: 1;
}

.activity-text {
  font-size: 14px;
  color: #333;
  margin-bottom: 4px;
}

.activity-time {
  font-size: 12px;
  color: #999;
}

/* 待处理列表 */
.pending-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.pending-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  background: #fff7e6;
  border-radius: 8px;
  border-left: 3px solid #FF9800;
}

.pending-content {
  flex: 1;
}

.pending-title {
  font-size: 14px;
  font-weight: 500;
  color: #333;
  margin-bottom: 4px;
}

.pending-desc {
  font-size: 13px;
  color: #999;
}

.badge {
  margin-right: 8px;
}

@media (max-width: 768px) {
  .admin-dashboard {
    padding: 16px;
  }
  
  .stats-row {
    gap: 12px;
  }
  
  .stat-card {
    margin-bottom: 12px;
  }
}
</style> 