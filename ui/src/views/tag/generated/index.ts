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

export class TagService {
  /**
   *
   */
  tagCreate(
    params: {
      /** request */
      req: tagInsertRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<tagGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/tag/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  tagDelete(
    params: {
      /** request */
      req: tagDeleteRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<any> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/tag/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  tagGet(
    params: {
      /** request */
      req: tagGetRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<tagGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/tag/get';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  tagList(
    params: {
      /** request */
      req: tagListRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<tagListResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/tag/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  tagUpdate(
    params: {
      /** request */
      req: tagUpdateRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<tagGetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/v1/tag/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['req'];

      configs.data = data;
      axios(configs, resolve, reject);
    });
  }
}

export interface tag_filter {
  /**  */
  namespaceId: number;

  /**  */
  page: number;

  /**  */
  pageSize: number;

  /**  */
  pattern?: string;
}

export interface tagDeleteRequest {
  /**  */
  id: number;

  /**  */
  namespaceId: number;
}

export interface tagGetRequest {
  /**  */
  id: number;
}

export interface tagGetResponse {
  /**  */
  createdAt: string;

  /**  */
  deletedAt?: string;

  /**  */
  id: number;

  /**  */
  name: string;

  /**  */
  namespaceId: number;

  /**  */
  updatedAt: string;
}

export interface tagInsertRequest {
  /**  */
  name: string;

  /**  */
  namespaceId: number;
}

export interface tagListRequest {
  /**  */
  filter?: tag_filter;
}

export interface tagListResponse {
  /**  */
  items?: tagGetResponse[];

  /**  */
  total?: number;
}

export interface tagUpdateRequest {
  /**  */
  id: number;

  /**  */
  name: string;

  /**  */
  namespaceId: number;
}
