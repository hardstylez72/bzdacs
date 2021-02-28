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
      req: groupInsertRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<groupGetResponse> {
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
      req: groupDeleteRequest;
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
      req: groupGetByCodeRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<groupGetResponse> {
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
      req: groupGetByIdRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<groupGetResponse> {
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
      req: groupListRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<groupListResponse> {
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
      req: groupUpdateRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<groupGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/group/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
}

export class GroupRouteService {
  /**
   *
   */
  groupRoutesCreate(
    params: {
      /** request */
      req: groupRouteInsertRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<groupRouteRoute[]> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/group/route/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  groupRoutesDelete(
    params: {
      /** request */
      req: groupRouteDeleteRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<any> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/group/route/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  groupRoutesList(
    params: {
      /** request */
      req: groupRouteListRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<groupRouteListResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/group/route/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
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

export interface groupDeleteRequest {
  /**  */
  id: number;

  /**  */
  namespaceId: number;
}

export interface groupGetByCodeRequest {
  /**  */
  code: string;

  /**  */
  namespaceId: number;
}

export interface groupGetByIdRequest {
  /**  */
  id: number;

  /**  */
  namespaceId: number;
}

export interface groupGetResponse {
  /**  */
  code: string;

  /**  */
  createdAt: string;

  /**  */
  deletedAt?: string;

  /**  */
  description: string;

  /**  */
  id: number;

  /**  */
  namespaceId: number;

  /**  */
  updatedAt: string;
}

export interface groupInsertRequest {
  /**  */
  baseGroupId?: number;

  /**  */
  code: string;

  /**  */
  description: string;

  /**  */
  namespaceId: number;
}

export interface groupListRequest {
  /**  */
  filter?: group_filter;
}

export interface groupListResponse {
  /**  */
  items: groupGetResponse[];

  /**  */
  total: number;
}

export interface groupRouteDeleteRequest {
  /**  */
  pairs: groupRoutePair[];
}

export interface groupRouteFilter {
  /**  */
  belongToGroup?: boolean;

  /**  */
  groupId: number;

  /**  */
  namespaceId: number;

  /**  */
  page: number;

  /**  */
  pageSize?: number;

  /**  */
  routePattern?: string;
}

export interface groupRouteInsertRequest {
  /**  */
  pairs: groupRoutePair[];
}

export interface groupRouteListRequest {
  /**  */
  filter?: groupRouteFilter;
}

export interface groupRouteListResponse {
  /**  */
  items: groupRouteRoute[];

  /**  */
  total: number;
}

export interface groupRoutePair {
  /**  */
  groupId: number;

  /**  */
  routeId: number;
}

export interface groupRouteRoute {
  /**  */
  createdAt: string;

  /**  */
  deletedAt?: string;

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
  tags: string[];

  /**  */
  updatedAt: string;
}

export interface groupUpdateRequest {
  /**  */
  code: string;

  /**  */
  description: string;

  /**  */
  id: number;

  /**  */
  namespaceId: number;
}
