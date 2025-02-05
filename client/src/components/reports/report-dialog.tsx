import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
} from '@/components/ui/dialog'
import { useNavigate } from '@tanstack/react-router'
import ReportConfigComponents from './config'
import { TReportRoute } from '@/modules/owner/pages/reports/dashboard/cards-data'

interface ReportDialogProps extends TReportRoute {
    onClose: () => void
}

const ReportDialog = (reportProps: ReportDialogProps) => {
    const { title, description, path: reportKey, onClose } = reportProps

    const navigate = useNavigate()
    const ReportComponent = ReportConfigComponents[reportKey]

    const handleSubmit = (config: any) => {
        console.log('config', config)

        // if (reportKey) {
        //     navigate({ to: `/owner/reports/document/${reportKey}` })
        //     onClose()
        // }
    }

    if (!reportKey) return null

    return (
        <Dialog open onOpenChange={onClose}>
            <DialogContent className="max-h-[42rem] overflow-auto">
                <DialogHeader>
                    <DialogTitle>{title} - Configuration</DialogTitle>
                </DialogHeader>
                <DialogDescription>{description}</DialogDescription>
                <div className="space-y-4">
                    {ReportComponent ? (
                        <ReportComponent
                            onSubmit={handleSubmit}
                            onClose={onClose}
                        />
                    ) : (
                        <p>No configuration available.</p>
                    )}
                </div>
            </DialogContent>
        </Dialog>
    )
}

export default ReportDialog
