import { Route, RouteConfig } from 'vue-router';

import Routes from './pages/Routes.vue';

export const getRoutesProps = (route: Route) => ({
  systemId: route.query.systemId,
  namespaceId: route.query.namespaceId,
});

export default [
  {
    path: '/routes',
    name: 'Routes',
    component: Routes,
    props: getRoutesProps,
  },
] as RouteConfig[];
