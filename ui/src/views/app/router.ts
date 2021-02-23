import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import RoutePages from '@/views/route/routes';
import GroupPages from '@/views/group/routes';
import UserPages from '@/views/user/routes';
import TagPages from '@/views/tag/routes';
import Home from './pages/Home.vue';
import Login from './pages/AdminPage.vue';

Vue.use(VueRouter);
const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
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
