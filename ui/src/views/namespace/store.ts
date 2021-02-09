/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { moduleActionContext } from '../base/store';
import NamespaceService, { Namespace } from './service';

export interface State {
  service: NamespaceService;
  namespaces: Namespace[];
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new NamespaceService({ host: '', baseUrl: '/api/v1/namespace' }),
    namespaces: [],
  } as State,
  getters: {
    getNamespaces(state) {
      return state.namespaces;
    },
  },
  mutations: {
    setMany(state, namespaces: Namespace[]) {
      state.namespaces = namespaces;
    },
    deleteById(state, id: number) {
      state.namespaces = state.namespaces.filter((namespace) => namespace.id !== id);
    },
    addOne(state, namespace: Namespace) {
      state.namespaces.push(namespace);
    },
    updateTag(state, namespace: Namespace) {
      state.namespaces = state.namespaces.map((r) => {
        if (namespace.id === r.id) {
          return namespace;
        }
        return r;
      });
    },
  },
  actions: {
    async GetById(context, id: number): Promise<Namespace> {
      const { state } = actionContext(context);

      return state.service.GetById(id);
    },
    async GetList(context): Promise<Namespace[]> {
      const { state, commit } = actionContext(context);
      const namespaces = await state.service.GetList();
      commit.setMany(namespaces);
      return namespaces;
    },
    async Create(context, route: Namespace): Promise<Namespace> {
      const { state, commit } = actionContext(context);
      const namespace = await state.service.Create(route);
      commit.addOne(namespace);
      return namespace;
    },
    async Update(context, route: Namespace): Promise<Namespace> {
      const { state, commit } = actionContext(context);
      const createdRoute = await state.service.Update(route);
      commit.updateTag(createdRoute);
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
