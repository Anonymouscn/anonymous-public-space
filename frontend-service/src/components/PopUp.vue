<!-- 弹窗组件 -->
<template>
  <!-- 弹窗容器 -->
  <div class="pop-up-container center--flex" v-show="popUpData.isShow" ref="popUp">
    <!-- 内容区 -->
    <div class="content">
      <!-- header 插槽 -->
      <slot name="header" v-if="showHeader">
        <div class="header center--flex">
          <div class="title center--flex">默认标题</div>
        </div>
      </slot>
      <slot name="body">
        <div class="body center--flex">默认弹窗内容</div>
      </slot>
      <slot name="buttons" v-if="showButtons">
        <div class="buttons-box center--flex">
          <a class="button center--flex" @click="cancel">取消</a>
          <a class="button center--flex" @click="confirm">确定</a>
        </div>
      </slot>
    </div>
  </div>
</template>

<script setup>
import {onMounted, onUnmounted, ref} from "vue";

// 弹窗数据
const popUpData = defineProps({
  // 弹窗宽度
  width: {
    type: String,
    default: () => '16rem'
  },
  // 弹窗高度
  height: {
    type: String,
    default: () => '16rem'
  },
  // 背景颜色
  backgroundColor: {
    type: String,
    default: () => '#fff',
  },
  // 是否展示弹窗
  isShow: {
    type: Boolean,
    default: true
  },
  // 是否展示 header
  showHeader: {
    type: Boolean,
    default: true
  },
  // 是否展示按钮组
  showButtons: {
    type: Boolean,
    default: true
  },
  // 点击空白关闭弹窗
  clickBlankToClose: {
    type: Boolean,
    default: true
  }
})

// 父事件传递
const emit = defineEmits(['close', 'cancel', 'confirm'])

// 取消事件 [父组件实现]
const cancel = () => {
  emit('cancel')
}

// 确定事件 [父组件实现]
const confirm = () => {
  emit('confirm')
}

// 关闭事件 [父组件实现]
const close = () => {
  emit('close')
}

const popUp = ref(null) // popUp 引用

// 空白点击处理
function handleClickBlank(e) {
  if(popUp.value === e.target && popUpData.clickBlankToClose) {
    close()
  }
}

// 挂载钩子函数
onMounted(() => {
  document.addEventListener("click", handleClickBlank)
})

// 卸载钩子函数
onUnmounted(() => {
  document.removeEventListener("click", handleClickBlank)
})
</script>

<style scoped>
/* 弹窗容器 */
.pop-up-container {
  width: 100vw;
  height: 100%;
  position: fixed;
  background-color: rgb(0, 0, 0, 0.3);
  z-index: 20;
}

/* 内容区 */
.content {
  width: v-bind('popUpData.width');
  height: v-bind('popUpData.height');
  background-color: v-bind('popUpData.backgroundColor');
  border-radius: 20px;
  position: relative;
}

/* 默认 header 区域 */
.header {
  height: 3rem;
  border-top: 20px;
  border-bottom: 2px solid #eee;
}

/* 默认标题 */
.title {
  font-size: 1.2rem;
}

/* 默认主体 */
.body {
  min-height: 8rem;
}

/* 默认按钮区域 */
.buttons-box {
  border-top: 2px solid #eee;
  position: absolute;
  bottom: 0;
  height: 3rem;
  width: 100%;
}

/* 默认按钮 */
.button {
  width: 50%;
  height: 100%;
  cursor: pointer;
  border-right: 2px solid #eee;
}

.button:hover {
  background-color: rgb(240, 240, 240);
}

.button:first-child {
  border-bottom-left-radius: 20px;
}

.button:last-child {
  border-bottom-right-radius: 20px;
  border-right: none;
}
</style>