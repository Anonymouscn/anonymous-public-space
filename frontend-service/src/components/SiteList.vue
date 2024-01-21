<!-- 站点列表 -->
<template>
  <div class="site-list-container">
    <a v-for="item in listData.list" class="site-item" :href="item.url">{{ item.name }}</a>
  </div>
</template>

<script setup>

// 列表数据
import {ref} from "vue";

const listData = defineProps({
  // 列表信息
  list: {
    validator(value) {
      if(value && value instanceof Array && value.length > 0) {
        let flag = true
        value.forEach(s => {
          if(!s.hasOwnProperty('url') || !s.hasOwnProperty('name')) {
            if(flag) flag = false
          }
        })
        return flag
      }
      return false
    },
    default: () => [
      {name: 'Space', url: '#'},
      {name: 'Message', url: '#'},
      {name: 'Chat', url: '/'}
    ]
  }
})

// 列表大小自适应
const resize = (info) => {
  // let itemWidth = itemSize.value
  // if(info) {
  //   column.value = parseInt(info.width / 53)
  //   listWidth.value = itemWidth * column.value
  // }
}

// 暴露函数调用
defineExpose({
  resize
})
</script>

<style scoped>
.site-item {
  color: #fff;
  text-decoration: none;
  font-size: 1.6rem;
  display: inline;
  margin: 0.5rem;
}
</style>