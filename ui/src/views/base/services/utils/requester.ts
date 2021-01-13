import axios, { AxiosError, Method } from 'axios';
import { v4 as uuid } from 'uuid';

const instance = axios.create();

export interface Request {
  method: Method;
  url: string;
  data: any;
  headers?: any;
}

const requester = (req: Request): Promise<any> => {
  const headers = {
    'x-request-id': uuid(),
    ...req.headers,
  };

  return instance({
    method: req.method,
    url: req.url,
    data: req.data,
    headers,
    withCredentials: true,
    timeout: 30000,
  })
    .then((res) => res.data)
    .catch(async (err: AxiosError) => {
      if (err.response) {
        if (err.response.status === 401) {
          const evt = new Event('req-status-401');
          window.dispatchEvent(evt);

          if (window.location.pathname === '/guest') {
            return;
          }
          if (window.location.pathname === '/admin') {
            return;
          }

          window.location.pathname = '/guest';
        } else if (err.response.status === 403) {
          const evt = new Event('req-status-403');
          window.dispatchEvent(evt);
        }
      }
      throw err;
    });
};

export const makeRequest = (req: Request) => requester(req);
