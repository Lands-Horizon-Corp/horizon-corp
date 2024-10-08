import { useEffect, useState, useCallback } from 'react'

export const useCameraDevices = () => {
    const [deviceId] = useState({})
    const [devices, setDevices] = useState<MediaDeviceInfo[]>([])

    const handleDevices = useCallback(
        (mediaDevices: MediaDeviceInfo[]) =>
            setDevices(
                mediaDevices.filter(({ kind }) => kind === 'videoinput')
            ),
        [setDevices]
    )

    useEffect(() => {
        navigator.mediaDevices.enumerateDevices().then(handleDevices)
    }, [handleDevices])

    return { devices, deviceId }
}
