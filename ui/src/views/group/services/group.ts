import { makeRequest, Request } from '@/views/base/services/utils/requester';
import { Entity } from '@/views/base/services/entity';
import DefaultService, { T } from '../../base/services/default';

export interface Group extends Entity {
  id: number;
  code: string;
  description: string;
}

export default class GroupService extends DefaultService<Group> {
  GetById(id: number): Promise<Group> {
    const req: Request = {
      data: { id },
      method: this.methodPost,
      url: `${this.baseUrl}/get`,
    };
    return makeRequest(req);
  }

  Update(t: Group): Promise<T> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/update`,
    };
    return makeRequest(req);
  }

  CreateBasedOnAnother(t: Group, baseGroupId: number): Promise<T> {
    const req: Request = {
      data: { ...t, baseGroupId },
      method: this.methodPost,
      url: `${this.baseUrl}/create`,
    };
    return makeRequest(req);
  }
}
