import {
    columns,
    FinancialRecords,
    tellerReportConfig,
} from '@/components/reports/config/user-config'
import ReportGenerator from '@/components/reports/report-generator'

const UserTeller = () => {
    // const { setConfig, setData, setColumns } = useTellerTransactionStore(
    //     (state) => state
    // )

    // const actionComponent = () => {
    //     return <div></div>
    // }

    // const columns = useMemo(
    //     () =>
    //         userReportsTableColumns({
    //             actionComponent,
    //         }),
    //     [actionComponent]
    // )

    // const { navigate } = useRouter()

    // useEffect(() => {
    //     setData(FinancialRecords)
    //     setConfig(tellerReportConfig)
    //     setColumns(columns)
    // }, [])

    // const goToTablePage = () => {
    //     navigate({ to: '/test/document' })
    // }

    // const { columnOrder, setColumnOrder } =
    //     useDataTableState<TellerTransactionRecord>({
    //         columnOrder: columns.map((c) => c.id!),
    //     })

    // const table = useReactTable({
    //     columns: columns,
    //     data: FinancialRecords,
    //     getCoreRowModel: getCoreRowModel(),
    //     state: {
    //         columnOrder,
    //     },
    // })

    const TellerReportConfig = {
        ...tellerReportConfig,
        data: FinancialRecords,
        columns: columns,
        footerColumns: columns,
        footerData: FinancialRecords
    }

    return <ReportGenerator reportConfig={TellerReportConfig} />
}

export default UserTeller
