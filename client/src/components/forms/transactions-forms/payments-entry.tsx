import { useMemo } from 'react'
import {
    getCoreRowModel,
    getSortedRowModel,
    useReactTable,
} from '@tanstack/react-table'

import { Button } from '@/components/ui/button'

import {
    F1Icon,
    F2Icon,
    F3Icon,
    HandDepositIcon,
    HandWithdrawIcon,
    PaymentsIcon,
    ReceiptIcon,
    XIcon,
} from '@/components/icons'

import { useDataTableSorting } from '@/hooks/data-table-hooks/use-datatable-sorting'
import useDataTableState from '@/hooks/data-table-hooks/use-datatable-state'
import { usePagination } from '@/hooks/use-pagination'
import useDatableFilterState from '@/hooks/use-filter-state'

import { IAccountingLedgerResource } from '@/server/types/accounts/accounting-ledger'

import MemberPicker from '@/components/pickers/member-picker'

import AccountsLedgerTable from '@/components/tables/transactions-table/accouting-ledger-table'
import IAccountingLedgerRequestTableColumns from '@/components/tables/transactions-table/accouting-ledger-table/columns'

import PaymentsEntryProfile from '@/components/transaction-profile/payments-entry-profile'
import CurrentPaymentsEntryList from '@/components/ledger/payments-entry/current-payment-entry'

import { toast } from 'sonner'
import { TransactionPaymentEntryFormModal } from './payments-entry-form'
import { cn } from '@/lib'
import { Input } from '@/components/ui/input'
import {
    usePaymentsDataStore,
    usePaymentsModalStore,
} from '@/store/transaction/payments-entry-store'
import { useFilteredPaginatedIAccountingLedger } from '@/hooks/api-hooks/transactions/use-accounting-ledger'
import { commaSeparators } from '@/helpers'
import { useCreatePaymentEntry } from '@/hooks/api-hooks/transactions/use-payments-entry'
import {
    IPaymentsEntry,
    IPaymentsEntryRequest,
} from '@/server/types/transactions/payments-entry'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import { ScrollArea } from '@/components/ui/scroll-area'
import usePaymentsShortcuts from '@/hooks/shortcut-hooks/payments-entry-shortcuts'
import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from '@/components/ui/resizable'

