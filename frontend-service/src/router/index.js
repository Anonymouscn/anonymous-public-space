import {createRouter, createWebHistory} from "vue-router";

const router = createRouter({
    // history: createWebHistory(import.meta.env.BASE_URL),
    history: createWebHistory('/v1'),
    routes: [
        // 生产环境组件映射
        {
            path: '/',
            name: 'about me',
            component: () => import('../views/public/AboutMe.vue'),
            meta: {
                title: 'Anonymous Public Space'
            }
        },
        // 后台登录
        {
            path: '/login',
            name: 'login to manage',
            component: () => import('../views/public/Login.vue'),
            meta: {
                title: 'Anonymous Public Space - Login to manage'
            }
        },
        // 后台管理
        {
            path: '/manage',
            name: 'manage console',
            component: () => import('../views/manage/Manage.vue'),
            meta: {
                title: 'Anonymous Public Space - Manage console'
            }
        },
        // 开发 debug 组件映射
        {
            path: '/dev',
            children: []
        }
    ]
})

router.beforeEach((to, from, next) => {
    // 动态获取网页标题
    window.document.title = to.meta.title
    next()
})

export default router