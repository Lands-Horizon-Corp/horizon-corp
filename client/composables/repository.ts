import type { AxiosInstance } from 'axios';
import axios from 'axios';
import { useRuntimeConfig } from '#app';

interface IRepository<T> {
  getAll(): Promise<T[]>;
  getById(id: string): Promise<T>;
  create(item: T): Promise<T>;
  update(id: string, item: T): Promise<T>;
  delete(id: string): Promise<void>;
}

export function useRepository<T>(endpoint: string): IRepository<T> {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.api;

  const axiosInstance: AxiosInstance = axios.create({
    baseURL: `${apiBaseUrl}/${endpoint}`,
  });

  async function getAll(): Promise<T[]> {
    const response = await axiosInstance.get('/');
    return response.data;
  }

  async function getById(id: string): Promise<T> {
    const response = await axiosInstance.get(`/${id}`);
    return response.data;
  }

  async function create(item: T): Promise<T> {
    const response = await axiosInstance.post('/', item);
    return response.data;
  }

  async function update(id: string, item: T): Promise<T> {
    const response = await axiosInstance.put(`/${id}`, item);
    return response.data;
  }

  async function deleteById(id: string): Promise<void> {
    await axiosInstance.delete(`/${id}`);
  }

  return {
    getAll,
    getById,
    create,
    update,
    delete: deleteById,
  };
}
