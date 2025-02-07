import { TEntityId } from './common'

export interface IMediaResource {
    id: TEntityId
    fileName: string
    fileSize: number
    fileType: string
    storageKey: string
    url: string
    bucketName: string
    createdAt: string
    updatedAt: string
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
