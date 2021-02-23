import { Namespace } from '@/views/namespace/entity';
import { System } from '@/views/system/entity';
import { NamespaceQuery } from '@/views/tree-menu/helper';

export type NodeType =
  'routes' |
  'groups' |
  'system' |
  'tags' |
  'namespace' |
  'users' |
  'createNewNamespaceBtn' |
  'createNewSystemBtn'

export interface Node {
  id?: string;
  name?: string;
  type: NodeType;
  color?: string;
  children?: Node[];
  system?: System;
  namespace?: Namespace;
  testId?: string;
}
