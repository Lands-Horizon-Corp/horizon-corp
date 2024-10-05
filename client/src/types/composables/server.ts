import { AxiosRequestConfig } from 'axios'

export interface RequestParams {
    [key: string]: unknown
}

export interface CustomAxiosRequestConfig extends AxiosRequestConfig {
    retryCount?: number
}
