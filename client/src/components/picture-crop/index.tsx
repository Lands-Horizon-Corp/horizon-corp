import { useState } from 'react'
import Cropper, { Area } from 'react-easy-crop'

import { Button } from '@/components/ui/button'
import { Slider } from '@/components/ui/slider'

import { IBaseCompNoChild } from '@/types'

interface Props extends IBaseCompNoChild {
    image: string
    onCancel: () => void
    onCrop: () => void
}

const PictureCrop = ({ image, onCancel, onCrop }: Props) => {
    const [zoom, setZoom] = useState(1)
    const [crop, setCrop] = useState({ x: 0, y: 0 })
    const [cropArea, setCropArea] = useState<Area | null>(null)

    // TODO : 
    // image + crop coord -> new image
    // new image -> new file

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
                    onZoomChange={setZoom}
                    onCropChange={setCrop}
                    onCropComplete={setCropArea}
                    style={{
                        cropAreaStyle: {
                            color: 'hsl(var(--background) / 0.7)',
                        },
                    }}
                    classes={{
                        cropAreaClassName: 'rounded-full',
                    }}
                />
            </div>
            <Slider
                min={1}
                max={3}
                step={0.08}
                value={[zoom]}
                defaultValue={[1]}
                className="group my-4"
                trackClassName="h-1"
                rangeClassName="duration-400 ease-in-out transition-colors bg-primary/70 group-hover:bg-primary"
                thumbClassName="size-4 duration-200 border-primary/50 bg-background cursor-grab"
                onValueChange={(val) => setZoom(val[0])}
            />
            <div className="flex items-center justify-center gap-x-2">
                <Button
                    variant="ghost"
                    className="rounded-full"
                    onClick={() => onCancel()}
                >
                    Cancel
                </Button>
                <Button className="rounded-full">
                    Crop
                </Button>
            </div>
        </div>
    )
}

export default PictureCrop
