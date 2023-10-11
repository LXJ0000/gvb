<script setup>

import UserInfo from "@/components/user_info.vue";
import FullScreen from "@/components/full_screen.vue";
import Theme from "@/components/theme.vue";
import {theme} from "ant-design-vue";
</script>

<template>
  <a-layout-header
      style="padding: 0"
  >
    <div class="header ">
      <div class="header_left flex">
        <a-breadcrumb>
          <a-breadcrumb-item>首页</a-breadcrumb-item>
          <a-breadcrumb-item><a href="">用户管理</a></a-breadcrumb-item>
          <a-breadcrumb-item><a href="">用户列表</a></a-breadcrumb-item>
          <a-breadcrumb-item>用户信息</a-breadcrumb-item>
        </a-breadcrumb>
      </div>
      <div class="header_right">
        <div class="icon_actions">
          <i class="iconfont icon-home"></i>
          <FullScreen></FullScreen>
          <Theme></Theme>
        </div>
        <UserInfo :isAvatar="true"></UserInfo>
      </div>
    </div>
  </a-layout-header>



</template>

<style scoped lang="scss">
.header {

  //height: 60px;
  padding: 0 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;

  .header_right {
    display: flex;
    align-items: center;
  }

  .icon_actions {
    margin-left: 20px;

    i {
      margin-right: 20px;
      cursor: pointer;
      font-size: 16px;
      color: black;
    }

    i:hover {
      color: var(--sys);
    }
  }
}
</style>