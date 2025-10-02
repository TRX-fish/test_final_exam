<template>
  <div class="favorites-container">
    <div class="page-header">
      <div>
        <h2 class="page-title">我的收藏</h2>
        <p class="page-subtitle">共收藏 {{ favorites.length }} 份试卷</p>
      </div>
      <el-button-group>
        <el-button :type="viewMode === 'grid' ? 'primary' : ''" @click="viewMode = 'grid'">
          <el-icon><Grid /></el-icon>
        </el-button>
        <el-button :type="viewMode === 'list' ? 'primary' : ''" @click="viewMode = 'list'">
          <el-icon><List /></el-icon>
        </el-button>
      </el-button-group>
    </div>

    <!-- 排序和筛选 -->
    <div class="toolbar">
      <el-select v-model="sortBy" placeholder="排序方式" style="width: 180px;">
        <el-option label="最近收藏" value="recent" />
        <el-option label="课程名称" value="course" />
        <el-option label="考试年份" value="year" />
      </el-select>
    </div>

    <!-- 空状态 -->
    <el-empty v-if="favorites.length === 0" description="还没有收藏任何试卷">
      <el-button type="primary" @click="$router.push('/papers')">
        去浏览试卷
      </el-button>
    </el-empty>

    <!-- 收藏列表 -->
    <div v-else>
      <!-- 网格视图 -->
      <div v-if="viewMode === 'grid'" class="favorites-grid">
        <div 
          v-for="item in sortedFavorites" 
          :key="item.id" 
          class="favorite-card"
        >
          <div class="card-header">
            <h4 class="card-title">{{ item.course }}</h4>
            <el-button 
              type="danger" 
              link 
              :icon="Delete"
              @click.stop="removeFavorite(item.id)"
            />
          </div>
          <div class="card-body" @click="$router.push(`/papers/${item.id}`)">
            <el-tag size="small">{{ item.year }}年 {{ item.term }}</el-tag>
            <p class="card-college">{{ item.college }}</p>
            <div class="card-footer">
              <span class="time">收藏于 {{ item.favoriteTime }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 列表视图 -->
      <div v-else class="favorites-list">
        <div 
          v-for="item in sortedFavorites" 
          :key="item.id" 
          class="list-item"
          @click="$router.push(`/papers/${item.id}`)"
        >
          <div class="list-content">
            <h4 class="list-title">{{ item.course }}</h4>
            <div class="list-meta">
              <span>{{ item.year }}年</span>
              <span>{{ item.term }}</span>
              <span>{{ item.college }}</span>
            </div>
          </div>
          <div class="list-actions">
            <span class="list-time">{{ item.favoriteTime }}</span>
            <el-button 
              type="danger" 
              link 
              :icon="Delete"
              @click.stop="removeFavorite(item.id)"
            >
              取消收藏
            </el-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Grid, List, Delete } from '@element-plus/icons-vue'

const router = useRouter()
const viewMode = ref('grid')
const sortBy = ref('recent')

// 模拟收藏数据（实际应从后端获取）
const favorites = ref([
  { id: 1, course: '高等数学', year: 2023, term: '秋季', college: '理学院', favoriteTime: '2天前' },
  { id: 2, course: '大学英语', year: 2022, term: '春季', college: '外国语学院', favoriteTime: '5天前' },
])

const sortedFavorites = computed(() => {
  const list = [...favorites.value]
  if (sortBy.value === 'course') {
    return list.sort((a, b) => a.course.localeCompare(b.course, 'zh-CN'))
  } else if (sortBy.value === 'year') {
    return list.sort((a, b) => b.year - a.year)
  }
  return list
})

const removeFavorite = (id) => {
  ElMessageBox.confirm('确定要取消收藏这份试卷吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    favorites.value = favorites.value.filter(item => item.id !== id)
    ElMessage.success('已取消收藏')
  }).catch(() => {})
}
</script>

<style scoped>
.favorites-container {
  padding: 32px;
  max-width: 1400px;
  margin: 0 auto;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 24px;
}

.page-title {
  font-size: 28px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 8px 0;
  line-height: 1.2;
}

.page-subtitle {
  font-size: 14px;
  color: #666;
  margin: 0;
}

.toolbar {
  margin-bottom: 24px;
}

/* 网格视图 */
.favorites-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
}

.favorite-card {
  background: #fff;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid #f0f0f0;
  transition: all 0.3s;
}

.favorite-card:hover {
  border-color: #3A7AFE;
  box-shadow: 0 4px 16px rgba(58, 122, 254, 0.1);
  transform: translateY(-2px);
}

.card-header {
  padding: 16px 20px;
  border-bottom: 1px solid #f0f0f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0;
}

.card-body {
  padding: 20px;
  cursor: pointer;
}

.card-college {
  font-size: 14px;
  color: #666;
  margin: 12px 0;
}

.card-footer {
  padding-top: 12px;
  border-top: 1px solid #f5f5f5;
}

.time {
  font-size: 12px;
  color: #999;
}

/* 列表视图 */
.favorites-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.list-item {
  background: #fff;
  border-radius: 8px;
  padding: 20px 24px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border: 1px solid #f0f0f0;
  cursor: pointer;
  transition: all 0.3s;
}

.list-item:hover {
  border-color: #3A7AFE;
  box-shadow: 0 2px 8px rgba(58, 122, 254, 0.08);
}

.list-content {
  flex: 1;
}

.list-title {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 8px 0;
}

.list-meta {
  display: flex;
  gap: 16px;
  font-size: 13px;
  color: #666;
}

.list-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.list-time {
  font-size: 12px;
  color: #999;
}

@media (max-width: 768px) {
  .favorites-container {
    padding: 16px;
  }
  
  .favorites-grid {
    grid-template-columns: 1fr;
  }
  
  .list-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
}
</style> 