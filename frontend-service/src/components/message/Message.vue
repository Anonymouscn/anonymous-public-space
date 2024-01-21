<template>
  <Transition name="message-m">
    <div class="message" :style="style[type]" v-show="show">
      <span class="text">{{ text }}</span>
    </div>
  </Transition>
</template>

<script setup>
import {onMounted, reactive, ref} from "vue";

// 消息数据
const messageData = defineProps({
  // 通知文本
  text: {
    type: String,
    default: 'text',
  },
  // 通知类型
  type: {
    type: String,
    default: 'normal'
  },
  // 持续时间 (ms)
  duration: {
    type: Number,
    default: 3000
  },
  // 自定义类型样式
  customStyle: {
    type: Object,
    default: null
  }
})

// 模版样式
const style = reactive({
  // 正常消息样式
  normal: {
    color: '#fff',
    backgroundColor: 'rgb(0,0,0, 0.3)',
    borderColor: 'rgb(60, 60, 60)',
    maxWidth: '86%',
    fontSize: '1.2rem',
    fontWeight: 500
  },
  // 警告消息样式
  warn: {
    color: '#E6A23C',
    backgroundColor: 'rgb(253, 246, 236)',
    borderColor: 'rgb(250, 236, 216)'
  },
  // 错误消息样式
  error: {

  },
  // 成功消息样式
  info: {

  },
  // 自定义消息样式
  custom: messageData.customStyle
})
// 是否展示消息弹窗
const show = ref(false)

// 挂载钩子函数
onMounted(() => {
  show.value = true
  setTimeout(() => {
    show.value = false
  }, messageData.duration)
})
</script>

<style scoped>
.message-m-enter-active {
  transition: all ease-out .3s;
}

.message-m-leave-active {
  transition: all cubic-bezier(1, 0.5, 0.8, 1) .8s;
}

.message-m-enter-from, .message-m-enter-to {
  transform: translateX(20px);
  opacity: 0;
}

.message-container {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  position: relative;
  z-index: 99 !important;
  visibility: visible;
}

.message {
  margin: 5px auto;
  padding: 0 1rem;
  color: #222;
  line-height: 1.8rem;
  border-radius: 4px;
  z-index: 3;
}

.message .text {
  vertical-align: auto;
  text-align: center;
}
</style>