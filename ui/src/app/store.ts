/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import Vue from 'vue';
import Vuex from 'vuex';
import { createDirectStore } from 'direct-vuex';

import { SysUserService } from '@/app/generated';
import routeModule from '../route/store';
import groupModule from '../group/store/group';
import userModule from '../user/store/store';
import groupRouteModule from '../group/store/grouproute';
import userGroupModule from '../user/store/usergroup';
import userRouteModule from '../user/store/userroute';
import tagModule from '../tag/store';
import namespaceModule from '../namespace/store';
import systemModule from '../system/store';
import sysUserModule from './user';

Vue.use(Vuex);

export interface State{
  showSnackbar: boolean;
  snackbarMessage: string;
}

const {
  store,
  rootActionContext,
  moduleActionContext,
  rootGetterContext,
  moduleGetterContext,
} = createDirectStore({
  state: {
    showSnackbar: false,
    snackbarMessage: '',
  } as State,
  getters: {
    snackbarMessage(state) {
      return state.snackbarMessage;
    },
  },
  mutations: {
    setShowSnackbar(state, show: boolean) {
      state.showSnackbar = show;
    },
    showError(state, msg: string) {
      state.snackbarMessage = msg;
      state.showSnackbar = true;
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
    namespace: namespaceModule,
    system: systemModule,
    sysUser: sysUserModule,
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
