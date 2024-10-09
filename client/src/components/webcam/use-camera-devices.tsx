import { useEffect, useState, useCallback } from 'react'

export const useCameraDevices = () => {
    const [devices, setDevices] = useState<MediaDeviceInfo[]>([])

    const handleDevices = useCallback(
        (mediaDevices: MediaDeviceInfo[]) => {
            const cameraDevices = mediaDevices.filter(
                ({ kind, deviceId }) =>
                    kind === 'videoinput' && deviceId.length !== 0
            )
            setDevices(cameraDevices)
        },
        [setDevices]
    )

    useEffect(() => {
        navigator.mediaDevices.enumerateDevices().then(handleDevices)
    }, [handleDevices])

    return { devices }
}
