import { useEffect, useState, useCallback } from 'react'

export const useCameraDevices = () => {
    const [devices, setDevices] = useState<MediaDeviceInfo[]>([])

    const handleDevices = useCallback(
        (mediaDevices: MediaDeviceInfo[]) =>
            setDevices(
                mediaDevices
                    .filter(
                        ({ kind, deviceId }) =>
                            kind === 'videoinput' && deviceId.length !== 0
                    )
                    .map((cam, i) => ({
                        ...cam,
                        label: cam.label
                            ? cam.label.split(' (')[0]
                            : `Camera ${i + 1}`,
                    }))
            ),
        [setDevices]
    )

    useEffect(() => {
        navigator.mediaDevices.enumerateDevices().then(handleDevices)
    }, [handleDevices])

    return { devices }
}
