abstract class Service<T> {
  abstract GetList(): Promise<T[]>

  abstract Delete(id: number): Promise<void>

  abstract Create(t: T): Promise<T>
}

export default Service;
