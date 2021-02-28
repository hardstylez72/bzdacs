import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import RoutePages from '@/route/routes';
import GroupPages from '@/group/routes';
import UserPages from '@/user/routes';
import TagPages from '@/tag/routes';
import Home from './pages/Home.vue';
import Login from './pages/Login.vue';
import Registration from './pages/Registration.vue';

Vue.use(VueRouter);
const routes: Array<RouteConfig> = [
  {
    path: '/register',
    name: 'Registration',
    component: Registration,
  },
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
