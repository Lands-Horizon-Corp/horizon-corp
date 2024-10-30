import MainMapContainer from '@/components/map'
import { ImagePreview, ImagePreviewContent } from '@/components/ui/image-preview'
// import { useMapStore } from '@/store/map-store'
import { LatLngExpression } from 'leaflet'

const TestLayout = () => {
    return (

        <div className="mx-auto h-[100vh] w-[80%]">
            {/* <MainMapContainer
                center={defaultCenter}
                zoom={defaultZoom}
                // multiplePins={true}
                // viewOnly={true}
            /> */}
            <ImagePreview open={true}>
                <ImagePreviewContent>Hello</ImagePreviewContent>
            </ImagePreview>
        </div>
    )
}

export default TestLayout
