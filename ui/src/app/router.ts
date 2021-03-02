import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import RoutePages from '@/route/routes';
import GroupPages from '@/group/routes';
import UserPages from '@/user/routes';
import TagPages from '@/tag/routes';
import Home from './pages/Home.vue';
import Login from './pages/Login.vue';
import Registration from './pages/Registration.vue';
import store from './store';

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

router.beforeEach((to, from, next) => {
  // Auth guard
  if (!store.getters.sysUser.isAuthorized) {
    if (to.name === 'Login' || to.name === 'Registration') {
      next();
    } else {
      next({ name: 'Login' });
    }
    // Tries to go 404 url => redirect home
  } else if (!to.name) {
    next('Home');
  } else if (store.getters.sysUser.isAuthorized) {
    if (to.name === 'Login' || to.name === 'Registration') {
      next(false);
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
