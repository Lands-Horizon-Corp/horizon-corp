import * as React from 'react'
import * as ImagePreviewPrimitive from '@radix-ui/react-dialog'

import {
    X,
    ChevronLeftIcon,
    ChevronRightIcon,
    ZoomIn,
    ZoomOut,
    FlipHorizontal,
    FlipVertical,
    DownloadIcon,
} from 'lucide-react'
import {
    Carousel,
    CarouselContent,
    CarouselItem,
    useCarousel,
} from '@/components/ui/carousel'
import { cn } from '@/lib/utils'
import { MediaResource } from '@/horizon-corp/types'
import { Card } from './card'
import { type CarouselApi } from '@/components/ui/carousel'
import { useCallback, useEffect, useRef, useState } from 'react'
import { useImagePreview } from '@/store/image-preview-store'
import { Button } from './button'
import useEmblaCarousel from 'embla-carousel-react'
import { MdRotateLeft, MdRotateRight } from 'react-icons/md'
import { AspectRatio } from '@/components/ui/aspect-ratio'

type UseCarouselParameters = Parameters<typeof useEmblaCarousel>
type CarouselOptions = UseCarouselParameters[0]

const ImagePreview = ImagePreviewPrimitive.Root

const ImagePreviewPortal = ImagePreviewPrimitive.Portal

interface ImagePreviewProps
    extends React.ComponentPropsWithoutRef<
        typeof ImagePreviewPrimitive.Content
    > {
    hideCloseButton?: boolean
    closeButtonClassName?: string
    overlayClassName?: string
    Images: MediaResource[]
}

interface ImageZoomContainerProps {
    url: string
    scale: number
    rotateDegree: number
    flipScale: string
}
interface Position {
    x: number
    y: number
}

interface ImagePreviewActionProps extends Partial<ImagePreviewProps> {
    handleZoomIn: () => void
    handleZoomOut: () => void
    handleRoteteLeft: () => void
    handleRoteteRight: () => void
    handleResetActionState: () => void
    handleFlipHorizontal: () => void
    handleFlipVertical: () => void
    downloadImage: DownloadProps
}
interface DownloadProps {
    fileUrl: string
    filename: string
}

const handleDownload = ({ fileUrl, filename }: DownloadProps) => {
    const link = document.createElement('a')
    link.href = fileUrl
    link.download = filename
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
}

interface DownloadButtonProps {
    downloadImage: DownloadProps
}

export const DownloadButton = ({ downloadImage }: DownloadButtonProps) => {
    return (
        <Button
            variant="outline"
            className={cn('size-14 rounded-full border-0 bg-transparent')}
            onClick={() => handleDownload(downloadImage)}
        >
            <DownloadIcon className="size-full" />
        </Button>
    )
}

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

const CarouselPrevious = React.forwardRef<
    HTMLButtonElement,
    React.ComponentProps<typeof Button>
>(({ className, variant = 'outline', size = 'icon', ...props }, ref) => {
    const { orientation, scrollPrev, canScrollPrev } = useCarousel()

    return (
        <Button
            ref={ref}
            variant={variant}
            size={size}
            className={cn(
                '!hover:scale-125 absolute size-10 bg-transparent ease-in-out hover:bg-transparent',
                orientation === 'horizontal'
                    ? '-left-12 top-1/2 -translate-y-1/2'
                    : '-top-12 left-1/2 -translate-x-1/2 rotate-90',
                !canScrollPrev ? 'hidden' : '',
                className
            )}
            onClick={scrollPrev}
            {...props}
        >
            <ChevronLeftIcon className="size-full" />
            <span className="sr-only">Previous slide</span>
        </Button>
    )
})
CarouselPrevious.displayName = 'CarouselPrevious'

const CarouselNext = React.forwardRef<
    HTMLButtonElement,
    React.ComponentProps<typeof Button>
