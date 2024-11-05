<<<<<<< HEAD
import { Outlet } from '@tanstack/react-router'
import { ImagePreview, ImagePreviewContent } from '@/components/ui/image-preview'
import { Button } from '@/components/ui/button'
import { sampleMediaResourceList } from './testSampleData'
import { useImagePreview } from '@/store/image-preview-store'
const TestLayout = () => {
    const { isOpen, setIsOpen } = useImagePreview()

    return (
        <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
            <main>
                <Outlet />
            </main>

        <div className="mx-auto h-[100vh] w-[80%]">
            <Button onClick={()=> setIsOpen(true)}>view</Button>
            <ImagePreview open={isOpen} onOpenChange={()=> setIsOpen(false)}>
                <ImagePreviewContent Images={sampleMediaResourceList} />
            </ImagePreview>
        </div>
=======
import MainMapContainer from '@/components/map'
// import { useMapStore } from '@/store/map-store'
import { LatLngExpression } from 'leaflet'

const TestLayout = () => {
    const defaultCenter: LatLngExpression = [14.5995, 120.9842]
    const defaultZoom = 13

    // if wanted to use current Position or List of Markers
    // const position = useMapStore((state) => state.markerPosition)
    // const markers = useMapStore((state) => state.Markers)
    // console.log(position, markers)

    return (

        <div className="mx-auto h-[100vh] w-[80%]">
            <MainMapContainer
                center={defaultCenter}
                zoom={defaultZoom}
                multiplePins={true}
                // viewOnly={true}
            />
>>>>>>> main
        </div>
    )
}

export default TestLayout
