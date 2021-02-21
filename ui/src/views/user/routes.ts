import { Route, RouteConfig } from 'vue-router';

import User from '@/views/user/pages/User.vue';
import Users from '@/views/user/pages/Users.vue';

export const getUserProps = (route: Route) => ({
  id: route.params.id,
});

export const getUsersProps = (route: Route) => ({
  systemId: route.query.systemId,
  systemName: route.query.systemName,
  namespaceId: route.query.namespaceId,
  namespaceName: route.query.namespaceName,
});

export default [
  {
    path: '/users',
    name: 'Users',
    component: Users,
    props: getUsersProps,
  },
  {
    path: '/user/:id',
    name: 'User',
    component: User,
    props: getUserProps,
  },
] as RouteConfig[];
