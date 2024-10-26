import { toast } from 'sonner'
import { useState } from 'react'
import Cropper, { Area } from 'react-easy-crop'

import {
    RotateLeftIcon,
    RotateRightIcon,
    ZoomInIcon,
    ZoomOutIcon,
} from '@/components/icons'
import { Button } from '@/components/ui/button'
import { Slider } from '@/components/ui/slider'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn, withCatchAsync } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { getCroppedImg, ICroppedImageResult } from './picture-crop-utils'

interface Props extends IBaseCompNoChild {
    image: string
    onCancel: () => void
    onCrop: (cropResult: ICroppedImageResult) => void
}

const PictureCrop = ({ image, onCancel, onCrop }: Props) => {
    const [zoom, setZoom] = useState(1)
    const [rotation, setRotation] = useState(0)
    const [loading, setLoading] = useState(false)
    const [crop, setCrop] = useState({ x: 0, y: 0 })
    const [cropArea, setCropArea] = useState<Area | null>(null)

    const reset = () => {
        setZoom(1)
        setRotation(0)
        setCrop({ x: 0, y: 0 })
        setCropArea(null)
    }

    const handleCrop = async () => {
        if (!cropArea) return

        setLoading(true)

        const [error, croppedImageResult] = await withCatchAsync(
            getCroppedImg(image, cropArea, rotation)
        )

        setLoading(false)

        if (error) {
            toast.error(error.message)
            return
        }

        onCrop(croppedImageResult)
    }

    return (
        <div className="flex flex-col gap-y-2">
            <div className="relative h-[50vh] w-full overflow-clip rounded-2xl bg-background">
                <Cropper
                    crop={crop}
                    zoom={zoom}
                    image={image}
                    zoomWithScroll
                    aspect={1 / 1}
                    zoomSpeed={0.3}
                    showGrid={false}
                    rotation={rotation}
                    onZoomChange={setZoom}
                    onCropChange={setCrop}
                    onCropComplete={(_, cropAreaPixel) =>
                        setCropArea(cropAreaPixel)
                    }
                    onRotationChange={setRotation}
                    cropShape="round"
                />
                <p className="absolute bottom-1 left-1/2 -translate-x-1/2 rounded-md bg-popover/50 p-1 text-xs text-popover-foreground backdrop-blur">
                    {`${cropArea?.width.toLocaleString(undefined, { maximumFractionDigits: 1 }) ?? 0} x ${cropArea?.height.toLocaleString(undefined, { maximumFractionDigits: 1 }) ?? 0} rotate: ${rotation}°`}
                </p>
            </div>
            <fieldset
                disabled={loading}
                className={cn(
                    'space-y-2 rounded-2xl bg-background p-4 text-foreground/60',
                    loading && 'cursor-not-allowed'
                )}
            >
                <div className="flex items-center justify-between">
                    <p className="text-sm">Adjustment</p>
                    <Button variant="ghost" size="sm" onClick={() => reset()}>
                        Reset
                    </Button>
                </div>
                <div className="group flex items-center gap-x-2">
                    <ZoomOutIcon className="size-4 duration-200 ease-in-out group-hover:text-foreground" />
                    <Slider
                        min={1}
                        max={3}
                        step={0.08}
                        value={[zoom]}
                        disabled={loading}
                        defaultValue={[1]}
                        className="group my-4"
                        trackClassName="h-1"
                        rangeClassName="duration-400 ease-in-out transition-colors bg-primary/50 group-hover:bg-primary/80"
                        thumbClassName="size-4 duration-200 border-primary/50 bg-background shadow"
                        onValueChange={(val) => setZoom(val[0])}
                    />
                    <ZoomInIcon className="size-4 duration-200 ease-in-out group-hover:text-foreground" />
                </div>
                <div className="group flex items-center gap-x-2">
                    <RotateLeftIcon className="size-4 duration-200 ease-in-out group-hover:text-foreground" />
                    <Slider
                        min={-360}
                        max={360}
                        step={1}
                        disabled={loading}
                        value={[rotation]}
                        className="group my-4"
                        trackClassName="h-1"
                        rangeClassName="duration-400 ease-in-out transition-colors bg-primary/50 group-hover:bg-primary/80"
                        thumbClassName="size-4 duration-200 border-primary/50 bg-background"
                        onValueChange={(val) => setRotation(val[0])}
                    />
                    <RotateRightIcon className="size-4 duration-200 ease-in-out group-hover:text-foreground" />
                </div>
            </fieldset>
            <fieldset
                disabled={loading}
                className="flex items-center justify-center gap-x-2"
            >
                <Button
                    variant="ghost"
                    onClick={onCancel}
                    className="rounded-full"
                >
                    Cancel
                </Button>
                <Button onClick={handleCrop} className="rounded-full">
                    {loading ? <LoadingSpinner /> : 'Crop'}
                </Button>
            </fieldset>
        </div>
    )
}

export default PictureCrop
