// import { columns, FinancialRecords, tellerReportConfig } from '@/components/reports/config/user-config'
import DataTable from '@/components/data-table'
import { TellerTransactionRecord } from '@/components/reports/types'
import { getCoreRowModel, useReactTable } from '@tanstack/react-table'
import { useMemo } from 'react'
import { userReportsTableColumns } from './columns'
import { FinancialRecords } from '@/components/reports/config/user-config'
import useDataTableState from '@/components/data-table/hooks/use-datatable-state'

const UserTellerTable = () => {

    // const { setConfig, setData, setColumns } = useTellerTransactionStore(
    //     (state) => state
    // )

    const actionComponent = () => {
        return <div></div>
    }

    const columns = useMemo(
        () =>
            userReportsTableColumns({
                actionComponent,
            }),
        [actionComponent]
    )

    const { columnOrder, setColumnOrder } =
        useDataTableState<TellerTransactionRecord>({
            columnOrder: columns.map((c) => c.id!),
        })

    const table = useReactTable({
        columns: columns,
        data: FinancialRecords,
        getCoreRowModel: getCoreRowModel(),
        state: {
            columnOrder,
        },
    })

  
    return (
        <div className="flex w-full max-w-full flex-col px-4 pb-6 sm:px-8">
            <DataTable
                table={table}
                setColumnOrder={setColumnOrder}
                isScrollable
            />
        </div>
    )
}

export default UserTellerTable
