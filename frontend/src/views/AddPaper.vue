<template>
  <el-row justify="center" style="margin-top: 40px;">
    <el-col :span="12">
      <el-card>
        <template #header>
          <div style="display: flex; justify-content: space-between; align-items: center;">
            <h2 style="margin: 0;">添加试卷</h2>
            <el-button text @click="$router.back()">返回</el-button>
          </div>
        </template>
        <el-form :model="form" :rules="rules" ref="paperForm" label-width="100px">
          <el-form-item label="课程名称" prop="course">
            <el-input v-model="form.course" placeholder="如：高等数学" />
          </el-form-item>
          <el-form-item label="考试年份" prop="year">
            <el-input-number v-model="form.year" :min="2000" :max="2030" style="width: 100%;" />
          </el-form-item>
          <el-form-item label="学期" prop="term">
            <el-select v-model="form.term" placeholder="请选择学期" style="width: 100%;">
              <el-option label="春季" value="春季" />
              <el-option label="秋季" value="秋季" />
              <el-option label="夏季" value="夏季" />
            </el-select>
          </el-form-item>
          <el-form-item label="所属学院" prop="college">
            <el-input v-model="form.college" placeholder="如：计算机学院" />
          </el-form-item>
          <el-form-item label="试卷类型">
            <el-radio-group v-model="fileType">
              <el-radio label="pdf">PDF文件</el-radio>
              <el-radio label="image">图片</el-radio>
            </el-radio-group>
          </el-form-item>

          <el-form-item :label="fileType === 'pdf' ? '上传PDF' : '上传图片'">
            <el-upload
              class="upload-demo"
              :action="uploadUrl"
              :headers="uploadHeaders"
              :on-success="handleUploadSuccess"
              :on-error="handleUploadError"
              :before-upload="beforeUpload"
              :show-file-list="false"
            >
              <el-button type="primary" :loading="uploading">
                <el-icon><Upload /></el-icon>
                {{ fileType === 'pdf' ? '点击上传PDF文件' : '点击上传图片' }}
              </el-button>
              <template #tip>
                <div class="el-upload__tip">
                  {{ fileType === 'pdf' ? '支持PDF格式，不超过50MB' : '支持jpg/png格式，不超过10MB' }}
                </div>
              </template>
            </el-upload>
            <div v-if="form.image_path && fileType === 'image'" style="margin-top: 10px;">
              <el-image :src="form.image_path" style="width: 200px;" fit="cover" />
            </div>
            <div v-if="form.file_url && fileType === 'pdf'" style="margin-top: 10px;">
              <el-tag type="success">
                <el-icon><Document /></el-icon>
                PDF已上传
              </el-tag>
              <el-button type="primary" link size="small" @click="previewPDF">预览</el-button>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit" :loading="loading">提交</el-button>
            <el-button @click="onReset">重置</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import api from '../api/request'
import { ElMessage } from 'element-plus'
import { Upload, Document } from '@element-plus/icons-vue'

const router = useRouter()
const fileType = ref('pdf') // 默认PDF
const form = ref({
  course: '',
  year: new Date().getFullYear(),
  term: '',
  college: '',
  image_path: '',
  file_url: '',
  file_type: 'pdf',
  file_size: 0
})
const loading = ref(false)
const uploading = ref(false)
const paperForm = ref(null)

const rules = {
  course: [{ required: true, message: '请输入课程名称', trigger: 'blur' }],
  year: [{ required: true, message: '请选择年份', trigger: 'change' }],
  college: [{ required: true, message: '请输入学院名称', trigger: 'blur' }]
}

const uploadUrl = computed(() => {
  return fileType.value === 'pdf' 
    ? 'http://47.93.252.105:5000/api/upload/paper-pdf'
    : 'http://47.93.252.105:5000/api/upload/paper-image'
})

const uploadHeaders = {
  'Authorization': 'Bearer ' + localStorage.getItem('token')
}

const beforeUpload = (file) => {
  if (fileType.value === 'pdf') {
    // PDF验证
    const isPDF = file.type === 'application/pdf'
    const isLt50M = file.size / 1024 / 1024 < 50

    if (!isPDF) {
      ElMessage.error('只能上传 PDF 格式的文件!')
      return false
    }
    if (!isLt50M) {
      ElMessage.error('PDF文件不能超过 50MB!')
      return false
    }
  } else {
    // 图片验证
    const isImage = file.type === 'image/jpeg' || file.type === 'image/png'
    const isLt10M = file.size / 1024 / 1024 < 10

    if (!isImage) {
      ElMessage.error('只能上传 JPG/PNG 格式的图片!')
      return false
    }
    if (!isLt10M) {
      ElMessage.error('图片大小不能超过 10MB!')
      return false
    }
  }
  
  uploading.value = true
  return true
}

const handleUploadSuccess = (response) => {
  uploading.value = false
  if (response.code === '2000') {
    if (fileType.value === 'pdf') {
      // PDF上传成功
      form.value.file_url = response.data.file_url
      form.value.file_type = 'pdf'
      form.value.file_size = response.data.file_size
      form.value.image_path = response.data.file_path
      ElMessage.success('PDF上传成功')
    } else {
      // 图片上传成功
      form.value.image_path = response.data.image_path
      form.value.file_type = 'image'
      ElMessage.success('图片上传成功')
    }
  } else {
    ElMessage.error(response.msg || '上传失败')
  }
}

const handleUploadError = () => {
  uploading.value = false
  ElMessage.error('图片上传失败')
}

const onSubmit = () => {
  paperForm.value.validate(async valid => {
    if (!valid) return
    loading.value = true
    try {
      const res = await api.post('/api/papers', form.value)
      if (res.data.code === '2000') {
        ElMessage.success('试卷添加成功')
        router.push('/papers')
      } else {
        ElMessage.error(res.data.msg)
      }
    } catch (e) {
      ElMessage.error('添加失败')
    } finally {
      loading.value = false
    }
  })
}

const onReset = () => {
  paperForm.value.resetFields()
  form.value.image_path = ''
  form.value.file_url = ''
  form.value.file_size = 0
}

const previewPDF = () => {
  if (form.value.file_url) {
    window.open(form.value.file_url, '_blank')
  }
}
</script>

<style scoped>
.el-upload__tip {
  color: #999;
  font-size: 12px;
}
</style> 