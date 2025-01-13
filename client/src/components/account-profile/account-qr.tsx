import { useState } from 'react'

import Modal from '@/components/modals/modal'
import { QrCodeIcon } from '@/components/icons'
import { Button } from '@/components/ui/button'
import QrCodeDownloadable from '../qr-code/qr-code-downloadable'
import { QrCode } from '../qr-code'

// interface Props {}

const AccountQr = () => {
    const [toggle, setToggle] = useState(false)

    const stringContent = 'as0f97as09 f7as0f8 sa'

    return (
        <>
            <Button
                size="icon"
                variant="ghost"
                onClick={() => setToggle((val) => !val)}
                className="ml-2 absolute top-5 right-5 inline w-auto h-[40%] sm:h-[70%] bg-transparent text-foreground/60"
            >
                {!stringContent ? (
                    <QrCodeIcon className="size-full" />
                ) : (
                    <QrCode
                        value={stringContent}
                        className="size-full rounded-sm p-0.5"
                    />
                )}
            </Button>
            <Modal
                open={toggle}
                title="Profile QR"
                description="QR of account for easy processing/trasactions."
                onOpenChange={(val) => setToggle(val)}
                className="p-4 pb-8"
            >
                <QrCodeDownloadable
                    fileName="profile-qr"
                    value={stringContent}
                    className="p-3 size-80"
                    containerClassName="mx-auto"
                />
            </Modal>
        </>
    )
}

export default AccountQr
