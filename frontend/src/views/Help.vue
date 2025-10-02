<template>
  <div class="help-container">
    <div class="help-hero">
      <h1 class="hero-title">帮助中心</h1>
      <p class="hero-subtitle">快速找到您需要的答案</p>
      <el-input 
        v-model="searchKeyword" 
        placeholder="搜索问题..." 
        size="large"
        :prefix-icon="Search"
        class="search-box"
      />
    </div>

    <!-- 快速导航 -->
    <div class="quick-nav">
      <div 
        v-for="cat in categories" 
        :key="cat.id"
        class="nav-item"
        :class="{ active: activeCategory === cat.id }"
        @click="activeCategory = cat.id"
      >
        <el-icon :size="24"><component :is="cat.icon" /></el-icon>
        <span>{{ cat.name }}</span>
      </div>
    </div>

    <!-- FAQ列表 -->
    <div class="faq-section">
      <h3 class="section-title">常见问题</h3>
      <el-collapse v-model="activeNames" class="faq-collapse">
        <el-collapse-item 
          v-for="faq in filteredFaqs" 
          :key="faq.id"
          :name="faq.id"
        >
          <template #title>
            <div class="faq-title">
              <el-icon color="#3A7AFE"><QuestionFilled /></el-icon>
              <span>{{ faq.question }}</span>
            </div>
          </template>
          <div class="faq-answer">
            {{ faq.answer }}
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>

    <!-- 联系支持 -->
    <div class="contact-section">
      <div class="contact-card">
        <div class="contact-icon">
          <el-icon :size="48" color="#3A7AFE"><Service /></el-icon>
        </div>
        <h3 class="contact-title">还有其他问题？</h3>
        <p class="contact-desc">我们的团队随时为您提供帮助</p>
        <el-button type="primary" size="large" @click="contactSupport">
          联系客服
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { Search, QuestionFilled, Service, User, Document, Setting, QuestionFilled as Question } from '@element-plus/icons-vue'

const searchKeyword = ref('')
const activeCategory = ref('all')
const activeNames = ref([1])

const categories = [
  { id: 'all', name: '全部', icon: 'Document' },
  { id: 'account', name: '账号相关', icon: 'User' },
  { id: 'paper', name: '试卷使用', icon: 'Document' },
  { id: 'other', name: '其他问题', icon: 'Setting' }
]

const faqs = ref([
  {
    id: 1,
    category: 'account',
    question: '如何注册账号？',
    answer: '点击登录页面的"立即注册"按钮，填写用户名、密码和邮箱（选填）即可完成注册。注册后可以立即登录使用系统。'
  },
  {
    id: 2,
    category: 'account',
    question: '忘记密码怎么办？',
    answer: '目前密码找回功能正在开发中，如忘记密码请联系管理员重置。'
  },
  {
    id: 3,
    category: 'paper',
    question: '如何添加试卷？',
    answer: '登录后，进入"试卷列表"页面，点击"添加试卷"按钮，填写课程名称、年份、学期、学院等信息，还可以上传试卷图片。'
  },
  {
    id: 4,
    category: 'paper',
    question: '如何为试卷添加题目？',
    answer: '打开试卷详情页，点击"添加题目"按钮，选择题目类型（单选/多选/判断/填空/简答），填写题目内容、选项、答案和解析即可。'
  },
  {
    id: 5,
    category: 'paper',
    question: '支持哪些题目类型？',
    answer: '系统支持5种题目类型：单选题、多选题、判断题、填空题和简答题。选择题可以添加多个选项，所有题型都支持添加解析和图片。'
  },
  {
    id: 6,
    category: 'other',
    question: '如何搜索题目？',
    answer: '进入"题目搜索"页面，可以通过关键词搜索题目内容、选项、答案等，也可以按题目类型筛选。'
  },
  {
    id: 7,
    category: 'other',
    question: '收藏和错题本功能如何使用？',
    answer: '在浏览试卷时可以点击收藏按钮收藏试卷；做题时如果答错，系统会自动记录到错题本，方便后续复习。'
  }
])

