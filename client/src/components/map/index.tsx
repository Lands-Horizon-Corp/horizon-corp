import 'leaflet/dist/leaflet.css'
import axios from 'axios'
import { useState, useEffect, useCallback, useRef } from 'react'
import {
    MapContainer,
    TileLayer,
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
    Pin,
} from '@/types/custom-component'

import { LoadingCircleIcon, VscLocationIcon } from '../icons'
import { Button } from '../ui/button'

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

const MapWithClick = ({ onLocationFound }: TMapWithClickProps) => {
    useMapEvent('click', async (e) => {
        onLocationFound( e.latlng)
    })
    return null
}

const CustomSearch = ({ onLocationFound }: TCustomSearchProps) => {
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

    const handleOnLocationFound = (lat: number, lng: number) => {
        const latLng = new L.LatLng(lat, lng)
        onLocationFound(latLng)
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
                        {results.length > 0 && (
                            <>
                                {results.map((location, index) => (
                                    <div
                                        key={index}
                                        className="cursor-pointer hover:rounded-lg hover:bg-slate-200/40 focus:rounded-lg focus:bg-slate-200/40 focus:outline-none focus:ring-0"
                                        onClick={() =>
                                            handleOnLocationFound(
                                                parseFloat(location.lat),
                                                parseFloat(location.lng)
                                            )
                                        }
                                        tabIndex={0}
                                        onKeyDown={(e) => {
                                            if (
                                                e.key === 'Enter' ||
                                                e.key == ' '
                                            ) {
                                                handleOnLocationFound(
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
                            </>
                        )}
                    </div>
                )}
            </div>
        </div>
    )
}

const Maps = ({ handleMapCreated }: TMapProps) => {
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
        </>
    )
}

const MainMapContainer = ({
    disabledSearch = false,
    center,
    zoom = 13,
    zoomControl = false,
    className = '',
    style,
    scrollWheelZoom = true,
    minZoom,
    maxZoom,
    whenReady,
    children,
    multiplePins = false,
}: TMainMapProps) => {
    const [_, setSearchedAddress] = useState('')
    const [selectedPins, setSelectedPins] = useState<Pin[]>([])
    const [map, setMap] = useState<L.Map | null>(null)
    const markerRefs = useRef<{ [key: string]: L.Marker }>({})

    const handleMapReady = (mapInstance: L.Map) => {
        setMap(mapInstance)
    }

    const deletePin = useCallback(
        (id: number) => {
            
            setSelectedPins((pins) => pins.filter((pin) => pin.id !== id))
            const pin = selectedPins.find((pin) => pin.id === id)

            if (pin) {
                const { lat, lng } = pin.position as L.LatLngLiteral
                const markerKey = `${lat},${lng}`
                markerRefs.current[markerKey]?.remove()
                delete markerRefs.current[markerKey]
            }
        },
        [selectedPins]
    )

    const addMarker = useCallback(
        async (latLng: LatLngExpression) => {
            const { lat, lng } = latLng as L.LatLngLiteral
            const markerKey = `${lat},${lng}`

            if (markerRefs.current[markerKey])
                markerRefs.current[markerKey].remove()

            const marker = L.marker(latLng).addTo(map as L.Map)
            markerRefs.current[markerKey] = marker
            const address = await getLocationDescription(latLng)
            marker.bindPopup(address).openPopup()
        },
        [map]
    )

    const handleLocationFound = useCallback(
        async (latLng: LatLngExpression) => {
            const newPin: Pin = { id: Date.now(), position: latLng }

            setSelectedPins((prevPins) =>
                multiplePins ? [...prevPins, newPin] : [newPin]
            )

            if (!multiplePins && map) {
                map.eachLayer((layer: any) => {
                    if (layer instanceof L.Marker) {
                        map.removeLayer(layer)
                    }
                })
            }

            addMarker(latLng)
        },
        [addMarker, map, multiplePins]
    )

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
                center={center}
                zoom={zoom}
                className={`-z-0 h-[500px] w-full min-w-[500px] rounded-lg ${className ?? ''}`}
                ref={setMap}
                style={style}
                zoomControl={zoomControl}
                scrollWheelZoom={scrollWheelZoom}
                minZoom={minZoom}
                maxZoom={maxZoom}
                whenReady={whenReady}
            >
                <Maps handleMapCreated={handleMapReady} />
                <MapWithClick onLocationFound={handleLocationFound} />
                <ZoomControl position="bottomright" />
                {children}
            </MapContainer>
            <ul className="mt-4">
                {selectedPins.map((pin) => {
                    const { lat, lng } = pin.position as L.LatLngLiteral
                    return (
                        <div key={pin.id}>
                            <li>
                                Latitude: {lat}, Longitude: {lng}
                            </li>
                            <Button
                                onClick={() => deletePin(pin.id)}
                                variant="secondary"
                            >
                                Delete
                            </Button>
                        </div>
                    )
                })}
            </ul>
        </div>
    )
}

export default MainMapContainer
