import { makeRequest, Request } from '@/views/base/services/utils/requester';
import { Entity } from '@/views/base/services/entity';
import DefaultService from '@/views/base/services/default';

export interface User extends Entity {
  id: number;
  externalId: string;
  isSystem: boolean;
  description: string;
}

export default class UserService extends DefaultService<User> {
  GetById(id: number): Promise<User> {
    const req: Request = {
      data: { id },
      method: this.methodPost,
      url: `${this.baseUrl}/get`,
    };
    return makeRequest(req);
  }
}
