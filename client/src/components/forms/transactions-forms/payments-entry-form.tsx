import { useMemo, useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import * as z from 'zod'
import {
    getCoreRowModel,
    getSortedRowModel,
    useReactTable,
} from '@tanstack/react-table'

import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Checkbox } from '@/components/ui/checkbox'
import { Textarea } from '@/components/ui/textarea'
import { Card, CardContent } from '@/components/ui/card'
import FormErrorMessage from '@/components/ui/form-error-message'
import FormFieldWrapper from '@/components/ui/form-field-wrapper'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from '@/components/ui/form'
import Modal from '@/components/modals/modal'

import { XIcon } from '@/components/icons'

import { useCreatePaymentEntry } from '@/hooks/api-hooks/transactions/use-payments-entry'
import { useFilteredPaginatedIAccountingLedger } from '@/hooks/api-hooks/transactions/use-accounting-ledger'
import { useDataTableSorting } from '@/hooks/data-table-hooks/use-datatable-sorting'
import useDataTableState from '@/hooks/data-table-hooks/use-datatable-state'
import { usePagination } from '@/hooks/use-pagination'
import useDatableFilterState from '@/hooks/use-filter-state'

import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import { IPaymentsEntryRequest } from '@/server/types/transactions/payments-entry'
import { IMemberResource, TEntityId } from '@/server'
import { IAccountingLedgerResource } from '@/server/types/accounts/accounting-ledger'

import AccountsPicker from '@/components/pickers/accounts-picker'
import MemberPicker from '@/components/pickers/member-picker'
import TransactionPaymentTypesPicker from '@/components/pickers/transaction-payment-types-picker'

import AccountsLedgerTable from '@/components/tables/transactions-table/accouting-ledger-table'
import IAccountingLedgerRequestTableColumns from '@/components/tables/transactions-table/accouting-ledger-table/columns'

import PaymentsEntryProfile from '@/components/transaction-profile/payments-entry-profile'
import CurrentPaymentAccountingTransactionLedger from '@/components/ledger/payments-entry/current-transaction-ledger'

import CheckBankFormModal from '../check-bank-forms/check-bank-forms'
import { toast } from 'sonner'

export const paymentsEntrySchema = z.object({
    ORNumber: z.string().min(1, 'OR Number is required'),
    amount: z.coerce.number().positive('Amount must be greater than 0'),
    accountsId: z.string().min(1, 'Accounts is required'),
    isPrinted: z.boolean().optional(),
    notes: z.string().optional(),
    transactionType: z.string(),
})

type TPaymentFormValues = z.infer<typeof paymentsEntrySchema>

export interface IPaymentFormProps
    extends IBaseCompNoChild,
        IForm<Partial<IPaymentsEntryRequest>, unknown, string> {
    open: boolean
    onOpenChange: (open: boolean) => void
    selectedMemberId?: TEntityId
}

export const PaymentsEntryForm = () => {
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

    return (
        <>
            <PaymentsEntryFormModal
                selectedMemberId={selectedMember?.id}
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
                                onClick={() => {
                                    setSelectedMember(null)
                                }}
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

const PaymentsEntryFormModal = ({
    open,
    onOpenChange,
    onSuccess,
    onError,
    selectedMemberId,
}: IPaymentFormProps) => {
    const [openCheckClearingFormModal, setIsCheckClearingFormModal] =
        useState(false)

    const form = useForm<TPaymentFormValues>({
        resolver: zodResolver(paymentsEntrySchema),
        reValidateMode: 'onChange',
        mode: 'onSubmit',
        defaultValues: {
            ORNumber: '',
            amount: 0,
            accountsId: '',
            notes: '',
            isPrinted: false,
        },
    })

    const createMutation = useCreatePaymentEntry({
        onSuccess: (data) => {
            onSuccess?.(data)
            onOpenChange(false)
        },
        onError,
    })

    const onSubmit = form.handleSubmit((data) => {
        if (data.transactionType === 'trans-pay-002') {
            setIsCheckClearingFormModal(true)
            return
        }
        createMutation.mutate({
            memberId: selectedMemberId ?? '',
            ...data,
            isPrinted: data.isPrinted ?? false,
        })
    })

    const onCancel = () => {
        onOpenChange(false)
        form.reset()
    }

    return (
        <Modal
            open={open}
            onOpenChange={onOpenChange}
            title="Payment Entry Form"
            description="Fill in the details to create a payment entry."
            footer={
                <div className="flex justify-end gap-2 p-4">
                    <Button variant="outline" onClick={() => onCancel()}>
                        Cancel
                    </Button>
                    <Button
                        onClick={onSubmit}
                        type="submit"
                        disabled={createMutation.isPending}
                    >
                        {createMutation.isPending ? (
                            <LoadingSpinner />
                        ) : (
                            'Submit'
                        )}
                    </Button>
                </div>
            }
        >
            <CheckBankFormModal
                open={openCheckClearingFormModal}
                onOpenChange={setIsCheckClearingFormModal}
            />
            <Form {...form}>
                <form onSubmit={onSubmit} className="space-y-4">
                    <FormFieldWrapper
                        control={form.control}
                        name="ORNumber"
                        label="OR Number"
                        render={({ field }) => (
                            <Input
                                {...field}
                                id={field.name}
                                placeholder="OR Number"
                                autoComplete="off"
                            />
                        )}
                    />
                    <FormFieldWrapper
                        control={form.control}
                        name="amount"
                        label="Amount"
                        render={({ field }) => (
                            <Input
                                {...field}
                                id={field.name}
                                type="number"
                                step="0.01"
                                placeholder="Amount"
                                autoComplete="off"
                            />
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="transactionType"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel htmlFor={field.name}>
                                    Payment Types
                                </FormLabel>
                                <FormControl>
                                    <TransactionPaymentTypesPicker
                                        value={field.value}
                                        onSelect={(transactionPaymenttypes) => {
                                            field.onChange(
                                                transactionPaymenttypes.id
                                            )
                                        }}
                                        placeholder="Select Payment types"
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="accountsId"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel htmlFor={field.name}>
                                    Accounts
                                </FormLabel>
                                <FormControl>
                                    <AccountsPicker
                                        value={field.value}
                                        onSelect={(accounts) => {
                                            field.onChange(accounts.id)
                                        }}
                                        placeholder="Select Accounts"
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormField
                        name="notes"
                        control={form.control}
                        render={({ field }) => (
                            <FormItem className="col-span-1 space-y-1">
                                <FormLabel>Add some feedback/notes.</FormLabel>
                                <FormControl>
                                    <Textarea
                                        {...field}
                                        placeholder="Enter some notes."
                                        className="resize-none"
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="isPrinted"
                        render={({ field }) => (
                            <FormItem className="flex flex-row items-center space-x-2">
                                <FormLabel htmlFor={field.name}>
                                    Allow to make print
                                </FormLabel>
                                <FormControl>
                                    <Checkbox
                                        className="-translate-y-1"
                                        checked={field.value}
                                        onCheckedChange={field.onChange}
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormErrorMessage errorMessage={createMutation.error} />
                </form>
            </Form>
        </Modal>
    )
}
