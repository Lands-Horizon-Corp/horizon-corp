import ReportWrapper from './report-wrapper'
import ReportHeader from './report-header'
import ReportsDataTable from './reports-data-table'
import { ColumnDefinition } from './types'
import ReportFooter from './report-footer'

export interface ReportGeneratorConfig<TData> {
    title: string
    subtitle?: string
    filters?: Record<string, string>
    footer?: { totals: { label: string; value: number } }
    columns: ColumnDefinition<TData>[]
    data: TData[]
    footerData: TData[]
    footerColumns: ColumnDefinition<TData>[]
}

export interface ReportGeneratorProps<TData> {
    reportConfig: ReportGeneratorConfig<TData>
}

const ReportGenerator = <TData,>(props: ReportGeneratorProps<TData>) => {
    
    const { reportConfig } = props

    const {
        title,
        subtitle,
        filters,
        footer,
        data,
        columns,
        footerColumns,
        footerData,
    } = reportConfig
   
    return (
        <ReportWrapper
            title={title}
            subtitle={subtitle}
            header={filters ? <ReportHeader metadata={filters} /> : null}
            table={<ReportsDataTable<TData> columns={columns} data={data} />}
            footer={
                footerData && footer ? (
                    <ReportFooter<TData>
                        totals={footer.totals}
                        footerColumns={footerColumns}
                        footerData={footerData}
                    />
                ) : null
            }
        />
    )
}

export default ReportGenerator
