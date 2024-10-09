import Webcam from 'react-webcam'
import { forwardRef, useCallback, useRef, useState } from 'react'

import { CameraOffIcon } from '@/components/icons'

import { useCameraGrant } from './use-camera-grant'
import CameraDevicePicker from './camera-device-picker'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {
    enableBleed?: boolean
    bleedClassName?: string
}

const WebCam = forwardRef<Webcam, Props>(
    ({ className, bleedClassName, enableBleed = false }: Props, ref) => {
        const bleedRef = useRef<HTMLVideoElement>(null)
        const [selectedCamera, setSelectedCamera] =
            useState<MediaDeviceInfo | null>()
        const granted = useCameraGrant()

        const handleOnStream = useCallback(
            (stream: MediaStream) => {
                if (!bleedRef || !bleedRef.current || !enableBleed) return
                const bleedVid = bleedRef.current
                bleedVid.disablePictureInPicture
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
                {!granted && (
                    <CameraOffIcon
                        className="size-24 text-secondary-foreground/30 dark:text-secondary"
                        strokeWidth={1}
                    />
                )}
                {granted && (
                    <div className="relative max-h-none w-full max-w-full">
                        <div className="overflow-hidden rounded-xl">
                            <Webcam
                                ref={ref}
                                audio={false}
                                disablePictureInPicture
                                screenshotFormat="image/jpeg"
                                videoConstraints={{
                                    facingMode: 'user',
                                    deviceId: selectedCamera?.deviceId,
                                }}
                                onUserMedia={handleOnStream}
                            />
                        </div>
                        {enableBleed && (
                            <video
                                ref={bleedRef}
                                className={cn(
                                    'fade-in-anims absolute inset-0 left-1/2 -z-10 -translate-x-1/2 scale-x-110 blur-md',
                                    bleedClassName
                                )}
                            />
                        )}
                    </div>
                )}
                {granted && (
                    <CameraDevicePicker
                        className="absolute bottom-2 right-2"
                        autoPick
                        currentDevice={selectedCamera}
                        onPick={(devicePicked) =>
                            setSelectedCamera(devicePicked)
                        }
                    />
                )}
            </div>
        )
    }
)

export default WebCam
