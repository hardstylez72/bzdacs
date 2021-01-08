import { Entity } from '@/views/base/services/entity';
import { makeRequest, Request } from '@/views/base/services/utils/requester';
import DefaultService from '../base/services/default';

export interface Route extends Entity {
  id: number;
  route: string;
  method: string;
  description: string;
  tags?: string[];
}

export interface Filter {
  tags: {
    names: string[];
    exclude: boolean;
  };
}

export default class RouteService extends DefaultService<Route> {
  Update(t: Route): Promise<Route> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/update`,
    };
    return makeRequest(req);
  }

  GetById(id: number): Promise<Route> {
    const req: Request = {
      data: { id },
      method: this.methodPost,
      url: `${this.baseUrl}/get`,
    };
    return makeRequest(req);
  }

  GetListV2(filter: Filter): Promise<Route[]> {
    const req: Request = {
      data: { filter },
      method: this.methodPost,
      url: `${this.baseUrl}/list`,
    };
    return makeRequest(req);
  }
}
