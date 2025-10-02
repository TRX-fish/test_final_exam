<template>
  <div class="paper-browse-container">
    <el-row :gutter="24">
      <!-- 左侧筛选区 -->
      <el-col :span="6" :xs="24" :sm="8" :md="6" class="filter-col">
        <div class="filter-card">
          <div class="filter-header">
            <h3 class="filter-title">筛选条件</h3>
            <el-button type="primary" link size="small" @click="resetSearch">
              重置
            </el-button>
          </div>

          <el-form :model="search" label-position="top">
            <el-form-item label="课程名称">
              <el-input 
                v-model="search.course" 
                placeholder="搜索课程"
                clearable
                :prefix-icon="Search"
              />
            </el-form-item>

            <el-form-item label="考试年份">
              <el-select v-model="search.year" placeholder="选择年份" clearable style="width: 100%;">
                <el-option 
                  v-for="year in yearOptions" 
                  :key="year" 
                  :label="`${year}年`" 
                  :value="year" 
                />
              </el-select>
            </el-form-item>

            <el-form-item label="学期">
              <el-select v-model="search.term" placeholder="选择学期" clearable style="width: 100%;">
                <el-option label="春季" value="春季" />
                <el-option label="秋季" value="秋季" />
                <el-option label="夏季" value="夏季" />
              </el-select>
            </el-form-item>

            <el-form-item label="所属学院">
              <el-input 
                v-model="search.college" 
                placeholder="搜索学院"
                clearable
              />
            </el-form-item>

            <el-button 
              type="primary" 
              size="large"
              style="width: 100%; margin-top: 8px;"
              @click="fetchPapers"
            >
              <el-icon><Search /></el-icon>
              搜索试卷
            </el-button>
          </el-form>
        </div>
      </el-col>

      <!-- 右侧内容区 -->
      <el-col :span="18" :xs="24" :sm="16" :md="18">
        <div class="content-header">
          <div>
            <h2 class="content-title">试卷列表</h2>
            <p class="content-subtitle">共找到 {{ total }} 份试卷</p>
          </div>
          <el-button 
            v-if="isAdmin"
            type="primary" 
            @click="$router.push('/papers/add')"
          >
            <el-icon><Plus /></el-icon>
            添加试卷
          </el-button>
        </div>

        <!-- 空状态 -->
        <el-empty v-if="papers.length === 0" description="暂无试卷数据">
          <el-button type="primary" @click="$router.push('/papers/add')">
            添加第一份试卷
          </el-button>
        </el-empty>

        <!-- 试卷卡片网格 -->
        <div v-else class="papers-grid">
          <div 
            v-for="paper in papers" 
            :key="paper.id" 
            class="paper-item"
            @click="goDetail(paper)"
          >
            <div class="paper-content">
              <div class="paper-top">
                <h4 class="paper-name">{{ paper.course }}</h4>
                <el-tag size="small" type="primary">{{ paper.year }}年</el-tag>
              </div>
              
              <div class="paper-info">
                <div class="info-row">
                  <el-icon :size="14" color="#666"><School /></el-icon>
                  <span>{{ paper.college || '未分类' }}</span>
                </div>
                <div class="info-row">
                  <el-icon :size="14" color="#666"><Calendar /></el-icon>
                  <span>{{ paper.term || '未指定' }}</span>
                </div>
              </div>

              <div class="paper-bottom">
                <span class="time-text">{{ formatDate(paper.created_at) }}</span>
                <el-icon class="enter-icon"><ArrowRight /></el-icon>
              </div>
            </div>
          </div>
        </div>

        <!-- 分页器 -->
        <div class="pagination-wrapper" v-if="total > 0">
          <el-pagination
            background
            layout="prev, pager, next, jumper"
            :total="total"
            :page-size="pageSize"
            :current-page="page"
            @current-change="handlePageChange"
          />
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import api from '../api/request'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { Search, School, Calendar, Plus, ArrowRight } from '@element-plus/icons-vue'

