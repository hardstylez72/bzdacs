/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import Vue from 'vue';
import Vuex from 'vuex';
import { createDirectStore } from 'direct-vuex';
import routeModule from '../route/store';
import groupModule from '../group/store/group';
import userModule from '../user/store/store';
import groupRouteModule from '../group/store/grouproute';
import userGroupModule from '../user/store/usergroup';
import userRouteModule from '../user/store/userroute';
import tagModule from '../tag/store';
import AuthService, { Session } from './user';

Vue.use(Vuex);

export interface State{
  isAuthorized: boolean;
  showSnackbar: boolean;
  snackbarMessage: string;
  service: AuthService;
  login: string;
  isAdmin: boolean;
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
    showSnackbar: false,
    snackbarMessage: '',
    service: new AuthService(),
    login: '',
    isAdmin: false,
  } as State,
  getters: {
    login(state) {
      return state.login;
    },
    isAuthorized(state) {
      return state.isAuthorized;
    },
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

    setAuthorized(state, isAuthorized: boolean) {
      state.isAuthorized = isAuthorized;
    },
    setUser(state, payload: {login: string; isAdmin: boolean}) {
      state.login = payload.login;
      state.isAdmin = payload.isAdmin;
    },
  },
  actions: {

    async logout(context): Promise<Error> {
      const { commit, state } = actionContext(context);

      return state.service.logout()
        .then((s: Session) => {
          commit.setAuthorized(false);
          return s;
        })
        .catch((err) => {
          commit.setAuthorized(true);
          return err;
        })
        .finally(() => {
          window.location.reload();
        });
    },
    async userSession(context): Promise<Session | Error> {
      const { commit, state } = actionContext(context);

      return state.service.userSession()
        .then((s: Session) => {
          if (s.login !== '') {
            commit.setUser({ login: s.login, isAdmin: s.isAdmin });
            commit.setAuthorized(true);
            return s;
          }
        })
        .catch((err) => {
          commit.setAuthorized(false);
          return err;
        });
    },

    guestLogin(context, login): Promise<void> {
      const { state, commit } = actionContext(context);

      return state.service.guestLogin(login).then(() => {
        commit.setAuthorized(true);
      })
        .catch((err) => {
          commit.setAuthorized(false);
        });
    },
    adminLogin(context, payload?: {login: string; password: string}): Promise<void> {
      const { state, commit } = actionContext(context);

      return state.service.adminLogin(payload)
        .then(() => {
          commit.setAuthorized(true);
        })
        .catch((err) => {
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
