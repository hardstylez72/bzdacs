import { systemGetResponse } from '@/system/generated';

export type SystemSwg = systemGetResponse

export type System = {
  createdAt: string;
  deletedAt?: string;
  id: number;
  name: string;
  updatedAt: string;
}
