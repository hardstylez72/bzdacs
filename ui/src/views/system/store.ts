/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { moduleActionContext } from '../base/store';
import SystemService, { System } from './service';

export interface State {
  service: SystemService;
  items: System[];
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new SystemService({ host: '', baseUrl: '/api/v1/system' }),
    items: [],
  } as State,
  getters: {
    getItems(state) {
      return state.items;
    },
  },
  mutations: {
    setMany(state, items: System[]) {
      state.items = items;
    },
    deleteById(state, id: number) {
      state.items = state.items.filter((system) => system.id !== id);
    },
    addOne(state, system: System) {
      state.items.push(system);
    },
    updateOne(state, system: System) {
      state.items = state.items.map((r) => {
        if (system.id === r.id) {
          return system;
        }
        return r;
      });
    },
  },
  actions: {
    async GetById(context, id: number): Promise<System> {
      const { state } = actionContext(context);

      return state.service.GetById(id);
    },
    async GetList(context): Promise<System[]> {
      const { state, commit } = actionContext(context);
      const items = await state.service.GetList();
      commit.setMany(items);
      return items;
    },
    async Create(context, route: System): Promise<System> {
      const { state, commit } = actionContext(context);
      const system = await state.service.Create(route);
      commit.addOne(system);
      return system;
    },
    async Update(context, route: System): Promise<System> {
      const { state, commit } = actionContext(context);
      const createdRoute = await state.service.Update(route);
      commit.updateOne(createdRoute);
      return createdRoute;
    },
    async Delete(context, id: number): Promise<void> {
      const { state, commit } = actionContext(context);

      await state.service.Delete(id);
      commit.deleteById(id);
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
