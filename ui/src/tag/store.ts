/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { client } from '@/base/services/utils/requester';
import {
  serviceOptions, tagDeleteRequest, tagInsertRequest, tagListRequest, tagUpdateRequest,
  TagService,
} from '@/tag/generated';
import { ListResponse } from '@/common/helpers/types';
import { Route } from '@/route/entity';
import { moduleActionContext } from '../base/store';
import { TagSwg as Tag } from './entity';

serviceOptions.axios = client;

export interface State {
  service: TagService;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new TagService(),
  } as State,
  actions: {
    async GetById(context, id: number): Promise<Tag> {
      const { state } = actionContext(context);
      return state.service.tagGet({ req: { id } });
    },
    async GetList(context, req: tagListRequest): Promise<ListResponse<Tag>> {
      const { state } = actionContext(context);
      // @ts-ignore
      return state.service.tagList({ req });
    },
    async Create(context, req: tagInsertRequest): Promise<Tag> {
      const { state } = actionContext(context);
      return state.service.tagCreate({ req });
    },
    async GetByPattern(context, req: tagListRequest): Promise<Tag[]> {
      const { state } = actionContext(context);
      const res = await state.service.tagList({ req });
      // @ts-ignore
      return res.items;
    },
    async Update(context, req: tagUpdateRequest): Promise<Tag> {
      const { state } = actionContext(context);
      return state.service.tagUpdate({ req });
    },
    async Delete(context, req: tagDeleteRequest): Promise<void> {
      const { state } = actionContext(context);
      await state.service.tagDelete({ req });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
