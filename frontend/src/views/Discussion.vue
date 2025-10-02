<template>
  <div class="discussion-container">
    <div class="page-header">
      <div>
        <h2 class="page-title">讨论区</h2>
        <p class="page-subtitle">{{ discussions.length }} 个话题正在讨论中</p>
      </div>
      <el-button type="primary" @click="showNewTopic = true">
        <el-icon><ChatDotRound /></el-icon>
        发起讨论
      </el-button>
    </div>

    <!-- 话题分类 -->
    <div class="category-tabs">
      <el-radio-group v-model="category">
        <el-radio-button label="all">全部</el-radio-button>
        <el-radio-button label="question">题目讨论</el-radio-button>
        <el-radio-button label="experience">学习经验</el-radio-button>
        <el-radio-button label="resource">资源分享</el-radio-button>
      </el-radio-group>
    </div>

    <!-- 空状态 -->
    <el-empty v-if="filteredDiscussions.length === 0" description="暂无讨论">
      <el-button type="primary" @click="showNewTopic = true">
        发起第一个话题
      </el-button>
    </el-empty>

    <!-- 讨论列表 -->
    <div v-else class="discussions-list">
      <div 
        v-for="topic in filteredDiscussions" 
        :key="topic.id" 
        class="discussion-item"
        @click="viewTopic(topic.id)"
      >
        <div class="item-left">
          <div class="avatar">
            <el-icon :size="24"><UserFilled /></el-icon>
          </div>
        </div>
        <div class="item-content">
          <div class="topic-header">
            <h4 class="topic-title">{{ topic.title }}</h4>
            <el-tag size="small" :type="getCategoryType(topic.category)">
              {{ getCategoryLabel(topic.category) }}
            </el-tag>
          </div>
          <p class="topic-preview">{{ topic.content }}</p>
          <div class="topic-meta">
            <span class="meta-item">
              <el-icon><User /></el-icon>
              {{ topic.author }}
            </span>
            <span class="meta-item">
              <el-icon><ChatLineRound /></el-icon>
              {{ topic.replies }} 回复
            </span>
            <span class="meta-item">
              <el-icon><View /></el-icon>
              {{ topic.views }} 浏览
            </span>
            <span class="meta-time">{{ topic.time }}</span>
          </div>
        </div>
        <div class="item-right">
          <el-button 
            :type="topic.liked ? 'primary' : ''"
            circle
            @click.stop="toggleLike(topic.id)"
          >
            <el-icon><StarFilled v-if="topic.liked" /><Star v-else /></el-icon>
          </el-button>
          <span class="like-count">{{ topic.likes }}</span>
        </div>
      </div>
    </div>

    <!-- 新建话题对话框 -->
    <el-dialog 
      v-model="showNewTopic" 
      title="发起新讨论"
      width="600px"
    >
      <el-form :model="topicForm" label-width="80px">
        <el-form-item label="话题分类">
          <el-select v-model="topicForm.category" placeholder="选择分类" style="width: 100%;">
            <el-option label="题目讨论" value="question" />
            <el-option label="学习经验" value="experience" />
            <el-option label="资源分享" value="resource" />
          </el-select>
        </el-form-item>
        <el-form-item label="话题标题">
          <el-input v-model="topicForm.title" placeholder="简短描述你的问题或分享" />
        </el-form-item>
        <el-form-item label="详细内容">
          <el-input 
            v-model="topicForm.content" 
            type="textarea" 
            :rows="6"
            placeholder="详细描述你想讨论的内容..."
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showNewTopic = false">取消</el-button>
        <el-button type="primary" @click="publishTopic">发布</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ChatDotRound, UserFilled, User, ChatLineRound, View, Star, StarFilled } from '@element-plus/icons-vue'

const router = useRouter()
const category = ref('all')
const showNewTopic = ref(false)

const topicForm = ref({
  category: '',
  title: '',
  content: ''
})

// 模拟讨论数据
const discussions = ref([
  {
    id: 1,
    category: 'question',
    title: '这道英语题的B选项为什么不对？',
    content: 'What is the capital of England? 我选的Paris，但答案是London...',
    author: 'user1',
    replies: 5,
    views: 123,
    likes: 8,
    liked: false,
    time: '2小时前'
  },
  {
    id: 2,
    category: 'experience',
    title: '分享一个记忆高数公式的好方法',
    content: '我发现通过绘制思维导图可以更好地记忆积分公式...',
    author: 'admin',
    replies: 12,
    views: 256,
    likes: 23,
    liked: true,
    time: '1天前'
  },
  {
    id: 3,
    category: 'resource',
    title: '2023年计算机学院各科真题汇总',
    content: '整理了去年的所有专业课真题，分享给大家...',
    author: 'user2',
    replies: 8,
    views: 189,
    likes: 15,
    liked: false,
    time: '2天前'
  }
])

const filteredDiscussions = computed(() => {
  if (category.value === 'all') return discussions.value
  return discussions.value.filter(d => d.category === category.value)
})

const getCategoryType = (cat) => {
  const types = {
    question: 'primary',
    experience: 'success',
    resource: 'warning'
  }
  return types[cat] || ''
}

const getCategoryLabel = (cat) => {
  const labels = {
    question: '题目讨论',
    experience: '学习经验',
    resource: '资源分享'
  }
  return labels[cat] || cat
}

const toggleLike = (id) => {
  const topic = discussions.value.find(d => d.id === id)
  if (topic.liked) {
    topic.likes--
    topic.liked = false
  } else {
    topic.likes++
    topic.liked = true
  }
}

const viewTopic = (id) => {
  ElMessage.info('话题详情页面开发中...')
}

const publishTopic = () => {
  if (!topicForm.value.title || !topicForm.value.content) {
    ElMessage.warning('请填写完整信息')
    return
  }
  
  discussions.value.unshift({
    id: Date.now(),
    ...topicForm.value,
    author: '我',
    replies: 0,
    views: 0,
    likes: 0,
    liked: false,
    time: '刚刚'
  })
  
  ElMessage.success('话题发布成功')
  showNewTopic.value = false
  topicForm.value = { category: '', title: '', content: '' }
}
</script>

<style scoped>
.discussion-container {
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

.category-tabs {
  margin-bottom: 24px;
}

.discussions-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.discussion-item {
  background: #fff;
  border-radius: 8px;
  padding: 20px;
  display: flex;
  gap: 16px;
  border: 1px solid #f0f0f0;
  cursor: pointer;
  transition: all 0.3s;
}

.discussion-item:hover {
  border-color: #3A7AFE;
  box-shadow: 0 2px 12px rgba(58, 122, 254, 0.08);
}

.item-left {
  flex-shrink: 0;
}

.avatar {
  width: 48px;
  height: 48px;
  background: #E8F1FF;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #3A7AFE;
}

.item-content {
  flex: 1;
  min-width: 0;
}

.topic-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.topic-title {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0;
  flex: 1;
}

.topic-preview {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  margin: 0 0 12px 0;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.topic-meta {
  display: flex;
  gap: 16px;
  font-size: 13px;
  color: #999;
  flex-wrap: wrap;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.meta-time {
  margin-left: auto;
}

.item-right {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.like-count {
  font-size: 13px;
  color: #666;
  font-weight: 500;
}

@media (max-width: 768px) {
  .discussion-container {
    padding: 16px;
  }
  
  .discussion-item {
    flex-wrap: wrap;
  }
  
  .item-right {
    flex-direction: row;
    width: 100%;
    justify-content: flex-end;
  }
}
</style> 