/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import {
  defineActions, defineModule, defineMutations, defineGetters,
} from 'direct-vuex';
import UserGroupService, { UserGroup } from '@/views/user/services/usergroup';

import { Group } from '@/views/group/services/group';
import { moduleActionContext } from '../../base/store';

export interface State{
  service: UserGroupService;
  groupsBelongToUser: Group[];
  groupsNotBelongToUser: Group[];
  userId: number;
}

const state1 = {
  service: new UserGroupService({ host: '', baseUrl: '/api/v1/user/group' }),
  groupsBelongToUser: [],
  groupsNotBelongToUser: [],
  userId: -1,
} as State;

const mutations = defineMutations < State >()({

  setUserId(state, groupId: number) {
    state.userId = groupId;
  },
  setRoutesBelongToGroup(state, entities: Group[]) {
    state.groupsBelongToUser = entities;
  },
  deleteRoutesBelongToGroup(state, pairs: UserGroup[]) {
    state.groupsBelongToUser = state.groupsBelongToUser.filter((r: Group) => {
      const exist = pairs.some((pair) => pair.groupId === r.id);
      return !exist;
    });
  },
  addRoutesBelongToGroup(state, entities) {
    state.groupsBelongToUser.push(...entities);
  },

  setRoutesNotBelongToGroup(state, entities: Group[]) {
    state.groupsNotBelongToUser = entities;
  },
  deleteRoutesNotBelongToGroup(state, pairs: UserGroup[]) {
    state.groupsNotBelongToUser = state.groupsNotBelongToUser.filter((r: Group) => {
      const exist = pairs.some((pair) => pair.groupId === r.id);
      return !exist;
    });
  },
  addRoutesNotBelongToGroup(state, entities) {
    state.groupsNotBelongToUser.push(...entities);
  },
});

const actions = defineActions({

  async GetListNotBelongToUser(context, groupId: number): Promise<Group[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList(groupId, false);
    commit.setRoutesNotBelongToGroup(entities);
    commit.setUserId(groupId);
    return entities;
  },
  async GetListBelongToUser(context, groupId: number): Promise<Group[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList(groupId, true);
    commit.setRoutesBelongToGroup(entities);
    commit.setUserId(groupId);
    return entities;
  },
  async Create(context, pairs: UserGroup[]): Promise<Group[]> {
    const { state, commit } = actionContext(context);
    const createdEntity = await state.service.Create(pairs);
    commit.addRoutesBelongToGroup(createdEntity);
    commit.deleteRoutesNotBelongToGroup(pairs);
    return createdEntity;
  },
  async Delete(context, pairs: UserGroup[]): Promise<void> {
    const { state, commit } = actionContext(context);
    await state.service.Delete(pairs);
    commit.deleteRoutesBelongToGroup(pairs);
    commit.addRoutesNotBelongToGroup(pairs);
  },
});

const getters = defineGetters<State>()({
  getGroupsBelongToUser(state): Group[] {
    return state.groupsBelongToUser;
  },
  getGroupsNotBelongToUser(state): Group[] {
    return state.groupsNotBelongToUser;
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
