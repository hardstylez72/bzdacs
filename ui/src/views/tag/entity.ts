import { tagGetResponse } from '@/views/tag/generated';

export type TagSwg = tagGetResponse

export type Tag = {
  createdAt?: string;
  deletedAt?: string;
  id: number;
  name: string;
  updatedAt?: string;
}
