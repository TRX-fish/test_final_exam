<template>
  <div class="history-container">
    <div class="page-header">
      <div>
        <h2 class="page-title">学习记录</h2>
        <p class="page-subtitle">最近7天学习了 {{ todayCount }} 份试卷</p>
      </div>
      <el-button @click="clearHistory">
        <el-icon><Delete /></el-icon>
        清空记录
      </el-button>
    </div>

    <!-- 统计卡片 -->
    <el-row :gutter="16" class="stats-row">
      <el-col :span="8" :xs="24">
        <div class="stat-box">
          <div class="stat-icon" style="background: #E8F1FF;">
            <el-icon :size="24" color="#3A7AFE"><TrendCharts /></el-icon>
          </div>
          <div>
            <div class="stat-number">{{ totalDays }}</div>
            <div class="stat-label">累计学习天数</div>
          </div>
        </div>
      </el-col>
      <el-col :span="8" :xs="24">
        <div class="stat-box">
          <div class="stat-icon" style="background: #E7F9F6;">
            <el-icon :size="24" color="#00C2A8"><Document /></el-icon>
          </div>
          <div>
            <div class="stat-number">{{ totalPapers }}</div>
            <div class="stat-label">浏览试卷总数</div>
          </div>
        </div>
      </el-col>
      <el-col :span="8" :xs="24">
        <div class="stat-box">
          <div class="stat-icon" style="background: #FFF5E6;">
            <el-icon :size="24" color="#FF9800"><Clock /></el-icon>
          </div>
          <div>
            <div class="stat-number">{{ todayCount }}</div>
            <div class="stat-label">今日学习数</div>
          </div>
        </div>
      </el-col>
    </el-row>

    <!-- 按日期分组的历史记录 -->
    <div class="history-timeline">
      <div v-for="group in groupedHistory" :key="group.date" class="date-group">
        <div class="date-label">{{ group.date }}</div>
        <div class="records-list">
          <div 
            v-for="record in group.records" 
            :key="record.id"
            class="record-item"
            @click="$router.push(`/papers/${record.paperId}`)"
          >
            <div class="record-icon">
              <el-icon :size="20" color="#3A7AFE"><Document /></el-icon>
            </div>
            <div class="record-content">
              <h4 class="record-title">{{ record.paperName }}</h4>
              <div class="record-meta">
                <span>{{ record.college }}</span>
                <span>·</span>
                <span>{{ record.year }}年{{ record.term }}</span>
              </div>
            </div>
            <div class="record-time">
              {{ record.time }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 空状态 -->
    <el-empty v-if="history.length === 0" description="还没有学习记录">
      <el-button type="primary" @click="$router.push('/papers')">
        开始学习
      </el-button>
    </el-empty>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Delete, TrendCharts, Document, Clock } from '@element-plus/icons-vue'

const router = useRouter()

// 模拟学习记录数据
const history = ref([
  { id: 1, paperId: 2, paperName: '大学英语', college: '外国语学院', year: 2022, term: '春季', date: '2025-10-02', time: '14:30' },
  { id: 2, paperId: 1, paperName: '高等数学', college: '理学院', year: 2023, term: '秋季', date: '2025-10-02', time: '10:15' },
  { id: 3, paperId: 3, paperName: '线性代数', college: '理学院', year: 2023, term: '春季', date: '2025-10-01', time: '16:20' },
  { id: 4, paperId: 4, paperName: 'C语言程序设计', college: '计算机学院', year: 2021, term: '秋季', date: '2025-10-01', time: '09:30' },
])

const totalDays = ref(7)
const totalPapers = ref(15)
const todayCount = computed(() => {
  const today = new Date().toISOString().split('T')[0]
  return history.value.filter(h => h.date === today).length
})

const groupedHistory = computed(() => {
  const groups = {}
  history.value.forEach(record => {
    const dateLabel = formatDateLabel(record.date)
    if (!groups[dateLabel]) {
      groups[dateLabel] = []
    }
    groups[dateLabel].push(record)
  })
  
  return Object.keys(groups).map(date => ({
    date,
    records: groups[date]
  }))
})

const formatDateLabel = (dateStr) => {
  const date = new Date(dateStr)
  const today = new Date()
  const yesterday = new Date(today)
  yesterday.setDate(yesterday.getDate() - 1)
  
  if (dateStr === today.toISOString().split('T')[0]) {
    return '今天'
  } else if (dateStr === yesterday.toISOString().split('T')[0]) {
    return '昨天'
  } else {
    return dateStr
  }
}

const clearHistory = () => {
  ElMessageBox.confirm('确定要清空所有学习记录吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    history.value = []
    ElMessage.success('学习记录已清空')
  }).catch(() => {})
}
</script>

<style scoped>
.history-container {
  padding: 32px;
  max-width: 1000px;
  margin: 0 auto;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
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
  margin-bottom: 32px;
}

.stat-box {
  background: #fff;
  border-radius: 8px;
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 16px;
  border: 1px solid #f0f0f0;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.stat-number {
  font-size: 24px;
  font-weight: 600;
  color: #1a1a1a;
  line-height: 1;
  margin-bottom: 4px;
}

.stat-label {
  font-size: 13px;
  color: #666;
}

/* 时间轴 */
.history-timeline {
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.date-group {
  position: relative;
}

.date-label {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a1a;
  margin-bottom: 16px;
  padding-left: 24px;
  position: relative;
}

.date-label::before {
  content: '';
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 8px;
  height: 8px;
  background: #3A7AFE;
  border-radius: 50%;
}

.records-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  padding-left: 24px;
}

.record-item {
  background: #fff;
  border-radius: 8px;
  padding: 16px 20px;
  display: flex;
  align-items: center;
  gap: 16px;
  border: 1px solid #f0f0f0;
  cursor: pointer;
  transition: all 0.3s;
}

.record-item:hover {
  border-color: #3A7AFE;
  box-shadow: 0 2px 8px rgba(58, 122, 254, 0.08);
  transform: translateX(4px);
}

.record-icon {
  width: 40px;
  height: 40px;
  background: #E8F1FF;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.record-content {
  flex: 1;
}

.record-title {
  font-size: 15px;
  font-weight: 500;
  color: #1a1a1a;
  margin: 0 0 6px 0;
}

.record-meta {
  font-size: 13px;
  color: #999;
  display: flex;
  gap: 8px;
}

.record-time {
  font-size: 13px;
  color: #999;
  white-space: nowrap;
}

@media (max-width: 768px) {
  .history-container {
    padding: 16px;
  }
  
  .stats-row {
    gap: 12px;
  }
  
  .stat-box {
    margin-bottom: 12px;
  }
  
  .record-item {
    flex-wrap: wrap;
  }
}
</style> 