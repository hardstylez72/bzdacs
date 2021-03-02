import { userGetResponse } from '@/user/generated';

export interface User {
  /**  */
  createdAt?: string;

  /**  */
  deletedAt?: string;

  /**  */
  externalId?: string;

  /**  */
  id: number;

  /**  */
  namespaceId?: number;

  /**  */
  updatedAt?: string;
}