>(({ className, variant = 'outline', size = 'icon', ...props }, ref) => {
    const { orientation, scrollNext, canScrollNext } = useCarousel()

    return (
        <Button
            ref={ref}
            variant={variant}
            size={size}
            className={cn(
                '!hover:scale-125 absolute size-10 bg-transparent ease-in-out hover:bg-transparent',
                orientation === 'horizontal'
                    ? '-right-12 top-1/2 -translate-y-1/2'
                    : '-bottom-12 left-1/2 -translate-x-1/2 rotate-90',
                !canScrollNext ? 'hidden' : '',
                className
            )}
            onClick={scrollNext}
            {...props}
        >
            <ChevronRightIcon className="size-full" />
            <span className="sr-only">Next slide</span>
        </Button>
    )
})

CarouselNext.displayName = 'CarouselNext'

const ImageZoomContainer = ({
    url,
    scale,
    rotateDegree,
    flipScale,
}: ImageZoomContainerProps) => {
    const [position, setPosition] = useState<Position>({ x: 0, y: 0 })
    const ImageRef = useRef<HTMLImageElement | null>(null)

    useEffect(() => {
        const image = ImageRef.current
        let isDragging = false
        let previousPosition = { x: 0, y: 0 }

        const handleMouseDown = (e: MouseEvent) => {
            isDragging = true
            previousPosition = { x: e.clientX, y: e.clientY }
        }

        const handleMouseMove = (e: MouseEvent) => {
            if (!isDragging) return
            const deltaX = e.clientX - previousPosition.x
            const deltaY = e.clientY - previousPosition.y
            previousPosition = { x: e.clientX, y: e.clientY }
            setPosition((position) => ({
                x: position.x + deltaX,
                y: position.y + deltaY,
            }))
        }

        const handleMouseUp = () => {
            isDragging = false
        }

        if (image) {
            image.addEventListener('mousedown', handleMouseDown)
        }

        // Add mousemove and mouseup events to window
        window.addEventListener('mousemove', handleMouseMove)
        window.addEventListener('mouseup', handleMouseUp)

        return () => {
            if (image) {
                image.removeEventListener('mousedown', handleMouseDown)
            }
            window.removeEventListener('mousemove', handleMouseMove)
            window.removeEventListener('mouseup', handleMouseUp)
        }
    }, [ImageRef, scale])

    return (
        <div className="relative overflow-hidden rounded-lg">
            <img
                ref={ImageRef}
                style={{
                    width: '100vw',
                    height: 'auto',
                    cursor: 'move',
                    transform: `scale(${scale}) translate(${position.x}px, ${position.y}px) rotate(${rotateDegree}deg) ${flipScale}`,
                    transition: 'transform 0.2s ease-in-out',
                }}
                draggable="false"
                src={url}
                alt="Zoomable"
            />
        </div>
    )
}

