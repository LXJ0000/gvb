<script setup>
import {useStore} from "@/stores/store";
import {useRoute, useRouter} from "vue-router";

const store = useStore()
const route = useRoute()
const router = useRouter()
const isActive = item => {
//   判断当前路由是否和item匹配
//   判断是不是home
  if (route.name === item.name) {
    return `tabs_item active`
  }

  return `tabs_item`
}

const checkTab = (item) => {
  router.push({name: item.name})
}
store.loadTabs()
window.onbeforeunload = function () {
  store.saveTabs()
  return undefined
}

const removeTab = item => {
  let removeIndex = store.removeTab(item)
//   如果删除当前所在的 则向前走一步
  if(item.name === route.name){
    let beforeName = store.tabList[removeIndex - 1].name
    router.push({name:beforeName})
  }
}

const removeAllTab = ()=>{
  store.removeAllTab()
  router.push({name:"home"})
}
</script>

<template>
  <div class="tabs">
    <div @click="checkTab(item)" :class="isActive(item)" v-for="(item, index) in store.tabList" :key="index">
      {{ item.title }}
      <i @click.stop="removeTab(item)" v-if="item.name !== 'home'" class="tabs_item_icon iconfont icon-quxiao"></i>
    </div>

    <div class="tabs_item remove_all " @click="removeAllTab">全部关闭</div>

  </div>
</template>

<style scoped lang="scss">
.tabs {
  border-radius: 20px;
  height: 35px;
  background-color: var(--inbg);
  display: flex;
  justify-content: left;
  align-items: center;
  position: relative;

  .tabs_item {
    background-color: var(--outbg);
    border: 1px solid;
    padding: 2px 10px;
    margin-right: 10px;
    cursor: pointer;
    display: flex;
    align-items: center;
    border-radius: 10px;

    .tabs_item_icon {
      font-size: 15px;
      margin-left: 8px;
      height: 15px;
      width: 15px;
      cursor: pointer;
      display: flex;
      justify-content: center;
      align-items: center;

      &:hover {
        background-color: #CBE9BF;
        border-radius: 50%;
      }
    }

    &:first-child {
      margin-left: 20px;
    }

    &.active {
      background-color: var(--sys);
      border-color: var(--sys);
      color: white;

      &::before {
        display: inline-block;
        content: "";
        width: 12px;
        height: 12px;
        background-color: white;
        border-radius: 50%;
        margin-right: 5px;
      }

      .tabs_item_icon:hover {
        background-color: #8fbef3;
      }
    }

    &.remove_all {
      position: absolute;
      right: 20px;
    }

  }
}
</style>