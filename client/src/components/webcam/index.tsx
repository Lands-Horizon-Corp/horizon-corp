import Webcam from 'react-webcam'
import { forwardRef } from 'react'

import { CameraOffIcon } from '../icons'
import { useCameraGrant } from './use-camera-grant'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {}

const WebCam = forwardRef<Webcam, Props>(({ className }: Props, ref) => {
    const granted = useCameraGrant()

    return (
        <div
            className={cn(
                'grid size-80 relative place-content-center overflow-clip rounded-full bg-secondary dark:bg-background',
                className
            )}
        >
            {!granted ? (
                <CameraOffIcon
                    className="size-24 text-secondary-foreground/30 dark:text-secondary"
                    strokeWidth={1}
                />
            ) : (
                <Webcam
                    ref={ref}
                    audio={false}
                    screenshotFormat="image/jpeg"
                    videoConstraints={{ facingMode : 'user'}}
                    className="object-fit absolute h-full w-full left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2"
                />
            )}
        </div>
    )
})

export default WebCam
