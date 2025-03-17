import { IQrScanResult } from '@/types'
import { IScannerProps } from '@yudiel/react-qr-scanner'

// Just scanner component desu (no decoding)
export interface IQrScannerProps extends IScannerProps {}

// Scanne component desu (with decoding support)
export interface IQrCodeScannerProps<T, Err> extends IQrScannerProps {
    disableDecode?: boolean
    pauseOnDecoding?: boolean
    onErrorDecode?: (error: Err) => void
    onSuccessDecode?: (data: IQrScanResult<T>) => void
}
