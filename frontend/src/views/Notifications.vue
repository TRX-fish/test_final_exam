<template>
  <div class="notifications-container">
    <div class="page-header">
      <div>
        <h2 class="page-title">通知中心</h2>
        <p class="page-subtitle">{{ unreadCount }} 条未读消息</p>
      </div>
      <el-button @click="markAllRead" :disabled="unreadCount === 0">
        全部标为已读
      </el-button>
    </div>

    <!-- 消息分类 -->
    <div class="tabs-wrapper">
      <el-tabs v-model="activeTab">
        <el-tab-pane label="全部" name="all">
          <span slot="label">
            全部
            <el-badge :value="unreadCount" :hidden="unreadCount === 0" />
          </span>
        </el-tab-pane>
        <el-tab-pane label="系统公告" name="system">系统公告</el-tab-pane>
        <el-tab-pane label="互动消息" name="interaction">互动消息</el-tab-pane>
        <el-tab-pane label="评论回复" name="reply">评论回复</el-tab-pane>
      </el-tabs>
    </div>

    <!-- 空状态 -->
    <el-empty v-if="filteredNotifications.length === 0" description="暂无通知" />

    <!-- 通知列表 -->
    <div v-else class="notifications-list">
      <div 
        v-for="notif in filteredNotifications" 
        :key="notif.id"
        class="notification-item"
        :class="{ unread: !notif.read }"
        @click="markAsRead(notif.id)"
      >
        <div class="notif-icon" :style="{ background: getIconBg(notif.type) }">
          <el-icon :size="24" :color="getIconColor(notif.type)">
            <component :is="getIcon(notif.type)" />
          </el-icon>
        </div>
        <div class="notif-content">
          <h4 class="notif-title">{{ notif.title }}</h4>
          <p class="notif-desc">{{ notif.content }}</p>
          <span class="notif-time">{{ notif.time }}</span>
        </div>
        <div class="notif-status">
          <el-tag v-if="!notif.read" type="danger" size="small" round>未读</el-tag>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { Bell, UserFilled, ChatDotRound, Warning } from '@element-plus/icons-vue'

const activeTab = ref('all')

// 模拟通知数据
const notifications = ref([
  {
    id: 1,
    type: 'system',
    title: '系统维护通知',
    content: '系统将于今晚22:00-23:00进行升级维护，期间可能无法访问',
    time: '30分钟前',
    read: false
  },
  {
    id: 2,
    type: 'interaction',
    title: '用户 @user1 收藏了你的试卷',
    content: '《高等数学 2023秋季》获得了新的收藏',
    time: '2小时前',
    read: false
  },
  {
    id: 3,
    type: 'reply',
    title: '你的讨论有新回复',
    content: '用户 @admin 回复了你的话题《英语语法问题》',
    time: '3小时前',
    read: true
  },
  {
    id: 4,
    type: 'system',
    title: '新功能上线',
    content: '错题本功能已上线，快来体验吧！',
    time: '1天前',
    read: true
  }
])

const unreadCount = computed(() => {
  return notifications.value.filter(n => !n.read).length
})

const filteredNotifications = computed(() => {
  if (activeTab.value === 'all') return notifications.value
  return notifications.value.filter(n => n.type === activeTab.value)
})

const getIcon = (type) => {
  const icons = {
    system: 'Bell',
    interaction: 'UserFilled',
    reply: 'ChatDotRound'
  }
  return icons[type] || 'Bell'
}

const getIconBg = (type) => {
  const colors = {
    system: '#E8F1FF',
    interaction: '#E7F9F6',
    reply: '#FFF5E6'
  }
  return colors[type] || '#f0f0f0'
}

const getIconColor = (type) => {
  const colors = {
    system: '#3A7AFE',
    interaction: '#00C2A8',
    reply: '#FF9800'
  }
  return colors[type] || '#909399'
}

const markAsRead = (id) => {
  const notif = notifications.value.find(n => n.id === id)
  if (notif && !notif.read) {
    notif.read = true
  }
}

const markAllRead = () => {
  notifications.value.forEach(n => n.read = true)
  ElMessage.success('已全部标为已读')
}
</script>

<style scoped>
.notifications-container {
  padding: 32px;
  max-width: 900px;
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

.tabs-wrapper {
  background: #fff;
  border-radius: 8px;
  padding: 0 24px;
  margin-bottom: 16px;
}

.notifications-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.notification-item {
  background: #fff;
  border-radius: 8px;
  padding: 20px;
  display: flex;
  gap: 16px;
  border: 1px solid #f0f0f0;
  cursor: pointer;
  transition: all 0.3s;
  position: relative;
}

.notification-item.unread {
  background: #f8f9ff;
  border-left: 3px solid #3A7AFE;
}

.notification-item:hover {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.notif-icon {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.notif-content {
  flex: 1;
  min-width: 0;
}

.notif-title {
  font-size: 15px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 8px 0;
}

.notif-desc {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  margin: 0 0 8px 0;
}

.notif-time {
  font-size: 12px;
  color: #999;
}

.notif-status {
  flex-shrink: 0;
  display: flex;
  align-items: flex-start;
}

@media (max-width: 768px) {
  .notifications-container {
    padding: 16px;
  }
  
  .notification-item {
    flex-wrap: wrap;
  }
}
</style> 