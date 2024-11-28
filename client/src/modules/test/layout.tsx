// import MainMapContainer from '@/components/map'
import { Outlet } from '@tanstack/react-router'
// import { useMapStore } from '@/store/map-store'
// import { LatLngExpression } from 'leaflet'

const TestLayout = () => {
    // const defaultCenter: LatLngExpression = [14.5995, 120.9842]
    // const defaultZoom = 13

    // if wanted to use current Position or List of Markers
    // const position = useMapStore((state) => state.markerPosition)
    // const markers = useMapStore((state) => state.Markers)
    // console.log(position, markers)

    return (
        <div className="mx-auto h-[100vh] w-[80%]">
            {/* <MainMapContainer
                center={defaultCenter}
                zoom={defaultZoom}
                multiplePins={true}
                // viewOnly={true}
            /> */}
            <Outlet />
        </div>
    )
}

export default TestLayout
