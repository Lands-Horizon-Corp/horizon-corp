import { MediaResource } from "@/horizon-corp/types";
import { MutableRefObject } from "react";
import * as ImagePreviewPrimitive from '@radix-ui/react-dialog'

export interface ImagePreviewProps extends React.ComponentPropsWithoutRef<typeof ImagePreviewPrimitive.Content> {
          hideCloseButton?: boolean;
          closeButtonClassName?: string;
          overlayClassName?: string;
          Images: MediaResource[];
          scaleInterval?: number;
      }
      
export interface ImageContainerProps extends Partial<DownloadProps> {
          media: MediaResource
          scale: number
          rotateDegree: number
          flipScale: string
      }

export  interface ImagePreviewActionProps extends Partial<DownloadProps> {
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
          Images: MediaResource[]
          focusIndex: number
          scrollToIndex: (index: number) => void
      }

export interface DownloadProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
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