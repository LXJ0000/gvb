<script setup>
import {DownOutlined} from '@ant-design/icons-vue';

import {useRouter, useRoute} from "vue-router";

const props = defineProps({
  // 是否显示图像
  isAvatar: {
    type: Boolean,
    default: false,
  }
})
const route = useRoute()
const router = useRouter()
const menuClick = ({key}) => {
  if (key === "logout") {
    console.log("logout")
    return
  }
  if (key === "login") {
    router.push({
      name: key,
      query: {
        redirect_url: route.path
      }
    })
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

    <div v-if="props.isAvatar" class="avatar">
      <img src="/static/avatar/6.jpg" alt="">
    </div>

    <div class="drop_menu">
      <a-dropdown placement="bottom">

        <a class="ant-dropdown-link" @click.prevent>
          Jannan
          <DownOutlined/>
        </a>

        <template #overlay>
          <a-menu @click="menuClick">
            <a-menu-item key="user_center">
              个人中心
            </a-menu-item>
            <a-menu-item key="my_message">
              我的消息
            </a-menu-item>
            <a-menu-item key="article_list">
              文章列表
            </a-menu-item>
            <a-menu-item key="login">
              用户登录
            </a-menu-item>
            <a-menu-item key="logout">
              注销退出
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </div>

  </div>

</template>

<style lang="scss">
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