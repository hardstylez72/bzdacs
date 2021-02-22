/** Generate by swagger-axios-codegen */
// tslint:disable
/* eslint-disable */
import axiosStatic, { AxiosInstance } from 'axios';

export interface IRequestOptions {
  headers?: any;
  baseURL?: string;
  responseType?: string;
}

export interface IRequestConfig {
  method?: any;
  headers?: any;
  url?: any;
  data?: any;
  params?: any;
}

// Add options interface
export interface ServiceOptions {
  axios?: AxiosInstance;
}

// Add default options
export const serviceOptions: ServiceOptions = {};

// Instance selector
export function axios(configs: IRequestConfig, resolve: (p: any) => void, reject: (p: any) => void): Promise<any> {
  if (serviceOptions.axios) {
    return serviceOptions.axios
      .request(configs)
      .then(res => {
        resolve(res.data);
      })
      .catch(err => {
        reject(err);
      });
  } else {
    throw new Error('please inject yourself instance like axios  ');
  }
}

export function getConfigs(method: string, contentType: string, url: string, options: any): IRequestConfig {
  const configs: IRequestConfig = { ...options, method, url };
  configs.headers = {
    ...options.headers,
    'Content-Type': contentType
  };
  return configs;
}

const basePath = '/api';

export interface IList<T> extends Array<T> {}
export interface List<T> extends Array<T> {}
export interface IDictionary<TValue> {
  [key: string]: TValue;
}
export interface Dictionary<TValue> extends IDictionary<TValue> {}

export interface IListResult<T> {
  items?: T[];
}

export class ListResultDto<T> implements IListResult<T> {
  items?: T[];
}

export interface IPagedResult<T> extends IListResult<T> {
  totalCount?: number;
  items?: T[];
}

export class PagedResultDto<T> implements IPagedResult<T> {
  totalCount?: number;
  items?: T[];
}

// customer definition
// empty

export class RouteService {
  /**
   *
   */
  static routeCreate(
    params: {
      /** request */
      req: route_insertRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<route_getResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/route/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  static routeDelete(
    params: {
      /** request */
      req: route_deleteRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<any> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/route/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  static routeGetById(
    params: {
      /** request */
      req: route_getByIdRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<route_getResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/route/getById';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  static routeGetByParams(
    params: {
      /** request */
      req: route_getByParamsRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<route_getResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/route/getByParams';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  static routeUpdate(
    params: {
      /** request */
      req: route_updateRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<route_getResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/route/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
}

export interface route_deleteRequest {
  /**  */
  id: number;

  /**  */
  namespaceId: number;
}

export interface route_getByIdRequest {
  /**  */
  id: number;

  /**  */
  namespaceId: number;
}

export interface route_getByParamsRequest {
  /**  */
  method: string;

  /**  */
  namespaceId: number;

  /**  */
  route: string;
}

export interface route_getResponse {
  /**  */
  createdAt?: string;

  /**  */
  deletedAt?: string;

  /**  */
  description?: string;

  /**  */
  id?: number;

  /**  */
  method?: string;

  /**  */
  namespaceId?: number;

  /**  */
  route?: string;

  /**  */
  tags?: string[];

  /**  */
  updatedAt?: string;
}

export interface route_insertRequest {
  /**  */
  description: string;

  /**  */
  method: string;

  /**  */
  namespaceId: number;

  /**  */
  route: string;

  /**  */
  tags?: string[];
}

export interface route_updateRequest {
  /**  */
  description: string;

  /**  */
  id: number;

  /**  */
  method: string;

  /**  */
  namespaceId: number;

  /**  */
  route: string;

  /**  */
  tags?: string[];
}
