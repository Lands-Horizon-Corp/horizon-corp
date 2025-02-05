import { useState } from 'react'
import { reportRoutes, TReportRoute } from './cards-data'
import ReportDialog from '@/components/reports/report-dialog'

const ReportsDashboard = () => {
    const [selectedReport, setSelectedReport] = useState<TReportRoute | null>(
        null
    )
    return (
        <div className="grid grid-cols-1 gap-4 p-10 md:grid-cols-2 lg:grid-cols-3">
            {reportRoutes.map((route: TReportRoute) => {
                const { icon: Icon } = route
                return (
                    <div
                        key={route.path}
                        onClick={() => setSelectedReport(route)}
                        className="flex cursor-pointer items-center gap-3 rounded-lg border p-4 shadow-md transition hover:bg-gray-100"
                    >
                        <Icon className="text-blue text-2xl" />
                        <div>
                            <h3 className="text-lg font-semibold">
                                {route.title}
                            </h3>
                            <p className="text-sm text-gray-600">
                                {route.description}
                            </p>
                        </div>
                    </div>
                )
            })}
            {selectedReport && (
                <ReportDialog
                    {...selectedReport}
                    onClose={() => setSelectedReport(null)}
                />
            )}
        </div>
    )
}

export default ReportsDashboard
