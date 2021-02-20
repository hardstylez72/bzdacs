import { Entity } from '@/views/base/services/entity';
import { makeRequest, Request } from '@/views/base/services/utils/requester';
import { Method } from 'axios';

export interface Namespace extends Entity {
  name: string;
}
interface Options {
  host: string;
  baseUrl: string;
}

export default class NamespaceService {
  private options: Options

  readonly methodPost: Method = 'POST'

  readonly baseUrl: string;

  constructor(options: Options) {
    this.options = options;
    this.baseUrl = `${this.options.host}${this.options.baseUrl}`;
  }

  Create(namespace: Namespace, systemId: number): Promise<Namespace> {
    const req: Request = {
      data: {
        name: namespace.name,
        systemId,
      },
      method: this.methodPost,
      url: `${this.baseUrl}/create`,
    };
    return makeRequest(req);
  }

  Delete(namespaceId: number, systemId: number): Promise<void> {
    const req: Request = {
      data: {
        namespaceId,
        systemId,
      },
      method: this.methodPost,
      url: `${this.baseUrl}/delete`,
    };
    return makeRequest(req);
  }

  GetListBySystemId(id: number): Promise<Namespace[]> {
    const req: Request = {
      data: { id },
      method: this.methodPost,
      url: `${this.baseUrl}/list`,
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

  Update(namespace: Namespace): Promise<Namespace> {
    const req: Request = {
      data: namespace,
      method: this.methodPost,
      url: `${this.baseUrl}/update`,
    };
    return makeRequest(req);
  }
}
