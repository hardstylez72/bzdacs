/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { client } from '@/views/base/services/utils/requester';
import {
  serviceOptions,
  GroupService,
  groupGetByIdRequest,
  groupListRequest,
  groupUpdateRequest,
  groupInsertRequest, groupDeleteRequest,
} from '@/views/group/generated';
import { ListResponse } from '@/views/common/helpers/types';
import { Group } from '../entity';

import { moduleActionContext } from '../../base/store';

serviceOptions.axios = client;

export interface State<T>{
  service: GroupService;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new GroupService(),
  } as State<Group>,
  actions: {
    async GetById(context, req: groupGetByIdRequest): Promise<Group> {
      const { state } = actionContext(context);
      return state.service.groupGetById({ req });
    },
    async GetList(context, req: groupListRequest): Promise<ListResponse<Group>> {
      const { state } = actionContext(context);
      return state.service.groupList({ req });
    },
    async Update(context, req: groupUpdateRequest): Promise<Group> {
      const { state } = actionContext(context);
      return state.service.groupUpdate({ req });
    },
    async Create(context, req: groupInsertRequest): Promise<Group> {
      const { state } = actionContext(context);
      return state.service.groupCreate({ req });
    },
    async Delete(context, req: groupDeleteRequest): Promise<void> {
      const { state } = actionContext(context);
      await state.service.groupDelete({ req });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
