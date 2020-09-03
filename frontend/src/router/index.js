import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'domain',
      component: () => import('../views/Home.vue')
    },
    {
      path: '/previous',
      name: 'Previous',
      component: () => import('../views/Previous.vue')
    }] 
})

export default router
