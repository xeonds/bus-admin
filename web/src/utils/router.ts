import { createRouter, createWebHashHistory } from 'vue-router'

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      //      redirect: "/login",
      component: () => import('@/views/Home.vue'),
    },
  ],
})
