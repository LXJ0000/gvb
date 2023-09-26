<script setup>

import {useRouter} from "vue-router";

const router = useRouter()
const menuClick = ({key}) => {
  if (key === "logout") {
    console.log("logout")
    return
  }
  router.push(
      {
        name: key,
      },
  )
}


</script>

<template>
  <div class="gvb_user_info flex">
    <div class="avatar">
      <img src="http://s0yvn46ks.hn-bkt.clouddn.com/gvb/1695639074__6.jpg" alt="">
    </div>
    <div class="drop_menu">
      <a-dropdown placement="bottom">
        <a class="ant-dropdown-link" @click.prevent>
          Janan
          <i class="fa fa-angle-down"></i>
        </a>
        <template #overlay>
          <a-menu @click="menuClick">
            <a-menu-item key="user_center">
              <a href="javascript:">个人中心</a>
            </a-menu-item>
            <a-menu-item key="my_message">
              <a href="javascript:">我的消息</a>
            </a-menu-item>
            <a-menu-item key="article_list">
              <a href="javascript:">文章列表</a>
            </a-menu-item>
            <a-menu-item key="logout">
              <a href="javascript:">注销退出</a>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </div>
  </div>

</template>

<style scoped lang="scss">
.avatar {
  img {
    width: 40px;
    height: 40px;
    border-radius: 50%;
  }
}

.drop_menu {
  margin-left: 15px;
}
</style>