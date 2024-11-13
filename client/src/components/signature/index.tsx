import {
    forwardRef,
    MutableRefObject,
    useCallback,
    useRef,
    useState,
} from 'react'

import ReactSignatureCanvas from 'react-signature-canvas'
import SignaturePad from 'react-signature-canvas'
import Webcam from 'react-webcam'

import { Button } from '../ui/button'
import WebCam from '../webcam'

import { dataUrlToFile } from '@/helpers'
import { useTheme } from '@/providers/theme-provider'
import FileUploader from '../ui/file-uploader'
import { FileWithPath } from 'react-dropzone'
import { useSignature } from '@/store/signature-store'
import {
    AiOutlineClearIcon,
    AiOutlineFullscreenExitIcon,
    AiOutlineFullscreenIcon,
    FiCameraIcon,
    LuHardDriveUploadIcon,
    MdOutlineDrawIcon,
} from '../icons'
import { cn } from '@/lib'
import { MdOutlineFileDownload } from 'react-icons/md'
import { toast } from 'sonner'
import {
    Tooltip,
    TooltipContent,
    TooltipProvider,
    TooltipTrigger,
} from '@/components/ui/tooltip'

interface UploadSignatureProps {
    /**
     * Callback function triggered when files are added or changed.
     * Receives an array of File objects.
     */
    onFileChange: (files: File[]) => void
}

interface CaptureSignatureProps {
    /**
     * Reference to the Webcam component for capturing signatures.
     * Used to access the webcam instance methods.
     */
    camRef: MutableRefObject<Webcam | null>
}

interface DrawSignatureProps {
    /**
     * Reference to the SignatureCanvas component for drawing signatures.
     * Used to access the canvas instance methods.
     */
    signatureRef: MutableRefObject<ReactSignatureCanvas | null>
}

interface SignatureProps {
    className?: string
}

const CaptureSignature = ({ camRef }: CaptureSignatureProps) => {
    return (
        <>
            <WebCam className="!mx-0 rounded-lg" ref={camRef} />
        </>
    )
}

const DrawSignature = ({ signatureRef }: DrawSignatureProps) => {
    const { resolvedTheme: theme } = useTheme()

    return (
        <>
            <div className="w-full">
                <SignaturePad
                    penColor={`${theme === 'dark' ? 'white' : 'black'}`}
                    ref={signatureRef}
                    canvasProps={{
                        className:
                            'sigCanvas w-full h-[200px] rounded-lg border shadow-md dark:bg-secondary',
                    }}
                />
            </div>
        </>
    )
}

const UploadSignature = ({ onFileChange }: UploadSignatureProps) => {
    return (
        <FileUploader
            className="!mx-0 w-full"
            maxFiles={1}
            accept={{
                'image/png': ['.png'],
                'image/jpeg': ['.jpg', '.jpeg'],
            }}
            onFileChange={onFileChange}
        />
    )
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

    const { file, setFile } = useSignature()
    const camRef = useRef<Webcam>(null)

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
    }

    const handleDownloadDrawSignature = () => {
        const canvas = document.createElement('canvas')
        const context = canvas.getContext('2d')
        let downloadImage = imageRef?.current

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
        setIsFullScreenMode((prev) => !prev)
    }

    return (
        <div
        
            className={cn(
                'h-fit max-w-xl rounded-lg border p-4 text-xs',
                isFullScreenMode ? 'w-screen fixed inset-0 left-0 top-0 max-w-none backdrop-blur-sm bg-background/90 z-50 h-[100vh]' : '',
                className
            )}
        >
            <div className="flex h-14 min-h-[50px] w-full justify-start gap-2 py-2">
                <Button
                    variant={'outline'}
                    className="text-xs"
                    onClick={() =>
                        handleSignaturePicking(SignatureModes.CAPTURE)
                    }
                >
                    <FiCameraIcon className="mr-2 size-4" />
                    Capture Signature
                </Button>
                <Button
                    variant={'outline'}
                    className="text-xs"
                    onClick={() =>
                        handleSignaturePicking(SignatureModes.UPLOAD)
                    }
                >
                    <LuHardDriveUploadIcon className="mr-2 size-4" />
                    Upload Signature
                </Button>
                <Button
                    variant={'outline'}
                    className="text-xs"
                    onClick={() => handleSignaturePicking(SignatureModes.DRAW)}
                >
                    <MdOutlineDrawIcon className="mr-2 size-4" />
                    Draw Signature
                </Button>
                <div className="flex h-full w-full items-center justify-center">
                    {isFullScreenMode ? (
                        <>
                            <TooltipProvider>
                                <Tooltip>
                                    <TooltipTrigger asChild>
                                        <Button size={'sm'} variant={'ghost'}>
                                            <AiOutlineFullscreenIcon
                                                onClick={handleIsFullScreenMode}
                                                size={24}
                                                className="ease-in-out hover:scale-105 hover:cursor-pointer"
                                            />
                                        </Button>
                                    </TooltipTrigger>
                                    <TooltipContent>
                                        <p>Enter Full Screen Mode</p>
                                    </TooltipContent>
                                </Tooltip>
                            </TooltipProvider>
                        </>
                    ) : (
                        <>
                            <TooltipProvider>
                                <Tooltip>
                                    <TooltipTrigger asChild>
                                        <Button variant={'ghost'}>
                                            <AiOutlineFullscreenExitIcon
                                                onClick={handleIsFullScreenMode}
                                                size={24}
                                                className="ease-in-out hover:scale-105 hover:cursor-pointer"
                                            />
                                        </Button>
                                    </TooltipTrigger>
                                    <TooltipContent>
                                        <p>Exit fullscreen mode</p>
                                    </TooltipContent>
                                </Tooltip>
                            </TooltipProvider>
                        </>
                    )}
                </div>
            </div>
            <div className="flex h-fit items-center justify-center">
                {currentMode === SignatureModes.DRAW && (
                    <DrawSignature signatureRef={signatureRef} />
                )}
                {currentMode === SignatureModes.UPLOAD && (
                    <UploadSignature
                        onFileChange={handleOnUploadSignatureFileChange}
                    />
                )}
                {currentMode === SignatureModes.CAPTURE && (
                    <CaptureSignature camRef={camRef} />
                )}
            </div>
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
                        <AiOutlineClearIcon className="mr-2 size-4" />
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
                        {trimmedData && (
                            <Button
                                className="text-xs"
                                size={'sm'}
                                onClick={handleDownloadDrawSignature}
                            >
                                <MdOutlineFileDownload className="mr-2 size-4" />
                                download
                            </Button>
                        )}
                    </>
                )}
                {currentMode === SignatureModes.CAPTURE && (
                    <Button
                        className="text-xs"
                        size={'sm'}
                        onClick={handleCaptureImage}
                    >
                        capture
                    </Button>
                )}
            </div>
        </div>
    )
}

export default Signature