import { Method } from 'axios';

import { Group } from '@/group/entity';
import { Route } from '@/route/entity';
import { makeRequest, Request } from '../../app/util/requester';

export interface UserRoute {
  routeId: number;
  userId: number;
}

export interface UpdateUserRoute {
  routeId: number;
  userId: number;
  isExcluded: boolean;
}
export interface CreateUserRoute {
  routeId: number;
  userId: number;
  isExcluded: boolean;
}
interface Groups {
  groups: Group[];
}

export interface RouteExt extends Route {
  isExcluded: boolean;
  isOverwritten: boolean;
  isIndependent: boolean;
}

export type RouteWithGroups = RouteExt & Groups
