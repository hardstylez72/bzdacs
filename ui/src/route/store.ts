/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { client } from '@/base/services/utils/requester';
import {
  serviceOptions,
  routeDeleteRequest, routeGetByIdRequest,
  routeInsertRequest, routeListRequest,
  routeUpdateRequest,
  RouteService,
} from '@/route/generated';
import { ListResponse } from '@/common/helpers/types';
import { moduleActionContext } from '../base/store';
import { Route } from './entity';

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
    async GetList(context, req: routeListRequest): Promise<ListResponse<Route>> {
      const { state, commit } = actionContext(context);
      const res = await state.service.routeList({ req });
      commit.setRoutes(res.items);
      // @ts-ignore
      return res;
    },
    async GetById(context, req: routeGetByIdRequest): Promise<Route> {
      const { state } = actionContext(context);
      return state.service.routeGetById({ req });
    },
    async Create(context, req: routeInsertRequest): Promise<Route> {
      const { state } = actionContext(context);
      return state.service.routeCreate({ req });
    },
    async Update(context, req: routeUpdateRequest): Promise<Route> {
      const { state } = actionContext(context);
      return state.service.routeUpdate({ req });
    },
    async Delete(context, req: routeDeleteRequest): Promise<void> {
      const { state } = actionContext(context);
      return state.service.routeDelete({ req });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
