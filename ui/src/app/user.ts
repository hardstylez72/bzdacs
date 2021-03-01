/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */
import store, { moduleActionContext } from '@/app/store';
import { defineModule } from 'direct-vuex';
import { client } from '@/app/util/requester';
import {
  serviceOptions,
  sysUserLoginRequest, sysUserRegisterRequest, SysUserService,
} from '@/app/generated';

serviceOptions.axios = client;

export interface State{
  isAuthorized: boolean;
  service: SysUserService;
  login: string;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    isAuthorized: false,
    service: new SysUserService(),
    login: '',
  } as State,
  getters: {
    login(state) {
      return state.login;
    },
    isAuthorized(state) {
      return state.isAuthorized;
    },
  },
  mutations: {
    setAuthorized(state, isAuthorized: boolean) {
      state.isAuthorized = isAuthorized;
    },
    setUser(state, login: string) {
      state.login = login;
    },
  },
  actions: {
    async logout(context): Promise<void> {
      const { state } = actionContext(context);

      return state.service.sysUserLogout()
        .then(() => {
          window.location.reload();
        });
    },
    async userSession(context): Promise<Error | void> {
      const { commit, state } = actionContext(context);

      await state.service.sysUserSession()
        .then((s) => {
          if (s.login) {
            commit.setUser(s.login);
            commit.setAuthorized(true);
            return s;
          }
        })
        .catch((err) => {
          commit.setAuthorized(false);
          throw err;
        });
    },
    async login(context, req: sysUserLoginRequest): Promise<void> {
      const { state } = actionContext(context);

      return state.service.sysUserLogin({ req })
        .catch((err) => {
          store.commit.showError('Invalid login or password');
        });
    },
    async register(context, req: sysUserRegisterRequest): Promise<void> {
      const { state } = actionContext(context);
      return state.service.sysUserRegister({ req });
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
