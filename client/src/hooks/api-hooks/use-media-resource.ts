import { toast } from 'sonner'
import { useMutation } from '@tanstack/react-query'

import { withCatchAsync } from '@/utils'
import { MediaResource } from '@/horizon-corp/types'
import { serverRequestErrExtractor } from '@/helpers'
import MediaService from '@/horizon-corp/server/common/MediaService'

export const useSinglePictureUpload = ({
    onUploadProgressChange,
    onUploadSuccess,
    onUploadError,
}: {
    onUploadSuccess?: (data: MediaResource) => void
    onUploadError?: (error: string) => void
    onUploadProgressChange?: (progress: number) => void
}) => {
    return useMutation<MediaResource, string, File>({
        mutationKey: ['upload-media-photo'],
        mutationFn: async (fileImage) => {
            if (!fileImage) {
                const errorMessage = 'No valid File Image to upload'
                toast.warning(errorMessage)
                throw errorMessage
            }

            onUploadProgressChange?.(0)

            const [error, data] = await withCatchAsync(
                MediaService.upload(fileImage, (progressEvent) => {
                    if (!progressEvent.total) return

                    const progress = Math.round(
                        (progressEvent.loaded * 100) / progressEvent.total
                    )

                    onUploadProgressChange?.(progress)
                })
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onUploadError?.(errorMessage)
                throw errorMessage
            }

            onUploadSuccess?.(data)
            return data
        },
    })
}
