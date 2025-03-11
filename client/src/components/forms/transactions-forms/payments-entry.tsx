import { useMemo, useState } from 'react'
import {
    getCoreRowModel,
    getSortedRowModel,
    useReactTable,
} from '@tanstack/react-table'

import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'

import { XIcon } from '@/components/icons'

import { useFilteredPaginatedIAccountingLedger } from '@/hooks/api-hooks/transactions/use-accounting-ledger'
import { useDataTableSorting } from '@/hooks/data-table-hooks/use-datatable-sorting'
import useDataTableState from '@/hooks/data-table-hooks/use-datatable-state'
import { usePagination } from '@/hooks/use-pagination'
import useDatableFilterState from '@/hooks/use-filter-state'

import { IMemberResource } from '@/server'
import { IAccountingLedgerResource } from '@/server/types/accounts/accounting-ledger'

import MemberPicker from '@/components/pickers/member-picker'

import AccountsLedgerTable from '@/components/tables/transactions-table/accouting-ledger-table'
import IAccountingLedgerRequestTableColumns from '@/components/tables/transactions-table/accouting-ledger-table/columns'

import PaymentsEntryProfile from '@/components/transaction-profile/payments-entry-profile'
import CurrentPaymentAccountingTransactionLedger from '@/components/ledger/payments-entry/current-transaction-ledger'

import { toast } from 'sonner'
import PaymentsEntryFormModal from './payments-entry-form'

