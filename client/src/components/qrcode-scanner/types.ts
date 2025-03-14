import { IQrScanResult } from '@/types'
import { IScannerProps } from '@yudiel/react-qr-scanner'

// Just scanner component desu (no decoding)
export interface IQrScannerProps extends IScannerProps {}

// Scanne component desu (with decoding support)
export interface IQrCodeScannerBaseProps extends IQrScannerProps {
    disableDecode?: boolean
    pauseOnDecoding?: boolean
}

// Case when disableDecode is true (Callbacks should not exist)
export interface IQrCodeScannerPropsDisabled extends IQrCodeScannerBaseProps {
    disableDecode: true
    onErrorDecode?: never
    onSuccessDecode?: never
}

// Case when disableDecode is false or not provided (Callbacks required)
export interface IQrCodeScannerPropsEnabled<T = unknown, Err = unknown>
    extends IQrCodeScannerBaseProps {
    disableDecode?: false
    onErrorDecode: (error: Err) => void
    onSuccessDecode: (data: IQrScanResult<T>) => void
}

// Final interface (Union of both)
export type IQrCodeScannerProps<T = unknown, Err = unknown> =
    | IQrCodeScannerPropsDisabled
    | IQrCodeScannerPropsEnabled<T, Err>
