/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { client } from '@/views/base/services/utils/requester';
import {
  serviceOptions, tag_deleteRequest, tag_insertRequest, tag_listRequest, tag_updateRequest,
  TagService,
} from '@/views/tag/generated';
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
    async GetList(context, req: tag_listRequest): Promise<{items: Tag[]; total: number}> {
      const { state } = actionContext(context);
      // @ts-ignore
      return state.service.tagList({ req });
    },
    async Create(context, req: tag_insertRequest): Promise<Tag> {
      const { state } = actionContext(context);
      return state.service.tagCreate({ req });
    },
    async GetByPattern(context, req: tag_listRequest): Promise<Tag[]> {
      const { state } = actionContext(context);
      const res = await state.service.tagList({ req });
      // @ts-ignore
      return res.items;
    },
    async Update(context, req: tag_updateRequest): Promise<Tag> {
      const { state } = actionContext(context);
      return state.service.tagUpdate({ req });
    },
    async Delete(context, req: tag_deleteRequest): Promise<void> {
      const { state } = actionContext(context);
      await state.service.tagDelete({ req });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
