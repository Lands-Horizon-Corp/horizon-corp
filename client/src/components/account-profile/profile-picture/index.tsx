import { useState } from 'react'

import { Button } from '@/components/ui/button'
import { CameraFillIcon } from '@/components/icons'
import ImageDisplay from '@/components/image-display'
import ActionTooltip from '@/components/action-tooltip'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import SingleImageUploaderModal from '@/components/single-image-uploader/single-image-uploader-modal'

import { IBaseCompNoChild } from '@/types'
import { IUserData } from '@/server/types'
import { useUserUpdateProfilePicture } from '@/hooks/api-hooks/use-user'
import { cn } from '@/lib'

interface Props extends IBaseCompNoChild {
    userData: IUserData
    onUploadSuccess?: (newUserData: IUserData) => void
}

const UserProfilePicture = ({
    userData,
    className,
    onUploadSuccess,
}: Props) => {
    const [modal, toggleModal] = useState(false)

    const {
        mutate: updateUserProfilePicture,
        isPending: isUpdatingUserProfilePicture,
    } = useUserUpdateProfilePicture({
        onSuccess: onUploadSuccess,
    })

    return (
        <div className={cn('relative size-24', className)}>
            <SingleImageUploaderModal
                modalState={modal}
                setModalState={toggleModal}
                modalTitle="Update Profile Image"
                defaultFileName={`user-${userData.id}`}
                onUploadComplete={(newMediaResource) => {
                    updateUserProfilePicture(newMediaResource)
                    toggleModal(false)
                }}
            />
            <ImageDisplay
                className="size-full border-4 border-popover shadow-sm"
                src={userData.media?.downloadURL}
                fallback={userData.username.charAt(0) ?? '-'}
            />
            <ActionTooltip tooltipContent="Change" align="center" side="right">
                <Button
                    variant="secondary"
                    disabled={isUpdatingUserProfilePicture}
                    onClick={() => toggleModal((prev) => !prev)}
                    className="absolute bottom-2 right-2 size-fit rounded-full border border-transparent p-1 hover:border-foreground/20"
                >
                    {isUpdatingUserProfilePicture ? (
                        <LoadingSpinner />
                    ) : (
                        <CameraFillIcon className="size-4 opacity-50 duration-300 ease-in-out group-hover:opacity-80" />
                    )}
                </Button>
            </ActionTooltip>
        </div>
    )
}

export default UserProfilePicture
