import { Entity } from '@/views/base/services/entity';
import { makeRequest, Request } from '@/views/base/services/utils/requester';
import DefaultService from '../base/services/default';

export interface Namespace extends Entity {
  id: number;
  name: string;
}

export default class NamespaceService extends DefaultService<Namespace> {
  Update(t: Namespace): Promise<Namespace> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/update`,
    };
    return makeRequest(req);
  }

  GetById(id: number): Promise<Namespace> {
    const req: Request = {
      data: { id },
      method: this.methodPost,
      url: `${this.baseUrl}/get`,
    };
    return makeRequest(req);
  }
}
