import { toast } from 'sonner'
import { useState } from 'react'
import { useMutation } from '@tanstack/react-query'

import { Button } from '@/components/ui/button'
import { AdjustIcon } from '@/components/icons'
import UserAvatar from '@/components/user-avatar'
import { Progress } from '@/components/ui/progress'
import PictureCrop from '@/components/picture-crop'
import ActionTooltip from '@/components/action-tooltip'
import ProfileUploadOption from './profile-upload-option'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { withCatchAsync } from '@/lib'
import { serverRequestErrExtractor } from '@/helpers'
import { MediaResource, UserData } from '@/horizon-corp/types'
import { base64ImagetoFile } from '@/helpers/picture-crop-helper'
import MediaService from '@/horizon-corp/server/common/MediaService'
import ProfileService from '@/horizon-corp/server/auth/ProfileService'

interface IProfileUploadProps {
    currentUser: UserData
    onUploadComplete: (newUserData: UserData) => void
}

const ProfileUpload = ({
    currentUser,
    onUploadComplete,
}: IProfileUploadProps) => {
    const [reAdjust, setReAdjust] = useState(false)
    const [newImage, setNewImage] = useState<string | null>(null)
    const [croppedImage, setCroppedImage] = useState<string | null>(null)
    const [uploadMediaProgress, setUploadMediaProgress] = useState<number>(0)

    const {
        data: newUserData,
        isPending: isUpdatingUserProfile,
        mutate: updateUserProfile,
    } = useMutation<UserData, string, MediaResource>({
        mutationKey: ['update-user-profile-photo'],
        mutationFn: async (data) => {
            const [error, response] = await withCatchAsync(
                ProfileService.ProfilePicture(data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            onUploadComplete(response.data)

            return response.data
        },
    })

    const {
        data: uploadedPhoto,
        isPending: isUploadingPhoto,
        mutate: uploadPhoto,
    } = useMutation<MediaResource>({
        mutationKey: ['upload-media-photo'],
        mutationFn: async () => {
            setUploadMediaProgress(0)

            if (!croppedImage) {
                toast.warning('Nothing to upload')
                throw 'Nothing to upload'
            }

            const fileImage = base64ImagetoFile(
                croppedImage,
                `${currentUser.id}-profile.jpg`
            )

            if (!fileImage) {
                toast.error('Sorry failed to process the cropped image')
                throw 'Failed to convert image into file'
            }

            const [error, data] = await withCatchAsync(
                MediaService.upload(fileImage, (progressEvent) => {
                    if (!progressEvent.total) return

                    const progress = Math.round(
                        (progressEvent.loaded * 100) / progressEvent.total
                    )

                    setUploadMediaProgress(progress)
                })
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            updateUserProfile(data)
            return data
        },
    })

    return (
        <div className="space-y-4">
            {newImage === null && (
                <ProfileUploadOption
                    onPhotoChoose={(base64Image) => setNewImage(base64Image)}
                />
            )}
            {newImage !== null && (!croppedImage || reAdjust) ? (
                <PictureCrop
                    image={newImage}
                    onCrop={(result) => {
                        setReAdjust(false)
                        setCroppedImage(result)
                    }}
                    onCancel={() => {
                        if (croppedImage) {
                            setReAdjust(false)
                        } else {
                            setNewImage(null)
                        }
                        setReAdjust(false)
                    }}
                />
            ) : null}
            {croppedImage && !reAdjust && (
                <div className="space-y-4">
                    <div className="relative mx-auto size-fit">
                        <UserAvatar
                            fallback="-"
                            src={croppedImage ?? ''}
                            className="size-48"
                        />
                        <ActionTooltip
                            tooltipContent="ReAdjust Image"
                            align="center"
                            side="right"
                        >
                            <Button
                                variant="secondary"
                                onClick={() => setReAdjust(true)}
                                className="absolute bottom-2 right-2 size-fit rounded-full border border-transparent p-1 hover:border-foreground/20"
                            >
                                <AdjustIcon className="size-4 opacity-50 duration-300 ease-in-out group-hover:opacity-80" />
                            </Button>
                        </ActionTooltip>
                    </div>
                    {(isUploadingPhoto || isUpdatingUserProfile) && (
                        <>
                            <Progress
                                value={uploadMediaProgress}
                                className="h-1"
                            />
                            <div className="flex items-center justify-center gap-x-1 text-center text-xs text-foreground/60">
                                <LoadingSpinner className="size-2" />
                                {isUploadingPhoto && 'uploading picture...'}
                                {isUpdatingUserProfile &&
                                    'saving user profile...'}
                            </div>
                        </>
                    )}
                    <fieldset
                        disabled={isUploadingPhoto || isUpdatingUserProfile}
                        className="flex w-full items-center justify-center gap-x-2"
                    >
                        {!uploadedPhoto && !isUploadingPhoto && (
                            <Button
                                variant="secondary"
                                onClick={() => {
                                    setNewImage(null)
                                    setCroppedImage(null)
                                }}
                            >
                                Replace
                            </Button>
                        )}
                        {!newUserData &&
                            !isUpdatingUserProfile &&
                            !isUploadingPhoto && (
                                <Button
                                    onClick={() =>
                                        uploadedPhoto
                                            ? updateUserProfile(uploadedPhoto)
                                            : uploadPhoto()
                                    }
                                >
                                    {!uploadedPhoto &&
                                        !isUpdatingUserProfile &&
                                        'Upload & Save'}
                                    {!isUpdatingUserProfile &&
                                        uploadedPhoto &&
                                        'Retry'}
                                </Button>
                            )}
                    </fieldset>
                </div>
            )}
        </div>
    )
}

export default ProfileUpload
