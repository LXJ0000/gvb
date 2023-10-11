import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({

    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/admin/",
            name: "admin",
            component: () => import('@/views/admin/admin.vue'),
            children:[
                {
                    path: "",
                    name: "admin_index",
                    redirect:"/admin/home/"

                },
                {
                    path: "home/",
                    name: "home",
                    component: () => import('@/views/admin/home/home.vue'),
                },
                {
                    path: "user_list/",
                    name: "user_list",
                    component: () => import('@/views/admin/user/user_list.vue'),
                },
                {
                    path: "system_info/",
                    name: "system_info",
                    component: () => import('@/views/admin/system/system_info.vue'),
                },
                {
                    path: "image_list/",
                    name: "image_list",
                    component: () => import('@/views/admin/image/image_list.vue'),
                },
                {
                    path: "test/",
                    name: "test",
                    component: () => import('@/components/test.vue'),
                },
            ]
        },
        {
            path: "/login/",
            name: "login",
            component: () => import('@/views/login.vue')
        },

    ]
})

export default router
