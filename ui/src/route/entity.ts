import { routeGetResponse } from '@/route/generated';

export type RouteSwg = routeGetResponse

export type Route = {
  createdAt?: string;
  deletedAt?: string;
  description?: string;
  id: number;
  method?: string;
  namespaceId?: number;
  route?: string;
  tags?: string[];
  updatedAt?: string;
}
