import { useState } from 'react'

import FileDropUpload from './picture-drop'
import UserAvatar from '@/components/user-avatar'
import PictureCrop from '@/components/picture-crop'

const ProfileUpload = ({
    currentUserPhotoUrl,
}: {
    currentUserPhotoUrl: string | null
}) => {
    const [newImage, setNewImage] = useState<string | null>(null)

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
            {newImage !== null && (
                <PictureCrop
                    image={newImage}
                    onCrop={() => {}}
                    onCancel={() => setNewImage(null)}
                />
            )}
            {
                // TODO : Upload Button & Upload Progress
            }
        </div>
    )
}

export default ProfileUpload
