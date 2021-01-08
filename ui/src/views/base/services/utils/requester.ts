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
          // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
          // @ts-ignore
          this.$router.push({ name: 'Guest' });
        } else if (err.response.status === 403) {
          // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
          // @ts-ignore
          this.$store.commit.showError('user does not have permissions to do this operation');
        } else {
          // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
          // @ts-ignore
          this.$store.commit.showError(err.response.data);
        }
      }
      throw err;
    });
};

export const makeRequest = (req: Request) => requester(req);
