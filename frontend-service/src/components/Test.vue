<script setup>
import {ref} from "vue";

const show = ref(false)

const handleClick = () => {
  show.value = !show.value
}
</script>

<template>
  <div class="container">
    <a @click="handleClick">Click to show</a>
    <Transition name="slide-fade">
<!--      <p v-if="show">Hello !</p>-->
      <div class="sidebar" v-show="show">

      </div>
    </Transition>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

a {
  margin-top: 1rem;
  padding: 0.5rem 1rem;
  border: 2px solid #aaa;
  cursor: pointer;
  border-radius: 5px;
}

p {
  line-height: 1.4rem;
  margin-top: 1rem;
}

.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}

.slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}

.sidebar {
  position: absolute;
  right: 0;
  width: 30rem;
  height: 100%;
  background-color: #666;
}
</style>