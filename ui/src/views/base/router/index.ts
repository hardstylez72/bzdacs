import Vue from 'vue';
import VueRouter, { RouteConfig, Route } from 'vue-router';
import Home from '../../Main.vue';
import Group from '../../group/pages/main.vue';
import User from '../../user/pages/main.vue';

Vue.use(VueRouter);
export const generateItemPageProps = (route: Route) => ({
  id: route.params.id,
});

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/group/:id',
    name: 'Group',
    component: Group,
    props: generateItemPageProps,
  },
  {
    path: '/user/:id',
    name: 'User',
    component: User,
    props: generateItemPageProps,
  },
];

const router = new VueRouter({
  mode: 'history',
  routes,
});

export default router;
