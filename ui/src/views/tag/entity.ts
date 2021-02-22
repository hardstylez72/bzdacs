import { tag_getResponse } from '@/views/tag/generated';

export type TagSwg = tag_getResponse

export type Tag = {
  createdAt?: string;
  deletedAt?: string;
  id: number;
  name: string;
  updatedAt?: string;
}
