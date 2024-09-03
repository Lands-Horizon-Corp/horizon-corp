export interface FileDetails {
  message: string;
  file_name: string;
  file_size: number;
  file_format: string;
  file_url: string;
  temp_url?: string;
}

export interface UploadProgress {
  file_name: string;
  file_format: string;
  file_size: number;
  progress_bytes: number;
  progress: number;
}

export interface PresignedURL {
  url: string;
  file_name: string;
}

export interface IFileRepository {
  uploadFile(file: File): Promise<FileDetails>;
  uploadFileWithProgress(file: File, onProgress: (progress: UploadProgress) => void): Promise<FileDetails>;
  deleteFile(key: string): Promise<void>;
  generatePresignedURL(key: string): Promise<PresignedURL>;
}