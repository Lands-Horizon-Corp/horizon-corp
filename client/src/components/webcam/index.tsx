import { toast } from 'sonner'
import Webcam from 'react-webcam'
import { forwardRef, useCallback, useRef, useState } from 'react'

import {
    CameraFlipIcon,
    CameraOffIcon,
    CameraSelectIcon,
} from '@/components/icons'
import { Button } from '@/components/ui/button'
import ActionTooltip from '@/components/action-tooltip'
import CameraDevicePicker from './camera-device-picker'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn } from '@/lib/utils'

import { IBaseCompNoChild } from '@/types/component'

interface Props extends IBaseCompNoChild {
    enableBleed?: boolean
    bleedClassName?: string
}

const WebCam = forwardRef<Webcam, Props>(
    ({ className, bleedClassName, enableBleed = false }: Props, ref) => {
        const bleedRef = useRef<HTMLVideoElement>(null)
        const [camActive, setCamActive] = useState(false)
        const [camId, setCamId] = useState<string | undefined>(undefined)
        const [error, setError] = useState<string | DOMException | null>(null)
        const [facingMode, setFacingMode] = useState<'user' | 'environment'>(
            'user'
        )

        const handleOnStream = useCallback(
            (stream: MediaStream) => {
                setError(null)
                setCamActive(true)
                if (!bleedRef || !bleedRef.current || !enableBleed) return

                const bleedVid = bleedRef.current
                bleedVid.srcObject = stream
                bleedVid.onloadedmetadata = () => {
                    bleedVid.play()
                }
            },
            [bleedRef, enableBleed]
        )

        return (
            <div
                className={cn(
                    'relative mx-4 flex min-h-64 min-w-64 items-center justify-center bg-secondary dark:bg-background',
                    className
                )}
            >
                {(error || !camActive) && (
                    <span
                        className={cn(
                            'pointer-events-none absolute left-1/2 z-30 -translate-x-1/2 opacity-0 duration-300',
                            !camActive && 'opacity-100'
                        )}
                    >
                        {!error ? (
                            <LoadingSpinner />
                        ) : (
                            <CameraOffIcon
                                className="size-24 text-secondary-foreground/30 dark:text-secondary"
                                strokeWidth={1}
                            />
                        )}
                    </span>
                )}
                <div
                    className={cn(
                        'relative max-h-none w-full max-w-full opacity-0 duration-200 ease-in-out',
                        camActive && 'opacity-100'
                    )}
                >
                    <div className="overflow-hidden rounded-xl">
                        <Webcam
                            ref={ref}
                            id={camId}
                            audio={false}
                            disablePictureInPicture
                            screenshotFormat="image/jpeg"
                            videoConstraints={{
                                facingMode,
                                deviceId: camId,
                            }}
                            onUserMediaError={(e) => {
                                setError(e)
                                setCamActive(false)
                                toast.error(
                                    (e instanceof DOMException
                                        ? e.message
                                        : e) + '. Try reopening camera'
                                )
                            }}
                            onUserMedia={handleOnStream}
                        />
                    </div>
                    {enableBleed && camActive && (
                        <video
                            ref={bleedRef}
                            muted
                            disablePictureInPicture
                            className={cn(
                                'fade-in-anims absolute inset-0 left-1/2 -z-10 -translate-x-1/2 scale-x-110 blur-md',
                                bleedClassName
                            )}
                        />
                    )}
                </div>
                <div className="absolute bottom-2 right-2 flex items-center gap-x-2">
                    <ActionTooltip
                        side="left"
                        tooltipContent={`switch to ${facingMode === 'user' ? 'back camera' : 'front camera'}`}
                    >
                        <Button
                            disabled={!camActive}
                            onClick={() => {
                                setFacingMode((prev) =>
                                    prev === 'user' ? 'environment' : 'user'
                                )
                                setCamActive(false)
                            }}
                            variant="secondary"
                            className="size-fit rounded-full p-2"
                        >
                            <CameraFlipIcon className="size-4" />
                        </Button>
                    </ActionTooltip>
                    <CameraDevicePicker
                        currentCamId={camId}
                        onPick={(pickedId) => {
                            setError(null)
                            setCamActive(false)
                            setCamId(pickedId)
                        }}
                    >
                        <ActionTooltip
                            side="left"
                            tooltipContent="Change Camera"
                        >
                            <Button
                                variant="secondary"
                                className={cn(
                                    'size-fit rounded-full p-2',
                                    className
                                )}
                            >
                                <CameraSelectIcon className="size-4" />
                            </Button>
                        </ActionTooltip>
                    </CameraDevicePicker>
                </div>
            </div>
        )
    }
)

export default WebCam
