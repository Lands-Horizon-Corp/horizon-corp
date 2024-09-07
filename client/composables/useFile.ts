import axios from "axios";
import type { CancelTokenSource, AxiosInstance } from "axios";
import type { IFileRepository, FileType, UploadProgress } from "~/types";

export class FileRepository implements IFileRepository {
  private axiosInstance: AxiosInstance;

  constructor(apiBaseUrl: string) {
    this.axiosInstance = axios.create({
      baseURL: apiBaseUrl,
    });
  }

  async uploadFileWithProgress(
    file: File,
    onProgress: (progress: UploadProgress) => void,
    cancelToken: CancelTokenSource,
    onCancel?: () => void,
  ): Promise<FileType> {
    const formData = new FormData();
    formData.append("file", file);
    const startTime = Date.now();

    try {
      const response = await this.axiosInstance.post<FileType>(
        "/file/upload",
        formData,
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
          onUploadProgress: (event) => {
            const total = event?.total ?? 0;
            const elapsedTime = (Date.now() - startTime) / 1000;
            const speed = event.loaded / elapsedTime;
            const remainingTime = (total - event.loaded) / speed;

            const progress: UploadProgress = {
              loaded: event.loaded,
              total: total,
              percentage: Math.round((event.loaded * 100) / total),
              elapsedTime,
              remainingTime: Math.max(0, remainingTime),
            };
            onProgress(progress);
          },
          cancelToken: cancelToken.token,
        },
      );
      return response.data;
    } catch (error) {
      if (axios.isCancel(error) && onCancel) {
        onCancel();
      }
      throw error;
    }
  }

  async uploadFile(file: File): Promise<FileType> {
    const formData = new FormData();
    formData.append("file", file);
    const response = await this.axiosInstance.post<FileType>(
      "/files/upload",
      formData,
    );
    return response.data;
  }

  async deleteFile(fileId: number): Promise<void> {
    await this.axiosInstance.delete(`/files/${fileId}`);
  }

  async downloadFile(fileId: number): Promise<string> {
    const response = await this.axiosInstance.get<string>(
      `/files/${fileId}/download`,
    );
    return response.data;
  }

  async getFile(fileId: number): Promise<FileType> {
    const response = await this.axiosInstance.get<FileType>(`/files/${fileId}`);
    return response.data;
  }
}