export const PaymentsEntry = () => {
    const [selectedMember, setSelectedMember] =
        useState<IMemberResource | null>(null)
    const [openPaymentsEntryModal, setIsOpenPaymentsEntryModal] =
        useState(false)

    const {
        isPending: isPendingCurentMemberLedger,
        isRefetching: isRefetchingCurentMemberLedger,
        data: { data: CurrentMemberLedger },
        refetch: refetchAccountingLedger,
    } = useFilteredPaginatedIAccountingLedger({
        memberProfileId: selectedMember?.memberProfile?.id,
    })

    const totalAmount = CurrentMemberLedger.reduce(
        (acc, ledger) => acc + ledger.credit,
        0
    )

    const isAllowedToCreatePayment = selectedMember !== null

    const handleOpenCreateModal = () => {
        if (isAllowedToCreatePayment) {
            setIsOpenPaymentsEntryModal(true)
        }
    }

    const { pagination, setPagination } = usePagination()
    const { sortingState, tableSorting, setTableSorting } =
        useDataTableSorting()

    const columns = useMemo(() => IAccountingLedgerRequestTableColumns(), [])

    const {
        getRowIdFn,
        columnOrder,
        setColumnOrder,
        isScrollable,
        setIsScrollable,
        columnVisibility,
        setColumnVisibility,
        rowSelectionState,
        createHandleRowSelectionChange,
    } = useDataTableState<IAccountingLedgerResource>({})

    const filterState = useDatableFilterState({
        onFilterChange: () => setPagination({ ...pagination, pageIndex: 0 }),
    })

    const {
        isPending: isPendingSelectedMemberLedger,
        isRefetching: isRefetchingSelectedMemberLedger,
        data: {
            data: SelectedMemberLedger,
            totalPage: selectedMemberTotalPage,
            pageSize: selectedMemberPageSize,
            totalSize: SelectedMemberTotalSize,
        },
        refetch: refetchSelectedMemberLedger,
    } = useFilteredPaginatedIAccountingLedger({
        pagination,
        sort: sortingState,
        filterPayload: filterState.finalFilterPayload,
    })

    const handleRowSelectionChange =
        createHandleRowSelectionChange(SelectedMemberLedger)

    const table = useReactTable({
        columns,
        data: SelectedMemberLedger,
        initialState: {
            columnPinning: { left: ['description', 'balance'] },
        },
        state: {
            sorting: tableSorting,
            pagination,
            columnOrder,
            rowSelection: rowSelectionState.rowSelection,
            columnVisibility,
        },
        rowCount: selectedMemberPageSize,
        manualSorting: true,
        pageCount: selectedMemberTotalPage,
        enableMultiSort: false,
        manualFiltering: true,
        manualPagination: true,
        getRowId: getRowIdFn,
        onSortingChange: setTableSorting,
        onPaginationChange: setPagination,
        getCoreRowModel: getCoreRowModel(),
        onColumnOrderChange: setColumnOrder,
        getSortedRowModel: getSortedRowModel(),
        onColumnVisibilityChange: setColumnVisibility,
        onRowSelectionChange: handleRowSelectionChange,
    })

    const resetSelectedMember = () => {
        setSelectedMember(null)
    }

    return (
        <>
            <PaymentsEntryFormModal
                formProps={{
                    resetSelectedMember: resetSelectedMember,
                    open: openPaymentsEntryModal,
                    onOpenChange: setIsOpenPaymentsEntryModal,
                    selectedMemberId: selectedMember?.id ?? '',
                }}
                open={openPaymentsEntryModal}
                onOpenChange={setIsOpenPaymentsEntryModal}
            />
            <div className="flex w-full flex-col gap-y-4 p-5">
                <div className="grid grid-cols-1 gap-x-4 gap-y-4 lg:grid-cols-4">
                    <legend className="text-lg font-semibold text-secondary-foreground">
                        Payment Transactions
                    </legend>
                    <div className="col-span-2 lg:col-span-4 lg:col-start-1">
                        <div className="flex space-x-2">
                            <MemberPicker
                                value={selectedMember?.id}
                                onSelect={(member) => {
                                    if (!member.memberProfile)
                                        return toast.warning(
                                            'Sorry, this member profile is not yet completed'
                                        )
                                    setSelectedMember(member)
                                    refetchSelectedMemberLedger()
                                }}
                                placeholder="Select Members"
                            />
                            <Button
                                size="icon"
                                className=""
                                variant="destructive"
                                disabled={!selectedMember}
                                onClick={resetSelectedMember}
                            >
                                <XIcon />
                            </Button>
                            <Button
                                onClick={(e) => {
                                    e.preventDefault()
                                    handleOpenCreateModal()
                                }}
                                disabled={!isAllowedToCreatePayment}
                            >
                                create
                            </Button>
                        </div>
                    </div>
                    <div className="col-span-2">
                        <PaymentsEntryProfile profile={selectedMember} />
                        <div className="max-h-96 space-y-4 overflow-auto py-4">
                            <CurrentPaymentAccountingTransactionLedger
                                isRefetching={isRefetchingCurentMemberLedger}
                                isPending={isPendingCurentMemberLedger}
                                data={CurrentMemberLedger}
                                onRefetch={refetchAccountingLedger}
                            />
                        </div>
                    </div>
                    <div className="col-span-1 md:col-span-2">
                        <AccountsLedgerTable
                            table={table}
                            isScrollable={isScrollable}
                            isPending={isPendingSelectedMemberLedger}
                            isRefetching={isRefetchingSelectedMemberLedger}
                            totalSize={SelectedMemberTotalSize}
                            refetch={refetchSelectedMemberLedger}
                            filterState={filterState}
                            setColumnOrder={setColumnOrder}
                            setIsScrollable={setIsScrollable}
                            className="h-full"
                        />
                    </div>
                    <div className="col-span-2 w-full">
                        <Card>
                            <CardContent className="flex items-center justify-between gap-x-2 py-5">
                                <label className="font-bold uppercase">
                                    Total Amount
                                </label>
                                <div className="text-right">
                                    <p className="text-2xl font-bold text-green-600 dark:text-green-400">
                                        ₱ {totalAmount}
                                    </p>
                                </div>
                            </CardContent>
                        </Card>
                    </div>
                </div>
            </div>
        </>
    )
}

export default PaymentsEntry
