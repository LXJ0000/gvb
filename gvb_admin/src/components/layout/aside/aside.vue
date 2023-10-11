<template>
    <a-layout-sider
        v-model:collapsed="collapsed"
        collapsible
        theme="light"

    >
<!--      :style="{ overflow: 'auto', height: '100vh', position: 'fixed', left: 0, top: 0, bottom: 0 }"-->

      <div class="gvb_aside_header flex">
        <div>
          <img class="gvb_aside_header_logo flex" src="/static/avatar/6.jpg"
               alt="">
        </div>
        <div v-if="!collapsed" class="gvb_aside_header_title">
          Jannan后台管理系统
        </div>
      </div>

      <div class="gvb_aside_body">
        <AsideMenu></AsideMenu>
      </div>
    </a-layout-sider>
</template>
<script lang="ts" setup>
import {ref} from 'vue';
import AsideMenu from "@/components/layout/aside/aside_menu.vue";

const collapsed = ref<boolean>(false);
</script>
<style scoped lang="scss">

.gvb_aside_header {
  width: 100%;
  height: 180px;
  flex-direction: column;

  .gvb_aside_header_logo {
    width: 60px;
    height: 60px;
    border-radius: 50%;
  }

  .gvb_aside_header_title {
    margin-top: 20px;
    font-size: 20px;
    color: var(--sys);
  }
}
</style>
