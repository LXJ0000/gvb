<script setup>
import {useRouter} from "vue-router";
import {ref} from "vue";

const router = useRouter()
const menuClick = ({key}) => {
  if (key === "logout") {
    console.log("logout")
    return
  }
  router.push({
    name: key,
  })
  // console.log(`Click on item ${key}`);
};


// true代表白天 false代表黑夜
const theme = ref(true)
const setTheme = () => {
  theme.value = !theme.value
  if (theme.value) {
    // 白天
    document.documentElement.classList.remove("dark")
  } else {
    // 黑夜
    document.documentElement.classList.add("dark")

  }
}


// 全屏
const screen = ref(false)
const setScreen = () =>{
  screen.value = !screen.value
  if(screen.value){

  }else {}

}


</script>

<template>
  <div class="gvb_admin">
    <aside>
      <!--侧边栏-->
      <div>

      </div>

    </aside>

    <div class="main">
      <header>
        <!--顶部导航栏-->
        <div class="left">
          <a-breadcrumb>
            <a-breadcrumb-item>首页</a-breadcrumb-item>
            <a-breadcrumb-item><a href="">个人中心</a></a-breadcrumb-item>
            <a-breadcrumb-item>用户列表</a-breadcrumb-item>
          </a-breadcrumb>
        </div>
        <div class="right">
          <div class="icon-action">
            <i class="iconfont icon-home"></i>
            <i v-if="theme" class="iconfont icon-yueliang" @click="setTheme"> </i>
            <i v-else class="iconfont icon-liangdu_huaban" @click="setTheme"> </i>
            <i v-if="screen" class="iconfont icon-tuichuquanping" @click="setScreen"> </i>
            <i v-else class="iconfont icon-quanping1" @click="setScreen"> </i>
          </div>

          <div class="avatar">
            <img src="http://s0yvn46ks.hn-bkt.clouddn.com/gvb/1695639074__2.jpg" alt="">
          </div>

          <div class="drop-menu">
            <a-dropdown placement="bottomLeft">
              <a class="ant-dropdown-link" @click.prevent>
                Jannan
                <span class="iconfont icon-xiala"></span>
              </a>
              <template #overlay>
                <a-menu @click="menuClick">
                  <a-menu-item key="user_center">
                    <a href="javascript:;">个人中心</a>
                  </a-menu-item>
                  <a-menu-item key="my_message">
                    <a href="javascript:;">我的消息</a>
                  </a-menu-item>
                  <a-menu-item key="article_list">
                    <a href="javascript:;">文章列表</a>
                  </a-menu-item>
                  <a-menu-item key="logout">
                    <a href="javascript:;">注销退出</a>
                  </a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </div>
        </div>
      </header>
      <div class="tabs"></div>
      <main></main>
    </div>

  </div>
</template>

<style scoped lang="scss">
.gvb_admin {
  width: 100%;
  display: flex;

  aside {
    width: 240px;
    height: 100vh;
    background-color: #2b3539;
  }

  .main {
    width: calc(100% - 240px);

    header {
      height: 60px;
      background-color: white;
      padding: 0 30px;

      display: flex;
      justify-content: space-between;
      align-items: center;

      .right {

        display: flex;
        justify-content: space-between;
        align-items: center;

        .icon-action {
          margin-right: 20px;

          i {
            margin-inline: 10px;
            cursor: pointer;
            font-size: 18px;
            color: var(--test);
          }

          i:hover {
            color: var(--active);
          }

        }

        .avatar {
          img {
            width: 40px;
            height: 40px;
            border-radius: 50%;
          }

        }

        .drop-menu {
          margin-left: 10px;
        }
      }
    }

    .tabs {
      height: 30px;
      border: 1px solid #f0eeee;
    }

    main {
      background-color: var(--bg);
      height: calc(100vh - 90px);
    }


  }

}

</style>