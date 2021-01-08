export interface Entity {
  id: number;
  createdAt?: Date;
  updatedAt?: Date | null;
  deletedAt?: Date | null;
}
