import { toast } from 'sonner'
import { useMutation } from '@tanstack/react-query'

import { withCatchAsync } from '@/utils'
import { IOperationCallbacks } from './types'
import { MediaResource } from '@/horizon-corp/types'
import { serverRequestErrExtractor } from '@/helpers'
import MediaService from '@/horizon-corp/server/common/MediaService'

export const useSinglePictureUpload = ({
    onUploadProgressChange,
    onSuccess,
    onError,
}: {
    onUploadProgressChange?: (progress: number) => void
} & IOperationCallbacks<MediaResource, string>) => {
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
                onError?.(errorMessage)
                throw errorMessage
            }

            onSuccess?.(data)
            return data
        },
    })
}