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

    const fileSize = file.size;
    const fileName = file.name;

    const source = new EventSource('/api/v1/file/upload-progress');

    source.onmessage = (event) => {
      if (event.data === 'done') {
        source.close();
      } else {
        const progressBytes = parseInt(event.data, 10);
        const percentage = (progressBytes / fileSize) * 100;
        onProgress({ fileName, fileSize, progressBytes, percentage });
      }
    };

    const response = await axiosInstance.post<FileDetails>('/api/v1/file/upload-progress', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    source.close();
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
