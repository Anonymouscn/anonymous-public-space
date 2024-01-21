<template>
  <svg :class="svgClass" aria-hidden="true" class=".svg">
    <use :xlink:href="iconName" :fill="color" :class="customClass" :width="width" :height="height"/>
  </svg>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  iconClass: {
    type: String,
    required: true
  },
  className: {
    type: String,
    default: ''
  },
  color: {
    type: String,
    default: '#fff'
  },
  width: {
    type: String,
    default:'98%'
  },
  height: {
    type: String,
    default: '98%'
  },
  customClass: {
    type: String,
    default: ''
  }
})

const iconName = computed(() => {
  return `#icon-${props.iconClass}`
})

const svgClass = computed(() => {
  if (props.className) {
    return `svg-icon ${props.className}`
  }
  return 'svg-icon'
})
</script>

<style scoped>
.svg {
  position: relative;
}
.icon {
  position: absolute;
  width: 98%;
  height: 98%;
}
</style>