// DataTable.tsx
import React from 'react'
import { Table, TableBody, TableHeader, TableRow } from '../ui/table'
import { TableProps } from './types'

/**
 * A flexible, reusable DataTable component using generics.
 * 
 * @example
 * <DataTable
 *   data={rows}
 *   columns={[
 *     {
 *       key: 'id',
 *       header: 'ID',
 *       width: '5%',
 *       renderCell: (row) => row.id.toString(),
 *     },
 *     {
 *       key: 'name',
 *       header: 'Name',
 *       renderCell: (row) => row.name || 'N/A',
 *     },
 *   ]}
 *   getRowKey={(row) => row.id}
 * />
 */
function ReportDataTable<TData>(props: TableProps<TData>) {
  const {
    columns,
    data,
    getRowKey,
    rowClassName,
    rowStyle,
    placeholderForUndefined = '',
  } = props


  return (
    <Table>
      <TableHeader>
        <TableRow>
          {columns.map((col) => {
            const headerText = col.header ?? String(col.key)
            return (
              <th
                key={String(col.key)}
                style={{ textAlign: col.align, width: col.width, ...col.headerStyle }}
                className={col.headerClassName}
              >
                {headerText}
              </th>
            )
          })}
        </TableRow>
      </TableHeader>

      <TableBody>
        {data.map((row, rowIndex) => {
          // Use the user-provided getRowKey() if available, else fallback
          const uniqueRowKey = getRowKey ? getRowKey(row) : rowIndex

          return (
            <TableRow
              key={uniqueRowKey}
              style={typeof rowStyle === 'function' ? rowStyle(row, rowIndex) : rowStyle}
              className={typeof rowClassName === 'function'
                ? rowClassName(row, rowIndex)
                : rowClassName}
            >
              {columns.map((col) => {
                let cellValue: React.ReactNode

                // If there's a custom renderCell, use it; else, use the raw property value
                if (col.renderCell) {
                  cellValue = col.renderCell(row, rowIndex)
                } else {
                  const dataKey = col.key as keyof TData
                  const value = row[dataKey]
                  cellValue = value === undefined
                    ? placeholderForUndefined
                    : String(value)
                }

                console.log('col', col.cellStyle)

                return (
                  <td
                    key={String(col.key)}
                    style={{
                      textAlign: col.align,
                      ...col.cellStyle,
                    }}
                    className={col.cellClassName}
                  >
                    {cellValue}
                  </td>
                )
              })}
            </TableRow>
          )
        })}
      </TableBody>
    </Table>
  )
}

export default ReportDataTable
