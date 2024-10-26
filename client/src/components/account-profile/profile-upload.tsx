import { useState } from 'react'

import { Button } from '@/components/ui/button'
import { AdjustIcon } from '@/components/icons'
import UserAvatar from '@/components/user-avatar'
import PictureCrop from '@/components/picture-crop'
import ActionTooltip from '@/components/action-tooltip'

import FileDropUpload from './picture-drop'
import { ICroppedImageResult } from '../picture-crop/picture-crop-utils'

const ProfileUpload = ({
    currentUserPhotoUrl,
}: {
    currentUserPhotoUrl: string | null
}) => {
    const [reAdjust, setReAdjust] = useState(false)
    const [newImage, setNewImage] = useState<string | null>(null)
    const [croppedImage, setCroppedImage] =
        useState<ICroppedImageResult | null>(null)

    const onFileSelect = (files: FileList) => {
        if (files && files.length > 0) {
            const reader = new FileReader()
            reader.addEventListener('load', () => {
                const newImgUrl = reader.result?.toString() ?? ''
                setNewImage(newImgUrl)
            })
            reader.readAsDataURL(files?.[0])
        }
    }

    return (
        <div className="space-y-4">
            {newImage === null && (
                <FileDropUpload onFileSelect={onFileSelect}>
                    <UserAvatar
                        fallback="-"
                        src={currentUserPhotoUrl ?? ''}
                        className="mx-auto size-24"
                    />
                </FileDropUpload>
            )}
            {newImage !== null && (!croppedImage || reAdjust) ? (
                <PictureCrop
                    image={newImage}
                    onCrop={(result) => {
                        setReAdjust(false)
                        setCroppedImage(result)
                    }}
                    onCancel={() => setNewImage(null)}
                />
            ) : null}
            {
                // TODO : Upload Button & Upload Progress
                croppedImage && !reAdjust && (
                    <div className="space-y-4">
                        <div className="relative mx-auto size-fit">
                            <UserAvatar
                                fallback="-"
                                src={croppedImage.base64Url ?? ''}
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
                        <div className="flex gap-x-2 justify-center items-center w-full">
                            <Button
                                variant="secondary"
                                onClick={() => {
                                    setNewImage(null)
                                    setCroppedImage(null)
                                }}
                            >
                                Replace
                            </Button>
                            <Button>Upload & Save</Button>
                        </div>
                    </div>
                )
            }
        </div>
    )
}

export default ProfileUpload
