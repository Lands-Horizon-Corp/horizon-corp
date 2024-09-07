import type { CancelTokenSource } from "axios";

export interface FileType {
  id: number;
  fileName: string;
  fileSize: number;
  fileType: string;
  storageKey: string;
  url: string;
  bucketName: string;
}

export interface UploadProgress {
  loaded: number;
  total: number;
  percentage: number;
  elapsedTime: number;
  remainingTime: number;
}


export interface IFileRepository {
  uploadFile(file: File): Promise<FileType>;
  deleteFile(fileId: number): Promise<void>;
  downloadFile(fileId: number): Promise<string>;
  getFile(fileId: number): Promise<FileType>;
  uploadFileWithProgress(
    file: File,
    onProgress: (progress: UploadProgress) => void,
    cancelToken: CancelTokenSource,
    onCancel?: () => void
  ): Promise<FileType>;
}
