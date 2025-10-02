<template>
  <el-row justify="center" style="margin-top: 40px;">
    <el-col :span="20">
      <el-card>
        <h2 style="text-align:center;">题目搜索</h2>
        <el-form :inline="true" :model="search" style="margin-bottom: 20px;">
          <el-form-item label="关键词">
            <el-input v-model="search.keyword" placeholder="搜索题目内容" clearable style="width: 200px;" />
          </el-form-item>
          <el-form-item label="题目类型">
            <el-select v-model="search.type" placeholder="选择类型" clearable style="width: 150px;">
              <el-option label="单选" value="单选" />
              <el-option label="多选" value="多选" />
              <el-option label="判断" value="判断" />
              <el-option label="填空" value="填空" />
              <el-option label="简答" value="简答" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="fetchQuestions">搜索</el-button>
            <el-button @click="resetSearch">重置</el-button>
          </el-form-item>
        </el-form>

        <el-empty v-if="questions.length === 0" description="暂无题目数据" />

        <div v-else>
          <el-card v-for="(q, index) in questions" :key="q.id" style="margin-bottom: 20px;" shadow="hover">
            <template #header>
              <div style="display: flex; justify-content: space-between; align-items: center;">
                <span>
                  <el-tag type="info">{{ q.type }}</el-tag>
                  <span style="margin-left: 10px;">题目 ID: {{ q.id }}</span>
                </span>
                <el-button type="primary" link @click="viewPaper(q.paper_id)">查看试卷</el-button>
              </div>
            </template>
            <div>
              <p><strong>题目：</strong>{{ q.content }}</p>
              <p v-if="q.options"><strong>选项：</strong></p>
              <ul v-if="q.options">
                <li v-for="(opt, i) in parseOptions(q.options)" :key="i">{{ opt }}</li>
              </ul>
              <p><strong>答案：</strong><el-tag type="success">{{ q.answer }}</el-tag></p>
              <p v-if="q.explanation"><strong>解析：</strong>{{ q.explanation }}</p>
              <div v-if="q.image_path">
                <p><strong>题目图片：</strong></p>
                <el-image :src="`http://47.93.252.105:5000/static/${q.image_path}`" style="width: 200px;" fit="cover" />
              </div>
              <p style="color: #999; font-size: 12px; margin-top: 10px;">创建时间: {{ q.created_at }}</p>
            </div>
          </el-card>

          <el-pagination
            style="margin-top: 20px; text-align: right;"
            background
            layout="prev, pager, next, jumper, ->, total"
            :total="total"
            :page-size="pageSize"
            :current-page="page"
            @current-change="handlePageChange"
          />
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '../api/request'
import { ElMessage } from 'element-plus'

const router = useRouter()
const questions = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = 10
const search = ref({ keyword: '', type: '' })

const fetchQuestions = async () => {
  try {
    const params = { 
      keyword: search.value.keyword,
      type: search.value.type,
      page: page.value, 
      per_page: pageSize 
    }
    const res = await api.get('/api/questions/search', { params })
    if (res.data.code === '2000') {
      questions.value = res.data.data.items || []
      total.value = res.data.data.total || 0
    } else {
      questions.value = []
      total.value = 0
    }
  } catch (e) {
    questions.value = []
    total.value = 0
    ElMessage.error('获取题目失败')
  }
}

const handlePageChange = (val) => {
  page.value = val
  fetchQuestions()
}

const resetSearch = () => {
  search.value = { keyword: '', type: '' }
  page.value = 1
  fetchQuestions()
}

const parseOptions = (options) => {
  try {
    return JSON.parse(options)
  } catch {
    return []
  }
}

const viewPaper = (paperId) => {
  router.push(`/papers/${paperId}`)
}

onMounted(fetchQuestions)
</script> 