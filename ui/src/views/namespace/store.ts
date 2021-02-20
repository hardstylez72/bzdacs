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
    service: new NamespaceService({ host: '', baseUrl: '/api/v1/namespace' }),
  } as State,
  actions: {
    async GetListBySystemId(context, id: number): Promise<Namespace[]> {
      const { state } = actionContext(context);
      return state.service.GetListBySystemId(id);
    },
    async GetById(context, id: number): Promise<Namespace> {
      const { state } = actionContext(context);
      return state.service.GetById(id);
    },
    async Create(context, payload: {namespace: Namespace; systemId: number}): Promise<Namespace> {
      const { state } = actionContext(context);
      return state.service.Create(payload.namespace, payload.systemId);
    },
    async Update(context, namespace: Namespace): Promise<Namespace> {
      const { state } = actionContext(context);
      return state.service.Update(namespace);
    },
    async Delete(context, payload: {namespaceId: number; systemId: number}): Promise<void> {
      const { state } = actionContext(context);
      await state.service.Delete(payload.namespaceId, payload.systemId);
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
