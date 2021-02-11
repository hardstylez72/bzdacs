import { Entity } from '@/views/base/services/entity';
import { Namespace } from '@/views/namespace/service';

export type NodeType =
  'routes' |
  'groups' |
  'system' |
  'namespace' |
  'users' |
  'createNewNamespaceBtn' |
  'createNewSystemBtn'

export interface System extends Entity{
  name: string;
}

export interface Node {
  id: string;
  name: string;
  type: NodeType;
  color?: string;
  children?: Node[];
  system?: System;
  namespace?: Namespace;
  testId?: string;
}
