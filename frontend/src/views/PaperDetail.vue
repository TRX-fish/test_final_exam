<template>
  <el-row justify="center" style="margin-top: 40px;">
    <el-col :span="20">
      <el-card>
        <template #header>
          <div style="display: flex; justify-content: space-between; align-items: center;">
            <h2 style="margin: 0;">试卷详情</h2>
            <div>
              <el-button 
                v-if="isAdmin"
                type="primary" 
                @click="goAddQuestion"
              >
                添加题目
              </el-button>
              <el-button @click="$router.back()">返回</el-button>
            </div>
          </div>
        </template>
        
        <!-- 加载状态 -->
        <el-skeleton v-if="loading" :rows="5" animated />

        <!-- 试卷信息 -->
        <el-descriptions v-else-if="paper && paper.id" :column="2" border>
          <el-descriptions-item label="课程名称">{{ paper.course }}</el-descriptions-item>
          <el-descriptions-item label="考试年份">{{ paper.year }}</el-descriptions-item>
          <el-descriptions-item label="学期">{{ paper.term }}</el-descriptions-item>
          <el-descriptions-item label="所属学院">{{ paper.college }}</el-descriptions-item>
          <el-descriptions-item label="录入时间" :span="2">{{ paper.created_at }}</el-descriptions-item>
          <el-descriptions-item label="文件类型" :span="2">
            <el-tag :type="paper.file_type === 'pdf' ? 'success' : 'info'">
              {{ paper.file_type === 'pdf' ? 'PDF文件' : '图片' }}
            </el-tag>
          </el-descriptions-item>
        </el-descriptions>

        <!-- PDF预览（如果是PDF文件） -->
        <div v-if="paper && paper.file_url && paper.file_type === 'pdf'" style="margin-top: 24px;">
          <el-divider content-position="left">
            <h3>试卷预览</h3>
          </el-divider>
          <PDFViewer :pdfUrl="paper.file_url" :isMember="isMember" :fileSize="paper.file_size" />
        </div>

        <!-- 图片预览（如果是图片） -->
        <div v-else-if="paper && paper.image_path && paper.file_type === 'image'" style="margin-top: 24px;">
          <el-divider content-position="left">
            <h3>试卷图片</h3>
          </el-divider>
          <el-image :src="paper.image_path" style="max-width: 100%;" fit="contain" />
        </div>

        <!-- 题目列表 -->
        <el-divider content-position="left">
          <h3>题目列表 (共 {{ questions.length }} 题)</h3>
        </el-divider>

        <el-empty v-if="questions.length === 0" description="该试卷暂无题目">
          <el-button v-if="isAdmin" type="primary" @click="goAddQuestion">立即添加</el-button>
        </el-empty>

        <div v-else>
          <!-- 非会员提示 -->
          <el-alert 
            v-if="!isMember && questions.length > 3" 
            type="info" 
            :closable="false"
            style="margin-bottom: 16px;"
          >
            <template #title>
              您当前是免费用户，仅可查看前3道题目
            </template>
            <el-button type="primary" size="small" @click="$router.push('/membership')">
              开通会员查看全部 {{ questions.length }} 题
            </el-button>
          </el-alert>

          <el-card v-for="(q, index) in displayQuestions" :key="q.id" style="margin-bottom: 20px;" shadow="hover">
            <template #header>
              <div style="display: flex; justify-content: space-between; align-items: center;">
                <span><strong>第 {{ index + 1 }} 题</strong> [{{ q.type }}]</span>
                <el-button type="primary" link @click="editQuestion(q.id)">编辑</el-button>
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
                <el-image :src="`http://47.93.252.105:5000/static/${q.image_path}`" style="width: 200px;" />
              </div>
            </div>
          </el-card>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import api from '../api/request'
import { ElMessage } from 'element-plus'
import PDFViewer from '../components/PDFViewer.vue'

const router = useRouter()
const route = useRoute()
const paper = ref(null)
const questions = ref([])
const paperId = route.params.id
const loading = ref(true)

// 检查会员状态
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

// 检查是否是管理员
const isAdmin = computed(() => localStorage.getItem('userRole') === 'admin')

// 显示的题目列表（非会员只显示前3题）
const displayQuestions = computed(() => {
  if (isMember.value || isAdmin.value) {
    return questions.value
  }
  return questions.value.slice(0, 3)
})

const fetchPaper = async () => {
  loading.value = true
  try {
    const res = await api.get(`/api/papers?page=1&per_page=100`)
    if (res.data.code === '2000') {
      const paperData = res.data.data.items.find(p => p.id == paperId)
      paper.value = paperData || {}
    }
  } catch (e) {
    ElMessage.error('获取试卷信息失败')
    paper.value = {}
  } finally {
    loading.value = false
  }
}

const fetchQuestions = async () => {
  try {
    const res = await api.get('/api/questions', { params: { paper_id: paperId } })
    if (res.data.code === '2000') {
      questions.value = res.data.data || []
    } else {
      questions.value = []
    }
  } catch (e) {
    questions.value = []
  }
}

const parseOptions = (options) => {
  try {
    return JSON.parse(options)
  } catch {
    return []
  }
}

const goAddQuestion = () => {
  router.push(`/papers/${paperId}/questions/add`)
}

const editQuestion = (questionId) => {
  router.push(`/questions/${questionId}/edit`)
}

onMounted(() => {
  fetchPaper()
  fetchQuestions()
})
</script> 