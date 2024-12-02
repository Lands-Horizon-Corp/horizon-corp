import { useCallback, useEffect, useRef, useState } from 'react'

import SignaturePad from 'react-signature-canvas'
import Webcam from 'react-webcam'

import { Button } from '../ui/button'

import { dataUrlToFile } from '@/helpers'
import { FileWithPath } from 'react-dropzone'
import { useSignature } from '@/store/signature-store'

import {
    SmallBrushIcon,
    FullscreenExitIcon,
    FullscreenIcon,
    CameraIcon,
    HardDriveUploadIcon,
    DrawIcon,
} from '../icons'
import { MdOutlineFileDownload } from 'react-icons/md'

import { toast } from 'sonner'

import {
    Tooltip,
    TooltipContent,
    TooltipProvider,
    TooltipTrigger,
} from '@/components/ui/tooltip'

import { Separator } from '../ui/separator'
import { cn } from '@/lib'
import UploadSignature from './upload-signature'
import DrawSignature from './draw-signature'
import CaptureSignature from './capture-signature'
import useConfirmModalStore from '@/store/confirm-modal-store'

interface SignatureProps {
    className?: string
}

enum SignatureModes {
    DRAW = 'draw',
    CAPTURE = 'capture',
    UPLOAD = 'upload',
}

type SignatureModeType = SignatureModes

