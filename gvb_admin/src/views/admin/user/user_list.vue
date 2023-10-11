<script setup>
import MyTable from '@/components/table.vue'
import {reactive, ref} from "vue";
import {userCreateApi, UserRemoveBatchApi, UserUpdateNickNameAndRoleApi} from "@/api/user";
import {message} from "ant-design-vue";

const data = reactive({
  options: [
    {value: 1, label: "管理员"},
    {value: 2, label: "普通用户"},
    {value: 3, label: "游客"},
  ],
  modalCreateVisible: false,
  modalUpdateVisible: false,
  columns: [
    {title: 'ID', dataIndex: 'id', key: 'id'},
    {title: '昵称', dataIndex: 'nick_name', key: 'nick_name'},
    {title: '头像', dataIndex: 'avatar', key: 'avatar'},
    {title: '邮箱', dataIndex: 'email', key: 'email'},
    {title: '角色 ', dataIndex: 'role', key: 'role'},
    {title: '注册来源', dataIndex: 'sign_status', key: 'sign_status'},
    {title: 'IP', dataIndex: 'ip', key: 'ip'},
    {title: '地址', dataIndex: 'addr', key: 'addr'},
    {title: '注册时间', dataIndex: 'created_at', key: 'created_at'},
    {title: '操作', dataIndex: 'action', key: 'action'},

  ],
  formState: {
    nick_name: "",
    user_name: "",
    password: "",
    rePassword: "",
    role: undefined,
  },
  _formState: {
    nick_name: "",
    user_name: "",
    password: "",
    rePassword: "",
    role: null,
  },
  formUpdateState: {
    nick_name: "",
    role: undefined,
    user_id: ""
  },
  filter: undefined,
})
const formRef = ref(null)
const MyTable_ = ref(null)

const removeBatch = async (list) => {
  let res = await UserRemoveBatchApi(list)
  if (res.code) {
    message.error(res.msg)
    return
  }
  message.success(res.msg)
}

const addUserModal = () => {
  data.modalCreateVisible = true
}
const handleOk = async () => {
  try {
    let valid = await formRef.value.validate()
    let res = await userCreateApi(valid)
    if (res.code) {
      message.error(res.msg)
    } else {
      message.success(res.msg)
    }
    Object.assign(data.formState, data._formState)
    await MyTable_.value.ExportList()

  } catch (e) {
  } finally {
    data.modalCreateVisible = false

  }
}
const validatePass = async (_rule, value) => {
  if (value === '') {
    return Promise.reject('请输入密码');
  }
  return Promise.resolve();
};
const validatePass2 = async (_rule, value) => {
  if (value === '') {
    return Promise.reject('请再次输入密码');
  } else if (value !== data.formState.password) {
    return Promise.reject("密码不一致");
  } else {
    return Promise.resolve();
  }
};

const updateUserInfoModal = record => {
  data.modalUpdateVisible = true
  data.formUpdateState.user_id = record.id
  data.formUpdateState.nick_name = record.nick_name
  data.formUpdateState.role = record.role_id
}
const UpdateUserInfo = async () => {
  let res = await UserUpdateNickNameAndRoleApi(data.formUpdateState)
  if (res.code) {
    message.error(res.msg)
    return
  }
  message.success(res.msg)
  await MyTable_.value.ExportList()
  data.modalUpdateVisible = false
}

const onFilter = () => {
  MyTable_.value.ExportList( {role: data.filter})
}
</script>

<template>
<div>

  <a-modal
      v-model:open="data.modalCreateVisible"
      title="添加用户"
      @ok="handleOk"
  >
    <a-form
        :model="data.formState"
        name="basic"
        ref="formRef"
        :label-col="{span:4}"
        :wrapper-col="{span:20}"
        autocomplete="off"
    >

      <a-form-item
          has-feedback
          label="用户名"
          name="user_name"
          :rules="[{ required: true, message: '请输入用户名!' , trigger: 'change'}]"
      >
        <a-input v-model:value="data.formState.user_name" placeholder="用户名"/>
      </a-form-item>

      <a-form-item
          has-feedback
          label="昵称"
          name="nick_name"
          :rules="[{ required: true, message: '请输入昵称!' , trigger: 'change'}]"
      >
        <a-input v-model:value="data.formState.nick_name" placeholder="昵称"/>
      </a-form-item>

      <a-form-item
          has-feedback
          label="密码"
          name="password"
          :rules="[{ required: true, validator: validatePass, trigger: 'change' }]"
      >
        <a-input-password v-model:value="data.formState.password" placeholder="密码"/>
      </a-form-item>
      <a-form-item
          has-feedback
          label="确认密码"
          name="rePassword"
          :rules="[{ required: true, validator: validatePass2, trigger: 'change' }]"
      >
        <a-input-password v-model:value="data.formState.rePassword" placeholder="确认密码"/>
      </a-form-item>

      <a-form-item
          label="权限"
          name="role"
          :rules="[{ required: true, message: '请选择权限!' }]"
      >
        <a-select
            placeholder="权限"
            v-model:value="data.formState.role"
            style="width: 200px"
            :options="data.options"
        ></a-select>
      </a-form-item>

    </a-form>
  </a-modal>
  <a-modal
      v-model:open="data.modalUpdateVisible"
      title="编辑用户"
      @ok="UpdateUserInfo"
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
          label="昵称"
          name="nick_name"
          :rules="[{ required: true, message: '请输入昵称!' , trigger: 'change'}]"
      >
        <a-input v-model:value="data.formUpdateState.nick_name" placeholder="昵称"/>
      </a-form-item>


      <a-form-item
          label="权限"
          name="role"
          :rules="[{ required: true, message: '请选择权限!' }]"
      >
        <a-select
            v-model:value="data.formUpdateState.role"
            placeholder="权限"
            style="width: 200px"
            :options="data.options"
        ></a-select>
      </a-form-item>

    </a-form>
  </a-modal>

  <MyTable
      baseUrl="/api/user"
      :columns="data.columns"
      :options="data.options"
      searchBtnName="搜索用户昵称"
      @removeBatch="removeBatch"
      is-edit
      is-add
      ref="MyTable_"
  >
    <template #addModal>
      <a-button type="primary" @click="addUserModal">添加</a-button>
    </template>
    <template #edit="{record}">
      <a-button @click="updateUserInfoModal(record)" type="primary" ghost disabled
                v-if="record.role === '管理员'">
        编辑
      </a-button>
      <a-button @click="updateUserInfoModal(record)" type="primary" ghost v-else>编辑</a-button>
    </template>
    <template #cell="{column, record}">
      <template v-if="column.key === 'avatar'">
        <img class="table_avatar" :src="record.avatar" alt="">
      </template>
    </template>
    <template #filter>
      <a-select
          @change="onFilter"
          v-model:value="data.filter"
          style="width: 200px; padding-left: 10px"
          :options="data.options"
          placeholder="权限"
          allowClear
      ></a-select>
    </template>
  </MyTable>
</div>
</template>

<style lang="scss" scoped>
.table_avatar{
  width: 40px;
  height: 40px;
  border-radius: 50%;
}
</style>