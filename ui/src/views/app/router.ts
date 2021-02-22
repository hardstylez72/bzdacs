import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import RoutePages from '@/views/route/routes';
import GroupPages from '@/views/group/routes';
import UserPages from '@/views/user/routes';
import TagPages from '@/views/tag/routes';
import Home from './pages/Home.vue';
import Admin from './pages/AdminPage.vue';
import Guest from './pages/GuestPage.vue';

Vue.use(VueRouter);
const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/guest',
    name: 'Guest',
    component: Guest,
  },
  {
    path: '/admin',
    name: 'Admin',
    component: Admin,
  },
  ...RoutePages,
  ...GroupPages,
  ...UserPages,
  ...TagPages,
];

const router = new VueRouter({
  mode: 'history',
  routes,
});

export default router;
