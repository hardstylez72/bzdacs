/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import {
  defineActions, defineModule, defineMutations, defineGetters,
} from 'direct-vuex';
import UserRouteService, {
  CreateUserRoute,
  RouteExt,
  RouteWithGroups, UpdateUserRoute,
  UserRoute,
} from '@/views/user/services/userroute';

import { Route } from '@/views/route/service';
import { moduleActionContext } from '../../base/store';

export interface State{
  service: UserRouteService;
  routesBelongToUser: RouteWithGroups[];
  routesNotBelongToUser: RouteWithGroups[];
  userId: number;
}

const state1 = {
  service: new UserRouteService({ host: '', baseUrl: '/api/v1/user/route' }),
  routesBelongToUser: [],
  routesNotBelongToUser: [],
  userId: -1,
} as State;

const mutations = defineMutations < State >()({

  updateRoute(state, r: RouteExt) {
    state.routesBelongToUser = state.routesBelongToUser.map((route) => {
      if (route.id === r.id) {
        return {
          ...route,
          isExcluded: r.isExcluded,
        };
      }
      return route;
    });
  },

  rewriteRoute(state, r: RouteExt) {
    state.routesBelongToUser = state.routesBelongToUser.map((route) => {
      if (route.id === r.id) {
        return {
          ...route,
          isExcluded: r.isExcluded,
          isIndependent: true, // todo: form at backend
          isOverwritten: true, // todo: form at backend
        };
      }
      return route;
    });
  },

  setUserId(state, userId: number) {
    state.userId = userId;
  },
  setRoutesBelongToGroup(state, entities: RouteWithGroups[]) {
    state.routesBelongToUser = entities;
  },
  deleteRoutesBelongToGroup(state, pairs: UserRoute[]) {
    state.routesBelongToUser = state.routesBelongToUser.filter((r: RouteWithGroups, index) => {
      const exist = pairs.some((pair) => {
        if ((pair.routeId === r.id) && (!r.isOverwritten)) {
          return true;
        }
        return false;
      });
      return !exist;
    }).map((r: RouteWithGroups) => {
      const route = {
        ...r,
      };
      pairs.forEach((pair) => {
        if ((pair.routeId === r.id)) {
          route.isOverwritten = false;
          route.isIndependent = false;
          route.isExcluded = false;
        }
      });
      return route;
    });
  },
  addRoutesBelongToGroup(state, entities) {
    state.routesBelongToUser.push(...entities);
  },

  setRoutesNotBelongToGroup(state, entities: RouteWithGroups[]) {
    state.routesNotBelongToUser = entities;
  },
  deleteRoutesNotBelongToGroup(state, pairs: UserRoute[]) {
    state.routesNotBelongToUser = state.routesNotBelongToUser.filter((r: RouteWithGroups) => {
      const exist = pairs.some((pair) => pair.routeId === r.id);
      return !exist;
    });
  },
  addRoutesNotBelongToGroup(state, entities) {
    state.routesNotBelongToUser.push(...entities);
  },

});

const actions = defineActions({

  async UpdateRoute(context, route: UpdateUserRoute): Promise<RouteExt> {
    const { state, commit } = actionContext(context);
    const r = await state.service.Update(route);
    commit.updateRoute(r);
    return r;
  },
  async GetListNotBelongToUser(context, userId: number): Promise<RouteWithGroups[]> {
    const { state, commit } = actionContext(context);
    const routes = await state.service.GetList(userId, false);
    const rs: RouteWithGroups[] = routes.map((route) => ({
      ...route,
      groupCodes: route.groups.map((group) => group.code).join(', '),
    }));
    commit.setRoutesNotBelongToGroup(rs);
    commit.setUserId(userId);
    return rs;
  },
  async GetListBelongToUser(context, groupId: number): Promise<RouteWithGroups[]> {
    const { state, commit } = actionContext(context);
    const routes = await state.service.GetList(groupId, true);
    const rs: RouteWithGroups[] = routes.map((route) => ({
      ...route,
      groupCodes: route.groups.map((group) => group.code).join(', '),
      isSelectable: route.isIndependent,
    }));
    commit.setRoutesBelongToGroup(rs);
    commit.setUserId(groupId);
    return rs;
  },
  async Create(context, pairs: CreateUserRoute[]): Promise<Route[]> {
    const { state, commit } = actionContext(context);
    const createdEntity = await state.service.Create(pairs);
    return createdEntity;
  },
  async DeleteOne(context, pairs: UserRoute): Promise<void> {
    const { state, commit } = actionContext(context);
    await state.service.Delete([pairs]);
    commit.deleteRoutesBelongToGroup([pairs]);
  },
  async RewriteOne(context, pairs: CreateUserRoute): Promise<Route[]> {
    const { state, commit } = actionContext(context);
    const createdEntity = await state.service.Create([pairs]);
    createdEntity.forEach((route) => {
      commit.rewriteRoute(route);
    });

    return createdEntity;
  },

});

const getters = defineGetters<State>()({
  getRoutesBelongToUser(state): RouteWithGroups[] {
    return state.routesBelongToUser.map((route) => ({
      ...route,
      groupCodes: route.groups.map((group) => group.code).join(', '),
    }));
  },
  getRoutesNotBelongToGroup(state): RouteWithGroups[] {
    return state.routesNotBelongToUser;
  },
});

const module = defineModule({
  namespaced: true as true,
  state: state1,
  getters,
  mutations,
  actions,
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
