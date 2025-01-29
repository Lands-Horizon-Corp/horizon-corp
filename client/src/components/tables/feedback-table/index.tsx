import { useMemo } from 'react'
import {
    useReactTable,
    getCoreRowModel,
    getSortedRowModel,
} from '@tanstack/react-table'

import { cn } from '@/lib'

import DataTable from '@/components/data-table'
import DataTablePagination from '@/components/data-table/data-table-pagination'
import DataTableToolbar, {
    IDataTableToolbarProps,
} from '@/components/data-table/data-table-toolbar'

import FeedbackColumns, {
    FeedbackGlobalSearchTargets,
    IFeedbackTableColumnProps,
} from './column'

import useDataTableState from '@/components/data-table/hooks/use-datatable-state'
import useDatableFilterState from '@/hooks/use-filter-state'
import FilterContext from '@/contexts/filter-context/filter-context'

import { TableProps } from '../types'
import { IFeedbackResource } from '@/server/types'

import FeedbackService from '@/server/api-service/feedback-service'

const data: IFeedbackResource[] = [
    {
        id: '7f76efd0-940a-42f9-afa9-8644453e20aa',
        email: 'user1@example.com',
        description: 'Great service, very satisfied.',
        feedbackType: 'positive',
        createdAt: '2025-01-01T10:00:00Z',
        updatedAt: '2025-01-01T10:00:00Z',
    },
    {
        id: '24d740f2-f57f-4579-a9a3-03d6ddec44b5',
        email: 'user2@example.com',
        description: 'Had some issues with the checkout process.',
        feedbackType: 'negative',
        createdAt: '2025-01-02T14:30:00Z',
        updatedAt: '2025-01-02T14:30:00Z',
    },
    {
        id: '9482cd20-df59-49a2-b6ce-bf88bed349e3',
        email: 'user3@example.com',
        description: 'Can we have more payment options?',
        feedbackType: 'neutral',
        createdAt: '2025-01-03T09:15:00Z',
        updatedAt: '2025-01-03T09:15:00Z',
    },
    {
        id: 'f13a44b8-d627-47db-8584-614eb016f6d6',
        email: 'user4@example.com',
        description: 'Loved the user interface. Very intuitive!',
        feedbackType: 'positive',
        createdAt: '2025-01-04T18:45:00Z',
        updatedAt: '2025-01-04T18:45:00Z',
    },
    {
        id: '88350bb7-29de-4e99-bcf0-2f668bcfae92',
        email: 'user5@example.com',
        description: 'The app is crashing frequently. Please fix.',
        feedbackType: 'negative',
        createdAt: '2025-01-05T11:20:00Z',
        updatedAt: '2025-01-05T11:20:00Z',
    },
]

export interface FeedbackTableProps
    extends TableProps<IFeedbackResource>,
        IFeedbackTableColumnProps {
    toolbarProps?: Omit<
        IDataTableToolbarProps<IFeedbackResource>,
        | 'table'
        | 'refreshActionProps'
        | 'globalSearchProps'
        | 'scrollableProps'
        | 'filterLogicProps'
        | 'exportActionProps'
        | 'deleteActionProps'
    >
}

const FeedBackTable = ({
    actionComponent,
    onSelectData,
    defaultFilter,
    className,
    toolbarProps,
}: FeedbackTableProps) => {

    const columns = useMemo(
        () =>
            FeedbackColumns({
                actionComponent,
            }),
        [actionComponent]
    )
    
    const {
        sorting,
        setSorting,
        getRowIdFn,
        pagination,
        setPagination,
        columnOrder,
        setColumnOrder,
        isScrollable,
        setIsScrollable,
        rowSelectionState,
        createHandleRowSelectionChange,
        columnVisibility,
        setColumnVisibility,
    } = useDataTableState<IFeedbackResource>({
        columnOrder: columns.map((c) => c.id!),
        onSelectData,
    })

    const filterState = useDatableFilterState({
        defaultFilter,
        onFilterChange: () => setPagination({ ...pagination, pageIndex: 0 }),
    })

    const handleRowSelectionChange = createHandleRowSelectionChange(data)

    const table = useReactTable({
        columns,
        data: data,
        initialState: {
            columnPinning: { left: ['select'] },
        },
        state: {
            sorting,
            pagination,
            columnOrder,
            rowSelection: rowSelectionState.rowSelection,
            columnVisibility,
        },
        rowCount: 50,
        pageCount: 1,
        manualSorting: true,
        manualFiltering: true,
        manualPagination: true,
        getRowId: getRowIdFn,
        onSortingChange: setSorting,
        onPaginationChange: setPagination,
        getCoreRowModel: getCoreRowModel(),
        onColumnOrderChange: setColumnOrder,
        getSortedRowModel: getSortedRowModel(),
        onColumnVisibilityChange: setColumnVisibility,
        onRowSelectionChange: handleRowSelectionChange,
    })

    return (
        <FilterContext.Provider value={filterState}>
            <div className={cn('flex h-full flex-col gap-y-2', className)}>
                <DataTableToolbar
                    hideCreateButton
                    globalSearchProps={{
                        defaultMode: 'equal',
                        targets: FeedbackGlobalSearchTargets,
                    }}
                    table={table}
                    refreshActionProps={{
                        onClick: () => {},
                        isLoading: false,
                    }}
                    scrollableProps={{ isScrollable, setIsScrollable }}
                    exportActionProps={{
                        pagination,
                        isLoading: false,
                        filters: filterState.finalFilterPayload,
                        disabled: false,
                        exportAll: FeedbackService.exportAll
                    }}
                    filterLogicProps={{
                        filterLogic: filterState.filterLogic,
                        setFilterLogic: filterState.setFilterLogic,
                    }}
                    {...toolbarProps}
                />
                <DataTable
                    table={table}
                    isStickyHeader
                    isStickyFooter
                    isScrollable={isScrollable}
                    setColumnOrder={setColumnOrder}
                    className="mb-2"
                />
                <DataTablePagination table={table} totalSize={5} />
            </div>
        </FilterContext.Provider>
    )
}

export default FeedBackTable
