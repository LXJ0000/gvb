<script setup>

</script>

<template>
  <aside class="gvb_aside">
    <div>
      <div class="gvb_aside_header flex">
        <div>
          <img class="gvb_aside_header_logo flex" src="http://s0yvn46ks.hn-bkt.clouddn.com/gvb/1695639074__6.jpg" alt="">
        </div>
        <div class="gvb_aside_header_title flex">
          Jannan后台管理系统
        </div>
      </div>
      <div class="gvb_aside_body"></div>
    </div>
  </aside>
</template>

<style scoped lang="scss">
.gvb_aside {
  width: 240px;
  height: 100vh;
  background-color: var(--slide);

  .gvb_aside_header {
    width: 100%;
    height: 180px;
    flex-direction: column;

    .gvb_aside_header_logo {
      width: 60px;
      height: 60px;
    }
    .gvb_aside_header_title{
      margin-top: 20px;
      font-size: 20px;
      color: white;
    }

  }
}
</style>