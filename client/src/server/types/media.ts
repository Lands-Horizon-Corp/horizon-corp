export interface IMediaResource {
    id: number
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
    id?: number
    fileName: string
    fileSize: number
    fileType: string
    storageKey: string
    key?: string
    url?: string
    bucketName?: string
}
