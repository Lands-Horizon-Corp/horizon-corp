import { toast } from 'sonner'
import { useState, forwardRef } from 'react'

import { ImageIcon } from '@/components/icons'
import { Button } from '@/components/ui/button'
import { SingleImageUploaderModal } from '@/components/uploaders/single-image-uploader'

import { IMediaResource, TEntityId } from '@/server'
import { abbreviateUUID } from '@/utils/formatting-utils'

export interface ImageUploadFieldProps {
    id?: string
    name?: string
    value?: TEntityId
    placeholder?: string
    onChange?: (media: IMediaResource | undefined) => void
}

export const ImageUploadField = forwardRef<
    HTMLButtonElement,
    ImageUploadFieldProps
>(({ value, placeholder, onChange, ...other }, ref) => {
    const [uploaderModal, setUploaderModal] = useState(false)

    // TODO: Implement loading preview of MediaResource based on value
    const handleOnUpload = (media: IMediaResource) => {
        toast.success(`Image Uploaded ${media.fileName} with ID: ${media.id}`)
        onChange?.({ ...media, id: '550e8400-e29b-41d4-a716-446655440000' })
        setUploaderModal(false)
    }

    return (
        <div>
            <SingleImageUploaderModal
                title="Upload Image"
                description="Choose/Upload single image. You may also capture using camera."
                open={uploaderModal}
                onOpenChange={setUploaderModal}
                singleImageUploaderProp={{
                    onSuccess: handleOnUpload,
                    className: 'p-0',
                }}
                className="max-w-xl bg-popover p-8"
            />
            <Button
                {...other}
                ref={ref}
                role="combobox"
                variant="outline"
                onClick={() => setUploaderModal(true)}
                className="w-full justify-between bg-background px-3 font-normal outline-offset-0 hover:bg-background focus-visible:border-ring focus-visible:outline-[3px] focus-visible:outline-ring/20"
            >
                {value ? (
                    <span>{abbreviateUUID(value, 14)} (Uploaded Image)</span>
                ) : (
                    (placeholder ?? 'Upload Image')
                )}
                <ImageIcon className="shrink-0 text-muted-foreground/80" />
            </Button>
        </div>
    )
})

ImageUploadField.displayName = 'ImageUploadField'
