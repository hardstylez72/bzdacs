import { groupGetResponse } from '@/group/generated';

export interface Group {
  code?: string;
  createdAt?: string;
  deletedAt?: string;
  description?: string;
  id: number;
  namespaceId: number;
  updatedAt?: string;
}
