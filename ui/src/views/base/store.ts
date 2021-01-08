/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import Vue from 'vue';
import Vuex from 'vuex';
import { createDirectStore } from 'direct-vuex';
import GroupRouteService from '@/views/group/services/grouproute';
import { Route } from '@/views/route/service';
import { makeRequest, Request } from '@/views/base/services/utils/requester';
import routeModule from '../route/store';
import groupModule from '../group/store/group';
import userModule from '../user/store/store';
import groupRouteModule from '../group/store/grouproute';
import userGroupModule from '../user/store/usergroup';
import userRouteModule from '../user/store/userroute';
import tagModule from '../tag/store';

Vue.use(Vuex);

export interface State{
  isAuthorized: boolean;
}

const {
  store,
  rootActionContext,
  moduleActionContext,
  rootGetterContext,
  moduleGetterContext,
} = createDirectStore({

  state: {
    isAuthorized: false,
  } as State,
  getters: {
    isAuthorized(state) {
      return state.isAuthorized;
    },
  },
  mutations: {
    setAuthorized(state, isAuthorized: boolean) {
      state.isAuthorized = isAuthorized;
    },
  },
  actions: {
    login(context, payload?: {login: string; password: string}): Promise<void> {
      const { state, commit } = actionContext(context);
      const req: Request = {
        data: {},
        method: 'POST',
        url: '/api/v1/user/login',
      };

      if (payload) {
        req.headers = { token: window.btoa(`${payload.login}:${payload.password}`) };
      }

      return makeRequest(req)
        .then(() => {
          commit.setAuthorized(true);
        })
        .catch((err) => {
          console.error(err);
          commit.setAuthorized(false);
        });
    },
  },
  modules: {
    route: routeModule,
    group: groupModule,
    groupRoute: groupRouteModule,
    user: userModule,
    userGroup: userGroupModule,
    userRoute: userRouteModule,
    tag: tagModule,
  },
});

export default store;

export {
  rootActionContext,
  moduleActionContext,
  rootGetterContext,
  moduleGetterContext,
};

// @typescript-eslint/no-use-before-define
// eslint-disable-next-line @typescript-eslint/ban-ts-ignore
// @ts-ignore
const actionContext = (context: any) => rootActionContext(context);

export type AppStore = typeof store
declare module 'vuex' {
  interface Store<S> {
    direct: AppStore;
  }
}
