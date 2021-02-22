/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { client } from '@/views/base/services/utils/requester';
import { moduleActionContext } from '../base/store';
import { SystemSwg as System } from './entity';
import {
  serviceOptions,
  system_insertRequest,
  system_updateRequest,
  SystemService,
} from './generated/index';

serviceOptions.axios = client;

export interface State {
  service: SystemService;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new SystemService(),
  } as State,
  actions: {
    async GetById(context, id: number): Promise<System> {
      const { state } = actionContext(context);
      return state.service.systemGet({ req: { id } });
    },
    async GetList(context): Promise<System[]> {
      const { state } = actionContext(context);
      return state.service.systemList();
    },
    async Create(context, req: system_insertRequest): Promise<System> {
      const { state } = actionContext(context);
      return state.service.systemCreate({ req });
    },
    async Update(context, req: system_updateRequest): Promise<System> {
      const { state } = actionContext(context);
      return state.service.systemUpdate({ req });
    },
    async Delete(context, id: number): Promise<void> {
      const { state } = actionContext(context);
      return state.service.systemDelete({ req: { id } });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
