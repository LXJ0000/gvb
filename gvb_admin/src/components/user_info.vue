<template>
  <a-space size="small">
    <a-avatar
        imageUrl="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/3ee5f13fb09879ecb5185e440cef6eb9.png~tplv-uwbnlip3yd-webp.webp"
    >
    </a-avatar>

    <a-dropdown trigger="hover" @select="menuClick">
      <a-button style="background-color: white; padding: 0">
        Jannan
        <i class="iconfont icon-xiala" ></i>
      </a-button>
      <template #content>
        <a-doption :value="{value:'user_info'}">个人中心</a-doption>
        <a-doption :value="{value:'user_message'}">我的消息</a-doption>
        <a-doption :value="{value:'article_list'}">文章列表</a-doption>
        <a-doption :value="{value:'logout'}">注销退出</a-doption>
      </template>
    </a-dropdown>


  </a-space>
</template>
<script setup>
const menuClick = (v) =>{
  console.log(v.value)
}

</script>

<style scoped>
.arco-dropdown-open .iconfont {
  transform: rotate(180deg);
}
</style>