const filteredFaqs = computed(() => {
  let result = faqs.value
  
  if (activeCategory.value !== 'all') {
    result = result.filter(faq => faq.category === activeCategory.value)
  }
  
  if (searchKeyword.value) {
    result = result.filter(faq => 
      faq.question.includes(searchKeyword.value) || 
      faq.answer.includes(searchKeyword.value)
    )
  }
  
  return result
})

const contactSupport = () => {
  ElMessage.info('客服功能开发中，请通过邮箱 support@example.com 联系我们')
}
</script>

<style scoped>
.help-container {
  min-height: calc(100vh - 60px);
}

.help-hero {
  background: linear-gradient(135deg, #3A7AFE 0%, #5B8DFF 100%);
  padding: 64px 32px;
  text-align: center;
  color: white;
}

.hero-title {
  font-size: 36px;
  font-weight: 600;
  margin: 0 0 16px 0;
  line-height: 1.2;
}

.hero-subtitle {
  font-size: 16px;
  margin: 0 0 32px 0;
  opacity: 0.9;
}

.search-box {
  max-width: 600px;
  margin: 0 auto;
}

.search-box :deep(.el-input__wrapper) {
  background: white;
  border-radius: 24px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
}

/* 快速导航 */
.quick-nav {
  max-width: 1000px;
  margin: -40px auto 48px;
  padding: 0 32px;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
}

.nav-item {
  background: #fff;
  border-radius: 8px;
  padding: 24px;
  text-align: center;
  cursor: pointer;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  transition: all 0.3s;
  border: 2px solid transparent;
}

.nav-item:hover, .nav-item.active {
  border-color: #3A7AFE;
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(58, 122, 254, 0.12);
}

.nav-item span {
  display: block;
  margin-top: 12px;
  font-size: 14px;
  font-weight: 500;
  color: #333;
}

/* FAQ部分 */
.faq-section {
  max-width: 1000px;
  margin: 0 auto;
  padding: 0 32px 48px;
}

.section-title {
  font-size: 24px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 24px 0;
}

.faq-collapse {
  border: none;
}

.faq-collapse :deep(.el-collapse-item) {
  background: #fff;
  border-radius: 8px;
  margin-bottom: 12px;
  border: 1px solid #f0f0f0;
  overflow: hidden;
}

.faq-collapse :deep(.el-collapse-item__header) {
  padding: 20px 24px;
  background: #fff;
  border: none;
  font-size: 15px;
  font-weight: 500;
  transition: all 0.3s;
}

.faq-collapse :deep(.el-collapse-item__header:hover) {
  background: #f8f9fa;
}

.faq-collapse :deep(.el-collapse-item__wrap) {
  border: none;
}

.faq-title {
  display: flex;
  align-items: center;
  gap: 12px;
}

.faq-answer {
  padding: 0 24px 20px 48px;
  font-size: 14px;
  color: #666;
  line-height: 1.8;
}

/* 联系支持 */
.contact-section {
  background: #f8f9fa;
  padding: 48px 32px;
}

.contact-card {
  max-width: 500px;
  margin: 0 auto;
  background: #fff;
  border-radius: 12px;
  padding: 48px 32px;
  text-align: center;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
}

.contact-icon {
  width: 80px;
  height: 80px;
  background: #E8F1FF;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 24px;
}

.contact-title {
  font-size: 24px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 12px 0;
}

.contact-desc {
  font-size: 14px;
  color: #666;
  margin: 0 0 24px 0;
}

@media (max-width: 768px) {
  .quick-nav {
    grid-template-columns: repeat(2, 1fr);
    padding: 0 16px;
  }
  
  .faq-section {
    padding: 0 16px 32px;
  }
}
</style> 