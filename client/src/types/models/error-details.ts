export interface ErrorDetails {
  message: string;
  name: string;
  stack?: string;
  response?: unknown;
  status?: number;
}