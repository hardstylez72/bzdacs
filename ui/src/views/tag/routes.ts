import { Route, RouteConfig } from 'vue-router';

import Tags from '@/views/tag/pages/Tags.vue';

export const getGroupsProps = (route: Route) => ({
  systemId: route.query.systemId,
  namespaceId: route.query.namespaceId,
});

export default [
  {
    path: '/tags',
    name: 'Tags',
    component: Tags,
    props: getGroupsProps,
  },
] as RouteConfig[];
