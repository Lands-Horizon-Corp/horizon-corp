import QrCodeScanner from '@/components/qrcode-scanner'

import logger from '@/helpers/loggers/logger'

const OwnerDashboard = () => {
    return (
        <div>
            <QrCodeScanner
                allowMultiple
                onScan={(results) => {
                    logger.log(results)
                }}
                classNames={{ container: 'size-[400px]' }}
            />
        </div>
    )
}

export default OwnerDashboard
