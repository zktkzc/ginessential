import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import userRoutes from "@/router/module/user";
import store from "@/store";

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'Home',
        component: HomeView
    },
    {
        path: '/about',
        name: 'About',
        component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
    },
    ...userRoutes
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

router.beforeEach((to, from, next) => {
    if (to.meta.auth) { // 判断是否需要登录
        // 判断用户是否登录
        if (store.state.userModule.token) {
            // 判断token的有效性，比如有没有过期，需要后台发放token的时候带上token的有效期
            // 如果token无效，需要请求token
            next()
        } else {
            // 跳转登录页面
            router.push({name: 'login'})
        }
    } else {
        next()
    }
})

export default router
