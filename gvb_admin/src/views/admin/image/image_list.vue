<script setup>
import MyTable from '@/components/table.vue'
import {reactive, ref} from "vue";
import {message} from "ant-design-vue";
import {imageRemoveBatchApi, imageUpdateApi} from "@/api/image";
import { PlusOutlined } from '@ant-design/icons-vue';

const data = reactive({
  columns: [
    {title: 'ID ', dataIndex: 'id', key: 'id'},
    {title: '图片昵称', dataIndex: 'name', key: 'name'},
    {title: '路径', dataIndex: 'path', key: 'path'},
    {title: '上传时间', dataIndex: 'created_at', key: 'created_at'},
    {title: '存储类型', dataIndex: 'image_type', key: 'image_type'},
    {title: '操作', dataIndex: 'action', key: 'action'},
  ],
  modalUpdateVisible: false,
  modalCreateVisible: false,
  formUpdateState: {
    name: "",
    id: "",
    path: "",
  },
})
const formRef = ref(null)
const MyTable_ = ref(null)


const updateInfoModal = record => {
  data.modalUpdateVisible = true
  data.formUpdateState.id = record.id
  data.formUpdateState.name = record.name
  data.formUpdateState.path = record.path
}
const UpdateInfo = async () => {
  let res = await imageUpdateApi(data.formUpdateState)
  if (res.code) {

    message.error(res.msg)
    return
  }

  message.success(res.msg)

  await MyTable_.value.ExportList()
  data.modalUpdateVisible = false

}

const removeBatch = async (list) => {
  let res = await imageRemoveBatchApi(list)
  if (res.code) {
    message.error(res.msg)
    return
  }
  message.success(res.msg)
  await MyTable_.value.ExportList()

}

const handleUploadOk = async ()=>{
  message.success("图片上传成功")
  await MyTable_.value.ExportList()

  data.modalCreateVisible=false
  fileList.value = []

}

//--------------------------------------------------------------

function getBase64(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
}

const previewVisible = ref(false);
const previewImage = ref('');
const previewTitle = ref('');
const fileList = ref([]);
const handleCancel = () => {
  previewVisible.value = false;
  previewTitle.value = '';
};
const handlePreview = async file => {
  if (!file.url && !file.preview) {
    file.preview = await getBase64(file.originFileObj);
  }
  previewImage.value = file.url || file.preview;
  previewVisible.value = true;
  previewTitle.value = file.name || file.url.substring(file.url.lastIndexOf('/') + 1);
};


// --------------------------------------------------------------
</script>

<template>
<div>
  <a-modal
      v-model:open="data.modalUpdateVisible"
      title="编辑图片"
      @ok="UpdateInfo"
  >
    <a-form
        :model="data.formUpdateState"
        name="basic"
        ref="formRef"
        :label-col="{span:4}"
        :wrapper-col="{span:20}"
        autocomplete="off"
    >


      <a-form-item
          has-feedback
          label="图片名称"
          name="name"
          :rules="[{ required: true, message: '请输入图片名称!' , trigger: 'blur'}]"
      >
        <a-input v-model:value="data.formUpdateState.name" placeholder="图片名称"/>
      </a-form-item>

      <a-form-item
          label="图片预览"
      >
        <img :src="data.formUpdateState.path" alt="" height="210" style="border-radius: 8px">
      </a-form-item>

    </a-form>
  </a-modal>
  <a-modal
      v-model:open="data.modalCreateVisible"
      title="上传图片"
      @ok="handleUploadOk"
  >
    <div class="clearfix">
      <a-upload
          v-model:file-list="fileList"
          action="/api/images/"
          list-type="picture-card"
          @preview="handlePreview"
          name="images"
          multiple
      >
        <div v-if="fileList.length < 8">
          <plus-outlined/>
          <div style="margin-top: 8px">上传图片</div>
        </div>
      </a-upload>
      <a-modal :open="previewVisible" :title="previewTitle" :footer="null" @cancel="handleCancel">
        <img alt="" style="width: 100%" :src="previewImage"/>
      </a-modal>
    </div>

  </a-modal>


  <MyTable
      baseUrl="/api/images/"
      :columns="data.columns"
      :options="data.options"
      searchBtnName="搜索图片昵称"
      @removeBatch="removeBatch"
      is-edit
      is-add
      :pageSize="4"
      ref="MyTable_"
  >


    <template #cell="{column, record}">
      <template v-if="column.key === 'path'">
        <img class="table_avatar" :src="record.path" alt="" height="60" style="border-radius: 5px">
      </template>
    </template>
    <template #edit="{record}">

      <a-button @click="updateInfoModal(record)" type="primary" ghost>编辑</a-button>
    </template>
    <template #addModal>
      <a-button type="primary" @click="data.modalCreateVisible = true">上传</a-button>
    </template>

  </MyTable>
</div>
</template>
<style scoped>
/* you can make up upload button and sample style by using stylesheets */
.ant-upload-select-picture-card i {
  font-size: 32px;
  color: #999;
}

.ant-upload-select-picture-card .ant-upload-text {
  margin-top: 8px;
  color: #666;
}
</style>
