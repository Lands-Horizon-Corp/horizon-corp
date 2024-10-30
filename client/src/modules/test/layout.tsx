import { Outlet } from '@tanstack/react-router'
// import MainMapContainer from '@/components/map'
import { ImagePreview, ImagePreviewContent } from '@/components/ui/image-preview'
// import { useMapStore } from '@/store/map-store'
import { useState } from 'react'
import { Button } from '@/components/ui/button'
import { sampleMediaResourceList } from '@/components/image-preview/sampleImageData'
const TestLayout = () => {
    const [isOpen, setIsOpen] = useState(false)

    return (
        <div className="grid min-h-[100dvh] bg-red-500 grid-cols-[auto_1fr]">
            <main>
                <Outlet />
            </main>

        <div className="mx-auto h-[100vh] w-[80%]">
            {/* <MainMapContainer
                center={defaultCenter}
                zoom={defaultZoom}
                // multiplePins={true}
                // viewOnly={true}
            /> */}
            <Button onClick={()=> setIsOpen(true)}></Button>
            <ImagePreview open={isOpen} onOpenChange={()=> setIsOpen(false)}>
                <ImagePreviewContent Images={sampleMediaResourceList} />
            </ImagePreview>
        </div>
        </div>
    )
}

export default TestLayout
