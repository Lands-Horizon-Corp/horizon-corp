import { MapContainerProps } from 'react-leaflet'
import { create } from 'zustand'

import { Pin } from '@/types/custom-component'

interface MapStoreState {
    markerPosition?: { x: number; y: number }
    disabledSearch?: boolean
    Markers?: Pin[]
}

interface MapStoreActions {
    setMarkerPosition: (nextPosition: MapStoreState['markerPosition']) => void
    setMarkers: (newPin: Pin[]) => void
}

interface MapStoreProps
    extends MapStoreState,
        MapStoreActions,
        MapContainerProps {}

export const useMapStore = create<MapStoreProps>()((set) => ({
    setMarkerPosition: (nextPosition) =>
        set((state) => ({
            ...state,
            markerPosition: nextPosition,
        })),
    disabledSearch: false,
    setMarkers: (newPin) =>
        set((state) => ({
            ...state,
            Markers: newPin,
        })),
}))
