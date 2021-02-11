/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { moduleActionContext } from '../base/store';
import NamespaceService, { Namespace } from './service';

export interface State {
  service: NamespaceService;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new NamespaceService({ host: '', baseUrl: '/api/v1/system/namespace' }),
  } as State,
  getters: {
  },
  mutations: {},
  actions: {
    async GetListBySystemId(context, id: number): Promise<Namespace[]> {
      const { state } = actionContext(context);
      return state.service.GetListBySystemId(id);
    },
    // async GetList(context): Promise<Namespace[]> {
    //   const { state, commit } = actionContext(context);
    //   const namespaces = await state.service.GetList();
    //   commit.setMany(namespaces);
    //   return namespaces;
    // },
    async Create(context, payload: {namespace: Namespace; systemId: number}): Promise<Namespace> {
      const { state } = actionContext(context);
      return state.service.Create(payload.namespace, payload.systemId);
    },
    // async Update(context, route: Namespace): Promise<Namespace> {
    //   const { state, commit } = actionContext(context);
    //   const createdRoute = await state.service.Update(route);
    //   commit.updateTag(createdRoute);
    //   return createdRoute;
    // },
    // async Delete(context, id: number): Promise<void> {
    //   const { state, commit } = actionContext(context);
    //
    //   await state.service.Delete(id);
    //   commit.deleteById(id);
    // },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
