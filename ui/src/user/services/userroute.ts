import { Group } from '@/group/entity';
import { Route } from '@/route/entity';

export interface UpdateUserRoute {
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
