import React from 'react'

interface ReportWrapperProps {
    title: string
    subtitle?: string
    header?: React.ReactNode
    table: React.ReactNode
    footer?: React.ReactNode
    footerTable?: React.ReactNode
}

const ReportWrapper: React.FC<ReportWrapperProps> = ({
    title,
    subtitle,
    header,
    table,
    footer,
    footerTable
}) => {

    return (
        <div className="report-wrapper">
            <header>
                <h3 className="text-xl font-bold">{title}</h3>
                {subtitle && <h4>{subtitle}</h4>}
            </header>
            {/* Additional "header" content, e.g. filters or metadata */}
            {header && <section className="report-subheader">{header}</section>}
            <main className="py-2">{table}</main>
            {footerTable && footerTable}
            {/* {footer && <footer>{footer}</footer>} */}
        </div>
    )
}

export default ReportWrapper
