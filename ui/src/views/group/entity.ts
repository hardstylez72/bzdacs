import { group_getResponse } from '@/views/group/generated';

export type GroupSwg = group_getResponse

export interface Group {
  code: string;
  createdAt: string;
  deletedAt?: string;
  description: string;
  id: number;
  namespaceId: number;
  updatedAt: string;
}
