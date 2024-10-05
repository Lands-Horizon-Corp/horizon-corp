import { AxiosRequestConfig } from "axios";

export interface RequestParams {
  [key: string]: unknown;
}

export interface CustomAxiosRequestConfig extends AxiosRequestConfig {
  retryCount?: number;
}

export interface ErrorDetails {
  message: string;
  name: string;
  stack?: string;
  response?: unknown;
  status?: number;
}
