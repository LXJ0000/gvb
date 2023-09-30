<template>
  <a-layout-header style="padding-left: 20px; padding-right: 20px; display: flex; justify-content: flex-end;">
    <div style="display: flex; align-items: center;">
      <div class="icon_actions">
        <i class="iconfont icon-quanping1"></i>
        <i class="iconfont icon-yueliang"></i>

      </div>

      <UserInfo></UserInfo>
    </div>
  </a-layout-header>
</template>

<script setup>
import UserInfo from "@/components/user_info.vue";
</script>

<style >
.icon_actions {
  margin-right: 20px;

  i {
    margin-left: 10px;
    cursor: pointer;
    font-size: 16px;
    color: #7BC0FC;
  }

  i:hover {
    color: #7BC0FC;
  }
}
</style>