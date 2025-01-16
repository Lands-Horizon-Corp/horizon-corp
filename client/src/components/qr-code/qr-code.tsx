import { forwardRef } from 'react'
import { QRCodeSVG } from 'qrcode.react'

import { QrCodeIcon } from '@/components/icons'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types'

export interface IQrCodeProps extends IBaseCompNoChild {
    value: string
    themeResponsive?: boolean
}

const QrCode = forwardRef<HTMLDivElement, IQrCodeProps>(
    ({ value, className }, qrRef) => {
        return (
            <div
                ref={qrRef}
                className={cn(
                    'size-48 rounded-xl bg-white p-5',
                    className
                )}
            >
                {value.length === 0 ? (
                    <div className="flex flex-col items-center gap-y-4 text-gray-700/70">
                        <QrCodeIcon className="size-36" />
                        <p className="text-center text-sm lg:text-lg">
                            QR No Content
                        </p>
                    </div>
                ) : (
                    <QRCodeSVG
                        level={'M'}
                        imageSettings={{
                            src: '/favicon.ico',
                            x: undefined,
                            y: undefined,
                            height: 24,
                            width: 24,
                            opacity: 1,
                            excavate: true,
                        }}
                        value={value}
                        className="size-full rounded-sm duration-300"
                    />
                )}
            </div>
        )
    }
)

export default QrCode
