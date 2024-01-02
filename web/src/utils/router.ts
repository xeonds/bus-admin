import { createRouter, createWebHashHistory } from 'vue-router'

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      component: () => import('@/views/Home.vue'),
    },
    {
      path: '/company',
      component: () => import('@/views/Company.vue'),
    },
    {
      path: '/vehicle',
      component: () => import('@/views/Vehicle.vue'),
    },
    {
      path: '/team',
      component: () => import('@/views/Team.vue'),
    },
    {
      path: '/route',
      component: () => import('@/views/Route.vue'),
    },
    {
      path: '/driver',
      component: () => import('@/views/Driver.vue'),
    },
    {
      path: '/violation',
      component: () => import('@/views/Violation.vue'),
    },
  ],
})
