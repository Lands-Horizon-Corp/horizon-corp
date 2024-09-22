import { BaseModel } from './base'

export interface Media extends BaseModel {
    id: string
    url: string
    fileName: string
    fileType: string
    fileSize: number
    uploadTime: Date
    description: string
    bucketName: string
    temporaryUrl?: string
}
