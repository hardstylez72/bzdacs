import { Entity } from '@/views/base/services/entity';
import DefaultService from '../base/services/default';

export interface System extends Entity {
  name: string;
}

export default class NamespaceService extends DefaultService<System> {}
