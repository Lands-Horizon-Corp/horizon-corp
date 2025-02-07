import { MutableRefObject } from 'react'
import * as ImagePreviewPrimitive from '@radix-ui/react-dialog'

import { IMediaResource } from '@/server/types'

export interface ImagePreviewProps
    extends React.ComponentPropsWithoutRef<
        typeof ImagePreviewPrimitive.Content
    > {
    hideCloseButton?: boolean
    closeButtonClassName?: string
    overlayClassName?: string
    Images: IMediaResource[]
    scaleInterval?: number
}

export interface ImageContainerProps extends Partial<DownloadProps> {
    media: IMediaResource
    scale: number
    rotateDegree: number
    flipScale: string
}

export interface ImagePreviewActionProps extends Partial<DownloadProps> {
    handleZoomIn: () => void
    handleZoomOut: () => void
    handleRotateLeft: () => void
    handleRotateRight: () => void
    handleResetActionState: () => void
    handleFlipHorizontal: () => void
    handleFlipVertical: () => void
    downloadImage: DownloadProps
    className: string
}

export interface ImagePreviewPanelProps {
    Images: IMediaResource[]
    focusIndex: number
    scrollToIndex: (index: number) => void
}

export interface DownloadProps
    extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    fileUrl?: string
    fileName: string
    fileType: string
    imageRef?: MutableRefObject<HTMLImageElement | null>
    name?: string
}

export interface ImagePreviewButtonActionProps {
    onClick: () => void
    className?: string
    Icon?: React.ReactNode
    name?: string
    iconClassName?: string
}
