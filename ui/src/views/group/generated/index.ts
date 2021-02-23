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

export class GroupService {
  /**
   *
   */
  groupCreate(
    params: {
      /** request */
      req: group_insertRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<group_getResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/group/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  groupDelete(
    params: {
      /** request */
      req: group_deleteRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<any> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/group/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  groupGetByCode(
    params: {
      /** request */
      req: group_getByCodeRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<group_getResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/group/getByCode';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  groupGetById(
    params: {
      /** request */
      req: group_getByIdRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<group_getResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/group/getById';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  groupList(
    params: {
      /** request */
      req: group_listRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<group_listResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/group/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  groupUpdate(
    params: {
      /** request */
      req: group_updateRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<group_getResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/group/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
}

export interface group_deleteRequest {
  /**  */
  id: number;

  /**  */
  namespaceId: number;
}

export interface group_filter {
  /**  */
  namespaceId: number;

  /**  */
  page: number;

  /**  */
  pageSize: number;

  /**  */
  pattern?: string;
}

export interface group_getByCodeRequest {
  /**  */
  code: string;

  /**  */
  namespaceId: number;
}

export interface group_getByIdRequest {
  /**  */
  id: number;

  /**  */
  namespaceId: number;
}

export interface group_getResponse {
  /**  */
  code?: string;

  /**  */
  createdAt?: string;

  /**  */
  deletedAt?: string;

  /**  */
  description?: string;

  /**  */
  id?: number;

  /**  */
  namespaceId?: number;

  /**  */
  updatedAt?: string;
}

export interface group_insertRequest {
  /**  */
  baseGroupId?: number;

  /**  */
  code: string;

  /**  */
  description: string;

  /**  */
  namespaceId: number;
}

export interface group_listRequest {
  /**  */
  filter?: group_filter;
}

export interface group_listResponse {
  /**  */
  items?: group_getResponse[];

  /**  */
  total?: number;
}

export interface group_updateRequest {
  /**  */
  code: string;

  /**  */
  description: string;

  /**  */
  id: number;

  /**  */
  namespaceId: number;
}
