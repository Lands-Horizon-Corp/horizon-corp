import React, { useState } from 'react'
import Cropper from 'react-easy-crop'

import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from '@/components/ui/dialog'

interface Crop {
    x: number
    y: number
}

interface CroppedAreaPixels {
    x: number
    y: number
    width: number
    height: number
}

const Demo: React.FC = () => {
    const [imageSrc, setImageSrc] = useState<string | null>(null)
    const [crop, setCrop] = useState<Crop>({ x: 0, y: 0 })
    const [rotation, setRotation] = useState<number>(0)
    const [zoom, setZoom] = useState<number>(1)
    const [croppedAreaPixels, setCroppedAreaPixels] =
        useState<CroppedAreaPixels | null>(null)

    const onCropComplete = (_: any, croppedAreaPixels: CroppedAreaPixels) => {
        setCroppedAreaPixels(croppedAreaPixels)
    }

    const onFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0]
        if (file) {
            const imageDataUrl = await readFile(file)
            setImageSrc(imageDataUrl as string)
        }
    }

    return (
        <div className="flex h-screen w-screen flex-col">
            <Dialog>
                <DialogTrigger>Open</DialogTrigger>
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle>Are you absolutely sure?</DialogTitle>
                        <DialogDescription>
                            This action cannot be undone. This will permanently
                            delete your account and remove your data from our
                            servers.
                        </DialogDescription>
                        {imageSrc ? (
                            <>
                                <div className="relative h-[50vh] w-full">
                                    <Cropper
                                        image={imageSrc}
                                        crop={crop}
                                        rotation={rotation}
                                        zoom={zoom}
                                        aspect={4 / 3}
                                        onCropChange={setCrop}
                                        onRotationChange={setRotation}
                                        onCropComplete={onCropComplete}
                                        onZoomChange={setZoom}
                                    />
                                </div>
                                <p>
                                    {` x ${croppedAreaPixels?.x.toLocaleString(
                                        undefined,
                                        {
                                            maximumFractionDigits: 1,
                                        }
                                    )} y ${croppedAreaPixels?.y.toLocaleString(
                                        undefined,
                                        {
                                            maximumFractionDigits: 1,
                                        }
                                    )} (${
                                        croppedAreaPixels?.width.toLocaleString(
                                            undefined,
                                            {
                                                maximumFractionDigits: 1,
                                            }
                                        ) ?? 0
                                    } x ${
                                        croppedAreaPixels?.height.toLocaleString(
                                            undefined,
                                            {
                                                maximumFractionDigits: 1,
                                            }
                                        ) ?? 0
                                    }) rotate: ${rotation}°`}
                                </p>
                            </>
                        ) : (
                            <input
                                type="file"
                                onChange={onFileChange}
                                accept="image/*"
                            />
                        )}
                    </DialogHeader>
                </DialogContent>
            </Dialog>
        </div>
    )
}

// Utility function to read file data as a Data URL
const readFile = (file: File): Promise<string | ArrayBuffer | null> => {
    return new Promise((resolve) => {
        const reader = new FileReader()
        reader.addEventListener('load', () => resolve(reader.result), false)
        reader.readAsDataURL(file)
    })
}

// These functions (getCroppedImg, getOrientation, getRotatedImage) are assumed to be available
// Please ensure to import them correctly from where they are defined
export default Demo
