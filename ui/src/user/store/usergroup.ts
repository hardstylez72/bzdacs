/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */
import { moduleActionContext } from '@/base/store';
import { defineModule } from 'direct-vuex';

import { Group } from '@/group/entity';
import { client } from '@/base/services/utils/requester';
import {
  serviceOptions, userGroupListRequest, userGroupPair, UserGroupService,
} from '@/user/generated';
import { ListResponse } from '@/common/helpers/types';

serviceOptions.axios = client;
export interface State{
  service: UserGroupService;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new UserGroupService(),
  } as State,
  actions: {
    async GetList(context, req: userGroupListRequest): Promise<ListResponse<Group>> {
      const { state } = actionContext(context);
      return state.service.userGroupList({ req });
    },
    async Create(context, pairs: userGroupPair[]): Promise<Group[]> {
      const { state } = actionContext(context);
      return state.service.userGroupCreate({ req: { pairs } });
    },
    async Delete(context, pairs: userGroupPair[]): Promise<void> {
      const { state } = actionContext(context);
      await state.service.userGroupDelete({ req: { pairs } });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
