<!-- 技术栈列表 -->
<template>
  <div class="tech-stack-list-container center--flex">
    <div class="tech-stack-list">
      <div 
          v-for="item in listData.list" 
          class="tech-stack-box center--flex"
          :ref="setList"
      >
        <a class="tech-stack-icon-box">
          <svg-icon :icon-class="item.icon" class="tech-stack-icon"></svg-icon>
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import {onMounted, ref} from "vue";
import SvgIcon from "./SvgIcon.vue";

// 列表数据
const listData = defineProps({
  // 列表信息
  list: {
    // 技术栈列表数据校验
    validator(value) {
      if(value && value instanceof Array && value.length > 0) {
        let flag = true
        value.forEach(s => {
          if(!s.hasOwnProperty('icon') || !s.hasOwnProperty('name')) {
            if(flag) flag = false
          }
        })
        return flag
      }
      return false
    },
    default: () => []
  },
  // 列表项大小
  itemSize: {
    type: String,
    default: '2.4rem'
  },
  // 列表项行间距
  itemRowGap: {
    type: String,
    default: '0.2rem'
  },
  // 列表项列间距
  itemColumnGap: {
    type: String,
    default: '0.2rem'
  }
});

const row = ref(3) // 列表行数
const column = ref(8) // 列表列数
const listWidth = ref(24) // 列表宽度
const setList = (el) => {
  list.value.push(el)
}
const list = ref([]) // 列表项 dom 数组

// 列表大小自适应
const resize = (columns, rows) => {
  // update size
  row.value = rows
  column.value = columns
  // reset items
  let visitable = rows * columns
  let len = list.value.length
  visitable = visitable > len ? len : visitable
  for(let i = 0; i < visitable; ++i) {
    list.value[i].style.display = 'inline-flex'
  }
  for(let i = visitable; i < len; ++i) {
    list.value[i].style.display = 'none'
  }
}

// 列表项自适应
const resizeItem = () => {
  let len = list.value.length
  let str = listData.itemSize
  let size = (Number(str.substr(0, str.length - 3)) + 0.2) + 'rem'
  for(let i = 0; i < len; ++i) {
    list.value[i].style.width = size
    list.value[i].style.height = size
  }
}

// 挂载钩子函数
onMounted(() => {
  resizeItem()
})

// 暴露函数调用
defineExpose({
  resize
})
</script>

<style scoped>
/* 技术栈列表容器 */
.tech-stack-list-container {
  width: 100%;
  justify-content: flex-start;
}

/* 技术栈列表 */
.tech-stack-list {
  display: grid;
  grid-template-columns: repeat(v-bind(column), 1fr);
  width: v-bind(listWidth)rem;
  column-gap: v-bind('listData.itemColumnGap');
  row-gap: v-bind('listData.itemRowGap');
}

/* 技术栈列表项 */
.tech-stack-box {
  text-align: center;
  width: v-bind('listData.itemSize');
  height: v-bind('listData.itemSize');
}

/* 技术栈列表图标项 */
.tech-stack-icon-box {
  display: inline-flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  width: v-bind('listData.itemSize');
  height: v-bind('listData.itemSize');
  border-radius: 50%;
}

/* 技术栈列表图标 */
.tech-stack-icon {
  width: 100%;
  height: 100%;
}
</style>