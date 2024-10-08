import { AxiosProgressEvent } from 'axios'
import UseServer from '../../request/server'
import { MediaResource } from '../../types'

export default class MediaService {
  private static readonly BASE_ENDPOINT = '/media'

  /**
   * Uploads a new media file with progress tracking.
   *
   * @param {FormData} formData - The FormData object containing the file to upload.
   * @param {(progressEvent: AxiosProgressEvent) => void} onProgress - The callback function for tracking progress.
   * @returns {Promise<MediaResource>} - A promise that resolves to the uploaded media resource.
   */
  public static async upload(
    file: File,
    onProgress?: (progressEvent: AxiosProgressEvent) => void
  ): Promise<MediaResource> {
    const formData = new FormData()
    formData.append('file', file)
    const response = await UseServer.uploadFile<MediaResource>(
      `${MediaService.BASE_ENDPOINT}`,
      formData,
      {},
      {
        onUploadProgress: onProgress
      }
    )
    return response.data
  }

  /**
   * Deletes a media file by ID.
   *
   * @param {string} id - The ID of the media file to delete.
   * @returns {Promise<void>} - A promise that resolves when the media file is deleted.
   */
  public static async delete(id: number): Promise<void> {
    await UseServer.delete(`${MediaService.BASE_ENDPOINT}/${id}`)
  }
}
