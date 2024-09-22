import { Filter } from './filter';
import { Pagination } from './pagination';

export interface ListRequest {
  filters: Filter[];
  pagination: Pagination;
}

export interface ListResponse<T> {
  data: T[];
  pagination: Pagination;
}

export abstract class ModelRepository<T> {
  abstract list(req: ListRequest, eagerLoads?: string[]): Promise<ListResponse<T>>;
  abstract findAll(eagerLoads?: string[]): Promise<T[]>;
  abstract findById(id: string, eagerLoads?: string[]): Promise<T | null>;
  abstract create(entity: T): Promise<void>;
  abstract update(entity: T): Promise<void>;
  abstract delete(entity: T): Promise<void>;
}