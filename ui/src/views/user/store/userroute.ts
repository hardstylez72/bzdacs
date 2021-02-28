/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';

import { client } from '@/views/base/services/utils/requester';
import {
  serviceOptions, userGroupPair,
  userRouteListRequest,
  userRoutePair, userRoutePairToDelete, userRouteRoute,
  userRouteRouteWithGroups,
  UserRouteService,
  userRouteUpdateResponse,
} from '@/views/user/generated';

import { ListResponse } from '@/views/common/helpers/types';
import { moduleActionContext } from '../../base/store';

serviceOptions.axios = client;

export interface State{
  service: UserRouteService;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new UserRouteService(),
  } as State,
  actions: {
    async UpdateRoute(context, req: userRoutePair): Promise<userRouteUpdateResponse> {
      const { state } = actionContext(context);
      return state.service.userRouteUpdate({ req });
    },
    async GetList(context, req: userRouteListRequest): Promise<ListResponse<userRouteRouteWithGroups>> {
      const { state } = actionContext(context);
      return state.service.userRouteList({ req });
    },
    async Create(context, pairs: userRoutePair[]): Promise<userRouteRoute[]> {
      const { state } = actionContext(context);
      return state.service.userRouteCreate({ req: { pairs } });
    },
    async Delete(context, pairs: userRoutePairToDelete[]): Promise<void> {
      const { state } = actionContext(context);
      await state.service.userRouteDelete({ req: { pairs } });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
