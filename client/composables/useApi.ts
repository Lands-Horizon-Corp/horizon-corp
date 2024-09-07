import type { AxiosInstance } from "axios";
import axios from "axios";
import { useRuntimeConfig } from "#app";

export function useApi<T = unknown>(endpoint?: string) {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.api;

  const axiosInstance: AxiosInstance = axios.create({
    baseURL: `${apiBaseUrl}/${endpoint}`,
  });

  async function get<U = T>(route?: string): Promise<U> {
    const value = route ?? "/";
    const response = await axiosInstance.get(value);
    return response.data;
  }

  async function getAll<U = T>(route?: string): Promise<U[]> {
    const value = route ?? "/";
    const response = await axiosInstance.get(value);
    return response.data;
  }

  async function getById<U = T>(id: string, route?: string): Promise<U> {
    const value = route ? `${route}/${id}` : `/${id}`;
    const response = await axiosInstance.get(value);
    return response.data;
  }

  async function create<U = T, V = U>(item: U, route?: string): Promise<V> {
    const value = route ?? "/";
    const response = await axiosInstance.post<V>(value, item);
    return response.data;
  }

  async function update<U = T>(
    id: string,
    item: U,
    route?: string,
  ): Promise<U> {
    const value = route ? `${route}/${id}` : `/${id}`;
    const response = await axiosInstance.put(value, item);
    return response.data;
  }

  async function deleteById(id: string, route?: string): Promise<void> {
    const value = route ? `${route}/${id}` : `/${id}`;
    await axiosInstance.delete(value);
  }

  return { get, getAll, getById, create, update, delete: deleteById };
}
