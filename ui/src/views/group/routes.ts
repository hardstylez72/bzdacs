import { Route, RouteConfig } from 'vue-router';

import Group from '@/views/group/pages/Group.vue';
import Groups from '@/views/group/pages/Groups.vue';

export const getGroupProps = (route: Route) => ({
  id: route.params.id,
});

export const getGroupsProps = (route: Route) => ({
  systemId: route.query.systemId,
  namespaceId: route.query.namespaceId,
});

export default [
  {
    path: '/groups',
    name: 'Groups',
    component: Groups,
    props: getGroupsProps,
  },
  {
    path: '/group/:id',
    name: 'Group',
    component: Group,
    props: getGroupProps,
  },
] as RouteConfig[];
