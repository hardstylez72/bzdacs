import { Namespace } from '@/namespace/entity';
import { System } from '@/system/entity';
import { NamespaceQuery } from '@/tree-menu/helper';

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
