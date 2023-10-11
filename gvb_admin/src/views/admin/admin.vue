<template>
  <a-config-provider
      :theme="{
      algorithm: theme.defaultAlgorithm,
      components:{
        Layout:{
          colorBgHeader:'#ffffff',

        }
      }
    }"
  >


    <a-layout style="min-height: 100vh; ">

      <Aside></Aside>

      <a-layout
      >
        <Header></Header>


        <a-layout-content
            style="padding-top: 10px"
        >
          <Tabs></Tabs>
          <main>
            <div class="view">
              <router-view v-slot="{Component}">
                <transition name="fade" mode="out-in">
                  <component :is="Component"></component>
                </transition>
              </router-view>
            </div>
          </main>
        </a-layout-content>

        <a-layout-footer style="text-align: center">
          Ant Design ©2018 Created by Ant UED
        </a-layout-footer>

      </a-layout>
    </a-layout>
  </a-config-provider>

</template>
<script lang="ts" setup>
import Header from '@/components/layout/header.vue'
import Aside from "@/components/layout/aside/aside.vue";
import Tabs from '@/components/layout/tabs.vue'
import {theme} from 'ant-design-vue';

</script>
<style scoped lang="scss">


//.site-layout .site-layout-background {
//  background: #fff;
//}
//
//[data-theme='dark'] .site-layout .site-layout-background {
//  background: #141414;
//}



main {
  padding: 10px 20px 0 20px;
}


</style>

<style>
.fade-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.fade-enter-active {
  transform: translateX(-30px);
  opacity: 0;
}

.fade-enter-to {
  transform: translateX(0px);
  opacity: 1;
}

.fade-leave-active, .fade-enter-active {
  transition: all 0.3s ease-out;
}
</style>