 import 'leaflet/dist/leaflet.css'
import axios from 'axios'
import { useState, useEffect, useCallback } from 'react'
import {
    MapContainer,
    TileLayer,
    Marker,
    useMapEvent,
    useMap,
    ZoomControl,
} from 'react-leaflet'
import { useMutation } from '@tanstack/react-query'
import debounce from 'lodash-es/debounce'
import L, { LatLngExpression, latLng } from 'leaflet'

import { Input } from '../ui/input'
import LayerControl from './LayerControl'
import {
    TCustomSearchProps,
    TMainMapProps,
    TLatLngExpressionWithDesc,
    TMapProps,
    TMapWithClickProps,
} from '@/types/custom-component'

import { LoadingCircleIcon, VscLocationIcon } from '../icons'

const getLocationDescription = async (latlng: LatLngExpression) => {
    const { lat, lng } = latLng(latlng)
    const reverseGeocodeUrl = `https://nominatim.openstreetmap.org/reverse?lat=${lat}&lon=${lng}&format=json`
    try {
        const response = await fetch(reverseGeocodeUrl)
        const data = await response.json()
        return data.display_name || 'Address not found'
    } catch (error) {
        console.error('Error fetching reverse geocode data:', error)
    }
}

const MapWithClick = ({
    onLocationFound,
    map,
    setSearchedAddress,
}: TMapWithClickProps) => {
    useMapEvent('click', async (e) => {
        const latLng = e.latlng
        onLocationFound(latLng)

        map?.eachLayer((layer: any) => {
            if (layer instanceof L.Marker) {
                map.removeLayer(layer)
            }
        })
        const marker = L.marker([latLng.lat, latLng.lng]).addTo(map as L.Map)
        const address = await getLocationDescription(latLng)
        setSearchedAddress(address)
        marker.bindPopup(address || 'No address found').openPopup()
    })

    return null
}

// Custom search component
const CustomSearch = ({
    onLocationFound,
    map,
    setSearchedAddress,
}: TCustomSearchProps) => {
    const [results, setResults] = useState<TLatLngExpressionWithDesc[]>([])
    const [query, setQuery] = useState('')

    const { mutate: handleSearchMutation, isPending } = useMutation<
        void,
        unknown,
        string
    >({
        mutationKey: ['handleSearch'],
        mutationFn: async (query: string) => {
            if (!query) return
            const searchLocation = await axios.get(
                `https://nominatim.openstreetmap.org/search?q=${query}&format=json&addressdetails=1&limit=5`
            )
            const locations: TLatLngExpressionWithDesc[] =
                searchLocation.data.map((result: any) => ({
                    lat: result.lat,
                    lng: result.lon,
                    desc: result.display_name,
                }))
            setResults(locations)
        },
    })

    const debouncedSearch = useCallback(
        debounce((query: string) => handleSearchMutation(query), 500),
        [handleSearchMutation]
    )

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value ?? ''
        if (value === '') {
            setResults([])
        }
        setQuery(value)
        debouncedSearch(value)
    }

    const handleLocationClick = async (lat: number, lng: number) => {
        if (map) {
            map?.eachLayer((layer) => {
                if (layer instanceof L.Marker) {
                    map.removeLayer(layer)
                }
            })
            const marker = L.marker([lat, lng]).addTo(map as L.Map)
            const location = latLng(lat, lng)
            const address = await getLocationDescription(location)
            setSearchedAddress(address)
            marker.bindPopup(address || 'No address found').openPopup()
        }
        onLocationFound([lat, lng])
        map?.setView([lat, lng], 13)
    }

    const isResultEmpty = results.length === 0

    return (
        <div className="absolute top-5 z-[100] w-[30rem]">
            <Input
                type="text"
                value={query}
                onChange={handleInputChange}
                placeholder="Search Google Maps"
                className="rounded-lg border border-gray-300 px-4 py-2 focus:border-none focus:outline-none"
            />
            <div
                className={`mt-1 flex flex-col space-y-2 bg-white/90 dark:bg-secondary/90 dark:text-white ${isResultEmpty ? '' : 'p-5'} rounded-lg`}
            >
                {isPending ? (
                    <div className="flex w-full justify-center">
                        <LoadingCircleIcon className="animate-spin" />
                    </div>
                ) : (
                    <div>
                        {results.map((location, index) => (
                            <div
                                key={index}
                                className="cursor-pointer hover:rounded-lg hover:bg-slate-200/40 focus:rounded-lg focus:bg-slate-200/40 focus:outline-none focus:ring-0"
                                onClick={() =>
                                    handleLocationClick(
                                        parseFloat(location.lat),
                                        parseFloat(location.lng)
                                    )
                                }
                                tabIndex={0}
                                onKeyDown={(e) => {
                                    if (e.key === 'Enter' || e.key == ' ') {
                                        handleLocationClick(
                                            parseFloat(location.lat),
                                            parseFloat(location.lng)
                                        )
                                    }
                                }}
                            >
                                <div className="flex p-2">
                                    <div className="w-9">
                                        <VscLocationIcon className="size-6 text-slate-600 dark:text-destructive-foreground" />
                                    </div>
                                    <p className="truncate text-sm">
                                        {location.desc}
                                    </p>
                                </div>
                            </div>
                        ))}
                    </div>
                )}
            </div>
        </div>
    )
}

const Maps = ({
    defaultCenter,
    searchedAddress,
    handleMapCreated,
    setSearchedAddress,
}: TMapProps) => {
    const [position, setPosition] = useState<LatLngExpression>(defaultCenter)
    const map = useMap()

    useEffect(() => {
        if (handleMapCreated) {
            handleMapCreated(map)
        }
    }, [handleMapCreated, map])

    return (
        <>
            <TileLayer
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                attribution="&copy; OpenStreetMap contributors"
            />
            <LayerControl />
            <MapWithClick
                map={map}
                setSearchedAddress={setSearchedAddress}
                searchedAddress={searchedAddress}
                onLocationFound={setPosition}
            />
            <Marker position={position}>
                <p>{searchedAddress}</p>
            </Marker>
        </>
    )
}

const MainMapContainer = ({
    defaultCenter,
    defaultZoom = 13,
    disabledSearch = false,
    className,
    style,
}: TMainMapProps) => {
    const [position, setPosition] = useState<LatLngExpression>(defaultCenter)
    const [searchedAddress, setSearchedAddress] = useState('')
    const [map, setMap] = useState<L.Map | null>(null)

    const handleLocationFound = (latLng: LatLngExpression) =>
        setPosition(latLng)

    const handleMapReady = (mapInstance: L.Map) => {
        setMap(mapInstance)
    }

    return (
        <div className={`relative w-full p-5 pt-20 shadow-sm ${style ?? ''}`}>
            {!disabledSearch && (
                <CustomSearch
                    map={map}
                    setSearchedAddress={setSearchedAddress}
                    onLocationFound={handleLocationFound}
                />
            )}
            <MapContainer
                center={position}
                zoom={defaultZoom}
                className={`-z-0 h-[500px] w-full min-w-[500px] rounded-lg ${className ?? ''}`}
                ref={setMap}
                zoomControl={false}
            >
                <Maps
                    map={map}
                    setSearchedAddress={setSearchedAddress}
                    searchedAddress={searchedAddress}
                    defaultCenter={defaultCenter}
                    handleMapCreated={handleMapReady}
                />
                <ZoomControl position="bottomright" />
            </MapContainer>
        </div>
    )
}

export default MainMapContainer
