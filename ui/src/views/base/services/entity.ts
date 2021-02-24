export interface Entity {
  id: number;
  createdAt?: Date | string;
  updatedAt?: Date | string | null;
  deletedAt?: Date | string | null;
}

export interface SimpleEntity extends Entity {
  name: string;
}
