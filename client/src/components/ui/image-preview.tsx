import * as React from 'react'
import * as ImagePreviewPrimitive from '@radix-ui/react-dialog'
import { X } from 'lucide-react'
import {
    Carousel,
    CarouselContent,
    CarouselNext,
    CarouselPrevious,
    CarouselItem,
} from '@/components/ui/carousel'
import { cn } from '@/lib/utils'
import { MediaResource } from '@/horizon-corp/types'
import { Card } from './card'
import { type CarouselApi } from "@/components/ui/carousel"
import { useState } from 'react'
import { Button } from './button'

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
    extends React.ComponentPropsWithoutRef<
        typeof ImagePreviewPrimitive.Content
    > {
    hideCloseButton?: boolean
    closeButtonClassName?: string
    overlayClassName?: string
    Images: MediaResource[]
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
        const isMultipleImage = Images.length > 0

        const [api, setApi] = useState<CarouselApi>()
        const [selectedScrollSnap, setSelectedScrollSnap] = useState<number>()
        
        React.useEffect(() => {
          if (!api) {
            return
          }
       
          api.on("select", () => {
                    const selectedIndex = api.selectedScrollSnap();
                    setSelectedScrollSnap(selectedIndex)

                    console.log(`Selected slide index: ${selectedIndex}`);
                    const slidesInView = api.slidesInView();
                    console.log("Slides in view: ", slidesInView);
                    
          })
        }, [api])

        const scrollToIndex = (index: number) => {
          if (api) {
            api.scrollTo(index);
            setSelectedScrollSnap(api.selectedScrollSnap())
          }
        };

        return (
         <div>
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
                        'fixed left-[50%] top-[50%] z-50 flex h-full px-20 w-full translate-x-[-50%] translate-y-[-50%] items-center justify-center gap-4 shadow-lg backdrop-blur duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%]',
                        className
                    )}
                    {...props}
                >   
                    <div className=' w-[20%] bg-secondary p-5 overflow-y-auto space-y-2 left-0 h-full'>
                    {Images.map((data, index) => (
                                <Card onClick={()=> scrollToIndex(index)} className={cn(`border h-fit ${selectedScrollSnap === index ? "border-primary":"border-none"}`)} key={index}>
                                    <img  className=" rounded-lg" src={data.url}
                                    ></img>
                                </Card>
                    ))}
                    </div>
                    <div className='w-[70%] h-full flex items-center justify-center'>
                    <Carousel  setApi={setApi} className="flex h-fit w-full max-w-2xl bg-transparent">
                        <CarouselContent className="">
                            {/* {children} */}
                            {Images.map((data, index) => (
                                <CarouselItem className='flex justify-center items-center' key={index}>
                                    <img className="" src={data.url}
                                    
                                    ></img>
                                </CarouselItem>
                            ))}
                        </CarouselContent>
                        <CarouselPrevious
                    //         onClick={handleFocusTo}
                            className={`${isMultipleImage ? '' : 'hidden'} border-0`}
                        />
                        <CarouselNext
                            className={`${isMultipleImage ? '' : 'hidden'} border-0`}
                        />
                       
                    </Carousel>
                    </div>
                    {!hideCloseButton && (
                    <ImagePreviewPrimitive.Close
                        className={cn(
                            'absolute right-4 top-4 z-[1000] size-4 cursor-pointer rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100  disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground',
                            closeButtonClassName
                        )}
                    >
                        <X className="size-full" />
                        <span className="sr-only">Close</span>
                    </ImagePreviewPrimitive.Close>
                )}
                </ImagePreviewPrimitive.Content> 
            </ImagePreviewPortal>
         </div>
        )
    }
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
