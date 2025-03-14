import { toast } from 'sonner'
import { useCallback } from 'react'
import { IDetectedBarcode } from '@yudiel/react-qr-scanner'

import QrScanner from './qr-scanner'

import { IQrScanResult } from '@/types'
import { IQrCodeScannerProps } from './types'
import { useQrDecryptData } from '@/hooks/api-hooks/use-qr-crypto'

const QrCodeScanner = <TData, TErr = string>({
    paused,
    scanDelay = 1000,
    disableDecode = false,
    pauseOnDecoding = false,
    onScan,
    onErrorDecode,
    onSuccessDecode,
    ...props
}: Omit<IQrCodeScannerProps<TData, TErr>, 'onError'>) => {
    const { mutateAsync, isPending } = useQrDecryptData<TData>()

    const handleOnScan = useCallback(
        (results: IDetectedBarcode[]) => {
            if (results.length === 0) return

            if (disableDecode) return onScan(results)

            try {
                const parsedQrContent = JSON.parse(
                    results[0].rawValue
                ) as IQrScanResult

                if (!parsedQrContent.data)
                    return toast.error(
                        'Sorry, QR Code is Not ECoop generated QR'
                    )

                toast.promise(mutateAsync(parsedQrContent.data as string), {
                    loading: 'Verifying QR...',
                    success: (data) => {
                        const finalQrResult: IQrScanResult<TData> = {
                            ...parsedQrContent,
                            data,
                        }
                        onSuccessDecode?.(finalQrResult as IQrScanResult<TData>)
                        return 'QRCode verified'
                    },
                    error: (err) => {
                        onErrorDecode?.(err)
                        return err
                    },
                })

                // eslint-disable-next-line @typescript-eslint/no-unused-vars
            } catch (e) {
                toast.error('Invalid QR')
                onErrorDecode?.('Invalid QR code content' as TErr)
            }
        },
        [disableDecode, mutateAsync, onErrorDecode, onScan, onSuccessDecode]
    )

    return (
        <QrScanner
            onScan={handleOnScan}
            scanDelay={scanDelay}
            paused={paused || pauseOnDecoding ? isPending : undefined}
            {...props}
        />
    )
}

export default QrCodeScanner
