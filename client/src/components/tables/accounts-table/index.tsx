import DataTable from "@/components/data-table";
import DataTablePagination from "@/components/data-table/data-table-pagination";
import DataTableToolbar, { IDataTableToolbarProps } from "@/components/data-table/data-table-toolbar";
import FilterContext from "@/contexts/filter-context/filter-context";
import { useDataTableSorting } from "@/hooks/data-table-hooks/use-datatable-sorting";
import useDataTableState from "@/hooks/data-table-hooks/use-datatable-state";
import { usePagination } from "@/hooks/use-pagination";
import { cn } from "@/lib";
import { IAccountResource } from "@/server/types/accounts/accounts";
// import { useQueryClient } from "@tanstack/react-query";
import { useReactTable, getCoreRowModel, getSortedRowModel } from "@tanstack/react-table";
import { useMemo } from "react";
import { TableProps } from "../types";
import useDatableFilterState from '@/hooks/use-filter-state'

import accountTableColumns, {
  accountsGlobalSearchTargets,
  IAccountsTableColumnProps,
} from './columns'
import { DummyAccountsData } from "./dummy-accounts";


export interface AccountsTableProps
    extends TableProps<IAccountResource>,
        IAccountsTableColumnProps {
    toolbarProps?: Omit<
        IDataTableToolbarProps<IAccountResource>,
        | 'table'
        | 'refreshActionProps'
        | 'globalSearchProps'
        | 'scrollableProps'
        | 'filterLogicProps'
        | 'exportActionProps'
        | 'deleteActionProps'
    >;
}

const AccountsTable = ({
  className,
  toolbarProps,
  defaultFilter,
  onSelectData,
  actionComponent,
}: AccountsTableProps) => {
//   const queryClient = useQueryClient();
  const { pagination, setPagination } = usePagination();
  const { tableSorting, setTableSorting } = useDataTableSorting();

  const columns = useMemo(
      () =>
          accountTableColumns({
              actionComponent,
          }),
      [actionComponent]
  );

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
  } = useDataTableState<IAccountResource>({
      columnOrder: columns.map((c) => c.id!),
      onSelectData,
  });

  const filterState = useDatableFilterState({
      defaultFilter,
      onFilterChange: () => setPagination({ ...pagination, pageIndex: 0 }),
  });

  // const {
  //     isPending,
  //     isRefetching,
  //     data: { data, totalPage, pageSize, totalSize },
  //     refetch,
  // } = useFilteredPaginatedAccounts({
  //     pagination,
  //     sort: sortingState,
  //     filterPayload: filterState.finalFilterPayload,
  // });

  const handleRowSelectionChange = createHandleRowSelectionChange(DummyAccountsData);

  const table = useReactTable({
      columns,
      data: DummyAccountsData,
      initialState: {
          columnPinning: { left: ['select'] },
      },
      state: {
          sorting: tableSorting,
          pagination,
          columnOrder,
          rowSelection: rowSelectionState.rowSelection,
          columnVisibility,
      },
      // rowCount: pageSize,
      manualSorting: true,
      // pageCount: totalPage,
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
  });

  return (
      <FilterContext.Provider value={filterState}>
          <div className={cn('flex h-full flex-col gap-y-2', className)}>
              <DataTableToolbar
                  className=""
                  globalSearchProps={{
                      defaultMode: 'equal',
                      targets: accountsGlobalSearchTargets,
                  }}
                  table={table}
                  refreshActionProps={{
                      onClick: () => {},
                      isLoading: false,
                  }}
                  // deleteActionProps={{
                  //     onDeleteSuccess: () =>
                  //         queryClient.invalidateQueries({
                  //             queryKey: ['account', 'resource-query'],
                  //         }),
                  //     onDelete: (selectedData) =>
                  //         AccountService.deleteMany(
                  //         //     selectedData.map((data) => data.id)
                  //         // ),
                  // }}
                  scrollableProps={{ isScrollable, setIsScrollable }}
                  exportActionProps={{
                      pagination,
                      // isLoading: isPending,
                      filters: filterState.finalFilterPayload,
                      // disabled: isPending || isRefetching,
                      // exportAll: AccountService.exportAll,
                      // exportAllFiltered: AccountService.exportAllFiltered,
                      // exportCurrentPage: (ids) =>
                      //     AccountService.exportSelected(
                      //         ids.map((data) => data.id)
                      //     ),
                      // exportSelected: (ids) =>
                      //     AccountService.exportSelected(
                      //         ids.map((data) => data.id)
                      //     ),
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
              <DataTablePagination table={table} totalSize={50} />
          </div>
      </FilterContext.Provider>
  );
};

export default AccountsTable;
