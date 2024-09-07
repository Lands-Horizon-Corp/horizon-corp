export interface IApiRepository<T = unknown> {
  get<U = T>(route?: string): Promise<U>
  getAll<U = T>(route?: string): Promise<U[]>;
  getById<U = T>(id: string, route?: string): Promise<U>;
  create<U = T, V = U>(item: U, route?: string): Promise<V>;
  update<U = T>(id: string, item: U, route?: string): Promise<U>;
  delete(id: string, route?: string): Promise<void>;
}
