<script setup>
import {message} from 'ant-design-vue';
import {reactive} from "vue";
import {emailLoginApi} from "@/api/user";
import jwt_decode from "jwt-decode";
import {useStore} from "@/stores/store";
import {useRoute, useRouter} from "vue-router";

const router = useRouter()
const route = useRoute()
const store = useStore()
const data = reactive({
  user_name: "",
  password: "",
})

async function emailLogin() {
  if (data.user_name.trim() === "") {
    message.error("请输入用户名")
    return
  }
  if (data.password.trim() === "") {
    message.error("请输入用户密码")
  }

  let res = await emailLoginApi(data)
  if (res.code) {
    message.error(res.msg)
    return
  }
  message.success(res.msg)
  const userinfo = jwt_decode(res.data); // 使用jwt_decode解析token
  userinfo.token = res.data
  // 维护store
  store.setUserInfo(userinfo)

//   页面跳转
//   原页面
//   home
  const redirect_url = route.query.redirect_url

  if (redirect_url === undefined) {
    setTimeout(() => {
      router.push({
        name: "home"
      })
    }, 1000)
    return
  }
  setTimeout(() => {
    router.push({
      path: redirect_url,
    })
  }, 1000)
}
</script>

<template>
  <div class="login_bg">
    <div class="login_bg_container flex">
      <div class="main">
        <div class="title">用户登录</div>
        <div class="form">
          <div class="form_item">
            <a-input v-model:value="data.user_name" placeholder="User Name" type="username">
              <template #prefix>
                <i class="iconfont icon-username"></i>
              </template>
            </a-input>
          </div>

          <div class="form_item">
            <a-input-password @pressEnter="emailLogin" v-model:value="data.password" type="password" placeholder="Password">
              <template #prefix>
                <i class="iconfont icon-password"></i>
              </template>
            </a-input-password>
          </div>

          <div class="form_item">
            <a-button @click="emailLogin" html-type="submit" type="primary">Login</a-button>
          </div>

        </div>
        <div class="other flex">
          第三方登陆
        </div>

        <div class="icons flex">
          <img class="icon" src="../assets/img/QQ.png" alt="">
        </div>

      </div>
    </div>

  </div>

</template>

<style scoped>
.login_bg {
  background: url(/static/file/1696250922__20.jpg) 50%/cover no-repeat;
  width: 100%;
  height: 100vh;

  .login_bg_container {
    position: fixed;
    right: 0;
    top: 0;
    width: 300px;
    height: 100vh;
    background-color: rgba(256, 256, 256, 0.5);

    .main {
      width: 70%;

      .title {
        font-size: 22px;
        font-weight: 600;
        color: var(--sys);
        text-align: center;
        margin-bottom: 10px;
      }

      .form {

        .form_item {
          margin-bottom: 10px;

          .ant-btn {
            width: 100%;
            background-color: var(--sys);
          }
        }
      }

      .other {
        font-size: 4px;
        color: var(--sys);

        &::before {
          width: 30%;
          height: 1px;
          background-color: var(--sys);
          content: "";
          display: inline-block;
          margin-right: 5px;
        }

        &::after {
          width: 30%;
          height: 1px;
          background-color: var(--sys);
          content: "";
          display: inline-block;
          margin-inline: 5px;
        }
      }

      .icons {
        margin-top: 10px;

        .icon {
          width: 40px;
          height: 40px;
          cursor: pointer;
          margin-right: 10px;
        }

        &:last-child {
          margin-right: 0;
        }

      }
    }
  }
}
</style>