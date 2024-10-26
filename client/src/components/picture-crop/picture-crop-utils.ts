import { Area } from 'react-easy-crop'

export interface ICroppedImageResult {
    file: Blob
    base64Url: string
}

export const createImage = (url: string): Promise<HTMLImageElement> =>
    new Promise((resolve, reject) => {
        const image = new Image()
        image.addEventListener('load', () => resolve(image))
        image.addEventListener('error', (error) => reject(error))
        image.setAttribute('crossOrigin', 'anonymous')
        image.src = url
    })

export const getRadianAngle = (degreeValue: number): number =>
    (degreeValue * Math.PI) / 180

export const rotateSize = (width: number, height: number, rotation: number) => {
    const rotRad = getRadianAngle(rotation)

    return {
        width:
            Math.abs(Math.cos(rotRad) * width) +
            Math.abs(Math.sin(rotRad) * height),
        height:
            Math.abs(Math.sin(rotRad) * width) +
            Math.abs(Math.cos(rotRad) * height),
    }
}

export const getCroppedImg = async (
    imageSrc: string,
    pixelCrop: Area,
    rotation = 0
    // flip: FlipOptions = { horizontal: false, vertical: false }
): Promise<ICroppedImageResult> => {
    const image = await createImage(imageSrc)
    const canvas = document.createElement('canvas')
    const ctx = canvas.getContext('2d')

    if (!ctx) {
        throw new Error(
            'Crop Failed, unexpected problem while creating a canvas'
        )
    }

    const rotRad = getRadianAngle(rotation)

    const { width: bBoxWidth, height: bBoxHeight } = rotateSize(
        image.width,
        image.height,
        rotation
    )

    canvas.width = bBoxWidth
    canvas.height = bBoxHeight

    ctx.translate(bBoxWidth / 2, bBoxHeight / 2)
    ctx.rotate(rotRad)
    ctx.scale(1, 1)
    ctx.translate(-image.width / 2, -image.height / 2)

    ctx.drawImage(image, 0, 0)

    const croppedCanvas = document.createElement('canvas')
    const croppedCtx = croppedCanvas.getContext('2d')

    if (!croppedCtx) {
        throw new Error(
            'Crop Failed, unexpected problem while cropping the canvas'
        )
    }

    croppedCanvas.width = pixelCrop.width
    croppedCanvas.height = pixelCrop.height

    croppedCtx.drawImage(
        canvas,
        pixelCrop.x,
        pixelCrop.y,
        pixelCrop.width,
        pixelCrop.height,
        0,
        0,
        pixelCrop.width,
        pixelCrop.height
    )

    return new Promise((resolve, reject) => {
        croppedCanvas.toBlob((file) => {
            if (file) {
                resolve({
                    file,
                    base64Url: URL.createObjectURL(file),
                })
            } else {
                reject(new Error('Failed to produce final crop result'))
            }
        }, 'image/jpeg')
    })
}
