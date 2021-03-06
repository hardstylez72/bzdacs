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

export class NamespaceService {
  /**
   *
   */
  namespaceCreate(
    params: {
      /** request */
      req: namespaceInsertRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<namespaceGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/namespace/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  namespaceDelete(
    params: {
      /** request */
      req: namespaceDeleteRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<any> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/namespace/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  namespaceGet(
    params: {
      /** request */
      req: namespaceGetRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<namespaceGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/namespace/get';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  namespaceList(
    params: {
      /** request */
      req: namespaceListRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<namespaceGetResponse[]> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/namespace/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  namespaceUpdate(
    params: {
      /** request */
      req: namespaceUpdateRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<namespaceGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/namespace/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
}

export interface namespaceDeleteRequest {
  /**  */
  namespaceId: number;
}

export interface namespaceGetRequest {
  /**  */
  id?: number;

  /**  */
  name?: string;

  /**  */
  systemId?: number;
}

export interface namespaceGetResponse {
  /**  */
  createdAt: string;

  /**  */
  deletedAt?: string;

  /**  */
  id: number;

  /**  */
  name: string;

  /**  */
  systemId: number;

  /**  */
  updatedAt: string;
}

export interface namespaceInsertRequest {
  /**  */
  name: string;

  /**  */
  systemId: number;
}

export interface namespaceListRequest {
  /**  */
  id: number;
}

export interface namespaceUpdateRequest {
  /**  */
  id: number;

  /**  */
  name: string;
}
