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

export interface Tree {
  nodes: Node[];
  children: Node[];
}
export interface Node {
  id?: number; // todo: сдеать обязательным
  name: string;
  type: NodeType;
  color?: string;
  system?: string | null;
  namespace?: string | null;
  children?: Node[];
  origin?: System | Namespace;
}
