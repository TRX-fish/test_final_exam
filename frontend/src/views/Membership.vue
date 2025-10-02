<template>
  <div class="membership-container">
    <!-- Hero区域 -->
    <div class="membership-hero">
      <h1 class="hero-title">升级会员，解锁全部资源</h1>
      <p class="hero-subtitle">海量真题 · 详细解析 · 无限下载</p>
    </div>

    <!-- 套餐选择 -->
    <div class="plans-section">
      <el-row :gutter="24" justify="center">
        <el-col :span="7" :xs="24" :sm="12" :md="7">
          <div class="plan-card">
            <div class="plan-header">
              <h3 class="plan-name">月卡</h3>
              <div class="plan-price">
                <span class="currency">¥</span>
                <span class="amount">19.9</span>
                <span class="unit">/月</span>
              </div>
            </div>
            <ul class="plan-features">
              <li>
                <el-icon color="#00C2A8"><Check /></el-icon>
                查看全部试卷
              </li>
              <li>
                <el-icon color="#00C2A8"><Check /></el-icon>
                完整答案解析
              </li>
              <li>
                <el-icon color="#00C2A8"><Check /></el-icon>
                下载打印
              </li>
              <li>
                <el-icon color="#00C2A8"><Check /></el-icon>
                30天有效期
              </li>
            </ul>
            <el-button class="plan-btn" @click="selectPlan('monthly', 19.9)">
              选择月卡
            </el-button>
          </div>
        </el-col>

        <el-col :span="7" :xs="24" :sm="12" :md="7">
          <div class="plan-card featured">
            <div class="recommend-badge">最划算</div>
            <div class="plan-header">
              <h3 class="plan-name">年卡</h3>
              <div class="plan-price">
                <span class="currency">¥</span>
                <span class="amount">99.9</span>
                <span class="unit">/年</span>
              </div>
              <div class="save-tip">省 ¥139</div>
            </div>
            <ul class="plan-features">
              <li>
                <el-icon color="#00C2A8"><Check /></el-icon>
                月卡所有权益
              </li>
              <li>
                <el-icon color="#00C2A8"><Check /></el-icon>
                365天有效期
              </li>
              <li>
                <el-icon color="#00C2A8"><Check /></el-icon>
                专属客服支持
              </li>
              <li>
                <el-icon color="#00C2A8"><Check /></el-icon>
                优先更新通知
              </li>
            </ul>
            <el-button type="primary" class="plan-btn" @click="selectPlan('yearly', 99.9)">
              立即开通
            </el-button>
          </div>
        </el-col>

        <el-col :span="7" :xs="24" :sm="12" :md="7">
          <div class="plan-card">
            <div class="plan-header">
              <h3 class="plan-name">季卡</h3>
              <div class="plan-price">
                <span class="currency">¥</span>
                <span class="amount">49.9</span>
                <span class="unit">/季</span>
              </div>
            </div>
            <ul class="plan-features">
              <li>
                <el-icon color="#00C2A8"><Check /></el-icon>
                月卡所有权益
              </li>
              <li>
                <el-icon color="#00C2A8"><Check /></el-icon>
                90天有效期
              </li>
              <li>
                <el-icon color="#00C2A8"><Check /></el-icon>
                优惠续费
              </li>
              <li>
                <el-icon color="#999"><Close /></el-icon>
                <span style="color: #999;">专属客服</span>
              </li>
            </ul>
            <el-button class="plan-btn" @click="selectPlan('quarterly', 49.9)">
              选择季卡
            </el-button>
          </div>
        </el-col>
      </el-row>
    </div>

    <!-- 会员权益 -->
    <div class="benefits-section">
      <h2 class="section-title">会员专享特权</h2>
      <el-row :gutter="24" class="benefits-grid">
        <el-col :span="8" :xs="24" :sm="12" :md="8">
          <div class="benefit-card">
            <div class="benefit-icon" style="background: #E8F1FF;">
              <el-icon :size="32" color="#3A7AFE"><View /></el-icon>
            </div>
            <h4 class="benefit-title">无限查看</h4>
            <p class="benefit-desc">查看平台所有试卷，不限次数</p>
          </div>
        </el-col>
        <el-col :span="8" :xs="24" :sm="12" :md="8">
          <div class="benefit-card">
            <div class="benefit-icon" style="background: #E7F9F6;">
              <el-icon :size="32" color="#00C2A8"><Download /></el-icon>
            </div>
            <h4 class="benefit-title">下载打印</h4>
            <p class="benefit-desc">高清PDF下载，支持打印学习</p>
          </div>
        </el-col>
        <el-col :span="8" :xs="24" :sm="12" :md="8">
          <div class="benefit-card">
            <div class="benefit-icon" style="background: #FFF5E6;">
              <el-icon :size="32" color="#FF9800"><Reading /></el-icon>
            </div>
            <h4 class="benefit-title">完整解析</h4>
            <p class="benefit-desc">查看所有题目的详细答案解析</p>
          </div>
        </el-col>
      </el-row>
    </div>

    <!-- 支付对话框 -->
    <el-dialog v-model="showPayDialog" title="选择支付方式" width="500px">
      <div class="payment-info">
        <div class="order-summary">
          <div class="summary-item">
            <span>套餐类型</span>
            <span>{{ selectedPlanName }}</span>
          </div>
          <div class="summary-item total">
            <span>支付金额</span>
            <span class="price">¥{{ selectedPrice }}</span>
          </div>
        </div>

        <div class="payment-methods">
          <div 
            class="payment-method"
            :class="{ active: paymentMethod === 'alipay' }"
            @click="paymentMethod = 'alipay'"
          >
            <el-icon :size="32" color="#1677FF"><CreditCard /></el-icon>
            <span>支付宝</span>
          </div>
          <div 
            class="payment-method"
            :class="{ active: paymentMethod === 'wechat' }"
            @click="paymentMethod = 'wechat'"
          >
            <el-icon :size="32" color="#07C160"><ChatDotRound /></el-icon>
            <span>微信支付</span>
          </div>
        </div>
      </div>
      <template #footer>
        <el-button @click="showPayDialog = false">取消</el-button>
        <el-button type="primary" @click="confirmPayment" :loading="paying">
          确认支付
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Check, Close, View, Download, Reading, CreditCard, ChatDotRound } from '@element-plus/icons-vue'

