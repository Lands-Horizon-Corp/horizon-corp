import { useState } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import { PaginationState } from '@tanstack/react-table'

import GenericPicker from './generic-picker'
import { Button } from '@/components/ui/button'
import { ChevronDownIcon } from '@/components/icons'
import ImageDisplay from '@/components/image-display'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import MiniPaginationBar from '@/components/pagination-bars/mini-pagination-bar'

import {
    useOwner,
    useFilteredPaginatedOwners,
} from '@/hooks/api-hooks/use-owner'

import {
    PAGINATION_INITIAL_INDEX,
    PAGINATION_INITIAL_PAGE_SIZE,
} from '@/constants'
import useFilterState from '@/hooks/use-filter-state'
import { OwnerResource } from '@/horizon-corp/types'

interface Props {
    value? : number
    placeholder?: string
    onSelect?: (selectedOwner: OwnerResource) => void
}

const OwnerPicker = ({ value, placeholder, onSelect }: Props) => {
    const queryClient = useQueryClient()
    const [pickerState, setPickerState] = useState(false)
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
        useFilteredPaginatedOwners({
            filterPayload: finalFilterPayload,
            pagination,
        })

    const owner = useOwner({ ownerId: value as number })

    return (
        <>
            <GenericPicker
                items={data.data}
                open={pickerState}
                listHeading={`Matched Results (${data.totalSize})`}
                searchPlaceHolder="Search owner name..."
                isLoading={isPending || isLoading || isFetching}
                onSelect={(selectedOwner) => {
                    queryClient.setQueryData(
                        ['owner', value],
                        selectedOwner
                    )
                    onSelect?.(selectedOwner)
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
                renderItem={(owner) => (
                    <div className="flex w-full items-center justify-between py-1">
                        <div className="flex items-center gap-x-2">
                            <ImageDisplay src={owner.media?.downloadURL} />
                            <span className="text-ellipsis text-foreground/80">
                                {owner.name}
                            </span>
                        </div>
                        <span className="mr-2 font-mono text-xs italic text-foreground/40">
                            #{owner.id}
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
                onClick={() => setPickerState((val) => !val)}
                className="w-full items-center justify-between rounded-md border bg-background p-0 px-2"
            >
                <span className="justify-betweentext-sm inline-flex w-full items-center text-foreground/90">
                    <span className="inline-flex w-full items-center gap-x-2">
                        <div>
                            {owner.isFetching ? (
                                <LoadingSpinner />
                            ) : (
                                <ImageDisplay
                                    src={owner.data?.media?.downloadURL}
                                />
                            )}
                        </div>
                        {!value ? (
                            <span className="text-foreground/70">
                                {placeholder}
                            </span>
                        ) : (
                            <span>{owner.data?.name}</span>
                        )}
                    </span>
                    <span className="mr-1 font-mono text-sm text-foreground/30">
                        #{value ?? '?'}
                    </span>
                </span>
                <ChevronDownIcon />
            </Button>
        </>
    )
}

export default OwnerPicker
