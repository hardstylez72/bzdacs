import { Entity } from '@/views/base/services/entity';
import { makeRequest, Request } from '@/views/base/services/utils/requester';
import DefaultService from '../base/services/default';

export interface Tag extends Entity {
  id: number;
  name: string;
}

export default class TagService extends DefaultService<Tag> {
  Update(t: Tag): Promise<Tag> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/update`,
    };
    return makeRequest(req);
  }

  GetById(id: number): Promise<Tag> {
    const req: Request = {
      data: { id },
      method: this.methodPost,
      url: `${this.baseUrl}/get`,
    };
    return makeRequest(req);
  }

  async GetByPattern(name: string): Promise<Tag[]> {
    if (!name) {
      const out: Tag[] = [];
      return out;
    }
    const req: Request = {
      data: { pattern: name },
      method: this.methodPost,
      url: `${this.baseUrl}/suggest`,
    };
    return makeRequest(req);
  }
}
