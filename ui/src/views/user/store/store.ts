/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { moduleActionContext } from '@/views/base/store';
import { client } from '@/views/base/services/utils/requester';
import {
  serviceOptions,
  userCreateRequest, userDeleteRequest,
  userGetByIdRequest,
  userListRequest,
  UserService,
  userUpdateRequest,
} from '@/views/user/generated';
import { ListResponse } from '@/views/common/helpers/types';
import { User } from '../entity';

serviceOptions.axios = client;
export interface State<T>{
  service: UserService;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new UserService(),
  } as State<User>,
  actions: {
    async GetById(context, req: userGetByIdRequest): Promise<User> {
      const { state } = actionContext(context);
      return state.service.userGetById({ req });
    },
    async GetList(context, req: userListRequest): Promise<ListResponse<User>> {
      const { state } = actionContext(context);
      return state.service.userList({ req });
    },
    async Create(context, req: userCreateRequest): Promise<User> {
      const { state } = actionContext(context);
      return state.service.userCreate({ req });
    },
    async Update(context, req: userUpdateRequest): Promise<User> {
      const { state } = actionContext(context);
      return state.service.userUpdate({ req });
    },
    async Delete(context, req: userDeleteRequest): Promise<void> {
      const { state, commit } = actionContext(context);
      await state.service.userDelete({ req });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
