import { SetStateAction, Dispatch } from 'react'
import L, { LatLngExpression } from 'leaflet'

interface TLocationProps {
    defaultCenter: LatLngExpression
    defaultZoom?: number
}

interface TSearchableProps {
    searchedAddress?: string
    setSearchedAddress: Dispatch<SetStateAction<string>>
}

export interface TCustomSearchProps extends TSearchableProps {
    onLocationFound: (latLng: LatLngExpression) => void
    map: L.Map | null
}

export interface TMainMapProps
    extends TLocationProps,
        Partial<TSearchableProps> {
    handleMapCreated?: (map: L.Map) => void
    map?: L.Map | null
    disabledSearch?: boolean
    className?: string
    style?: string
}

export interface TLatLngExpressionWithDesc {
    lat: string
    lng: string
    desc: string
}

export interface TMapProps extends TLocationProps, TSearchableProps {
    handleMapCreated?: (map: L.Map) => void
    map?: L.Map | null
    className?: string
}

export interface TMapWithClickProps extends TSearchableProps {
    onLocationFound: (latLng: LatLngExpression) => void
    map: L.Map
}
