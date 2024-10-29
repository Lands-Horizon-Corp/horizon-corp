import MainMapContainer from '@/components/map'
import { LatLngExpression } from 'leaflet'

const TestLayout = () => {
    const defaultCenter: LatLngExpression = [14.5995, 120.9842]
    const defaultZoom = 13

    return (
        <div className="mx-auto h-[100vh] w-[80%]">
            <MainMapContainer
                center={defaultCenter}
                zoom={defaultZoom}
                // multiplePins={true}
            />
        </div>
    )
}

export default TestLayout
