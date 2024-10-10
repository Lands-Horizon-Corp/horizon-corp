import { toast } from 'sonner'
import { useEffect, useState } from 'react'

export const useCameraGrant = () => {
    const [granted, setGranted] = useState(false)

    const checkPermissions = async () => {
        try {
            const stream = await navigator.mediaDevices.getUserMedia({
                video: true,
                audio: false,
            })
            stream.getTracks().forEach((track) => track.stop())
            setGranted(true)
        } catch (error) {
            setGranted(false)
            if (
                error instanceof DOMException &&
                error.name === 'NotAllowedError'
            ) {
                toast.warning(
                    'Camera access has been denied. Please check your browser settings and allow camera access to use this feature.'
                )
            } else {
                toast.warning(
                    'An unexpected error occurred while accessing the camera. Please try again later.'
                )
            }
        }
    }

    useEffect(() => {
        checkPermissions()
    }, [])

    return granted
}
