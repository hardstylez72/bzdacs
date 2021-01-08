/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import {
  defineActions, defineModule, defineMutations, defineGetters,
} from 'direct-vuex';
import { moduleActionContext } from '@/views/base/store';
import UserService, { User } from '../services/user';

export interface State<T>{
  service: UserService;
  entities: T[];
}

const state1 = {
  service: new UserService({ host: '', baseUrl: '/api/v1/user' }),
  entities: [],
} as State<User>;

const mutations = defineMutations < State < User >>()({
  setEntities(state, entities) {
    state.entities = entities;
  },
  deleteEntity(state, id: number) {
    state.entities = state.entities.filter((route) => route.id !== id);
  },
  addEntity(state, entities) {
    state.entities.push(entities);
  },
});

const actions = defineActions({
  async GetById(context, id: number): Promise<User> {
    const { state } = actionContext(context);
    return state.service.GetById(id);
  },
  async GetList(context): Promise<User[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList();
    commit.setEntities(entities);
    return entities;
  },
  async Create(context, entity: User): Promise<User> {
    const { state, commit } = actionContext(context);
    const createdEntity = await state.service.Create(entity);
    commit.addEntity(createdEntity);
    return createdEntity;
  },
  async Delete(context, id: number): Promise<void> {
    const { state, commit } = actionContext(context);
    await state.service.Delete(id);
    commit.deleteEntity(id);
  },
});

const getters = defineGetters<State<User>>()({
  getEntities(state): User[] {
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
