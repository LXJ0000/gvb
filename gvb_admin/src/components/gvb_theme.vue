<script setup>
import {ref} from "vue";
import {useStore} from "@/stores/store";

const store = useStore()
store.loadTheme()
const setTheme = () =>{
  store.setTheme()
}
</script>

<template>
  <i @click="setTheme" v-if="store.theme" class="iconfont icon-yueliang"></i>
  <i @click="setTheme" v-else class="iconfont icon-taiyang1"></i>
</template>

<style scoped lang="scss">

</style>