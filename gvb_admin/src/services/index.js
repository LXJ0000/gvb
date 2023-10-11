import axios from "axios";
import {useStore} from "@/stores/store";

export const Service = axios.create({
    timeout: 7000, baseURL: "", headers: {
        "Content-Type": "application/json"
    }
})
// 发请求使用的中间件
Service.interceptors.request.use(request => {
// 一般用于添加用户token
    const store = useStore()
    request.headers["token"] = store.UserInfo.token
    return request
})
// 响应中间件
Service.interceptors.response.use(response => {
    return response.data
})

