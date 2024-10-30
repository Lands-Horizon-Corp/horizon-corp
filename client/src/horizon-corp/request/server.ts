import { RequestParams } from '../types'
import axios, {
  AxiosInstance,
  AxiosRequestConfig,
  AxiosResponse,
  InternalAxiosRequestConfig,
} from 'axios'



/**
 * Class to handle HTTP requests using Axios with static methods.
 */
export default class UseServer {
  private static httpClient: AxiosInstance = axios.create({
    baseURL: UseServer.getDefaultUrl(),
    headers: {
      'Content-Type': 'application/json',
    },
    withCredentials: true,
  });

  /**
   * Retrieves the API key from environment variables.
   *
   * @returns {string | undefined} - The Horizon corp API key.
   */
  // private static getApiKey(): string | undefined {
  //     return (
  //         import.meta.env.VITE_HORIZON_CORP_API_KEY ||
  //         process.env.HORIZON_CORP_API_KEY
  //     ) // Support for various environments
  // }

  /**
   * Retrieves the default base URL based on the environment.
   *
   * @returns {string} - The default base URL.
   */
  private static getDefaultUrl(): string {
    const environment =
      import.meta.env.VITE_HORIZON_CORP_ENVIRONMENT ||
      process.env.HORIZON_CORP_ENVIRONMENT
    switch (environment) {
      case 'local':
      case 'development':
      case 'staging':
      case 'production':
        return 'http://localhost:8080/api/v1/'
      default:
        return 'http://localhost:8080/api/v1/'
    }
  }

  /**
   * Returns the internal Axios instance for customizations.
   *
   * @returns {AxiosInstance} - The Axios instance.
   */
  public static getHttpClient(): AxiosInstance {
    return UseServer.httpClient
  }

  /**
   * Adds a request interceptor to the Axios instance.
   *
   * @param {Function} [onFulfilled] - The function to handle fulfilled requests.
   * @param {Function} [onRejected] - The function to handle rejected requests.
   * @returns {number} - The interceptor ID.
   */
  public static addRequestInterceptor(
    onFulfilled?: (
      value: InternalAxiosRequestConfig
    ) => InternalAxiosRequestConfig | Promise<InternalAxiosRequestConfig>,
    onRejected?: (error: unknown) => Promise<never> | void
  ): number {
    return UseServer.httpClient.interceptors.request.use(
      onFulfilled,
      onRejected
    )
  }

  /**
   * Adds a response interceptor to the Axios instance.
   *
   * @param {Function} [onFulfilled] - The function to handle fulfilled responses.
   * @param {Function} [onRejected] - The function to handle rejected responses.
   * @returns {number} - The interceptor ID.
   */
  public static addResponseInterceptor(
    onFulfilled?: (
      value: AxiosResponse
    ) => AxiosResponse | Promise<AxiosResponse>,
    onRejected?: (error: unknown) => Promise<never> | void
  ): number {
    return UseServer.httpClient.interceptors.response.use(
      onFulfilled,
      onRejected
    )
  }

  /**
   * Makes a GET request.
   *
   * @param {string} url - The request URL.
   * @param {RequestParams} [params] - The query parameters for the request.
   * @param {AxiosRequestConfig} [config] - Optional Axios configuration.
   * @returns {Promise<AxiosResponse<R>>} - The Axios response.
   */
  public static async get<R = unknown>(
    url: string,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return UseServer.httpClient.get<R>(url, { params, ...config })
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
  public static async post<D = unknown, R = unknown>(
    url: string,
    data?: D,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return UseServer.httpClient.post<R>(url, data, { params, ...config })
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
  public static async put<D = unknown, R = unknown>(
    url: string,
    data?: D,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return UseServer.httpClient.put<R>(url, data, { params, ...config })
  }

  /**
   * Makes a DELETE request.
   *
   * @param {string} url - The request URL.
   * @param {RequestParams} [params] - The query parameters for the request.
   * @param {AxiosRequestConfig} [config] - Optional Axios configuration.
   * @returns {Promise<AxiosResponse<R>>} - The Axios response.
   */
  public static async delete<R = unknown>(
    url: string,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return UseServer.httpClient.delete<R>(url, { params, ...config })
  }

  /**
   * Makes a POST request for file uploads.
   *
   * @param {string} url - The request URL.
   * @param {FormData} formData - The FormData object containing the file.
   * @param {RequestParams} [params] - The query parameters for the request.
   * @param {AxiosRequestConfig} [config] - Optional Axios configuration.
   * @returns {Promise<AxiosResponse<R>>} - The Axios response.
   */
  public static async uploadFile<R = unknown>(
    url: string,
    formData: FormData,
    params?: RequestParams,
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<R>> {
    return UseServer.httpClient.post<R>(url, formData, {
      params,
      headers: {
        'Content-Type': 'multipart/form-data',
      },
      ...config,
    })
  }
}