const router = useRouter()
const showPayDialog = ref(false)
const selectedPlan = ref('')
const selectedPrice = ref(0)
const paymentMethod = ref('alipay')
const paying = ref(false)

const selectedPlanName = computed(() => {
  const names = {
    monthly: '月卡会员',
    quarterly: '季卡会员',
    yearly: '年卡会员'
  }
  return names[selectedPlan.value] || ''
})

const selectPlan = (plan, price) => {
  selectedPlan.value = plan
  selectedPrice.value = price
  showPayDialog.value = true
}

const confirmPayment = () => {
  paying.value = true
  
  // 模拟支付流程
  setTimeout(() => {
    paying.value = false
    showPayDialog.value = false
    
    ElMessage.success('支付成功！会员已开通')
    
    // 存储会员信息
    localStorage.setItem('membership', JSON.stringify({
      type: selectedPlan.value,
      status: 'active',
      endDate: getEndDate(selectedPlan.value)
    }))
    
    setTimeout(() => {
      router.push('/')
    }, 1500)
  }, 2000)
}

const getEndDate = (plan) => {
  const now = new Date()
  if (plan === 'monthly') {
    now.setMonth(now.getMonth() + 1)
  } else if (plan === 'quarterly') {
    now.setMonth(now.getMonth() + 3)
  } else if (plan === 'yearly') {
    now.setFullYear(now.getFullYear() + 1)
  }
  return now.toISOString()
}
</script>

<style scoped>
.membership-container {
  min-height: calc(100vh - 60px);
  background: #f8f9fa;
}