export const PaymentsEntry = () => {
    const {
        selectedMember,
        ORNumber,
        selectedPayments,
        setSelectedMember,
        setORNumber,
        setSelectedPayments,
    } = usePaymentsDataStore()

    const {
        setOpenPaymentsEntryModal,
        openPaymentsEntryModal,
        setTransactionType,
    } = usePaymentsModalStore()

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
        columnResizeMode: 'onChange',
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

    const handleOpenCreateModal = (paymentType: string) => {
        if (isAllowedToCreatePayment) {
            setTransactionType(paymentType)
            setOpenPaymentsEntryModal(true)
        }
    }

    const totalAmount = useMemo(() => {
        return (
            selectedPayments?.reduce(
                (acc, payment) => acc + payment.amount,
                0
            ) ?? 0
        )
    }, [selectedPayments])

    const {
        mutate: createPayment,
        isPending: isPendingCreatePayments,
        error,
    } = useCreatePaymentEntry({
        onSuccess: () => {
            setSelectedMember(null)
            setSelectedPayments([])
            setORNumber('')
        },
        onError: (error) => {
            toast.error(error)
        },
    })

    const handleSubmitPayment = (payments: IPaymentsEntry[]) => {
        const payload: IPaymentsEntryRequest[] = payments.map((payment) => ({
            ...payment,
            accountsId: payment.account.id || '',
        }))

        const toPrintCount = payload.filter((item) => item.isPrinted)
        // To do:
        // call printFunction()
        toast.loading(
            `${toPrintCount.length} item transaction is printing now!`
        )
        // tentative will depend of print function loading
        setTimeout(() => {
            toast.dismiss()
        }, 2000 * toPrintCount.length)

        createPayment(payload)
    }

    const isAllowedToCreatePayment = selectedMember !== null
    const hadSelectedPayments = selectedPayments.length <= 0

    usePaymentsShortcuts({
        hadSelectedPayments,
        selectedPayments,
        setSelectedMember,
        handleSubmitPayment,
        handleOpenCreateModal,
    })

    return (
        <>
            <TransactionPaymentEntryFormModal
                formProps={{
                    resetSelectedMember: resetSelectedMember,
                    open: openPaymentsEntryModal,
                    onOpenChange: setOpenPaymentsEntryModal,
                    selectedMemberId: selectedMember?.id ?? '',
                    error: error ?? '',
                }}
                open={openPaymentsEntryModal}
                onOpenChange={setOpenPaymentsEntryModal}
            />
            <div className="flex w-full flex-col gap-y-4 space-y-2 p-4 pt-0 lg:space-y-0">
                <legend className="text-lg font-semibold text-secondary-foreground">
                    Payment Transactions
                </legend>
                <ResizablePanelGroup
                    direction="horizontal"
                    className="h-[100vh] min-h-[100vh] w-full"
                >
                    <ResizablePanel className="" defaultSize={40}>
                        <ResizablePanelGroup
                            className="min-h-[20vh]"
                            direction="vertical"
                        >
                            <ResizablePanel defaultSize={30} className="p-2">
                                <div className="flex w-full min-w-[30vw] flex-col gap-y-2">
                                    <div className="flex space-x-2">
                                        <Button
                                            size="icon"
                                            className="min-w-10"
                                            variant="secondary"
                                            disabled={!selectedMember}
                                            onClick={(e) => {
                                                e.preventDefault()
                                                resetSelectedMember()
                                            }}
                                        >
                                            <XIcon />
                                        </Button>
                                        <MemberPicker
                                            disabled={!hadSelectedPayments}
                                            allowShorcutCommand
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
                                        <div className="relative w-full">
                                            <Input
                                                value={ORNumber}
                                                onChange={(e) =>
                                                    setORNumber(e.target.value)
                                                }
                                                className={cn(
                                                    'pr-9 text-lg font-semibold placeholder:text-sm placeholder:font-normal placeholder:text-foreground/40'
                                                )}
                                                id="OR number"
                                                placeholder="Enter the OR Number"
                                                autoComplete="off"
                                            />
                                            <span className="absolute right-3 top-1/2 -translate-y-1/2 text-lg font-bold text-secondary">
                                                <ReceiptIcon />
                                            </span>
                                        </div>
                                    </div>

                                    <PaymentsEntryProfile
                                        profile={selectedMember}
                                    />
                                    <div className="flex space-x-2">
                                        <Button
                                            onClick={(e) => {
                                                e.preventDefault()
                                                handleOpenCreateModal('payment')
                                            }}
                                            className={cn(
                                                'flex grow justify-between hover:bg-primary'
                                            )}
                                            disabled={!isAllowedToCreatePayment}
                                            variant={
                                                isAllowedToCreatePayment
                                                    ? 'outline'
                                                    : 'secondary'
                                            }
                                        >
                                            <PaymentsIcon className="mr-2" />
                                            Pay
                                            <span className="ml-1 flex items-center justify-center gap-x-1">
                                                ⌘
                                                <F1Icon
                                                    className="text-muted-foreground"
                                                    size={23}
                                                />
                                            </span>
                                        </Button>
                                        <Button
                                            className={cn(
                                                'flex grow justify-between hover:bg-primary'
                                            )}
                                            disabled={!isAllowedToCreatePayment}
                                            onClick={(e) => {
                                                e.preventDefault()
                                                handleOpenCreateModal('deposit')
                                            }}
                                            variant={
                                                isAllowedToCreatePayment
                                                    ? 'outline'
                                                    : 'secondary'
                                            }
                                        >
                                            <HandDepositIcon className="mr-2" />
                                            Deposit
                                            <span className="ml-1 flex items-center justify-center gap-x-1">
                                                ⌘
                                                <F2Icon
                                                    className="text-muted-foreground"
                                                    size={23}
                                                />
                                            </span>
                                        </Button>
                                        <Button
                                            className={cn(
                                                'flex grow justify-between hover:bg-primary'
                                            )}
                                            disabled={!isAllowedToCreatePayment}
                                            variant={
                                                isAllowedToCreatePayment
                                                    ? 'outline'
                                                    : 'secondary'
                                            }
                                            onClick={(e) => {
                                                e.preventDefault()
                                                handleOpenCreateModal(
                                                    'withdraw'
                                                )
                                            }}
                                        >
                                            <HandWithdrawIcon className="mr-2" />
                                            widthdraw
                                            <span className="ml-1 flex items-center justify-center gap-x-1">
                                                ⌘
                                                <F3Icon
                                                    className="text-muted-foreground"
                                                    size={23}
                                                />
                                            </span>
                                        </Button>
                                    </div>
                                </div>
                            </ResizablePanel>
                            <ResizableHandle />
                            <ResizablePanel defaultSize={70} className="p-2">
                                <div className="w-full p-2">
                                    <AccountsLedgerTable
                                        table={table}
                                        isScrollable={isScrollable}
                                        isPending={
                                            isPendingSelectedMemberLedger
                                        }
                                        isRefetching={
                                            isRefetchingSelectedMemberLedger
                                        }
                                        totalSize={SelectedMemberTotalSize}
                                        refetch={refetchSelectedMemberLedger}
                                        filterState={filterState}
                                        setColumnOrder={setColumnOrder}
                                        setIsScrollable={setIsScrollable}
                                        className="h-full"
                                    />
                                </div>
                            </ResizablePanel>
                        </ResizablePanelGroup>
                    </ResizablePanel>
                    <ResizableHandle withHandle />
                    <ResizablePanel defaultSize={30}>
                        <div className="w-full min-w-[30vw] p-2">
                            <div className="flex items-center gap-x-2">
                                <div className="to-indigo-background/10 flex-grow rounded-xl border-[0.1px] border-primary/30 bg-gradient-to-br from-primary/10 p-2">
                                    <div className="flex items-center justify-between gap-x-2">
                                        <label className="text-sm font-bold uppercase text-muted-foreground">
                                            Total Amount
                                        </label>
                                        <p className="text-lg font-bold text-green-600 dark:text-green-400">
                                            ₱{' '}
                                            {totalAmount
                                                ? commaSeparators(
                                                      totalAmount.toString()
                                                  )
                                                : '0.00'}
                                        </p>
                                    </div>
                                </div>
                                <Button
                                    onClick={() =>
                                        handleSubmitPayment(selectedPayments)
                                    }
                                    disabled={
                                        isPendingCreatePayments ||
                                        hadSelectedPayments
                                    }
                                >
                                    {isPendingCreatePayments ? (
                                        <span className="flex gap-x-2">
                                            Saving <LoadingSpinner />
                                        </span>
                                    ) : (
                                        'Save'
                                    )}
                                </Button>
                            </div>
                            <ScrollArea className="ecoop-scroll w-full overflow-auto py-2">
                                <CurrentPaymentsEntryList
                                    data={selectedPayments}
                                />
                            </ScrollArea>
                        </div>
                    </ResizablePanel>
                </ResizablePanelGroup>
            </div>
        </>
    )
}

export default PaymentsEntry
