import { MediaResource } from '..'

export interface TimeInRequest {
    timeIn: Date
    mediaIn: MediaResource
}

export interface TimeOutRequest {
    timeOut: Date
    mediaOut: MediaResource
}
