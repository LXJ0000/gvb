<template>
  <a-layout-sider collapsible breakpoint="xl" theme="light">

    <div class="logo"/>

    <a-menu
        :default-open-keys="['1']"
        :default-selected-keys="['0_3']"
        :style="{ width: '100%' }"
        @menu-item-click="onClickMenuItem"
    >

      <a-menu-item key="0_1">
        <IconHome></IconHome>
        首页
      </a-menu-item>

      <a-menu-item key="0_2">
        <IconCalendar></IconCalendar>
        个人中心
      </a-menu-item>

      <a-sub-menu key="4">
        <template #title>
          <IconCalendar></IconCalendar>
          用户管理
        </template>
        <a-menu-item key="4_1">用户列表</a-menu-item>
        <a-menu-item key="4_2">添加用户</a-menu-item>
      </a-sub-menu>
    </a-menu>

    <!-- trigger -->
    <template #trigger="{ collapsed }">
      <IconCaretRight v-if="collapsed"></IconCaretRight>
      <IconCaretLeft v-else></IconCaretLeft>
    </template>

  </a-layout-sider>
</template>

<script setup>
import {Message} from "@arco-design/web-vue";
import {
  IconCaretRight,
  IconCaretLeft,
  IconHome,
  IconCalendar,
} from '@arco-design/web-vue/es/icon';

const onClickMenuItem = (key) => {
  Message.info({content: `You select ${key}`, showIcon: true});
}
</script>

