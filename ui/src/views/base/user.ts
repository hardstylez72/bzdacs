import { makeRequest, Request } from '@/views/base/services/utils/requester';

export interface Session {
  isAdmin: boolean;
  login: string;
}

export default class User {
  async logout(): Promise<Session> {
    const req: Request = {
      data: {},
      method: 'POST',
      url: '/api/v1/user/logout',
    };
    return makeRequest(req);
  }

  async userSession(): Promise<Session> {
    const req: Request = {
      data: {},
      method: 'POST',
      url: '/api/v1/user/session/get',
    };
    return makeRequest(req);
  }

  async guestLogin(login: string): Promise<void> {
    const req: Request = {
      data: {},
      method: 'POST',
      url: '/api/v1/user/guest/login',
    };
    if (login) {
      req.headers = {
        login:
          encodeURIComponent(login),
      };
    }

    return makeRequest(req);
  }

  adminLogin(payload?: {login: string; password: string}): Promise<void> {
    const req: Request = {
      data: {},
      method: 'POST',
      url: '/api/v1/user/admin/login',
    };
    if (payload) {
      req.headers = { token: window.btoa(`${payload.login}:${payload.password}`) };
    }

    return makeRequest(req);
  }
}
