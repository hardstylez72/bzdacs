import axios from 'axios';
import { v4 as uuid } from 'uuid';

const instance = axios.create();

instance.interceptors.request.use(
  ((request) => {
    request.headers['x-request-id'] = uuid();
    return request;
  }),
  ((error) => Promise.reject(error)),
);

instance.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.message === 'Request failed with status code 403') {
      const evt = new Event('req-status-403');
      window.dispatchEvent(evt);
    }

    if (error.message === 'Request failed with status code 401') {
      const evt = new Event('req-status-401');
      window.dispatchEvent(evt);
      if (window.location.pathname === '/login' || window.location.pathname === '/register') {
        return;
      }
      window.location.pathname = '/login';
    }
    return Promise.reject(error);
  },
);

// eslint-disable-next-line import/prefer-default-export
export const client = instance;
