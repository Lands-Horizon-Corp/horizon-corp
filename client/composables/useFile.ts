// repositories/FileRepository.ts
import axios, { type AxiosInstance } from 'axios';
import type { FileDetails, IFileRepository, PresignedURL, UploadProgress } from '~/types';

export function useFile(endpoint: string): IFileRepository {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.api;

  const axiosInstance: AxiosInstance = axios.create({
    baseURL: `${apiBaseUrl}/${endpoint}`,
  });

  async function uploadFile(file: File): Promise<FileDetails> {
    const formData = new FormData();
    formData.append('file', file);
    const response = await axiosInstance.post<FileDetails>(`/api/v1/file/upload`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    return response.data;
  }

  async function uploadFileWithProgress(
    file: File,
    onProgress: (progress: UploadProgress) => void
  ): Promise<FileDetails> {
    const formData = new FormData();
    formData.append('file', file);

    const response = await axiosInstance.post<FileDetails>('/api/v1/file/upload-progress', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
      onUploadProgress: (event) => {
        if (event.lengthComputable && event.total) {
          const progressBytes = event.loaded;
          const progressPercentage = (progressBytes / event.total) * 100;

          const progressData: UploadProgress = {
            file_name: file.name,
            file_size: event.total, // Using event.total for file_size
            progress_bytes: progressBytes,
            progress: progressPercentage,
          };

          onProgress(progressData);
        } else {
          // Handle cases where event.total is undefined
          const progressData: UploadProgress = {
            file_name: file.name,
            file_size: file.size, // Fallback to file.size
            progress_bytes: event.loaded,
            progress: 0, // Progress cannot be calculated
          };

          onProgress(progressData);
        }
      },
    });

    return response.data;
  }

  async function deleteFile(key: string): Promise<void> {
    await axiosInstance.delete(`/api/v1/file/delete/${key}`);
  }

  async function generatePresignedURL(key: string): Promise<PresignedURL> {
    const response = await axiosInstance.get(`/api/v1/file/presigned-url/${key}`);
    return response.data;
  }

  return {
    uploadFile,
    uploadFileWithProgress,
    deleteFile,
    generatePresignedURL,
  }

}
