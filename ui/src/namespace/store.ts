/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { client } from '@/base/services/utils/requester';
import {
  serviceOptions,
  NamespaceService,
  namespaceInsertRequest, namespaceUpdateRequest, namespaceDeleteRequest,
} from '@/namespace/generated';
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
    async Create(context, req: namespaceInsertRequest): Promise<Namespace> {
      const { state } = actionContext(context);
      return state.service.namespaceCreate({ req });
    },
    async Update(context, req: namespaceUpdateRequest): Promise<Namespace> {
      const { state } = actionContext(context);
      return state.service.namespaceUpdate({ req });
    },
    async Delete(context, req: namespaceDeleteRequest): Promise<void> {
      const { state } = actionContext(context);
      await state.service.namespaceDelete({ req });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
