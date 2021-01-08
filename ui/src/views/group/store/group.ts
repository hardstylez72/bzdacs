/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import {
  defineActions, defineModule, defineMutations, defineGetters,
} from 'direct-vuex';
import { Route } from '@/views/route/service';
import { moduleActionContext } from '../../base/store';
import Service from '../../base/services/default';
import GroupService, { Group } from '../services/group';

export interface State<T>{
  service: GroupService;
  entities: Group[];
}

const state1 = {
  service: new GroupService({ host: '', baseUrl: '/api/v1/group' }),
  entities: [],
} as State<Group>;

const mutations = defineMutations < State < Group >>()({
  setEntities(state, entities) {
    state.entities = entities;
  },
  deleteEntity(state, id: number) {
    state.entities = state.entities.filter((route) => route.id !== id);
  },
  addEntity(state, entities) {
    state.entities.push(entities);
  },
  updateGroup(state, group: Group) {
    state.entities = state.entities.map((r) => {
      if (group.id === r.id) {
        return group;
      }
      return r;
    });
  },
});

const actions = defineActions({
  async GetById(context, id: number): Promise<Group> {
    const { state } = actionContext(context);
    return state.service.GetById(id);
  },
  async GetList(context): Promise<Group[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList();
    commit.setEntities(entities);
    return entities;
  },
  async Update(context, route: Group): Promise<Group> {
    const { state, commit } = actionContext(context);
    const createdEntity = await state.service.Update(route);
    commit.updateGroup(createdEntity);
    return createdEntity;
  },
  async Create(context, route: Group): Promise<Group> {
    const { state, commit } = actionContext(context);
    const createdEntity = await state.service.Create(route);
    commit.addEntity(createdEntity);
    return createdEntity;
  },
  async CreateBasedOnAnother(context, payload: {group: Group; baseGroupId: number}): Promise<Group> {
    const { state, commit } = actionContext(context);
    const createdEntity = await state.service.CreateBasedOnAnother(payload.group, payload.baseGroupId);
    commit.addEntity(createdEntity);
    return createdEntity;
  },
  async Delete(context, id: number): Promise<void> {
    const { state, commit } = actionContext(context);

    await state.service.Delete(id);
    commit.deleteEntity(id);
  },
});

const getters = defineGetters<State<Group>>()({
  getEntities(state): Group[] {
    return state.entities;
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
