<script setup>
import {computed, reactive} from "vue";
import {baseListApi} from "@/api/base";
import {message} from "ant-design-vue";
import {SyncOutlined} from '@ant-design/icons-vue';
import {formatDateTime} from "@/utils/time";

//---------------------------------------------------------------------
// 发送
const emits = defineEmits(["removeBatch"])
// 接收
const props = defineProps({
  baseUrl: {
    type: String
  },
  options: {
    type: Array
  },
  columns: {
    type: Array
  },
  searchBtnName: {
    type: String
  },
  isAdd: {
    type: Boolean,
    default: false,
  },
  isEdit: {
    type: Boolean,
    default: false,
  },
  pageSize: {
    type: Number,
    default: 5
  }
})

//---------------------------------------------------------------------

const data = reactive({
  spinning: true, //加载ing
  dataSource: [],
  cnt: 0,
  page: {
    page: 1,
    limit: props.pageSize,
    key: "",
  },
})


const onSearch = () => {
  data.page.page = 1
  data.page.key = data.page.key.trim()
  GetData(data.page)
}

const GetData = async (params) => {
  data.spinning = true
  let res = await baseListApi(props.baseUrl, params)
  data.dataSource = res.data.list
  data.cnt = res.data.count
  data.spinning = false
}
GetData(data.page)

const refresh = () => {
  GetData(data.page)
  message.success("刷新成功")
}

const removeBatch = async () => {
  emits("removeBatch", state.selectedRowKeys)
  state.selectedRowKeys = []
}

const pageChange = (_page, _limit) => {
  GetData(data.page)
}


//---------------------------------------------------------------------
const state = reactive({
  selectedRowKeys: [], // 选择的数据 id
  loading: false,
});
const hasSelected = computed(() => state.selectedRowKeys.length > 0);
const start = () => {
  state.loading = true;
  setTimeout(() => {
    state.loading = false;
    state.selectedRowKeys = [];
  }, 1000);
};
const onSelectChange = selectedRowKeys => {
  state.selectedRowKeys = selectedRowKeys;
};
//---------------------------------------------------------------------
const ExportList = (params) => {

  if (params === undefined) {
    params = {}
  }
  data.page.key = data.page.key.trim()

  Object.assign(data.page, params)
  GetData(data.page)
}
defineExpose({
  ExportList
})
//---------------------------------------------------------------------

</script>

<template>
  <div class="container">
    <div class="container_searchAndRefresh">
      <div class="container_search">
        <a-input-search
            :placeholder="props.searchBtnName"
            style="width: 200px"
            @search="onSearch"
            v-model:value="data.page.key"
        />

        <slot name="filter"></slot>

      </div>
      <div class="container_refresh">
        <a-button @click="refresh">
          <SyncOutlined/>
        </a-button>
      </div>
    </div>

    <div class="container_table">
      <div class="container_table_action">
        <div class="container_table_action_left">
          <a-button type="primary" :disabled="!hasSelected" :loading="state.loading" @click="start">
            重置
          </a-button>
          <slot v-if="isAdd" name="addModal"/>
          <a-popconfirm
              title="是否确定删除"
              ok-text="删除"
              cancel-text="取消"
              @confirm="removeBatch"
          >
            <a-button type="primary" danger v-if="state.selectedRowKeys.length">批量删除</a-button>
          </a-popconfirm>
        </div>
      </div>

      <div class="container_table_action_body">
        <a-spin :spinning="data.spinning" tip="加载中. . ." :delay="300">
          <a-table
              :row-selection="{ selectedRowKeys: state.selectedRowKeys, onChange: onSelectChange }"
              :row-key="'id'"
              :dataSource="data.dataSource"
              :columns="props.columns"
              :pagination="false"
          >

            <template #bodyCell="{ column, record }">
              <slot name="cell" v-bind="{ column, record }">
                <template v-if="column.key === 'created_at'">
                  <div>{{ formatDateTime(record.created_at) }}</div>
                </template>
                <template v-if="column.key === 'action'">
                  <slot v-if="isEdit" name="edit" v-bind="{ column, record }"/>
                </template>
              </slot>
            </template>
          </a-table>
        </a-spin>
      </div>

      <div class="container_table_page flex">
        <a-pagination
            v-model:current="data.page.page"
            v-model:page-size="data.page.limit"
            :total="data.cnt"
            show-less-items
            @change="pageChange"
        />
      </div>
    </div>
  </div>
</template>

<style lang="scss">
.container {
  background-color: var(--inbg);
  border-radius: 10px;

  .container_searchAndRefresh {
    padding: 10px;
    border-bottom: 1px solid var(--outbg);
    position: relative;

    .container_refresh {
      position: absolute;
      right: 10px;
      top: 10px;
    }
  }

  .container_table {
    padding: 10px;

    .container_table_action {
      padding-bottom: 10px;

      .container_table_action_left {
        .ant-btn {
          margin-right: 10px;
        }
      }

    }


    .container_table_page {
      padding: 10px;
      margin-bottom: 0;
    }
  }
}
</style>