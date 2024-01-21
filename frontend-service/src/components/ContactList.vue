<!-- 联系方式列表 -->
<template>
  <div class="contact-list-container center--flex">
    <div class="contact-list">
      <div
          v-for="item in listData.list"
          class="contact-app-box"
          :ref="setList"
          @click="handleItemClick(item)"
      >
        <a class="contact-app-icon-box">
          <svg-icon :icon-class="item.icon" class="contact-app-icon"></svg-icon>
          <div
              class="contact-app-qrcode-box center--flex"
              v-show="item.hasOwnProperty('qrcode') && !item.hasOwnProperty('action')"
          >
            <img
                :src="item.qrcode"
                class="contact-app-qrcode"
                alt=""
            >
          </div>
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import {onMounted, reactive, ref} from "vue";
import SvgIcon from "./SvgIcon.vue";

// 列表数据
const listData = defineProps({
  // 列表信息
  list: {
    // 联系方式列表数据校验
    validator(value) {
      if(value && value instanceof Array && value.length > 0) {
        let flag = true
        value.forEach(s => {
          if(!s.hasOwnProperty('icon') ||
              !s.hasOwnProperty('name') ||
              !(s.hasOwnProperty('url') ||
                  s.hasOwnProperty('qrcode') ||
                  s.hasOwnProperty('action'))) {
            if(flag) flag = false
          }
          if(s.hasOwnProperty('action') && typeof s.action !== 'string') {
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
    default: '0.5rem'
  },
  // 列表项列间距
  itemColumnGap: {
    type: String,
    default: '0.5rem'
  }
})

const iconSettings = reactive({
  customClass: 'contact-app-i'
})

// 父事件传递
const emit = defineEmits(['action'])

const row = ref(3) // 列表行数
const column = ref(8) // 列表列数
const listWidth = ref(24) // 列表宽度
const setList = (el) => {
  list.value.push(el)
}
const list = ref([]) // 列表项 dom 数组

// 处理列表项点击事件
const handleItemClick = (item) => {
  if(!item) return
  // 存在绑定事件执行事件
  if(item.hasOwnProperty('action')) {
    emit('action', item.action)
  } else {
    // 存在链接在新标签页跳转
    if(item.url != null && item.url !== '') {
      window.open(item.url, '_blank')
    }
  }
}

// 列表大小自适应
const resize = (columns, rows) => {
  // update size
  row.value = rows
  column.value = columns
  // reset items
  let visitable = rows * column.value
  let len = list.value.length
  visitable = visitable > len ? len : visitable
  for(let i = 0; i < len; ++i) {
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
/* 联系方式列表容器 */
.contact-list-container {
  width: 100%;
  justify-content: flex-start;
}

/* 联系方式列表 */
.contact-list {
  display: grid;
  grid-template-columns: repeat(v-bind(column), 1fr);
  width: v-bind(listWidth)rem;
  column-gap: v-bind('listData.itemColumnGap');
  row-gap: v-bind('listData.itemRowGap');
}

/* 联系方式列表项 */
.contact-app-box {
  width: v-bind('listData.itemSize');
  height: v-bind('listData.itemSize');
  text-align: center;
  display: inline-flex;
  justify-content: center;
  align-items: center;
}

/* 联系方式列表图标项 */
.contact-app-icon-box {
  display: inline-flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  width: v-bind('listData.itemSize');
  height: v-bind('listData.itemSize');
  border-radius: 50%;
  position: relative;
  cursor: pointer;
}

/* 联系方式列表图标 */
.contact-app-icon {
  width: 100%;
  height: 100%;
}

/* 联系方式列表项二维码 */
.contact-app-qrcode-box {
  position: absolute;
  bottom: 0;
  transform: translateX(-25%) translateY(105%);
  width: 5rem;
  height: 5rem;
  background-color: rgb(255, 255, 255, 0.8);
  border-radius: 5px;
  z-index: 2;
  display: none;
  align-items: center;
  justify-content: center;
}

/* 联系方式列表项二维码 - 悬浮展示 */
.contact-app-icon-box:hover .contact-app-qrcode-box {
  display: flex;
}

/* 联系方式列表项二维码图片 */
.contact-app-qrcode {
  width: 98%;
  height: 98%;
  border-radius: 5px;
}
</style>