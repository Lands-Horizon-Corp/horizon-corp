import * as React from 'react'
import * as ImagePreviewPrimitive from '@radix-ui/react-dialog'
import { X } from 'lucide-react'

import { cn } from '@/lib/utils'

const ImagePreview = ImagePreviewPrimitive.Root

const ImagePreviewTrigger = ImagePreviewPrimitive.Trigger

const ImagePreviewPortal = ImagePreviewPrimitive.Portal

const ImagePreviewClose = ImagePreviewPrimitive.Close

const ImagePreviewOverlay = React.forwardRef<
    React.ElementRef<typeof ImagePreviewPrimitive.Overlay>,
    React.ComponentPropsWithoutRef<typeof ImagePreviewPrimitive.Overlay>
>(({ className, ...props }, ref) => (
    <ImagePreviewPrimitive.Overlay
        ref={ref}
        className={cn(
            'fixed inset-0 z-50 bg-black/80 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0',
            className
        )}
        {...props}
    />
))
ImagePreviewOverlay.displayName = ImagePreviewPrimitive.Overlay.displayName

interface ImagePreviewProps
    extends React.ComponentPropsWithoutRef<typeof ImagePreviewPrimitive.Content> {
    hideCloseButton?: boolean
    closeButtonClassName?: string
    overlayClassName?: string
}

const ImagePreviewContent = React.forwardRef<
    React.ElementRef<typeof ImagePreviewPrimitive.Content>,
    ImagePreviewProps
>(
    (
        {
            className,
            hideCloseButton = false,
            closeButtonClassName,
            overlayClassName,
            children,
            ...props
        },
        ref
    ) => (
        <ImagePreviewPortal>
            <ImagePreviewOverlay
                className={cn(
                    'bg-black/30 dark:bg-background/70',
                    overlayClassName
                )}
            />
            <ImagePreviewPrimitive.Content
                ref={ref}
                className={cn(
                    'fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border bg-popover/90 p-6 shadow-lg backdrop-blur duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%] sm:rounded-lg',
                    className
                )}
                {...props}
            >
                {children}
                {!hideCloseButton && (
                    <ImagePreviewPrimitive.Close
                        className={cn(
                            'absolute right-4 top-4 size-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground',
                            closeButtonClassName
                        )}
                    >
                        <X className="size-full" />
                        <span className="sr-only">Close</span>
                    </ImagePreviewPrimitive.Close>
                )}
            </ImagePreviewPrimitive.Content>
        </ImagePreviewPortal>
    )
)
ImagePreviewContent.displayName = ImagePreviewPrimitive.Content.displayName

export {
    ImagePreview,
    ImagePreviewPortal,
    ImagePreviewOverlay,
    ImagePreviewClose,
    ImagePreviewTrigger,
    ImagePreviewContent,
}