const ImagePreviewActions = ({
    handleZoomIn,
    handleZoomOut,
    handleRoteteLeft,
    handleRoteteRight,
    hideCloseButton,
    closeButtonClassName,
    handleResetActionState,
    handleFlipHorizontal,
    handleFlipVertical,
    downloadImage,
    className
}: ImagePreviewActionProps) => {
    return (
        <div className={cn('absolute right-5 top-4 flex items-center', className)}>
            <Button
                variant={'outline'}
                className={cn('size-14 rounded-full border-0 bg-transparent')}
                onClick={handleZoomIn}
            >
                <ZoomIn className="size-full cursor-pointer dark:text-white" />
                <span className="sr-only">Zoom In</span>
            </Button>
            <Button
                variant={'outline'}
                className={cn('size-14 rounded-full border-0 bg-transparent')}
                onClick={handleZoomOut}
            >
                <ZoomOut className="size-full cursor-pointer dark:text-white" />
                <span className="sr-only">Zoom Out</span>
            </Button>
            <Button
                variant={'outline'}
                className={cn('size-14 rounded-full border-0 bg-transparent')}
                onClick={handleRoteteLeft}
            >
                <MdRotateLeft className="size-full cursor-pointer dark:text-white" />
                <span className="sr-only">Rotate Left</span>
            </Button>
            <Button
                variant={'outline'}
                className={cn('size-14 rounded-full border-0 bg-transparent')}
                onClick={handleRoteteRight}
            >
                <MdRotateRight className="size-full cursor-pointer dark:text-white" />
                <span className="sr-only">Rotate Right</span>
            </Button>
            <Button
                variant={'outline'}
                className={cn('size-14 rounded-full border-0 bg-transparent')}
                onClick={handleFlipHorizontal}
            >
                <FlipHorizontal className="size-full cursor-pointer dark:text-white" />
                <span className="sr-only">Flip Horizontal</span>
            </Button>
            <Button
                variant={'outline'}
                className={cn('size-14 rounded-full border-0 bg-transparent')}
                onClick={handleFlipVertical}
            >
                <FlipVertical className="size-full cursor-pointer dark:text-white" />
                <span className="sr-only">Flip Horizontal</span>
            </Button>
            <DownloadButton downloadImage={downloadImage}></DownloadButton>
            {!hideCloseButton && (
                <ImagePreviewPrimitive.Close
                    onClick={handleResetActionState}
                    className={cn(
                        'ml-2 size-8 cursor-pointer rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground',
                        closeButtonClassName
                    )}
                >
                    <X className="size-full" />
                    <span className="sr-only">Close</span>
                </ImagePreviewPrimitive.Close>
            )}
        </div>
    )
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
            Images,
            ...props
        },
        ref
    ) => {
        const isMultipleImage = Images.length > 1
        const [api, setApi] = useState<CarouselApi | undefined>()
        const [scale, setScale] = useState(1)
        const [rotateDegree, setRotateDegree] = useState(0)
        const [flipScale, setFlipScale] = useState('')
        const [downloadImage, setDownLoadImage] = useState<DownloadProps>({
            filename: Images[0].fileName,
            fileUrl: Images[0].url,
        })
        const [options, _] = React.useState<CarouselOptions>({
            align: 'center',
            loop: false,
            watchDrag: false,
            dragFree: false,
        })

        const { onClose, focusIndex, setFocusIndex } = useImagePreview()

        useEffect(() => {
            if (api && focusIndex !== undefined) {
                api.scrollTo(focusIndex)
                const ImageToDownload = Images[focusIndex]
                setDownLoadImage({
                    filename: ImageToDownload.fileName,
                    fileUrl: ImageToDownload.url,
                })
            }
        }, [api, focusIndex])

        const handleSelect = useCallback(() => {
            if (api) {
                const selectedIndex = api.selectedScrollSnap()
                setFocusIndex(selectedIndex)
            }
        }, [api, setFocusIndex])

        console.log(downloadImage)

        useEffect(() => {
            if (!api) {
                return
            }
            api.on('select', handleSelect)
        }, [api, handleSelect])

        const scrollToIndex = useCallback(
            (index: number) => {
                if (api) {
                    api.scrollTo(index)
                }
            },
            [api]
        )

        const handleOverlayClick = (e: React.MouseEvent<HTMLDivElement>) => {
            if ((e.target as HTMLDivElement).id === 'overlay') {
                onClose()
            }
        }

        const handleZoomIn = () => {
            setScale((prevScale) => prevScale + 0.1)
        }

        const handleZoomOut = () => {
            setScale((prevScale) => Math.max(prevScale - 0.1, 1)) // Prevent scaling below 1
        }

        const handleRoteteLeft = () => {
            setRotateDegree((prevRotateDegree) => prevRotateDegree + 10)
        }

        const handleRoteteRight = () => {
            setRotateDegree((prevRotateDegree) => prevRotateDegree - 10)
        }

        const handleFlipHorizontal = () => {
            setFlipScale((prevScale) =>
                prevScale === 'scaleX(-1)' ? 'scaleX(1)' : 'scaleX(-1)'
            )
        }

        const handleFlipVertical = () => {
            setFlipScale((prevScale) =>
                prevScale === 'scaleY(-1)' ? 'scaleY(1)' : 'scaleY(-1)'
            )
        }

        const handleResetActionState = () => {
            setRotateDegree(0)
            setScale(1)
            setFlipScale('')
        }

        return (
            <div>
                <ImagePreviewPortal>
                    <ImagePreviewOverlay
                        className={cn(
                            'bg-black/30 dark:bg-background/70',
                            overlayClassName
                        )}
                    />
                    <div
                        ref={ref}
                        className={cn(
                            'flex flex-col-reverse lg:flex-row fixed left-[50%] top-[50%] z-50 h-full w-full translate-x-[-50%] translate-y-[-50%] border shadow-lg backdrop-blur duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%]',
                            className
                        )}
                        onClick={handleOverlayClick}
                        {...props}
                    >
                        <ImagePreviewActions 
                            className='w-full lg:w-fit'
                            downloadImage={downloadImage}
                            handleResetActionState={handleResetActionState}
                            handleRoteteRight={handleRoteteRight}
                            handleRoteteLeft={handleRoteteLeft}
                            handleZoomIn={handleZoomIn}
                            handleZoomOut={handleZoomOut}
                            handleFlipHorizontal={handleFlipHorizontal}
                            handleFlipVertical={handleFlipVertical}
                            hideCloseButton={hideCloseButton}
                            closeButtonClassName={closeButtonClassName}
                        />
                        {Images && Images.length > 1 && (
                            <div className="h-fit flex w-full max-w-full overflow-x-auto space-x-2 overflow-y-hidden items-center p-5 lg:h-full lg:overflow-y-auto
                                            border-r-[.5px] border-background/20 bg-transparent
                                            lg:flex-col lg:w-[10rem] lg:max-w-[10rem] lg:py-5 lg:space-y-2 lg:space-x-0
                                            backdrop-blur duration-100 ease-in-out dark:border-slate-400/20 ">
                                {Images.map((data, index) => {
                                    return (
                                        <Card
                                            onClick={() => scrollToIndex(index)}
                                            className={cn(
                                                `aspect-square flex whitespace-nowrap  size-20 ${focusIndex === index ? 'scale-110' : 'border-none'}`
                                            )}
                                            key={index}
                                            tabIndex={index}
                                           
                                          
                                        >
                                            <img
                                                className="h-full w-full cursor-pointer overflow-hidden rounded-lg object-cover"
                                                src={data.url}
                                            />
                                        </Card>
                                    )
                                })}
                            </div>
                        )}
                        <div
                            id="overlay"
                            className="flex h-full px-10 lg:px-0 w-full items-center justify-center"
                        >
                            <Carousel
                                opts={options}
                                setApi={setApi}
                                className="flex h-fit w-full max-w-4xl bg-transparent"
                            >
                                <CarouselContent className="">
                                    {/* {children} */}
                                    {Images.map((data, index) => {
                                        return (
                                            <CarouselItem
                                                className="flex items-center justify-center"
                                                key={index}
                                            >
                                                <ImageZoomContainer
                                                    flipScale={flipScale}
                                                    rotateDegree={rotateDegree}
                                                    scale={scale}
                                                    url={data.url}
                                                ></ImageZoomContainer>
                                            </CarouselItem>
                                        )
                                    })}
                                </CarouselContent>
                                <CarouselPrevious
                                    className={`${isMultipleImage ? '' : 'hidden'} border-0`}
                                />
                                <CarouselNext
                                    className={`${isMultipleImage ? '' : 'hidden'} border-0`}
                                />
                            </Carousel>
                        </div>
                    </div>
                </ImagePreviewPortal>
            </div>
        )
    }
)
ImagePreviewContent.displayName = ImagePreviewPrimitive.Content.displayName

export { ImagePreview, ImagePreviewContent }
