<template>
  <el-row justify="center" style="margin-top: 40px;">
    <el-col :span="14">
      <el-card>
        <template #header>
          <div style="display: flex; justify-content: space-between; align-items: center;">
            <h2 style="margin: 0;">添加题目</h2>
            <el-button text @click="$router.back()">返回</el-button>
          </div>
        </template>
        <el-form :model="form" :rules="rules" ref="questionForm" label-width="100px">
          <el-form-item label="题目类型" prop="type">
            <el-select v-model="form.type" placeholder="请选择题目类型" style="width: 100%;">
              <el-option label="单选" value="单选" />
              <el-option label="多选" value="多选" />
              <el-option label="判断" value="判断" />
              <el-option label="填空" value="填空" />
              <el-option label="简答" value="简答" />
            </el-select>
          </el-form-item>
          
          <el-form-item label="题目内容" prop="content">
            <el-input 
              v-model="form.content" 
              type="textarea" 
              :rows="4"
              placeholder="请输入题目内容"
            />
          </el-form-item>

          <el-form-item label="选项" v-if="['单选', '多选'].includes(form.type)">
            <div v-for="(opt, index) in options" :key="index" style="margin-bottom: 10px;">
              <el-input 
                v-model="options[index]" 
                :placeholder="`选项 ${String.fromCharCode(65 + index)}`"
              >
                <template #append>
                  <el-button @click="removeOption(index)" v-if="options.length > 2">删除</el-button>
                </template>
              </el-input>
            </div>
            <el-button @click="addOption" type="primary" plain>添加选项</el-button>
          </el-form-item>

          <el-form-item label="正确答案" prop="answer">
            <el-input v-model="form.answer" placeholder="如：A 或 ABC 或 正确答案内容" />
          </el-form-item>

          <el-form-item label="题目解析">
            <el-input 
              v-model="form.explanation" 
              type="textarea" 
              :rows="3"
              placeholder="请输入题目解析（选填）"
            />
          </el-form-item>

          <el-form-item label="题目图片">
            <el-upload
              :action="uploadUrl"
              :headers="uploadHeaders"
              :on-success="handleUploadSuccess"
              :on-error="handleUploadError"
              :before-upload="beforeUpload"
              :show-file-list="false"
            >
              <el-button type="primary" :loading="uploading">点击上传题目图片</el-button>
              <template #tip>
                <div class="el-upload__tip">只能上传 jpg/png 文件，且不超过 10MB</div>
              </template>
            </el-upload>
            <div v-if="form.image_path" style="margin-top: 10px;">
              <el-image :src="`http://47.93.252.105:5000/static/${form.image_path}`" style="width: 200px;" />
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
import { ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import api from '../api/request'
import { ElMessage } from 'element-plus'

const router = useRouter()
const route = useRoute()
const paperId = route.params.id

const form = ref({
  paper_id: paperId,
  type: '',
  content: '',
  answer: '',
  explanation: '',
  image_path: ''
})

const options = ref(['', '', '', '']) // 默认4个选项
const loading = ref(false)
const uploading = ref(false)
const questionForm = ref(null)

const rules = {
  type: [{ required: true, message: '请选择题目类型', trigger: 'change' }],
  content: [{ required: true, message: '请输入题目内容', trigger: 'blur' }],
  answer: [{ required: true, message: '请输入正确答案', trigger: 'blur' }]
}

const uploadUrl = 'http://47.93.252.105:5000/api/upload/question-image'
const uploadHeaders = {
  'Authorization': 'Bearer ' + localStorage.getItem('token')
}

// 监听题目类型变化
watch(() => form.value.type, (newType) => {
  if (!['单选', '多选'].includes(newType)) {
    options.value = ['']
  }
})

const addOption = () => {
  options.value.push('')
}

const removeOption = (index) => {
  options.value.splice(index, 1)
}

const beforeUpload = (file) => {
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
  uploading.value = true
  return true
}

const handleUploadSuccess = (response) => {
  uploading.value = false
  if (response.code === '2000') {
    form.value.image_path = response.data.image_path
    ElMessage.success('图片上传成功')
  } else {
    ElMessage.error(response.msg || '上传失败')
  }
}

const handleUploadError = () => {
  uploading.value = false
  ElMessage.error('图片上传失败')
}

const onSubmit = () => {
  questionForm.value.validate(async valid => {
    if (!valid) return
    
    loading.value = true
    try {
      // 如果是选择题，将选项转为JSON
      const data = { ...form.value }
      if (['单选', '多选'].includes(form.value.type)) {
        const optionsList = options.value.filter(o => o.trim())
        data.options = JSON.stringify(optionsList)
      }
      
      const res = await api.post('/api/questions', data)
      if (res.data.code === '2000') {
        ElMessage.success('题目添加成功')
        router.back()
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
  questionForm.value.resetFields()
  options.value = ['', '', '', '']
  form.value.image_path = ''
}
</script>

<style scoped>
.el-upload__tip {
  color: #999;
  font-size: 12px;
  margin-top: 5px;
}
</style> 