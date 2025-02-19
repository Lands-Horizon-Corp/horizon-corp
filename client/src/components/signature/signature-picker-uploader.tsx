import { useState } from 'react'

import Signature from '.'
import { Button } from '../ui/button'
import ImageDisplay from '../image-display'
import FileItem from '../uploaders/file-item'
import Modal, { IModalProps } from '../modals/modal'
import LoadingSpinner from '../spinners/loading-spinner'

import { IMediaResource } from '@/server'
import { useSinglePictureUpload } from '@/hooks/api-hooks/use-media-resource'

interface Props {
    onSignatureUpload: (signatureMedia: IMediaResource) => void
}

const SignaturePickerUploader = ({ onSignatureUpload }: Props) => {
    const [eta, setEta] = useState('')
    const [progress, setProgress] = useState(0)
    const [file, setFile] = useState<File | undefined>(undefined)

    const { isPending: isUploading, mutate: uploadSignature } =
        useSinglePictureUpload({
            onSuccess: (media) => {
                onSignatureUpload?.(media)
            },
            onUploadProgressChange: (progress) => setProgress(progress),
            onUploadETAChange: (ETA) => setEta(ETA),
        })

    return (
        <div className="space-y-2">
            {file ? (
                <>
                    <ImageDisplay
                        fallbackClassName="rounded-none"
                        imageClassName="object-contain rounded-none"
                        className="min-h-60 w-full rounded-lg border bg-background"
                        src={URL.createObjectURL(file)}
                    />
                    <FileItem
                        file={file}
                        onRemoveFile={() => setFile(undefined)}
                        uploadDetails={{ eta, progress, isUploading }}
                    />
                    <Button
                        type="button"
                        disabled={isUploading}
                        onClick={() => uploadSignature(file)}
                        className="w-full"
                    >
                        {isUploading ? <LoadingSpinner /> : 'Upload Signature'}
                    </Button>
                </>
            ) : (
                <>
                    <Signature
                        hideDownload
                        onSignatureChange={(signature) => setFile(signature)}
                    />
                </>
            )}
        </div>
    )
}

export const SignaturePickerUploaderModal = ({
    signatureUploadProps,
    ...other
}: IModalProps & { signatureUploadProps: Props }) => {
    return (
        <Modal {...other}>
            <SignaturePickerUploader {...signatureUploadProps} />
        </Modal>
    )
}

export default SignaturePickerUploader
