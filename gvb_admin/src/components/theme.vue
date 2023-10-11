<script setup>
import {ref} from "vue";
const theme = ref(true)
const setTheme = () =>{
  theme.value = !theme.value
}
</script>

<template>
  <i @click="setTheme" v-if="theme" class="iconfont icon-yueliang"></i>
  <i @click="setTheme" v-else class="iconfont icon-taiyang1"></i>
</template>

<style  lang="scss">

</style>