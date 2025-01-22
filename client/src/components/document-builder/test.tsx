import React from 'react'

interface ReportWrapperProps {
    title: string
    subtitle?: string
    filters?: React.ReactNode
    header: React.ReactNode
    table: React.ReactNode
    footer?: React.ReactNode
}

const ReportWrapper: React.FC<ReportWrapperProps> = ({
    title,
    subtitle,
    filters,
    header,
    table,
    footer,
}) => {
    return (
        <div className="report-wrapper">
            <header>
                <h1>{title}</h1>
                {subtitle && <h2>{subtitle}</h2>}
                {filters && <div className="filters">{filters}</div>}
            </header>
            <section>{header}</section>
            <main>{table}</main>
            {footer && <footer>{footer}</footer>}
        </div>
    )
}

interface ReportHeaderProps {
    metadata: { [key: string]: string }
}

const ReportHeader: React.FC<ReportHeaderProps> = ({ metadata }) => {
    return (
        <div className="report-header">
            {Object.entries(metadata).map(([key, value]) => (
                <div key={key} className="metadata">
                    <strong>{key}:</strong> {value}
                </div>
            ))}
        </div>
    )
}

interface ColumnDefinition {
    label: string
    key: string
    align?: 'left' | 'right' | 'center'
    width?: string
}

interface ReportTableProps {
    columns: ColumnDefinition[]
    data: { [key: string]: any }[]
}

const ReportTable: React.FC<ReportTableProps> = ({ columns, data }) => {
    return (
        <table className="report-table">
            <thead>
                <tr>
                    {columns.map((column) => (
                        <th
                            key={column.key}
                            style={{
                                textAlign: column.align,
                                width: column.width,
                            }}
                        >
                            {column.label}
                        </th>
                    ))}
                </tr>
            </thead>
            <tbody>
                {data.map((row, rowIndex) => (
                    <tr key={rowIndex}>
                        {columns.map((column) => (
                            <td
                                key={column.key}
                                style={{ textAlign: column.align }}
                            >
                                {row[column.key]}
                            </td>
                        ))}
                    </tr>
                ))}
            </tbody>
        </table>
    )
}

interface ReportFooterProps {
    totals: { label: string; value: number }
}

const ReportFooter: React.FC<ReportFooterProps> = ({ totals }) => {
    return (
        <div className="report-footer">
            <div>
                <strong>{totals.label}:</strong> {totals.value.toLocaleString()}
            </div>
        </div>
    )
}

const ReportGenerator: React.FC<{ reportConfig: any }> = ({ reportConfig }) => {
    const { title, subtitle, filters, columns, data, footer } = reportConfig

    return (
        <ReportWrapper
            title={title}
            subtitle={subtitle}
            filters={<ReportHeader metadata={filters} />}
            header={<ReportHeader metadata={filters} />}
            table={<ReportTable columns={columns} data={data} />}
            footer={<ReportFooter totals={footer.totals} />}
        />
    )
}

const reportConfig = {
    title: 'Sales Report',
    subtitle: 'January 2025',
    filters: {
        Region: 'North America',
        'Date Range': '01/01/2025 - 01/31/2025',
    },
    columns: [
        { label: 'ID', key: 'id', width: '10%' },
        { label: 'Product', key: 'product', width: '30%' },
        { label: 'Quantity', key: 'quantity', align: 'right' },
        { label: 'Price', key: 'price', align: 'right' },
        { label: 'Total', key: 'total', align: 'right' },
    ],
    data: [
        { id: 1, product: 'Laptop', quantity: 2, price: 1500, total: 3000 },
        { id: 2, product: 'Mouse', quantity: 5, price: 20, total: 100 },
    ],
    footer: { totals: { label: 'Grand Total', value: 3100 } },
}

const Test = () => <ReportGenerator reportConfig={reportConfig} />

export default Test
