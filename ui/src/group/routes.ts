import { Route, RouteConfig } from 'vue-router';

import Group from '@/group/pages/Group.vue';
import Groups from '@/group/pages/Groups.vue';

export const getGroupProps = (route: Route) => ({
  id: route.params.id,
  systemId: route.query.systemId,
  namespaceId: route.query.namespaceId,
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
