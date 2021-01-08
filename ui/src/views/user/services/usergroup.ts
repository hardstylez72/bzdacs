import { Method } from 'axios';

import { Group } from '@/views/group/services/group';
import { makeRequest, Request } from '../../base/services/utils/requester';

export interface UserGroup {
  groupId: number;
  userId: number;
}

interface Options {
  host: string;
  baseUrl: string;
}

export default class UserGroupService {
  private options: Options

  private readonly methodPost: Method = 'POST'

  private readonly baseUrl: string;

  constructor(options: Options) {
    this.options = options;
    this.baseUrl = `${this.options.host}${this.options.baseUrl}`;
  }

  Create(t: UserGroup[]): Promise<Group[]> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/create`,
    };
    return makeRequest(req);
  }

  Delete(t: UserGroup[]): Promise<void> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/delete`,
    };
    return makeRequest(req);
  }

  GetList(userId: number, belongToUser: boolean): Promise<Group[]> {
    const req: Request = {
      data: { userId, belongToUser },
      method: this.methodPost,
      url: `${this.baseUrl}/list`,
    };
    return makeRequest(req);
  }
}
