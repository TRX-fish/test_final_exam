<template>
  <div class="user-manage-container">
    <div class="page-header">
      <div>
        <h2 class="page-title">用户管理</h2>
        <p class="page-subtitle">共 {{ users.length }} 位用户</p>
      </div>
      <el-input 
        v-model="searchText" 
        placeholder="搜索用户名或邮箱" 
        style="width: 300px;"
        :prefix-icon="Search"
        clearable
      />
    </div>

    <!-- 工具栏 -->
    <div class="toolbar">
      <div class="toolbar-left">
        <el-button 
          type="danger" 
          :disabled="selectedUsers.length === 0"
          @click="batchDelete"
        >
          <el-icon><Delete /></el-icon>
          批量删除 ({{ selectedUsers.length }})
        </el-button>
        <el-button 
          :disabled="selectedUsers.length === 0"
          @click="batchToggleStatus"
        >
          批量启用/禁用
        </el-button>
      </div>
      <div class="toolbar-right">
        <el-select v-model="filterRole" placeholder="角色筛选" style="width: 140px;" clearable>
          <el-option label="管理员" value="admin" />
          <el-option label="普通用户" value="user" />
        </el-select>
      </div>
    </div>

    <!-- 用户表格 -->
    <div class="table-card">
      <el-table 
        :data="filteredUsers" 
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="username" label="用户名" min-width="120">
          <template #default="scope">
            <div class="user-cell">
              <div class="user-avatar">
                <el-icon><UserFilled /></el-icon>
              </div>
              <span>{{ scope.row.username }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="email" label="邮箱" min-width="180" />
        <el-table-column label="角色" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.role === 'admin' ? 'danger' : 'primary'" size="small">
              {{ scope.row.role === 'admin' ? '管理员' : '普通用户' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.is_active ? 'success' : 'info'" size="small">
              {{ scope.row.is_active ? '激活' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="注册时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="scope">
            <el-button type="primary" link size="small" @click="editUser(scope.row)">
              编辑
            </el-button>
            <el-button 
              :type="scope.row.is_active ? 'warning' : 'success'" 
              link 
              size="small"
              @click="toggleStatus(scope.row)"
            >
              {{ scope.row.is_active ? '禁用' : '启用' }}
            </el-button>
            <el-button 
              type="danger" 
              link 
              size="small"
              @click="deleteUser(scope.row.id)"
              :disabled="scope.row.role === 'admin'"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 编辑用户对话框 -->
    <el-dialog v-model="showEditDialog" title="编辑用户" width="500px">
      <el-form :model="editForm" label-width="80px">
        <el-form-item label="用户名">
          <el-input v-model="editForm.username" disabled />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="editForm.email" />
        </el-form-item>
        <el-form-item label="角色">
          <el-select v-model="editForm.role" style="width: 100%;">
            <el-option label="普通用户" value="user" />
            <el-option label="管理员" value="admin" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="editForm.is_active" active-text="激活" inactive-text="禁用" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showEditDialog = false">取消</el-button>
        <el-button type="primary" @click="saveUser">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import api from '../../api/request'
import { Search, Delete, UserFilled } from '@element-plus/icons-vue'

const searchText = ref('')
const filterRole = ref('')
const selectedUsers = ref([])
const showEditDialog = ref(false)
const editForm = ref({})

// 模拟用户数据
const users = ref([
  { id: 1, username: 'admin', email: 'admin@example.com', role: 'admin', is_active: true, created_at: '2025-07-27 18:31:33' },
  { id: 2, username: 'user1', email: 'user1@example.com', role: 'user', is_active: true, created_at: '2025-07-27 18:31:45' },
  { id: 3, username: 'user2', email: 'user2@example.com', role: 'user', is_active: true, created_at: '2025-07-27 18:31:45' },
  { id: 4, username: 'testuser', email: 'test@example.com', role: 'user', is_active: false, created_at: '2025-07-27 10:39:29' },
  { id: 5, username: 'test', email: '12345@qq.com', role: 'user', is_active: true, created_at: '2025-07-28 16:23:13' }
])

const filteredUsers = computed(() => {
  let result = users.value
  
  if (searchText.value) {
    result = result.filter(u => 
      u.username.includes(searchText.value) || 
      u.email.includes(searchText.value)
    )
  }
  
  if (filterRole.value) {
    result = result.filter(u => u.role === filterRole.value)
  }
  
  return result
})

const handleSelectionChange = (selection) => {
  selectedUsers.value = selection
}

const editUser = (user) => {
  editForm.value = { ...user }
  showEditDialog.value = true
}

const saveUser = () => {
  const index = users.value.findIndex(u => u.id === editForm.value.id)
  users.value[index] = { ...editForm.value }
  showEditDialog.value = false
  ElMessage.success('用户信息已更新')
}

const toggleStatus = (user) => {
  user.is_active = !user.is_active
  ElMessage.success(user.is_active ? '用户已激活' : '用户已禁用')
}

const deleteUser = (id) => {
  ElMessageBox.confirm('确定要删除此用户吗？', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    users.value = users.value.filter(u => u.id !== id)
    ElMessage.success('用户已删除')
  }).catch(() => {})
}

const batchDelete = () => {
  ElMessageBox.confirm(`确定要删除选中的 ${selectedUsers.value.length} 个用户吗？`, '批量删除', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    const ids = selectedUsers.value.map(u => u.id)
    users.value = users.value.filter(u => !ids.includes(u.id))
    selectedUsers.value = []
    ElMessage.success('批量删除成功')
  }).catch(() => {})
}

const batchToggleStatus = () => {
  selectedUsers.value.forEach(user => {
    const target = users.value.find(u => u.id === user.id)
    target.is_active = !target.is_active
  })
  selectedUsers.value = []
  ElMessage.success('批量操作成功')
}
</script>

<style scoped>
.user-manage-container {
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
}

.page-subtitle {
  font-size: 14px;
  color: #666;
  margin: 0;
}

.toolbar {
  display: flex;
  justify-content: space-between;
  margin-bottom: 16px;
  gap: 16px;
}

.toolbar-left {
  display: flex;
  gap: 12px;
}

.table-card {
  background: #fff;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
}

.table-card :deep(.el-table) {
  border-radius: 8px;
}

.table-card :deep(.el-table th) {
  background: #f8f9fa;
  color: #333;
  font-weight: 600;
}

.user-cell {
  display: flex;
  align-items: center;
  gap: 10px;
}

.user-avatar {
  width: 32px;
  height: 32px;
  background: #E8F1FF;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #3A7AFE;
}

@media (max-width: 768px) {
  .user-manage-container {
    padding: 16px;
  }
  
  .toolbar {
    flex-direction: column;
  }
  
  .table-card {
    overflow-x: auto;
  }
}
</style> 