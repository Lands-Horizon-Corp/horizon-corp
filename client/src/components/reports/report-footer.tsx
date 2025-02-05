import ReportsDataTable from "./reports-data-table"
import { ColumnDefinition } from "./types"

interface ReportFooterProps<TData> {
    totals: { label: string; value: number }
    footerColumns:ColumnDefinition<TData>[]
    footerData:TData[]
}

const ReportFooter =<TData,> (props: ReportFooterProps<TData>) => {
    const { totals ,footerColumns, footerData } = props
    console.log(props)
    return (
        <div className="report-footer">
            <ReportsDataTable<TData> columns={footerColumns} data={footerData}   /> 
            <strong>{totals.label}:</strong> {totals.value.toLocaleString()}
        </div>
    )
}

export default ReportFooter
