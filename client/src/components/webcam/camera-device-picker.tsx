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
    onPick: (device: MediaDeviceInfo) => void
}

const CameraDevicePicker = ({ onPick, className }: Props) => {
    const { devices } = useCameraDevices()

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
                    >
                        <CameraIcon className="mr-2 size-4" />
                        {camDevice.label}
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
