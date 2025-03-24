import { ITimeStamps, TEntityId } from './common'

export interface IMediaResource extends ITimeStamps {
    id: TEntityId
    fileName: string
    fileSize: number
    fileType: string
    storageKey: string
    url: string
    bucketName: string
    downloadURL: string
}

export interface IMediaRequest {
    id?: TEntityId
    fileName: string
    fileSize: number
    fileType: string
    storageKey: string
    key?: string
    url?: string
    bucketName?: string
}
