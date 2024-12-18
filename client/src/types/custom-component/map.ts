
import { SetStateAction, Dispatch } from 'react'
import L, { LatLngExpression } from 'leaflet'
import { MapContainerProps } from 'react-leaflet'

interface TSearchableProps {
    searchedAddress?: string
    setSearchedAddress?: Dispatch<SetStateAction<string>>
}

export interface Pin {
    id: number
    position: LatLngExpression
}

export interface TCustomSearchProps extends TSearchableProps {
    onLocationFound: (latLng: LatLngExpression) => void
    map: L.Map | null
}

export interface TLatLngExpressionWithDesc {
    lat: string
    lng: string
    desc: string
}

export interface TMainMapProps extends Partial<MapContainerProps> {
    disabledSearch?: boolean
    multiplePins?: boolean
    viewOnly?: boolean
    markerPosition?: { x: number; y: number }
    onCoordinatesChange?: (coords: Pin[]) => void;
    disabledCoordinatesView?: boolean,

}

export interface TMapProps {
    handleMapCreated?: (map: L.Map) => void
}

export interface TMapWithClickProps {
    onLocationFound: (latLng: LatLngExpression) => void
}