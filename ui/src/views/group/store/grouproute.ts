/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import {
  defineActions, defineModule, defineMutations, defineGetters,
} from 'direct-vuex';
import GroupRouteService, { GroupRoute } from '@/views/group/services/grouproute';
import { Route } from '@/views/route/service';
import { moduleActionContext } from '../../base/store';

export interface State{
  service: GroupRouteService;
  routesBelongToGroup: Route[];
  routesNotBelongToGroup: Route[];
  groupId: number;
}

const state1 = {
  service: new GroupRouteService({ host: '', baseUrl: '/api/v1/group/route' }),
  routesBelongToGroup: [],
  routesNotBelongToGroup: [],
  groupId: -1,
} as State;

const mutations = defineMutations < State >()({

  setGroupId(state, groupId: number) {
    state.groupId = groupId;
  },
  setRoutesBelongToGroup(state, entities: Route[]) {
    state.routesBelongToGroup = entities;
  },
  deleteRoutesBelongToGroup(state, pairs: GroupRoute[]) {
    state.routesBelongToGroup = state.routesBelongToGroup.filter((r: Route) => {
      const exist = pairs.some((pair) => pair.routeId === r.id);
      return !exist;
    });
  },
  setRoutesNotBelongToGroup(state, entities: Route[]) {
    state.routesNotBelongToGroup = entities;
  },
  deleteRoutesNotBelongToGroup(state, pairs: GroupRoute[]) {
    state.routesNotBelongToGroup = state.routesNotBelongToGroup.filter((r: Route) => {
      const exist = pairs.some((pair) => pair.routeId === r.id);
      return !exist;
    });
  },
  addRoutesBelongToGroup(state, entities) {
    state.routesBelongToGroup.push(...entities);
  },
  addRoutesNotBelongToGroup(state, entities) {
    state.routesNotBelongToGroup.push(...entities);
  },
});

const actions = defineActions({

  async GetListNotBelongToUser(context, groupId: number): Promise<Route[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList(groupId, false);
    commit.setRoutesNotBelongToGroup(entities);
    commit.setGroupId(groupId);
    return entities;
  },
  async GetListBelongToUser(context, groupId: number): Promise<Route[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList(groupId, true);
    commit.setRoutesBelongToGroup(entities);
    commit.setGroupId(groupId);
    return entities;
  },
  async Create(context, pairs: GroupRoute[]): Promise<Route[]> {
    const { state, commit } = actionContext(context);
    const createdEntity = await state.service.Create(pairs);
    commit.addRoutesBelongToGroup(createdEntity);
    commit.deleteRoutesNotBelongToGroup(pairs);
    return createdEntity;
  },
  async Delete(context, pairs: GroupRoute[]): Promise<void> {
    const { state, commit } = actionContext(context);
    await state.service.Delete(pairs);
    commit.deleteRoutesBelongToGroup(pairs);
    commit.addRoutesNotBelongToGroup(pairs);
  },
});

const getters = defineGetters<State>()({
  getRoutesBelongToGroup(state): Route[] {
    return state.routesBelongToGroup;
  },
  getRoutesNotBelongToGroup(state): Route[] {
    return state.routesNotBelongToGroup;
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
