<template>
  <el-row justify="center" style="margin-top: 40px;">
    <el-col :span="20">
      <el-card>
        <h2 style="text-align:center;">试卷列表</h2>
        <el-form :inline="true" :model="search" style="margin-bottom: 20px;">
          <el-form-item label="课程名">
            <el-input v-model="search.course" placeholder="输入课程名" clearable />
          </el-form-item>
          <el-form-item label="年份">
            <el-input v-model="search.year" placeholder="输入年份" clearable />
          </el-form-item>
          <el-form-item label="学院">
            <el-input v-model="search.college" placeholder="输入学院" clearable />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="fetchPapers">搜索</el-button>
            <el-button @click="resetSearch">重置</el-button>
          </el-form-item>
        </el-form>
        <el-table :data="papers" style="width: 100%;" @row-click="goDetail">
          <el-table-column prop="course" label="课程名" width="180" />
          <el-table-column prop="year" label="年份" width="100" />
          <el-table-column prop="term" label="学期" width="100" />
          <el-table-column prop="college" label="学院" width="180" />
          <el-table-column prop="created_at" label="录入时间" width="180" />
        </el-table>
        <el-pagination
          style="margin-top: 20px; text-align: right;"
          background
          layout="prev, pager, next, jumper, ->, total"
          :total="total"
          :page-size="pageSize"
          :current-page="page"
          @current-change="handlePageChange"
        />
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../api/request'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'

const router = useRouter()
const papers = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = 10
const search = ref({ course: '', year: '', college: '' })

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
      ElMessage.warning(res.data.msg)
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
}

const resetSearch = () => {
  search.value = { course: '', year: '', college: '' }
  page.value = 1
  fetchPapers()
}

const goDetail = (row) => {
  router.push(`/papers/${row.id}`)
}

onMounted(fetchPapers)
</script> 