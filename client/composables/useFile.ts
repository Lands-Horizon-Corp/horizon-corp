// repositories/FileRepository.ts
import axios, { type AxiosInstance } from 'axios';
import type { FileDetails, IFileRepository, PresignedURL, UploadProgress } from '~/types';

export function useFile(endpoint?: string): IFileRepository {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.api;

  const axiosInstance: AxiosInstance = axios.create({
    baseURL: endpoint ? `${apiBaseUrl}/${endpoint}` : apiBaseUrl,
  });

  async function uploadFile(file: File): Promise<FileDetails> {
    const formData = new FormData();
    formData.append('file', file);
    const response = await axiosInstance.post<FileDetails>(`/file/upload`, formData, {
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

    try {
      const response = await axiosInstance.post<FileDetails>('/file/upload-progress', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
        onUploadProgress: (event) => {
          if (event.lengthComputable && event.total) {
            const progressBytes = event.loaded;
            const progressPercentage = (progressBytes / event.total) * 100;

            const progressData: UploadProgress = {
              file_name: file.name,
              file_format: file.type,
              file_size: event.total,
              progress_bytes: progressBytes,
              progress: progressPercentage,
            };

            onProgress(progressData);
          } else if (event.total) {
            const progressData: UploadProgress = {
              file_name: file.name,
              file_format: file.type,
              file_size: event.total,
              progress_bytes: event.loaded,
              progress: 0,
            };
            onProgress(progressData);
          }
        },
        timeout: 0,
      });

      return response.data;
    } catch (error) {
      console.error("Failed to upload file:", error);
      throw error; // Re-throw the error after logging or handling it
    }
  }

  async function deleteFile(key: string): Promise<void> {
    await axiosInstance.delete(`/file/delete/${key}`);
  }

  async function generatePresignedURL(key: string): Promise<PresignedURL> {
    const response = await axiosInstance.get(`/file/presigned-url/${key}`);
    return response.data;
  }

  return {
    uploadFile,
    uploadFileWithProgress,
    deleteFile,
    generatePresignedURL,
  }

}