const Signature = ({ className }: SignatureProps) => {
    const [currentMode, setCurrentMode] = useState<SignatureModeType | null>(
        SignatureModes.DRAW
    )
    const signatureRef = useRef<SignaturePad | null>(null)
    const [_, setCurrentFile] = useState<FileWithPath | null>()
    const [trimmedData, setTrimmedData] = useState<string | null>('')
    const [isFullScreenMode, setIsFullScreenMode] = useState(false)

    const imageRef = useRef<HTMLImageElement | null>(null)
    const camRef = useRef<Webcam>(null)

    const { setFile } = useSignature()
    const { onOpen } = useConfirmModalStore()

    const handleSignaturePicking = (mode: SignatureModeType) => {
        setCurrentMode(mode)
        handleClear()
    }

    const handleOnUploadSignatureFileChange = (file: FileWithPath[]) => {
        setCurrentFile(file[0])
        setFile(file[0])
    }

    const handleClear = () => {
        setCurrentFile(null)
        setTrimmedData(null)
        handleClearCanvas()
    }

    const handleClearCanvas = () => {
        if (signatureRef.current) {
            signatureRef.current.clear()
        }
    }

    const handleGetSignatureTrimmedData = () => {
        if (signatureRef.current) {
            if (signatureRef.current.isEmpty()) {
                toast.warning(
                    'Drawing required: Please create a drawing before rendering.'
                )
                return
            }
            const trimmedData = signatureRef.current
                .getTrimmedCanvas()
                .toDataURL('image/png')
            const convertedData = dataUrlToFile(trimmedData, 'signature')
            setTrimmedData(trimmedData)
            setFile(convertedData)
        }
    }

    const captureImage = useCallback(() => {
        if (!camRef.current) return null
        const imageSrc = camRef.current.getScreenshot()
        return imageSrc
    }, [camRef])

    const handleCaptureImage = () => {
        const imageSrc: string | null = captureImage() ?? ''
        setTrimmedData(imageSrc)
        const convertedImageToData = dataUrlToFile(
            imageSrc,
            'capture-signature'
        )
        setFile(convertedImageToData)
        if (imageSrc) {
            toast.success('Capture Image')
        }
    }

    const handleDownloadDrawSignature = () => {
        const canvas = document.createElement('canvas')
        const context = canvas.getContext('2d')
        const downloadImage = imageRef?.current

        if (!context) return
        if (!downloadImage) return

        canvas.width = downloadImage.naturalWidth
        canvas.height = downloadImage.naturalHeight

        context.drawImage(downloadImage, 0, 0)

        canvas.toBlob((blob) => {
            if (blob) {
                const url = URL.createObjectURL(blob)
                const link = document.createElement('a')
                link.href = url
                link.download = 'signature'

                document.body.appendChild(link)
                link.click()
                document.body.removeChild(link)
                URL.revokeObjectURL(url)
            }
        }, 'png')
    }

    const handleIsFullScreenMode = () => {
        if (!isFullScreenMode) {
            document.body.setAttribute('signature-data-scroll-locked', '1')
        } else {
            document.body.removeAttribute('signature-data-scroll-locked')
        }
        setIsFullScreenMode((prev) => !prev)
    }

    const applyScrollLock = (body: HTMLElement) => {
        body.classList.add(
            'signature-with-scroll-bars-hidden',
            'signature-right-scroll-bar-position',
            'signature-width-before-scroll-bar'
        )
        body.setAttribute('signature-data-scroll-locked', '')
    }

    const removeScrollLock = (body: HTMLElement) => {
        body.classList.remove(
            'signature-with-scroll-bars-hidden',
            'signature-right-scroll-bar-position',
            'signature-width-before-scroll-bar'
        )
        body.removeAttribute('signature-data-scroll-locked')
    }

    useEffect(() => {
        const body = document.body

        if (isFullScreenMode) {
            applyScrollLock(body)
        } else {
            removeScrollLock(body)
        }

        return () => removeScrollLock(body)
    }, [isFullScreenMode])

    const onKeyPress = (event: KeyboardEvent) => {
        if (event.key === 'Escape') {
            setIsFullScreenMode(false)
        }
    }

    useEffect(() => {
        document.addEventListener('keydown', onKeyPress)

        return () => {
            document.removeEventListener('keydown', onKeyPress)
        }
    }, [])

    const isCurrentMode = (mode: SignatureModeType) => {
        return mode === currentMode ? 'bg-secondary' : ''
    }

    return (
        <div
            className={cn(
                'h-fit max-w-xl rounded-lg border bg-background/90 p-4 text-xs backdrop-blur-sm',
                isFullScreenMode
                    ? 'fixed inset-0 left-0 top-0 z-50 h-screen w-screen max-w-none bg-background/90 backdrop-blur-sm'
                    : '',
                className
            )}
        >
            <div
                className={cn(
                    'flex h-14 min-h-[50px] w-full justify-between gap-2 py-2'
                )}
            >
                <div
                    className={cn(
                        'flex',
                        isFullScreenMode ? 'space-x-2' : 'space-x-1'
                    )}
                >
                    <Button
                        variant={'outline'}
                        className={cn(
                            'text-xs',
                            isCurrentMode(SignatureModes.DRAW)
                        )}
                        onClick={() =>
                            handleSignaturePicking(SignatureModes.DRAW)
                        }
                    >
                        <DrawIcon className="size-4 lg:mr-2" />
                        <span className="hidden lg:block">Draw Signature</span>
                    </Button>
                    <Button
                        variant={'outline'}
                        className={cn(
                            'text-xs',
                            isCurrentMode(SignatureModes.CAPTURE)
                        )}
                        onClick={() =>
                            handleSignaturePicking(SignatureModes.CAPTURE)
                        }
                    >
                        <CameraIcon className="size-4 lg:mr-2" />
                        <span className="hidden lg:block">
                            Capture Signature
                        </span>
                    </Button>
                    <Button
                        variant={'outline'}
                        className={cn(
                            'text-xs',
                            isCurrentMode(SignatureModes.UPLOAD)
                        )}
                        onClick={() =>
                            handleSignaturePicking(SignatureModes.UPLOAD)
                        }
                    >
                        <HardDriveUploadIcon className="size-4 lg:mr-2" />
                        <span className="hidden lg:block">
                            Upload Signature
                        </span>
                    </Button>
                </div>
                <div className="flex w-fit items-center justify-center">
                    {isFullScreenMode ? (
                        <>
                            <TooltipProvider>
                                <Tooltip>
                                    <TooltipTrigger asChild>
                                        <Button
                                            onClick={handleIsFullScreenMode}
                                            size={'sm'}
                                            variant={'ghost'}
                                        >
                                            <FullscreenIcon
                                                size={24}
                                                className="ease-in-out hover:scale-105 hover:cursor-pointer"
                                            />
                                        </Button>
                                    </TooltipTrigger>
                                    <TooltipContent>
                                        <p className="text-xs">
                                            Exit Full Screen Mode
                                        </p>
                                    </TooltipContent>
                                </Tooltip>
                            </TooltipProvider>
                        </>
                    ) : (
                        <>
                            <TooltipProvider>
                                <Tooltip>
                                    <TooltipTrigger asChild>
                                        <Button
                                            onClick={handleIsFullScreenMode}
                                            variant={'ghost'}
                                        >
                                            <FullscreenExitIcon
                                                size={24}
                                                className="ease-in-out hover:scale-105 hover:cursor-pointer"
                                            />
                                        </Button>
                                    </TooltipTrigger>
                                    <TooltipContent>
                                        <p className="text-xs">
                                            Enter fullscreen mode
                                        </p>
                                    </TooltipContent>
                                </Tooltip>
                            </TooltipProvider>
                        </>
                    )}
                </div>
            </div>
            <div
                className={cn(
                    'mb-4 flex justify-center',
                    isFullScreenMode ? 'h-[70%]' : 'h-fit'
                )}
            >
                {currentMode === SignatureModes.DRAW && (
                    <DrawSignature
                        isFullScreenMode={isFullScreenMode}
                        signatureRef={signatureRef}
                    />
                )}
                {currentMode === SignatureModes.UPLOAD && (
                    <UploadSignature
                        isFullScreenMode={isFullScreenMode}
                        onFileChange={handleOnUploadSignatureFileChange}
                    />
                )}
                {currentMode === SignatureModes.CAPTURE && (
                    <CaptureSignature camRef={camRef} />
                )}
            </div>
            <Separator className="" />
            {trimmedData && (
                <div className="my-2 w-fit rounded-lg border-[3px] border-primary p-2">
                    <img
                        ref={imageRef}
                        src={trimmedData ?? ''}
                        alt="signature"
                        className="h-auto w-16"
                    ></img>
                </div>
            )}
            <div className="flex space-x-2 py-2">
                {currentMode !== SignatureModes.UPLOAD && (
                    <Button
                        variant={'outline'}
                        className="text-xs"
                        size={'sm'}
                        onClick={handleClear}
                    >
                        <SmallBrushIcon className="mr-2 size-4" />
                        Clear
                    </Button>
                )}
                {currentMode === SignatureModes.DRAW && (
                    <>
                        <Button
                            className="text-xs"
                            size={'sm'}
                            onClick={handleGetSignatureTrimmedData}
                        >
                            render
                        </Button>
                    </>
                )}
                {currentMode === SignatureModes.CAPTURE && (
                    <Button
                        className="text-xs"
                        size={'sm'}
                        onClick={() => {
                            if (trimmedData) {
                                onOpen({
                                    title: 'Replace Captured Image',
                                    description:
                                        'Do you want to replace the current captured image with the new one?',
                                    onConfirm: () => handleCaptureImage(),
                                })
                                return
                            } else {
                                handleCaptureImage()
                            }
                        }}
                    >
                        capture
                    </Button>
                )}
                {currentMode !== SignatureModes.UPLOAD && (
                    <Button
                        className={cn('text-xs', trimmedData ? '' : 'hidden')}
                        size={'sm'}
                        onClick={handleDownloadDrawSignature}
                    >
                        <MdOutlineFileDownload className="mr-2 size-4" />
                        download
                    </Button>
                )}
            </div>
        </div>
    )
}

export default Signature
