import {defineStore} from "pinia";
import {message} from "ant-design-vue";
import router from "@/router";

export const useStore = defineStore('gvb', {
    state: () => {
        return {
            UserInfo: {
                avatar: "/static/avatar/1.jpg",
                exp: 0,
                iss: "",
                nick_name: "",
                role: 0,
                user_id: 0,
                token: "",
            },
            tabList: []
        }
    },
    actions: {
        //     modify userinfo
        setUserInfo(info) {
            this.$patch({
                UserInfo: info
            })
            //     持久化
            localStorage.setItem("userinfo", JSON.stringify(info))
        },
        // 加载userinfo
        loadUserInfo() {
            let info = localStorage.getItem("userinfo")
            if (info === null) {
                return
            }
            //     解析Json
            let userinfo = JSON.parse(info)
            //     判断时间是否失效
            let exp = userinfo.exp * 1000
            let nowTime = new Date().getTime()
            if (exp <= nowTime) {
                //     过期
                message.warn("登陆信息失效")
                router.push({name: "login"})
                return
            }
            this.setUserInfo(userinfo)
        },
        addTab(tab) {
            if (this.tabList.findIndex((item) => item.name === tab.name) === -1) {
                this.tabList.push({
                    name: tab.name,
                    title: tab.title,
                })
            }
        },
        saveTabs() {
            localStorage.setItem("tabList", JSON.stringify(this.tabList))
        },
        loadTabs() {
            let tabs = localStorage.getItem("tabList")
            if (tabs === null) {
                this.tabList = [{title: "首页", name: "home",}]
                return
            }
            this.tabList = JSON.parse(tabs)
        },
        removeTab(tab) {
            let removeIndex = this.tabList.findIndex((item) => item.name === tab.name)
            this.tabList.splice(removeIndex, 1)
            return removeIndex
        },
        removeAllTab() {
            this.tabList = [{title: "首页", name: "home",}]
        }
    }
})