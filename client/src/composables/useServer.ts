import { CustomAxiosRequestConfig, ErrorDetails, RequestParams } from '@/types/composables/server';
import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, AxiosError, InternalAxiosRequestConfig } from 'axios';
import { errorMessages } from './constants';

type ErrorLogger = (errorDetails: ErrorDetails) => Promise<void>;
type RetryCondition = (error: AxiosError, retryCount: number) => boolean;
type ErrorEnhancer = (error: AxiosError) => Error;

export class UseServer {
  private httpClient: AxiosInstance;
  private maxRetryCount: number;
  private errorUrl: string;
  private errorLogger: ErrorLogger;
  private retryCondition: RetryCondition;
  private errorEnhancer: ErrorEnhancer;

  constructor(
    baseURL: string = import.meta.env.VITE_CLIENT_SERVER_URL ?? 'http://localhost:8080/api/v1',
    defaultConfig?: AxiosRequestConfig,
    maxRetryCount: number = 3,
    errorUrl: string = '/error-handler',
    errorLogger?: ErrorLogger,
    retryCondition?: RetryCondition,
    errorEnhancer?: ErrorEnhancer
  ) {
    this.httpClient = axios.create({
      baseURL,
      headers: {
        'Content-Type': 'application/json',
        ...(defaultConfig?.headers || {}),
      },
      ...defaultConfig,
    });
    this.maxRetryCount = maxRetryCount;
    this.errorUrl = errorUrl;
    this.errorLogger = errorLogger || this.defaultErrorLogger.bind(this);
    this.retryCondition = retryCondition || this.defaultRetryCondition.bind(this);
    this.errorEnhancer = errorEnhancer || this.defaultErrorEnhancer.bind(this);

    this.httpClient.interceptors.response.use(
      (response: AxiosResponse) => response,
      (error: AxiosError) => this.handleError(error)
    );
  }

  // Generalized method to handle errors
  private async handleError(
    error: AxiosError
  ): Promise<AxiosResponse<unknown>> {
    const config = error.config as CustomAxiosRequestConfig;

    if (config) {
      config.retryCount = config.retryCount || 0;

      if (this.retryCondition(error, config.retryCount)) {
        config.retryCount += 1;
        return this.httpClient(config); // Retry request
      }
    }

    await this.errorLogger(this.formatErrorDetails(error));
    return Promise.reject(this.errorEnhancer(error));
  }

  // Format error details for backend logging
  private formatErrorDetails(error: AxiosError): ErrorDetails {
    return {
      message: error.message,
      name: error.name,
      stack: error.stack,
      response: error.response?.data || null,
      status: error.response?.status,
    };
  }

  // Default error logger that logs error to backend
  private async defaultErrorLogger(errorDetails: ErrorDetails): Promise<void> {
    try {
      // Use a separate Axios instance to avoid interceptors and baseURL
      const loggingClient = axios.create();
      await loggingClient.post(this.errorUrl, errorDetails);
    } catch (logError) {
      console.error('Failed to log error to backend:', logError);
    }
  }

  // Default retry condition
  private defaultRetryCondition(error: AxiosError, retryCount: number): boolean {
    return (
      retryCount < this.maxRetryCount &&
      (!error.response || error.response.status >= 500)
    );
  }

  // Default error enhancer
  private defaultErrorEnhancer(error: AxiosError): Error {
    if (error.response?.status) {
      const statusMessage = errorMessages[error.response.status];
      error.message = statusMessage
        ? `${statusMessage}: ${error.message}`
        : error.message;
    }
    return error;
  }

  // Expose the Axios instance for customizations
  getHttpClient(): AxiosInstance {
    return this.httpClient;
  }

  // Methods to add interceptors
  addRequestInterceptor(
    onFulfilled?: (
      value: InternalAxiosRequestConfig
    ) => InternalAxiosRequestConfig | Promise<InternalAxiosRequestConfig>,
    onRejected?: (error: unknown) => Promise<never> | void
  ): number {
    return this.httpClient.interceptors.request.use(onFulfilled, onRejected);
  }

  addResponseInterceptor(
    onFulfilled?: (
      value: AxiosResponse
    ) => AxiosResponse | Promise<AxiosResponse>,
    onRejected?: (error: unknown) => Promise<never> | void
  ): number {
    return this.httpClient.interceptors.response.use(onFulfilled, onRejected);
  }

  // GET request method
  async get<R = unknown>(
    url: string,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return this.httpClient.get<R>(url, { params, ...config });
  }

  // POST request method
  async post<D = unknown, R = unknown>(
    url: string,
    data?: D,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return this.httpClient.post<R>(url, data, { params, ...config });
  }

  // PUT request method
  async put<D = unknown, R = unknown>(
    url: string,
    data?: D,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return this.httpClient.put<R>(url, data, { params, ...config });
  }

  // DELETE request method
  async delete<R = unknown>(
    url: string,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return this.httpClient.delete<R>(url, { params, ...config });
  }
}