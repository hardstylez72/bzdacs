import { System } from '@/system/entity';
import { Namespace } from '@/namespace/entity';

export type QueryParams = NamespaceQuery | SystemQuery
// eslint-disable-next-line import/prefer-default-export
export const formQueryParams = (system: System, namespace: Namespace): QueryParams => ({
  systemId: system.id,
  systemName: system.name,
  namespaceId: namespace.id,
  namespaceName: namespace.name,
});

export interface NamespaceQuery {
  namespaceName?: string| null;
  namespaceId?: number;
}
export interface SystemQuery {
  systemName?: string | undefined;
  systemId?: number| undefined;
}
