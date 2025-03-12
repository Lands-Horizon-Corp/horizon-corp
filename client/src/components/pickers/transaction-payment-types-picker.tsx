import { useState } from 'react'
import { PaginationState } from '@tanstack/react-table'

import GenericPicker from './generic-picker'
import { Button } from '@/components/ui/button'
import { ChevronDownIcon } from '@/components/icons'
import MiniPaginationBar from '@/components/pagination-bars/mini-pagination-bar'

import {
    PAGINATION_INITIAL_INDEX,
    PAGINATION_INITIAL_PAGE_SIZE,
} from '@/constants'
import useFilterState from '@/hooks/use-filter-state'
import { TEntityId } from '@/server/types'

import { useFilteredPaginatedTransactionPaymentTypes } from '@/hooks/api-hooks/transactions/use-transaction-payment-types'
import { ITransactionPaymentTypesResource } from '@/server/types/transactions/transaction-payment-types'
import React from 'react'
import { DEFAULT_TRANSACTION_TYPE } from '@/validations/transactions/payments-entry'

interface Props {
    value?: TEntityId
    placeholder?: string
    disabled?: boolean
    leftIcon?: React.ReactNode
    onSelect?: (
        selectedTransactionType: ITransactionPaymentTypesResource
    ) => void
    defaultValue?: string
}

const TransactionPaymentTypesPicker = ({
    value,
    disabled,
    placeholder,
    onSelect,
    leftIcon,
    defaultValue,
}: Props) => {
    const [pickerState, setPickerState] = useState(false)
    const [selectedTransactionType, setSelectedTransactionType] =
        useState<ITransactionPaymentTypesResource>()
    const [pagination, setPagination] = useState<PaginationState>({
        pageIndex: PAGINATION_INITIAL_INDEX,
        pageSize: PAGINATION_INITIAL_PAGE_SIZE,
    })

    const { finalFilterPayload, setFilter } = useFilterState({
        onFilterChange: () =>
            setPagination((prev) => ({
                ...prev,
                pageIndex: PAGINATION_INITIAL_INDEX,
            })),
    })

    const { data, isPending, isLoading, isFetching } =
        useFilteredPaginatedTransactionPaymentTypes({
            filterPayload: finalFilterPayload,
            pagination,
            enabled: !disabled,
            showMessage: false,
        })

    const defaultSelected = data.data.find(
        (item) => item.id === DEFAULT_TRANSACTION_TYPE
    )?.name

    return (
        <>
            <GenericPicker
                items={data.data}
                open={pickerState}
                listHeading={`Matched Results (${data.totalSize})`}
                searchPlaceHolder="Search transaction types..."
                isLoading={isPending || isLoading || isFetching}
                onSelect={(selectedType) => {
                    onSelect?.(selectedType)
                    setSelectedTransactionType(selectedType)
                }}
                onOpenChange={setPickerState}
                onSearchChange={(searchValue) => {
                    setFilter('name', {
                        displayText: '',
                        mode: 'contains',
                        dataType: 'text',
                        value: searchValue,
                    })
                }}
                renderItem={(transactionType) => (
                    <div className="flex w-full items-center justify-between py-1">
                        <div className="flex flex-col">
                            <span className="font-medium text-foreground/80">
                                {transactionType.name}
                            </span>
                            <span className="text-sm text-foreground/60">
                                {transactionType.description}
                            </span>
                        </div>
                        <span className="text-sm text-foreground/50">
                            Cheque ID: {transactionType.chequeId}
                        </span>
                    </div>
                )}
            >
                <MiniPaginationBar
                    pagination={{
                        pageIndex: pagination.pageIndex,
                        pageSize: pagination.pageSize,
                        totalPage: data.totalPage,
                        totalSize: data.totalSize,
                    }}
                    disablePageMove={isFetching}
                    onNext={({ pageIndex }) =>
                        setPagination((prev) => ({ ...prev, pageIndex }))
                    }
                    onPrev={({ pageIndex }) =>
                        setPagination((prev) => ({ ...prev, pageIndex }))
                    }
                />
            </GenericPicker>
            <Button
                type="button"
                variant="secondary"
                disabled={disabled}
                onClick={() => setPickerState((val) => !val)}
                className="w-full items-center justify-between rounded-md border bg-background p-0 px-2"
            >
                <span className="flex w-full items-center justify-between text-sm text-foreground/90">
                    {leftIcon && (
                        <span className="mr-2 flex-shrink-0">{leftIcon}</span>
                    )}
                    <span className="flex w-full items-center gap-x-2 overflow-hidden">
                        {!value ? (
                            <span className="truncate text-xs text-foreground/40">
                                {placeholder}
                            </span>
                        ) : (
                            <span className="truncate font-mono text-sm text-foreground">
                                {defaultValue && !selectedTransactionType
                                    ? defaultSelected
                                    : selectedTransactionType?.name}
                            </span>
                        )}
                    </span>
                </span>
                <ChevronDownIcon />
            </Button>
        </>
    )
}

export default TransactionPaymentTypesPicker
