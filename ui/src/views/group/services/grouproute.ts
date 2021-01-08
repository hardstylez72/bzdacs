import { Method } from 'axios';
import { Route } from '@/views/route/service';

import { makeRequest, Request } from '../../base/services/utils/requester';

export interface GroupRoute {
  groupId: number;
  routeId: number;
}

interface Options {
  host: string;
  baseUrl: string;
}

export default class GroupRouteService {
  private options: Options

  private readonly methodPost: Method = 'POST'

  private readonly baseUrl: string;

  constructor(options: Options) {
    this.options = options;
    this.baseUrl = `${this.options.host}${this.options.baseUrl}`;
  }

  Create(t: GroupRoute[]): Promise<Route[]> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/create`,
    };
    return makeRequest(req);
  }

  Delete(t: GroupRoute[]): Promise<void> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/delete`,
    };
    return makeRequest(req);
  }

  GetList(groupId: number, belongToGroup: boolean): Promise<Route[]> {
    const req: Request = {
      data: { groupId, belongToGroup },
      method: this.methodPost,
      url: `${this.baseUrl}/list`,
    };
    return makeRequest(req);
  }
}
