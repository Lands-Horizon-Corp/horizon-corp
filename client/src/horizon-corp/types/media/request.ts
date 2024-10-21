export interface MediaRequest {
    id?: number
    fileName: string
    fileSize: number
    fileType: string
    storageKey: string
    key?: string
    url?: string
    bucketName?: string
}
