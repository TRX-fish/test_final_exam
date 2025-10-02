<template>
  <div class="pdf-viewer-container">
    <!-- 会员状态 -->
    <div class="member-status" v-if="!isMember">
      <el-alert type="warning" :closable="false" show-icon>
        <template #title>
          <span style="font-size: 14px;">您当前是免费用户，仅可预览前3题</span>
        </template>
        <el-button type="primary" size="small" @click="buyMembership">
          开通会员查看完整内容
        </el-button>
      </el-alert>
    </div>

    <!-- PDF预览区 -->
    <div class="pdf-preview-wrapper">
      <!-- 会员可完整查看 -->
      <div v-if="isMember" class="pdf-container">
        <div class="pdf-display">
          <!-- 直接在新标签页打开PDF（最可靠的方式） -->
          <div class="pdf-preview-card">
            <el-icon :size="120" color="#3A7AFE"><Document /></el-icon>
            <h3 class="preview-title">点击按钮查看完整PDF试卷</h3>
            <p class="preview-desc">为了最佳阅读体验，PDF将在新标签页打开</p>
            <el-button type="primary" size="large" @click="openInNewTab">
              <el-icon><View /></el-icon>
              查看PDF试卷
            </el-button>
            <el-button size="large" @click="downloadPDF" style="margin-top: 12px;">
              <el-icon><Download /></el-icon>
              下载到本地
            </el-button>
            <div class="file-info">
              <el-icon><InfoFilled /></el-icon>
              <span>文件大小：{{ formatSize(fileSize) }}</span>
            </div>
          </div>
        </div>
        <div class="watermark-static">{{ watermarkText }}</div>
      </div>

      <!-- 非会员模糊预览 -->
      <div v-else class="preview-locked">
        <div class="blur-preview">
          <iframe 
            :src="pdfUrl"
            width="100%"
            height="600px"
            frameborder="0"
            style="filter: blur(8px); pointer-events: none;"
          />
        </div>
        <div class="unlock-overlay">
          <div class="lock-icon">
            <el-icon :size="80" color="#3A7AFE"><Lock /></el-icon>
          </div>
          <h3 class="lock-title">开通会员解锁完整内容</h3>
          <p class="lock-desc">查看完整试卷、答案解析、下载打印</p>
          
          <div class="price-cards">
            <div class="price-card">
              <div class="price-label">月卡</div>
              <div class="price-value">¥19.9</div>
              <div class="price-desc">30天有效期</div>
            </div>
            <div class="price-card hot">
              <div class="hot-tag">最划算</div>
              <div class="price-label">年卡</div>
              <div class="price-value">¥99.9</div>
              <div class="price-desc">365天 · 省¥139</div>
            </div>
          </div>

          <el-button type="primary" size="large" class="buy-btn" @click="buyMembership">
            立即开通会员
          </el-button>
          <p class="trial-tip">新用户享3天免费试用</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Lock, Download, View, Document, InfoFilled } from '@element-plus/icons-vue'

const props = defineProps({
  pdfUrl: {
    type: String,
    required: true
  },
  isMember: {
    type: Boolean,
    default: false
  },
  fileSize: {
    type: Number,
    default: 0
  }
})

const router = useRouter()

const watermarkText = computed(() => {
  const username = localStorage.getItem('username') || '用户'
  return `${username} - 仅供个人学习使用`
})

const buyMembership = () => {
  router.push('/membership')
}

const downloadPDF = () => {
  const link = document.createElement('a')
  link.href = props.pdfUrl
  link.download = 'paper.pdf'
  link.click()
}

const openInNewTab = () => {
  window.open(props.pdfUrl, '_blank')
}

const formatSize = (bytes) => {
  if (!bytes || bytes === 0) return '未知'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return Math.round(bytes / Math.pow(k, i) * 100) / 100 + ' ' + sizes[i]
}
</script>

<style scoped>
.pdf-viewer-container {
  width: 100%;
}

.member-status {
  margin-bottom: 16px;
}

.pdf-preview-wrapper {
  position: relative;
}

/* 会员PDF容器 */
.pdf-container {
  position: relative;
  background: #fff;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
}

.pdf-display {
  position: relative;
}

.pdf-preview-card {
  background: #f8f9fa;
  border-radius: 12px;
  padding: 64px 32px;
  text-align: center;
  min-height: 400px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.preview-title {
  font-size: 22px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 24px 0 12px 0;
}

.preview-desc {
  font-size: 14px;
  color: #666;
  margin: 0 0 32px 0;
}

.file-info {
  margin-top: 24px;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #999;
}

.watermark-static {
  position: absolute;
  top: 20px;
  right: 20px;
  background: rgba(58, 122, 254, 0.15);
  color: #3A7AFE;
  padding: 8px 16px;
  border-radius: 20px;
  font-size: 12px;
  pointer-events: none;
  z-index: 1000;
  backdrop-filter: blur(10px);
  font-weight: 500;
  box-shadow: 0 2px 8px rgba(58, 122, 254, 0.2);
}

/* 非会员锁定状态 */
.preview-locked {
  position: relative;
  border-radius: 8px;
  overflow: hidden;
}

.blur-preview {
  position: relative;
}

.unlock-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0.95);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 48px;
  z-index: 100;
}

.lock-icon {
  width: 120px;
  height: 120px;
  background: #E8F1FF;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 32px;
}

.lock-title {
  font-size: 28px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 12px 0;
}

.lock-desc {
  font-size: 16px;
  color: #666;
  margin: 0 0 40px 0;
}

.price-cards {
  display: flex;
  gap: 16px;
  margin-bottom: 32px;
}

.price-card {
  background: #fff;
  border: 2px solid #e5e5e5;
  border-radius: 12px;
  padding: 24px;
  text-align: center;
  min-width: 150px;
  transition: all 0.3s;
  position: relative;
}

.price-card.hot {
  border-color: #3A7AFE;
  box-shadow: 0 4px 16px rgba(58, 122, 254, 0.2);
}

.hot-tag {
  position: absolute;
  top: -12px;
  right: 12px;
  background: linear-gradient(135deg, #FF6B6B 0%, #FF8E53 100%);
  color: white;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.price-label {
  font-size: 14px;
  color: #666;
  margin-bottom: 8px;
}

.price-value {
  font-size: 32px;
  font-weight: 600;
  color: #3A7AFE;
  margin-bottom: 8px;
}

.price-desc {
  font-size: 12px;
  color: #999;
}

.buy-btn {
  width: 100%;
  max-width: 300px;
  height: 48px;
  font-size: 16px;
  margin-bottom: 16px;
}

.trial-tip {
  font-size: 13px;
  color: #00C2A8;
  margin: 0;
}

@media (max-width: 768px) {
  .unlock-overlay {
    padding: 24px;
  }
  
  .lock-icon {
    width: 80px;
    height: 80px;
  }
  
  .lock-title {
    font-size: 22px;
  }
  
  .price-cards {
    flex-direction: column;
    width: 100%;
  }
}
</style> 