import { useEffect } from 'react'

import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'
import { CameraIcon, SwitchCameraIcon } from '@/components/icons'
import { useCameraDevices } from '@/components/webcam/use-camera-devices'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {
    autoPick?: boolean
    currentDevice?: MediaDeviceInfo | null
    onPick: (device: MediaDeviceInfo) => void
}

const CameraDevicePicker = ({
    onPick,
    className,
    currentDevice,
    autoPick = false,
}: Props) => {
    const { devices } = useCameraDevices()

    useEffect(() => {
        if (!autoPick || devices.length < 1) return
        onPick(devices[0])
    }, [devices])

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button
                    variant="secondary"
                    className={cn('size-fit rounded-full p-2', className)}
                >
                    <SwitchCameraIcon className="size-4" />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent side="bottom" align="end">
                <DropdownMenuLabel>Camera</DropdownMenuLabel>
                <DropdownMenuSeparator />
                {devices.map((camDevice, id) => (
                    <DropdownMenuItem
                        key={id}
                        onClick={() => onPick(camDevice)}
                        className={cn(
                            '',
                            currentDevice !== null &&
                                currentDevice?.deviceId ===
                                    camDevice.deviceId &&
                                'bg-primary'
                        )}
                    >
                        <CameraIcon className="mr-2 size-4" />
                        {camDevice.label
                            ? camDevice.label.split(' (')[0]
                            : `Camera ${id + 1}`}
                    </DropdownMenuItem>
                ))}
                {devices.length === 0 && (
                    <DropdownMenuLabel className="pointer-events-none font-normal text-foreground/70">
                        no camera available
                    </DropdownMenuLabel>
                )}
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

export default CameraDevicePicker
