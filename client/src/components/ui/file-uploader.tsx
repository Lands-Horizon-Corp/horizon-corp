import { formatBytes } from '@/helpers'
import { cn } from '@/lib'
import { useEffect, useState } from 'react'
import { DropzoneOptions, FileWithPath, useDropzone } from 'react-dropzone'
import { FaCloudUploadAlt } from 'react-icons/fa'
import { toast } from 'sonner'
import { Button } from './button'
import FileTypeIcon from './file-type'
import { RxTrashIcon } from '../icons'

interface FileUploaderProps extends DropzoneOptions {
    className?: string
    onFileChange: (file: File[]) => void
}

const FileUploader = ({
    className,
    onFileChange,
    ...props
}: FileUploaderProps) => {
    const [hasError, setHasError] = useState(false)
    const [uploadedFiles, setUploadedFiles] = useState<readonly FileWithPath[]>(
        []
    ) // Ensure it's initialized as an empty array

    const {
        getRootProps,
        getInputProps,
        acceptedFiles: files,
        open,
    } = useDropzone({
        onDrop: (acceptedFiles, fileRejections) => {
            if (fileRejections.length > 0) {
                const error = fileRejections[0].errors[0]
                const errorMessage =
                    error.code === 'too-many-files'
                        ? 'Only one image file is allowed. Please choose just one.'
                        : error.message
                toast.error(errorMessage)
                setHasError(true)
                return
            }

            setHasError(false)
            onFileChange(acceptedFiles)
            handleFilesChange(acceptedFiles)
        },
        onError: (error) => {
            toast.error(error.message)
        },
        ...props,
        noClick: true,
    })

    useEffect(() => {
        setUploadedFiles(files)
    }, [files])

    const handleFilesChange = (newFiles: FileWithPath[]) => {
        setUploadedFiles((prevFiles) => [...prevFiles, ...newFiles])
    }

    const handleDeleteFile = (index: number) => () => {
        setUploadedFiles((prevFiles) =>
            prevFiles.filter((_, idx) => idx !== index)
        )
    }

    return (
        <div
            className={cn(
                'relative mx-4 min-h-48 min-w-48 cursor-pointer',
                className
            )}
        >
            <div
                {...getRootProps()}
                className="mb-2 flex h-full min-h-64 w-full flex-col items-center justify-center gap-2 rounded-lg border-2 border-dashed border-slate-800/30 p-2 pb-4 text-sm dark:bg-background"
                aria-label="File upload area"
            >
                <input {...getInputProps()} aria-hidden="true" />
                <FaCloudUploadAlt className="size-24 text-primary" />
                <p>
                    <span
                        onClick={open}
                        className="cursor-pointer font-semibold underline"
                        aria-label="Click to upload files"
                    >
                        Click to upload
                    </span>{' '}
                    or Drag and Drop
                </p>
                <Button
                    onClick={open}
                    variant="outline"
                    className="text-xs"
                    aria-label="Select files"
                >
                    Select Files
                </Button>
            </div>
            <div className="space-y-2">
                {!hasError &&
                    uploadedFiles.map((file, idx) => (
                        <UploadedFileItem
                            key={idx}
                            file={file}
                            onDelete={handleDeleteFile(idx)}
                        />
                    ))}
            </div>
        </div>
    )
}

interface UploadedFileItemProps {
    file: File
    onDelete: () => void
}

const UploadedFileItem = ({ file, onDelete }: UploadedFileItemProps) => {
    return (
        <div className="flex w-full items-center space-x-3 rounded-lg border border-slate-300 p-3">
            <FileTypeIcon file={file} />
            <div className="flex-grow">
                <p className="text-xs font-semibold">{file.name}</p>
                <p className="text-xs">{formatBytes(file.size)}</p>
            </div>
            <RxTrashIcon
                onClick={onDelete}
                size={18}
                className="cursor-pointer text-destructive hover:scale-105"
            />
        </div>
    )
}

export default FileUploader
