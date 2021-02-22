import { system_getResponse } from '@/views/system/generated';

export type SystemSwg = system_getResponse

export type System = {
  createdAt: string;
  deletedAt?: string;
  id: number;
  name: string;
  updatedAt: string;
}
