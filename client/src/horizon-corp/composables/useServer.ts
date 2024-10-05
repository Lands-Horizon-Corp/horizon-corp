import { CustomAxiosRequestConfig, RequestParams } from '@/types/composables/server';
import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, AxiosError, InternalAxiosRequestConfig } from 'axios';
import { errorMessages } from './constants';
import { ErrorDetails } from '../types/error-details';


type ErrorLogger = (errorDetails: ErrorDetails) => Promise<void>;
type RetryCondition = (error: AxiosError, retryCount: number) => boolean;
type ErrorEnhancer = (error: AxiosError) => Error;


/**
 * Class to handle HTTP requests and errors using Axios with customizable options.
 */
export default class UseServer {
  private httpClient: AxiosInstance;
  private maxRetryCount: number;
  private errorUrl: string;
  private errorLogger: ErrorLogger;
  private retryCondition: RetryCondition;
  private errorEnhancer: ErrorEnhancer;


  /**
   * Initializes a new instance of UseServer class.
   * 
   * @param {string} [baseUrl=import.meta.env.VITE_CLIENT_SERVER_URL ?? 'http://localhost:8080/api/v1'] - The base URL for the Axios instance.
   * @param {AxiosRequestConfig} [defaultConfig] - The default Axios configuration.
   * @param {number} [maxRetryCount=3] - Maximum retry count for failed requests.
   * @param {string} [errorUrl='/error-handler'] - URL for logging errors.
   * @param {ErrorLogger} [errorLogger] - Custom error logger function.
   * @param {RetryCondition} [retryCondition] - Custom retry condition function.
   * @param {ErrorEnhancer} [errorEnhancer] - Custom error enhancer function.
   */
  constructor(
    baseUrl: string = import.meta.env.VITE_CLIENT_SERVER_URL ?? 'http://localhost:8080/api/v1',
    defaultConfig?: AxiosRequestConfig,
    maxRetryCount: number = 3,
    errorUrl: string = '/error-handler',
    errorLogger?: ErrorLogger,
    retryCondition?: RetryCondition,
    errorEnhancer?: ErrorEnhancer
  ) {
    this.httpClient = axios.create({
      baseURL: baseUrl,
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

  /**
   * Handles Axios errors, retrying requests if necessary and logging them.
   * 
   * @private
   * @param {AxiosError} error - The Axios error object.
   * @returns {Promise<AxiosResponse<unknown>>} - The Axios response after handling the error.
   */
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

  /**
  * Formats error details for logging purposes.
  * 
  * @private
  * @param {AxiosError} error - The Axios error object.
  * @returns {ErrorDetails} - The formatted error details.
  */
  private formatErrorDetails(error: AxiosError): ErrorDetails {
    return {
      message: error.message,
      name: error.name,
      stack: error.stack,
      response: error.response?.data || null,
      status: error.response?.status,
    };
  }

  /**
   * Default error logger that logs error details to the backend.
   * 
   * @private
   * @param {ErrorDetails} errorDetails - The details of the error to log.
   * @returns {Promise<void>} - A promise that resolves once the error is logged.
  */
  private async defaultErrorLogger(errorDetails: ErrorDetails): Promise<void> {
    try {
      // Use a separate Axios instance to avoid interceptors and baseURL
      const loggingClient = axios.create();
      await loggingClient.post(this.errorUrl, errorDetails);
    } catch (logError) {
      console.error('Failed to log error to backend:', logError);
    }
  }

  /**
   * Default retry condition to determine if a request should be retried.
   * 
   * @private
   * @param {AxiosError} error - The Axios error object.
   * @param {number} retryCount - The current retry count.
   * @returns {boolean} - True if the request should be retried, false otherwise.
   */
  private defaultRetryCondition(error: AxiosError, retryCount: number): boolean {
    return (
      retryCount < this.maxRetryCount &&
      (!error.response || error.response.status >= 500)
    );
  }

  /**
   * Default error enhancer to customize error messages.
   * 
   * @private
   * @param {AxiosError} error - The Axios error object.
   * @returns {Error} - The enhanced error object.
   */
  private defaultErrorEnhancer(error: AxiosError): Error {
    if (error.response?.status) {
      const statusMessage = errorMessages[error.response.status];
      error.message = statusMessage
        ? `${statusMessage}: ${error.message}`
        : error.message;
    }
    return error;
  }

  /**
   * Returns the internal Axios instance for customizations.
   * 
   * @returns {AxiosInstance} - The Axios instance.
   */
  getHttpClient(): AxiosInstance {
    return this.httpClient;
  }

  /**
   * Adds a request interceptor to the Axios instance.
   * 
   * @param {Function} [onFulfilled] - The function to handle fulfilled requests.
   * @param {Function} [onRejected] - The function to handle rejected requests.
   * @returns {number} - The interceptor ID.
  */
  addRequestInterceptor(
    onFulfilled?: (
      value: InternalAxiosRequestConfig
    ) => InternalAxiosRequestConfig | Promise<InternalAxiosRequestConfig>,
    onRejected?: (error: unknown) => Promise<never> | void
  ): number {
    return this.httpClient.interceptors.request.use(onFulfilled, onRejected);
  }

  /**
   * Adds a response interceptor to the Axios instance.
   * 
   * @param {Function} [onFulfilled] - The function to handle fulfilled responses.
   * @param {Function} [onRejected] - The function to handle rejected responses.
   * @returns {number} - The interceptor ID.
   */
  addResponseInterceptor(
    onFulfilled?: (
      value: AxiosResponse
    ) => AxiosResponse | Promise<AxiosResponse>,
    onRejected?: (error: unknown) => Promise<never> | void
  ): number {
    return this.httpClient.interceptors.response.use(onFulfilled, onRejected);
  }

  /**
   * Makes a GET request.
   * 
   * @param {string} url - The request URL.
   * @param {RequestParams} [params] - The query parameters for the request.
   * @param {AxiosRequestConfig} [config] - Optional Axios configuration.
   * @returns {Promise<AxiosResponse<R>>} - The Axios response.
   */
  async get<R = unknown>(
    url: string,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return this.httpClient.get<R>(url, { params, ...config });
  }

  /**
   * Makes a POST request.
   * 
   * @param {string} url - The request URL.
   * @param {D} [data] - The data to send with the request.
   * @param {RequestParams} [params] - The query parameters for the request.
   * @param {AxiosRequestConfig} [config] - Optional Axios configuration.
   * @returns {Promise<AxiosResponse<R>>} - The Axios response.
   */
  async post<D = unknown, R = unknown>(
    url: string,
    data?: D,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return this.httpClient.post<R>(url, data, { params, ...config });
  }

  /**
   * Makes a PUT request.
   * 
   * @param {string} url - The request URL.
   * @param {D} [data] - The data to send with the request.
   * @param {RequestParams} [params] - The query parameters for the request.
   * @param {AxiosRequestConfig} [config] - Optional Axios configuration.
   * @returns {Promise<AxiosResponse<R>>} - The Axios response.
   */
  async put<D = unknown, R = unknown>(
    url: string,
    data?: D,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return this.httpClient.put<R>(url, data, { params, ...config });
  }

  /**
   * Makes a DELETE request.
   * 
   * @param {string} url - The request URL.
   * @param {RequestParams} [params] - The query parameters for the request.
   * @param {AxiosRequestConfig} [config] - Optional Axios configuration.
   * @returns {Promise<AxiosResponse<R>>} - The Axios response.
   */
  async delete<R = unknown>(
    url: string,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return this.httpClient.delete<R>(url, { params, ...config });
  }
}
