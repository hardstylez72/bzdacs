/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { moduleActionContext } from '../base/store';
import TagService, { Tag } from './service';

export interface State {
  service: TagService;
  tags: Tag[];
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new TagService({ host: '', baseUrl: '/api/v1/tag' }),
    tags: [],
  } as State,
  getters: {
    getTags(state) {
      return state.tags;
    },
  },
  mutations: {
    setTags(state, routes: Tag[]) {
      state.tags = routes;
    },
    deleteTag(state, id: number) {
      state.tags = state.tags.filter((route) => route.id !== id);
    },
    addTag(state, routes: Tag) {
      state.tags.push(routes);
    },
    updateTag(state, route: Tag) {
      state.tags = state.tags.map((r) => {
        if (route.id === r.id) {
          return route;
        }
        return r;
      });
    },
  },
  actions: {
    async GetById(context, id: number): Promise<Tag> {
      const { state } = actionContext(context);

      return state.service.GetById(id);
    },
    async GetList(context): Promise<Tag[]> {
      const { state, commit } = actionContext(context);
      const routes = await state.service.GetList();
      commit.setTags(routes);
      return routes;
    },
    async Create(context, route: Tag): Promise<Tag> {
      const { state, commit } = actionContext(context);
      const createdRoute = await state.service.Create(route);
      commit.addTag(createdRoute);
      return createdRoute;
    },
    async GetByPattern(context, name: string): Promise<Tag[]> {
      const { state } = actionContext(context);
      return state.service.GetByPattern(name);
    },
    async Update(context, route: Tag): Promise<Tag> {
      const { state, commit } = actionContext(context);
      const createdRoute = await state.service.Update(route);
      commit.updateTag(createdRoute);
      return createdRoute;
    },
    async Delete(context, id: number): Promise<void> {
      const { state, commit } = actionContext(context);

      await state.service.Delete(id);
      commit.deleteTag(id);
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
