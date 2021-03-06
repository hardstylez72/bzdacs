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

export class UserService {
  /**
   *
   */
  userCreate(
    params: {
      /** request */
      req: userCreateRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<userGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  userDelete(
    params: {
      /** request */
      req: userDeleteRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<any> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  userGetById(
    params: {
      /** request */
      req: userGetByIdRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<userGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/getById';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  userGetByLogin(
    params: {
      /** request */
      req: userGetByLoginRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<userGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/getByLogin';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  userList(
    params: {
      /** request */
      req: userListRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<userListResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  userUpdate(
    params: {
      /** request */
      req: userUpdateRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<userGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
}

export class UserGroupService {
  /**
   *
   */
  userGroupCreate(
    params: {
      /** request */
      req: userGroupInsertRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<usergroup_Group[]> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/group/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  userGroupDelete(
    params: {
      /** request */
      req: userGroupDeleteRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<any> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/group/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  userGroupList(
    params: {
      /** request */
      req: userGroupListRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<userGroupListResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/group/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
}

export class UserRouteService {
  /**
   *
   */
  userRouteCreate(
    params: {
      /** request */
      req: userRouteInsertRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<userRouteRouteWithGroups[]> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/route/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  userRouteDelete(
    params: {
      /** request */
      req: userRouteDeleteRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<any> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/route/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  userRouteList(
    params: {
      /** request */
      req: userRouteListRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<userRouteListResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/route/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  userRouteUpdate(
    params: {
      /** request */
      req: userRoutePair;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<userRouteUpdateResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/user/route/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
}

export interface user_filter {
  /**  */
  namespaceId: number;

  /**  */
  page: number;

  /**  */
  pageSize?: number;

  /**  */
  pattern?: string;
}

export interface userCreateRequest {
  /**  */
  externalId: string;

  /**  */
  namespaceId: number;
}

export interface userDeleteRequest {
  /**  */
  id: number;

  /**  */
  namespaceId: number;
}

export interface userGetByIdRequest {
  /**  */
  id: number;

  /**  */
  namespaceId: number;
}

export interface userGetByLoginRequest {
  /**  */
  login: string;

  /**  */
  namespaceId: number;
}

export interface userGetResponse {
  /**  */
  createdAt: string;

  /**  */
  deletedAt?: string;

  /**  */
  externalId: string;

  /**  */
  id: number;

  /**  */
  namespaceId: number;

  /**  */
  updatedAt: string;
}

export interface userGroupDeleteRequest {
  /**  */
  pairs: userGroupPair[];
}

export interface userGroupFilter {
  /**  */
  belongToUser?: boolean;

  /**  */
  namespaceId: number;

  /**  */
  page: number;

  /**  */
  pageSize?: number;

  /**  */
  pattern?: string;

  /**  */
  userId: number;
}

export interface userGroupInsertRequest {
  /**  */
  pairs: userGroupPair[];
}

export interface userGroupListRequest {
  /**  */
  filter?: userGroupFilter;
}

export interface userGroupListResponse {
  /**  */
  items: usergroup_Group[];

  /**  */
  total: number;
}

export interface userGroupPair {
  /**  */
  groupId: number;

  /**  */
  userId: number;
}

export interface userListRequest {
  /**  */
  filter?: user_filter;
}

export interface userListResponse {
  /**  */
  items: userGetResponse[];

  /**  */
  total: number;
}

export interface userRouteDeleteRequest {
  /**  */
  pairs: userRoutePairToDelete[];
}

export interface userRouteFilter {
  /**  */
  belongToUser?: boolean;

  /**  */
  namespaceId: number;

  /**  */
  page: number;

  /**  */
  pageSize?: number;

  /**  */
  pattern?: string;

  /**  */
  userId: number;
}

export interface userRouteInsertRequest {
  /**  */
  pairs: userRoutePair[];
}

export interface userRouteListRequest {
  /**  */
  filter?: userRouteFilter;
}

export interface userRouteListResponse {
  /**  */
  items: userRouteRouteWithGroups[];

  /**  */
  total: number;
}

export interface userRoutePair {
  /**  */
  isExcluded?: boolean;

  /**  */
  routeId: number;

  /**  */
  userId: number;
}

export interface userRoutePairToDelete {
  /**  */
  routeId: number;

  /**  */
  userId: number;
}

export interface userRouteRouteWithGroups {
  /**  */
  createdAt: string;

  /**  */
  deletedAt?: string;

  /**  */
  description: string;

  /**  */
  groups: userroute_Group[];

  /**  */
  id: number;

  /**  */
  isExcluded: boolean;

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

export interface userRouteUpdateResponse {
  /**  */
  createdAt: string;

  /**  */
  deletedAt?: string;

  /**  */
  description: string;

  /**  */
  groups: userroute_Group[];

  /**  */
  id: number;

  /**  */
  isExcluded: boolean;

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

export interface userUpdateRequest {
  /**  */
  externalId: string;

  /**  */
  id: number;

  /**  */
  namespaceId: number;
}

export interface usergroup_Group {
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

export interface userroute_Group {
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
