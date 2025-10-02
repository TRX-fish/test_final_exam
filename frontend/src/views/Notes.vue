<template>
  <div class="notes-container">
    <div class="page-header">
      <div>
        <h2 class="page-title">我的笔记</h2>
        <p class="page-subtitle">共有 {{ notes.length }} 条笔记</p>
      </div>
      <el-button type="primary" @click="showAddDialog = true">
        <el-icon><Plus /></el-icon>
        新建笔记
      </el-button>
    </div>

    <!-- 标签筛选 -->
    <div class="tags-filter">
      <el-tag 
        v-for="tag in allTags" 
        :key="tag"
        :type="selectedTag === tag ? 'primary' : 'info'"
        @click="selectedTag = selectedTag === tag ? '' : tag"
        style="cursor: pointer; margin-right: 8px;"
      >
        {{ tag }}
      </el-tag>
    </div>

    <!-- 空状态 -->
    <el-empty v-if="filteredNotes.length === 0" description="还没有笔记">
      <el-button type="primary" @click="showAddDialog = true">
        创建第一条笔记
      </el-button>
    </el-empty>

    <!-- 笔记列表 -->
    <div v-else class="notes-grid">
      <div 
        v-for="note in filteredNotes" 
        :key="note.id" 
        class="note-card"
        @click="viewNote(note)"
      >
        <div class="note-header">
          <h4 class="note-title">{{ note.title }}</h4>
          <el-dropdown @command="handleCommand($event, note.id)">
            <el-icon class="more-icon"><MoreFilled /></el-icon>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="edit">编辑</el-dropdown-item>
                <el-dropdown-item command="delete">删除</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
        <div class="note-content">
          {{ note.content }}
        </div>
        <div class="note-footer">
          <div class="note-tags">
            <el-tag v-for="tag in note.tags" :key="tag" size="small" type="info">
              {{ tag }}
            </el-tag>
          </div>
          <span class="note-time">{{ note.time }}</span>
        </div>
      </div>
    </div>

    <!-- 新建/编辑笔记对话框 -->
    <el-dialog 
      v-model="showAddDialog" 
      :title="editingNote ? '编辑笔记' : '新建笔记'"
      width="600px"
    >
      <el-form :model="noteForm" label-width="80px">
        <el-form-item label="笔记标题">
          <el-input v-model="noteForm.title" placeholder="给笔记起个标题" />
        </el-form-item>
        <el-form-item label="笔记内容">
          <el-input 
            v-model="noteForm.content" 
            type="textarea" 
            :rows="8"
            placeholder="记录你的想法..."
          />
        </el-form-item>
        <el-form-item label="标签">
          <el-select 
            v-model="noteForm.tags" 
            multiple 
            placeholder="选择或创建标签"
            allow-create
            filterable
            style="width: 100%;"
          >
            <el-option 
              v-for="tag in allTags" 
              :key="tag" 
              :label="tag" 
              :value="tag" 
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showAddDialog = false">取消</el-button>
        <el-button type="primary" @click="saveNote">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, MoreFilled, Delete, TrendCharts, Document, Clock } from '@element-plus/icons-vue'

const router = useRouter()
const showAddDialog = ref(false)
const selectedTag = ref('')
const editingNote = ref(null)

const noteForm = ref({
  title: '',
  content: '',
  tags: []
})

// 模拟笔记数据
const notes = ref([
  {
    id: 1,
    title: '英语语法重点',
    content: '现在时态的用法：第一人称单数使用动词原形...',
    tags: ['英语', '语法'],
    time: '2小时前'
  },
  {
    id: 2,
    title: '高数积分公式',
    content: '常用积分公式整理：∫x^n dx = x^(n+1)/(n+1) + C...',
    tags: ['数学', '公式'],
    time: '1天前'
  }
])

const allTags = computed(() => {
  const tags = new Set()
  notes.value.forEach(note => {
    note.tags.forEach(tag => tags.add(tag))
  })
  return Array.from(tags)
})

const filteredNotes = computed(() => {
  if (!selectedTag.value) return notes.value
  return notes.value.filter(note => note.tags.includes(selectedTag.value))
})

const viewNote = (note) => {
  editingNote.value = note
  noteForm.value = { ...note }
  showAddDialog.value = true
}

const saveNote = () => {
  if (!noteForm.value.title) {
    ElMessage.warning('请输入笔记标题')
    return
  }
  
  if (editingNote.value) {
    // 编辑
    const index = notes.value.findIndex(n => n.id === editingNote.value.id)
    notes.value[index] = { ...noteForm.value, id: editingNote.value.id }
    ElMessage.success('笔记已更新')
  } else {
    // 新建
    notes.value.unshift({
      ...noteForm.value,
      id: Date.now(),
      time: '刚刚'
    })
    ElMessage.success('笔记已创建')
  }
  
  showAddDialog.value = false
  editingNote.value = null
  noteForm.value = { title: '', content: '', tags: [] }
}

const handleCommand = (command, noteId) => {
  if (command === 'edit') {
    const note = notes.value.find(n => n.id === noteId)
    viewNote(note)
  } else if (command === 'delete') {
    ElMessageBox.confirm('确定要删除这条笔记吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      notes.value = notes.value.filter(n => n.id !== noteId)
      ElMessage.success('笔记已删除')
    }).catch(() => {})
  }
}
</script>

<style scoped>
.notes-container {
  padding: 32px;
  max-width: 1200px;
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

.tags-filter {
  margin-bottom: 24px;
}

.notes-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 16px;
}

.note-card {
  background: #fff;
  border-radius: 8px;
  padding: 20px;
  border: 1px solid #f0f0f0;
  cursor: pointer;
  transition: all 0.3s;
}

.note-card:hover {
  border-color: #3A7AFE;
  box-shadow: 0 4px 16px rgba(58, 122, 254, 0.1);
  transform: translateY(-2px);
}

.note-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 12px;
}

.note-title {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0;
  flex: 1;
}

.more-icon {
  color: #999;
  cursor: pointer;
  padding: 4px;
  transition: color 0.3s;
}

.more-icon:hover {
  color: #3A7AFE;
}

.note-content {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  margin-bottom: 16px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
}

.note-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 12px;
  border-top: 1px solid #f5f5f5;
}

.note-tags {
  display: flex;
  gap: 6px;
  flex-wrap: wrap;
}

.note-time {
  font-size: 12px;
  color: #999;
}

@media (max-width: 768px) {
  .notes-container {
    padding: 16px;
  }
  
  .notes-grid {
    grid-template-columns: 1fr;
  }
}
</style> 