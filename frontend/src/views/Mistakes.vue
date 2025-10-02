<template>
  <div class="mistakes-container">
    <div class="page-header">
      <div>
        <h2 class="page-title">我的错题本</h2>
        <p class="page-subtitle">共有 {{ mistakes.length }} 道错题需要复习</p>
      </div>
      <el-button type="primary" @click="startReview" :disabled="mistakes.length === 0">
        <el-icon><RefreshRight /></el-icon>
        开始复习
      </el-button>
    </div>

    <!-- 筛选工具栏 -->
    <div class="toolbar">
      <el-radio-group v-model="filterType" @change="handleFilterChange">
        <el-radio-button label="">全部</el-radio-button>
        <el-radio-button label="单选">单选题</el-radio-button>
        <el-radio-button label="多选">多选题</el-radio-button>
        <el-radio-button label="判断">判断题</el-radio-button>
        <el-radio-button label="简答">简答题</el-radio-button>
      </el-radio-group>
      <el-select v-model="sortBy" placeholder="排序方式" style="width: 160px;">
        <el-option label="最近错误" value="recent" />
        <el-option label="错误次数" value="count" />
      </el-select>
    </div>

    <!-- 空状态 -->
    <el-empty v-if="filteredMistakes.length === 0" description="太棒了！暂无错题">
      <el-button type="primary" @click="$router.push('/papers')">
        去做题
      </el-button>
    </el-empty>

    <!-- 错题列表 -->
    <div v-else class="mistakes-list">
      <div 
        v-for="(mistake, index) in filteredMistakes" 
        :key="mistake.id" 
        class="mistake-card"
      >
        <div class="card-header">
          <div class="header-left">
            <el-tag :type="getTypeColor(mistake.type)" size="small">{{ mistake.type }}</el-tag>
            <span class="question-num">第 {{ index + 1 }} 题</span>
            <el-tag type="danger" size="small" v-if="mistake.wrongCount > 1">
              错 {{ mistake.wrongCount }} 次
            </el-tag>
          </div>
          <div class="header-right">
            <el-button 
              type="success" 
              link 
              size="small"
              @click="markAsMastered(mistake.id)"
            >
              <el-icon><Check /></el-icon>
              已掌握
            </el-button>
          </div>
        </div>

        <div class="card-body">
          <div class="question-content">
            <p class="question-text">{{ mistake.content }}</p>
            <ul v-if="mistake.options" class="options-list">
              <li v-for="(opt, i) in parseOptions(mistake.options)" :key="i">
                {{ opt }}
              </li>
            </ul>
          </div>

          <div class="answer-section">
            <div class="answer-row">
              <span class="label">我的答案：</span>
              <el-tag type="danger" size="small">{{ mistake.myAnswer }}</el-tag>
            </div>
            <div class="answer-row">
              <span class="label">正确答案：</span>
              <el-tag type="success" size="small">{{ mistake.correctAnswer }}</el-tag>
            </div>
          </div>

          <el-collapse class="explanation-collapse">
            <el-collapse-item title="查看解析" name="1">
              <div class="explanation-content">
                {{ mistake.explanation || '暂无解析' }}
              </div>
            </el-collapse-item>
          </el-collapse>

          <div class="card-footer">
            <span class="footer-text">来自《{{ mistake.paperName }}》</span>
            <span class="footer-time">{{ mistake.wrongTime }}</span>
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
import { RefreshRight, Check } from '@element-plus/icons-vue'

const router = useRouter()
const filterType = ref('')
const sortBy = ref('recent')

