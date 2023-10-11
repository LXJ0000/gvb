<script setup>
import {reactive} from "vue";

const data = reactive([
  {
    icon: "icon-fangwenliang",
    label: "访问量",
    value: 666,
  },
  {
    icon: "icon-yonghuliang",
    label: "用户量",
    value: 233,
  },
  {
    icon: "icon-wenzhang",
    label: "文章量",
    value: 6,
  },
  {
    icon: "icon-liuyan",
    label: "留言量",
    value: 123,
  },
])
</script>

<template>
  <div>
    <div class="data_preview">
      <div class="data_preview_item">
        <a-row :gutter="16">

          <a-col :span="6" v-for="(card) in data" :key="card.label">
            <a-card :bordered="false">
              <div class="flex">
                <div class="icon">
                  <i :class="'iconfont ' + card.icon"></i>
                </div>
                <div class="text">
                  <div class="data_title flex">{{ card.label }}</div>
                  <div class="data_num flex">{{ card.value }}</div>
                </div>
              </div>
            </a-card>
          </a-col>

        </a-row>

      </div>

    </div>
  </div>

</template>

<style scoped lang="scss">
.data_preview {
  .data_preview_item {

    .icon {
      width: 60%;

      i {
        font-size: 45px;
        color: var(--sys);
      }
    }

    .text {
      width: 40%;

      .data_title {
        font-size: 14px;
        color: #767C82;
      }

      .data_num {
        font-size: 24px;
        color: #333639;
        font-weight: 600;
      }
    }
  }
}

</style>