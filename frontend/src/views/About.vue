<template>
  <div class="about-container">
    <!-- Hero区域 -->
    <div class="about-hero">
      <div class="hero-content">
        <h1 class="hero-title">关于我们</h1>
        <p class="hero-subtitle">致力于打造最好用的期末考试题库平台</p>
      </div>
    </div>

    <!-- 系统介绍 -->
    <div class="intro-section">
      <div class="section-content">
        <h2 class="section-title">平台简介</h2>
        <p class="section-text">
          期末考试题库管理系统是一个专为大学生设计的学习辅助平台。我们收录历年各科目期末考试真题，
          提供详细的答案解析，帮助学生高效复习备考。平台采用现代化的技术栈（Go + Vue 3），
          提供流畅的用户体验和强大的功能支持。
        </p>
      </div>
    </div>

    <!-- 核心功能 -->
    <div class="features-section">
      <div class="section-content">
        <h2 class="section-title">核心功能</h2>
        <el-row :gutter="24" class="features-grid">
          <el-col :span="8" :xs="24" :sm="12" :md="8">
            <div class="feature-card">
              <div class="feature-icon" style="background: #E8F1FF;">
                <el-icon :size="32" color="#3A7AFE"><Document /></el-icon>
              </div>
              <h4 class="feature-title">海量真题</h4>
              <p class="feature-desc">收录各学科历年期末考试真题，覆盖主要专业课程</p>
            </div>
          </el-col>
          <el-col :span="8" :xs="24" :sm="12" :md="8">
            <div class="feature-card">
              <div class="feature-icon" style="background: #E7F9F6;">
                <el-icon :size="32" color="#00C2A8"><Search /></el-icon>
              </div>
              <h4 class="feature-title">智能搜索</h4>
              <p class="feature-desc">支持关键词、类型、年份等多维度搜索和筛选</p>
            </div>
          </el-col>
          <el-col :span="8" :xs="24" :sm="12" :md="8">
            <div class="feature-card">
              <div class="feature-icon" style="background: #FFF5E6;">
                <el-icon :size="32" color="#FF9800"><ChatDotRound /></el-icon>
              </div>
              <h4 class="feature-title">详细解析</h4>
              <p class="feature-desc">每道题目都配有详细的答案解析和思路点拨</p>
            </div>
          </el-col>
          <el-col :span="8" :xs="24" :sm="12" :md="8">
            <div class="feature-card">
              <div class="feature-icon" style="background: #FFE8E8;">
                <el-icon :size="32" color="#F56C6C"><Star /></el-icon>
              </div>
              <h4 class="feature-title">个性收藏</h4>
              <p class="feature-desc">收藏重点试卷，建立专属题库，随时复习</p>
            </div>
          </el-col>
          <el-col :span="8" :xs="24" :sm="12" :md="8">
            <div class="feature-card">
              <div class="feature-icon" style="background: #F0F0F0;">
                <el-icon :size="32" color="#909399"><Edit /></el-icon>
              </div>
              <h4 class="feature-title">错题管理</h4>
              <p class="feature-desc">自动记录错题，智能推荐复习，查漏补缺</p>
            </div>
          </el-col>
          <el-col :span="8" :xs="24" :sm="12" :md="8">
            <div class="feature-card">
              <div class="feature-icon" style="background: #E8F4FF;">
                <el-icon :size="32" color="#409EFF"><Connection /></el-icon>
              </div>
              <h4 class="feature-title">社区交流</h4>
              <p class="feature-desc">讨论区分享学习经验，互帮互助共同进步</p>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>

    <!-- 技术栈 -->
    <div class="tech-section">
      <div class="section-content">
        <h2 class="section-title">技术栈</h2>
        <div class="tech-grid">
          <div class="tech-item">
            <div class="tech-label">后端</div>
            <div class="tech-value">Go + Gin</div>
          </div>
          <div class="tech-item">
            <div class="tech-label">前端</div>
            <div class="tech-value">Vue 3 + Vite</div>
          </div>
          <div class="tech-item">
            <div class="tech-label">数据库</div>
            <div class="tech-value">MySQL/MariaDB</div>
          </div>
          <div class="tech-item">
            <div class="tech-label">UI框架</div>
            <div class="tech-value">Element Plus</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 联系我们 -->
    <div class="contact-section">
      <div class="section-content">
        <h2 class="section-title">联系我们</h2>
        <div class="contact-info">
          <div class="contact-item">
            <el-icon :size="20" color="#3A7AFE"><Message /></el-icon>
            <span>support@example.com</span>
          </div>
          <div class="contact-item">
            <el-icon :size="20" color="#3A7AFE"><Link /></el-icon>
            <span>https://github.com/TRX-fish/test_final_exam</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Document, Search, ChatDotRound, Star, Edit, Connection, Message, Link } from '@element-plus/icons-vue'
</script>

<style scoped>
.about-container {
  min-height: calc(100vh - 60px);
}

.about-hero {
  background: linear-gradient(135deg, #3A7AFE 0%, #5B8DFF 100%);
  padding: 80px 32px;
  text-align: center;
  color: white;
}

.hero-title {
  font-size: 42px;
  font-weight: 600;
  margin: 0 0 16px 0;
}

.hero-subtitle {
  font-size: 18px;
  opacity: 0.9;
}

.intro-section,
.features-section,
.tech-section,
.contact-section {
  padding: 64px 32px;
}

.intro-section {
  background: #fff;
}

.features-section {
  background: #f8f9fa;
}

.section-content {
  max-width: 1200px;
  margin: 0 auto;
}

.section-title {
  font-size: 32px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 32px 0;
  text-align: center;
}

.section-text {
  font-size: 16px;
  color: #666;
  line-height: 1.8;
  text-align: center;
  max-width: 800px;
  margin: 0 auto;
}

.features-grid {
  margin-top: 40px;
}

.feature-card {
  background: #fff;
  border-radius: 12px;
  padding: 32px 24px;
  text-align: center;
  border: 1px solid #f0f0f0;
  transition: all 0.3s;
  height: 100%;
}

.feature-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
  border-color: #3A7AFE;
}

.feature-icon {
  width: 72px;
  height: 72px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
}

.feature-title {
  font-size: 18px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 12px 0;
}

.feature-desc {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  margin: 0;
}

.tech-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 16px;
  margin-top: 32px;
}

.tech-item {
  background: #fff;
  border-radius: 8px;
  padding: 24px;
  text-align: center;
  border: 1px solid #f0f0f0;
}

.tech-label {
  font-size: 14px;
  color: #999;
  margin-bottom: 8px;
}

.tech-value {
  font-size: 18px;
  font-weight: 600;
  color: #1a1a1a;
}

.contact-info {
  display: flex;
  flex-direction: column;
  gap: 16px;
  max-width: 500px;
  margin: 32px auto 0;
}

.contact-item {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 16px;
  color: #333;
  padding: 16px;
  background: #f8f9fa;
  border-radius: 8px;
}

@media (max-width: 768px) {
  .about-hero {
    padding: 48px 24px;
  }
  
  .hero-title {
    font-size: 32px;
  }
  
  .intro-section,
  .features-section,
  .tech-section,
  .contact-section {
    padding: 48px 16px;
  }
  
  .features-grid {
    gap: 16px;
  }
}
</style> 