// 模拟错题数据
const mistakes = ref([
  {
    id: 1,
    type: '单选',
    content: 'What is the capital of England?',
    options: '["A. London", "B. Paris", "C. Berlin", "D. Rome"]',
    myAnswer: 'B',
    correctAnswer: 'A',
    explanation: 'London is the capital of England.',
    paperName: '大学英语 2022春季',
    wrongTime: '2天前',
    wrongCount: 2
  },
  {
    id: 2,
    type: '单选',
    content: 'Choose the correct form: "I _____ to school every day."',
    options: '["A. go", "B. goes", "C. going", "D. went"]',
    myAnswer: 'B',
    correctAnswer: 'A',
    explanation: 'First person singular present tense uses the base form "go".',
    paperName: '大学英语 2022春季',
    wrongTime: '3天前',
    wrongCount: 1
  }
])

const filteredMistakes = computed(() => {
  let list = [...mistakes.value]
  if (filterType.value) {
    list = list.filter(m => m.type === filterType.value)
  }
  if (sortBy.value === 'count') {
    list.sort((a, b) => b.wrongCount - a.wrongCount)
  }
  return list
})

const handleFilterChange = () => {
  // 筛选变化逻辑
}

const parseOptions = (options) => {
  try {
    return JSON.parse(options)
  } catch {
    return []
  }
}

const getTypeColor = (type) => {
  const colors = {
    '单选': 'primary',
    '多选': 'success',
    '判断': 'warning',
    '填空': 'info',
    '简答': ''
  }
  return colors[type] || ''
}

const startReview = () => {
  ElMessage.info('复习模式开发中...')
}

const markAsMastered = (id) => {
  ElMessageBox.confirm('确认已经掌握此题吗？', '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'success'
  }).then(() => {
    mistakes.value = mistakes.value.filter(m => m.id !== id)
    ElMessage.success('已标记为掌握')
  }).catch(() => {})
}
</script>

<style scoped>
.mistakes-container {
  padding: 32px;
  max-width: 1000px;
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
}

.page-subtitle {
  font-size: 14px;
  color: #666;
  margin: 0;
}

.toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  gap: 16px;
}

.mistakes-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.mistake-card {
  background: #fff;
  border-radius: 8px;
  border: 1px solid #f0f0f0;
  overflow: hidden;
  transition: all 0.3s;
}

.mistake-card:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.08);
}

.card-header {
  padding: 16px 20px;
  background: #fafafa;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.question-num {
  font-size: 14px;
  font-weight: 500;
  color: #333;
}

.card-body {
  padding: 24px;
}

.question-content {
  margin-bottom: 20px;
}

.question-text {
  font-size: 15px;
  color: #1a1a1a;
  line-height: 1.8;
  margin: 0 0 12px 0;
}

.options-list {
  list-style: none;
  padding: 0;
  margin: 12px 0;
}

.options-list li {
  padding: 8px 0;
  font-size: 14px;
  color: #666;
  line-height: 1.6;
}

.answer-section {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 16px;
}

.answer-row {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.answer-row:last-child {
  margin-bottom: 0;
}

.label {
  font-size: 14px;
  color: #666;
  font-weight: 500;
}

.explanation-collapse {
  margin-bottom: 16px;
  border: none;
}

.explanation-collapse :deep(.el-collapse-item__header) {
  background: #f8f9fa;
  padding: 12px 16px;
  border-radius: 8px;
  border: none;
  font-size: 14px;
  color: #3A7AFE;
}

.explanation-content {
  padding: 16px;
  background: #f8f9fa;
  border-radius: 8px;
  font-size: 14px;
  color: #666;
  line-height: 1.8;
  margin-top: 8px;
}

.card-footer {
  display: flex;
  justify-content: space-between;
  padding-top: 16px;
  border-top: 1px solid #f0f0f0;
}

.footer-text {
  font-size: 13px;
  color: #666;
}

.footer-time {
  font-size: 12px;
  color: #999;
}

@media (max-width: 768px) {
  .mistakes-container {
    padding: 16px;
  }
  
  .toolbar {
    flex-direction: column;
    align-items: stretch;
  }
  
  .toolbar :deep(.el-radio-group) {
    display: flex;
    flex-wrap: wrap;
  }
}
</style> 