const router = useRouter()
const papers = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = 12
const search = ref({ course: '', year: '', term: '', college: '' })

// 生成年份选项（最近10年）
const currentYear = new Date().getFullYear()
const yearOptions = Array.from({ length: 10 }, (_, i) => currentYear - i)

// 检查是否是管理员
const isAdmin = computed(() => localStorage.getItem('userRole') === 'admin')

const fetchPapers = async () => {
  try {
    const params = { ...search.value, page: page.value, per_page: pageSize }
    const res = await api.get('/api/papers', { params })
    if (res.data.code === '2000') {
      papers.value = res.data.data.items
      total.value = res.data.data.total
    } else {
      papers.value = []
      total.value = 0
    }
  } catch (e) {
    papers.value = []
    total.value = 0
    ElMessage.error('获取试卷失败')
  }
}

const handlePageChange = (val) => {
  page.value = val
  fetchPapers()
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const resetSearch = () => {
  search.value = { course: '', year: '', term: '', college: '' }
  page.value = 1
  fetchPapers()
}

const goDetail = (row) => {
  router.push(`/papers/${row.id}`)
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

onMounted(fetchPapers)
</script>

<style scoped>
.paper-browse-container {
  padding: 32px;
  max-width: 1400px;
  margin: 0 auto;
}

/* 左侧筛选卡片 */
.filter-card {
  background: #fff;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  position: sticky;
  top: 24px;
}

.filter-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.filter-title {
  font-size: 18px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0;
}

.filter-card :deep(.el-form-item) {
  margin-bottom: 20px;
}

.filter-card :deep(.el-form-item__label) {
  font-size: 14px;
  font-weight: 500;
  color: #333;
  line-height: 1.5;
  padding-bottom: 8px;
}

.filter-card :deep(.el-input__wrapper) {
  border-radius: 8px;
  box-shadow: 0 0 0 1px #e5e5e5 inset;
}

.filter-card :deep(.el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px #3A7AFE inset;
}

/* 内容头部 */
.content-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 24px;
}

.content-title {
  font-size: 24px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 4px 0;
  line-height: 1.2;
}

.content-subtitle {
  font-size: 14px;
  color: #666;
  margin: 0;
  line-height: 1.5;
}

/* 试卷网格 */
.papers-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
  margin-bottom: 32px;
}

.paper-item {
  background: #fff;
  border-radius: 8px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.3s;
  border: 1px solid #f0f0f0;
}

.paper-item:hover {
  border-color: #3A7AFE;
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(58, 122, 254, 0.12);
}

.paper-content {
  padding: 20px;
}

.paper-top {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 16px;
}

.paper-name {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0;
  line-height: 1.5;
  flex: 1;
  margin-right: 8px;
}

.paper-info {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 16px;
}

.info-row {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #666;
}

.paper-bottom {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 12px;
  border-top: 1px solid #f0f0f0;
}

.time-text {
  font-size: 12px;
  color: #999;
}

.enter-icon {
  color: #3A7AFE;
  transition: transform 0.3s;
}

.paper-item:hover .enter-icon {
  transform: translateX(4px);
}

/* 分页器 */
.pagination-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 32px;
}

.pagination-wrapper :deep(.el-pagination) {
  gap: 8px;
}

.pagination-wrapper :deep(.el-pager li) {
  border-radius: 6px;
  min-width: 32px;
  height: 32px;
  line-height: 32px;
}

/* 响应式 */
@media (max-width: 992px) {
  .filter-col {
    order: 2;
    margin-top: 24px;
  }
  
  .filter-card {
    position: static;
  }
}

@media (max-width: 768px) {
  .paper-browse-container {
    padding: 16px;
  }
  
  .papers-grid {
    grid-template-columns: 1fr;
  }
  
  .content-header {
    flex-direction: column;
    gap: 16px;
  }
}
</style> 