.membership-hero {
  background: linear-gradient(135deg, #3A7AFE 0%, #5B8DFF 100%);
  padding: 64px 32px;
  text-align: center;
  color: white;
}

.hero-title {
  font-size: 36px;
  font-weight: 600;
  margin: 0 0 16px 0;
}

.hero-subtitle {
  font-size: 18px;
  opacity: 0.9;
  margin: 0;
}

.plans-section {
  padding: 48px 32px;
  max-width: 1200px;
  margin: 0 auto;
  transform: translateY(-40px);
}

.plan-card {
  background: #fff;
  border-radius: 12px;
  padding: 32px 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  transition: all 0.3s;
  border: 2px solid transparent;
  position: relative;
}

.plan-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.12);
}

.plan-card.featured {
  border-color: #3A7AFE;
  transform: scale(1.05);
  box-shadow: 0 8px 32px rgba(58, 122, 254, 0.2);
}

.plan-card.featured:hover {
  transform: scale(1.05) translateY(-8px);
}

.recommend-badge {
  position: absolute;
  top: -12px;
  left: 50%;
  transform: translateX(-50%);
  background: linear-gradient(135deg, #FF6B6B 0%, #FF8E53 100%);
  color: white;
  padding: 6px 20px;
  border-radius: 16px;
  font-size: 13px;
  font-weight: 500;
}

.plan-header {
  text-align: center;
  margin-bottom: 24px;
  padding-bottom: 24px;
  border-bottom: 1px solid #f0f0f0;
}

.plan-name {
  font-size: 20px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 16px 0;
}

.plan-price {
  display: flex;
  align-items: baseline;
  justify-content: center;
  margin-bottom: 8px;
}

.currency {
  font-size: 20px;
  color: #3A7AFE;
  margin-right: 4px;
}

.amount {
  font-size: 48px;
  font-weight: 700;
  color: #3A7AFE;
  line-height: 1;
}

.unit {
  font-size: 16px;
  color: #666;
  margin-left: 4px;
}

.save-tip {
  color: #FF6B6B;
  font-size: 14px;
  font-weight: 500;
}

.plan-features {
  list-style: none;
  padding: 0;
  margin: 0 0 24px 0;
}

.plan-features li {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 0;
  font-size: 14px;
  color: #333;
}

.plan-btn {
  width: 100%;
  height: 44px;
  border-radius: 8px;
  font-size: 15px;
  font-weight: 500;
}

.benefits-section {
  padding: 64px 32px;
  max-width: 1200px;
  margin: 0 auto;
}

.section-title {
  font-size: 32px;
  font-weight: 600;
  text-align: center;
  margin: 0 0 48px 0;
}

.benefits-grid {
  margin-top: 32px;
}

.benefit-card {
  text-align: center;
  padding: 32px 24px;
}

.benefit-icon {
  width: 72px;
  height: 72px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
}

.benefit-title {
  font-size: 18px;
  font-weight: 600;
  margin: 0 0 12px 0;
}

.benefit-desc {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  margin: 0;
}

/* 支付对话框 */
.payment-info {
  padding: 16px 0;
}

.order-summary {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 24px;
}

.summary-item {
  display: flex;
  justify-content: space-between;
  padding: 8px 0;
  font-size: 14px;
  color: #666;
}

.summary-item.total {
  padding-top: 12px;
  margin-top: 12px;
  border-top: 1px solid #e5e5e5;
  font-weight: 600;
  color: #1a1a1a;
}

.summary-item .price {
  font-size: 24px;
  color: #FF6B6B;
}

.payment-methods {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.payment-method {
  background: #f8f9fa;
  border: 2px solid #e5e5e5;
  border-radius: 8px;
  padding: 24px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s;
}

.payment-method:hover {
  border-color: #3A7AFE;
}

.payment-method.active {
  border-color: #3A7AFE;
  background: #E8F1FF;
}

.payment-method span {
  display: block;
  margin-top: 8px;
  font-size: 14px;
  font-weight: 500;
}

@media (max-width: 768px) {
  .plans-section {
    padding: 24px 16px;
  }
  
  .plan-card {
    margin-bottom: 24px;
  }
  
  .plan-card.featured {
    transform: scale(1);
  }
}
</style> 