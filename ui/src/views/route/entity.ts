import { route_getResponse } from '@/views/route/generated';

export type RouteSwg = route_getResponse

export type Route = {
  createdAt: string;
  deletedAt?: string;
  description: string;
  id: number;
  method: string;
  namespaceId: number;
  route: string;
  tags: string[];
  updatedAt: string;
}
