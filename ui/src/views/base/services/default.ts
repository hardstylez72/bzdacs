import { Method } from 'axios';
import { makeRequest, Request } from './utils/requester';
import Service from './service';

export interface T {
  id: number;
  code: string;
  description: string;
}

interface Options {
  host: string;
  baseUrl: string;
}

export default class DefaultService<T> implements Service<T> {
  private options: Options

  readonly methodPost: Method = 'POST'

  readonly baseUrl: string;

  constructor(options: Options) {
    this.options = options;
    this.baseUrl = `${this.options.host}${this.options.baseUrl}`;
  }

  Create(t: T): Promise<T> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/create`,
    };
    return makeRequest(req);
  }

  Delete(id: number): Promise<void> {
    const req: Request = {
      data: { id },
      method: this.methodPost,
      url: `${this.baseUrl}/delete`,
    };
    return makeRequest(req);
  }

  GetList(): Promise<T[]> {
    const req: Request = {
      data: {},
      method: this.methodPost,
      url: `${this.baseUrl}/list`,
    };
    return makeRequest(req);
  }
}
