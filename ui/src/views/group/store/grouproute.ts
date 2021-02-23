/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { client } from '@/views/base/services/utils/requester';
import {
  serviceOptions,
  GroupRouteService, groupRoutePair, groupRouteListRequest,
} from '@/views/group/generated';

import { Route } from '@/views/route/entity';
import { ListResponse } from '@/views/common/helpers/types';
import { moduleActionContext } from '../../base/store';

serviceOptions.axios = client;
export interface State{
  service: GroupRouteService;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new GroupRouteService(),
  } as State,
  actions: {
    async GetList(context, req: groupRouteListRequest): Promise<ListResponse<Route>> {
      const { state } = actionContext(context);
      return state.service.groupRoutesList({ req });
    },
    async Create(context, pairs: groupRoutePair[]): Promise<Route[]> {
      const { state } = actionContext(context);
      return state.service.groupRoutesCreate({ req: { pairs } });
    },
    async Delete(context, pairs: groupRoutePair[]): Promise<void> {
      const { state } = actionContext(context);
      await state.service.groupRoutesDelete({ req: { pairs } });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
