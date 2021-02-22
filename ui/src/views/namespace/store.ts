/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { client } from '@/views/base/services/utils/requester';
import {
  serviceOptions,
  NamespaceService,
  namespace_insertRequest, namespace_updateRequest, namespace_deleteRequest,
} from '@/views/namespace/generated';
import { moduleActionContext } from '../base/store';
import { Namespace } from './entity';

serviceOptions.axios = client;

export interface State {
  service: NamespaceService;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new NamespaceService(),
  } as State,
  actions: {
    async GetListBySystemId(context, id: number): Promise<Namespace[]> {
      const { state } = actionContext(context);
      return state.service.namespaceList({ req: { id } });
    },
    async GetById(context, id: number): Promise<Namespace> {
      const { state } = actionContext(context);
      return state.service.namespaceGet({ req: { id } });
    },
    async Create(context, req: namespace_insertRequest): Promise<Namespace> {
      const { state } = actionContext(context);
      return state.service.namespaceCreate({ req });
    },
    async Update(context, req: namespace_updateRequest): Promise<Namespace> {
      const { state } = actionContext(context);
      return state.service.namespaceUpdate({ req });
    },
    async Delete(context, req: namespace_deleteRequest): Promise<void> {
      const { state } = actionContext(context);
      await state.service.namespaceDelete({ req });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
