import {Service} from "@/services";
// 邮箱登录接口
export const emailLoginApi = (data) => {
// 代理
    return Service.post("/api/email_login/", data)
}
//用户信息接口
export const userListApi = (params) => {
    return Service.get("/api/user/", {params})
}

//
export const userCreateApi = (data) => {
    return Service.post('/api/user/', data)
}

export const UserRemoveBatchApi = (id_list) => {
    return Service.delete('/api/user/', {data: {id_list}})
}
export const UserUpdateNickNameAndRoleApi = data => {
    return Service.post('/api/user_role/', data)
}
