/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { client } from '@/views/base/services/utils/requester';
import {
  serviceOptions,
  route_deleteRequest, route_getByIdRequest,
  route_insertRequest, route_listRequest,
  route_updateRequest,
  RouteService,
} from '@/views/route/generated';
import { moduleActionContext } from '../base/store';
import { RouteSwg as Route } from './entity';

serviceOptions.axios = client;

export interface State {
  service: RouteService;
  routes: Route[];
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new RouteService(),
    routes: [],
  } as State,
  getters: {
    getRoutes(state) {
      return state.routes;
    },
  },
  mutations: {
    setRoutes(state, routes?: Route[]) {
      if (routes) {
        state.routes = routes;
      }
    },
  },
  actions: {
    async GetList(context, req: route_listRequest): Promise<{items: Route[]; total: number}> {
      const { state, commit } = actionContext(context);
      const res = await state.service.routeList({ req });
      commit.setRoutes(res.items);
      // @ts-ignore
      return res;
    },
    async GetById(context, req: route_getByIdRequest): Promise<Route> {
      const { state } = actionContext(context);
      return state.service.routeGetById({ req });
    },
    async Create(context, req: route_insertRequest): Promise<Route> {
      const { state } = actionContext(context);
      return state.service.routeCreate({ req });
    },
    async Update(context, req: route_updateRequest): Promise<Route> {
      const { state } = actionContext(context);
      return state.service.routeUpdate({ req });
    },
    async Delete(context, req: route_deleteRequest): Promise<void> {
      const { state } = actionContext(context);
      return state.service.routeDelete({ req });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
