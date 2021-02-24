/** Generate by swagger-axios-codegen */
// tslint:disable
/* eslint-disable */
export interface IRequestOptions {
  headers?: any;
}

export interface IRequestPromise<T = any> extends Promise<IRequestResponse<T>> {}

export interface IRequestResponse<T = any> {
  data: T;
  status: number;
  statusText: string;
  headers: any;
  config: any;
  request?: any;
}

export interface IRequestInstance {
  (config: any): IRequestPromise;
  (url: string, config?: any): IRequestPromise;
  request<T = any>(config: any): IRequestPromise<T>;
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
  axios?: IRequestInstance;
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

export class SystemService {
  /**
   *
   */
  systemCreate(
    params: {
      /** request */
      req: systemInsertRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<systemGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/system/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  systemDelete(
    params: {
      /** request */
      req: systemDeleteRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<any> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/system/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  systemGet(
    params: {
      /** request */
      req: systemGetRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<systemGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/system/get';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  systemList(options: IRequestOptions = {}): Promise<systemGetResponse[]> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/system/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = null;

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  systemUpdate(
    params: {
      /** request */
      req: systemUpdateRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<systemGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/system/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
}

export interface systemDeleteRequest {
  /**  */
  id: number;
}

export interface systemGetRequest {
  /**  */
  id?: number;

  /**  */
  name?: string;
}

export interface systemGetResponse {
  /**  */
  createdAt: string;

  /**  */
  deletedAt?: string;

  /**  */
  id: number;

  /**  */
  name: string;

  /**  */
  updatedAt: string;
}

export interface systemInsertRequest {
  /**  */
  name: string;
}

export interface systemUpdateRequest {
  /**  */
  id: number;

  /**  */
  name: string;
}
