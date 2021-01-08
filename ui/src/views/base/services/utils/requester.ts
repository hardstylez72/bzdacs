import axios, { AxiosError, Method } from 'axios';
import { v4 as uuid } from 'uuid';
// eslint-disable-next-line import/no-cycle
import store from '@/views/base/store';

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
        store.commit.showError(err.response.data.error.message);
      }
      throw err;
    });
};

export const makeRequest = (req: Request) => requester(req);
