/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { moduleActionContext } from '../base/store';
import DefaultService from '../base/services/default';
import RouteService, { Filter, Route } from './service';

export interface State {
  service: RouteService;
  routes: Route[];
  filter: Filter;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    filter: {
      tags: {
        exclude: false,
        names: [],
      },
    },
    service: new RouteService({ host: '', baseUrl: '/api/v1/route' }),
    routes: [],
  } as State,
  getters: {
    getRoutes(state) {
      return state.routes;
    },
    getFilter(state) {
      return state.filter;
    },
  },
  mutations: {
    setFilter(state, filter: Filter) {
      state.filter = filter;
    },
    setFilterTagNames(state, names: string[]) {
      state.filter.tags.names = names;
    },
    setRoutes(state, routes: Route[]) {
      state.routes = routes;
    },
    deleteRoute(state, id: number) {
      state.routes = state.routes.filter((route) => route.id !== id);
    },
    addRoute(state, routes: Route) {
      state.routes.push(routes);
    },
    updateRoute(state, route: Route) {
      state.routes = state.routes.map((r) => {
        if (route.id === r.id) {
          return route;
        }
        return r;
      });
    },
  },
  actions: {
    async GetList(context, filter: Filter): Promise<Route[]> {
      const { state, commit } = actionContext(context);
      const routes = await state.service.GetListV2(filter);
      commit.setRoutes(routes);
      return routes;
    },
    async GetById(context, id: number): Promise<Route> {
      const { state } = actionContext(context);
      return state.service.GetById(id);
    },
    async Create(context, route: Route): Promise<Route> {
      const { state, commit } = actionContext(context);
      const createdRoute = await state.service.Create(route);
      commit.addRoute(createdRoute);
      return createdRoute;
    },
    async Update(context, route: Route): Promise<Route> {
      const { state, commit } = actionContext(context);
      const createdRoute = await state.service.Update(route);
      commit.updateRoute(createdRoute);
      return createdRoute;
    },
    async Delete(context, id: number): Promise<void> {
      const { state, commit } = actionContext(context);

      await state.service.Delete(id);
      commit.deleteRoute(id);
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
