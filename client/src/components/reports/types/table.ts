// types.ts or table.ts (wherever you keep your table types)

import React from 'react'

/**
 * ColumnDefinition<TData> is the shape for individual column configurations.
 */
export interface ColumnDefinition<TData> {
  key: keyof TData | string
  header?: string
  renderCell?: (row: TData, rowIndex: number) => React.ReactNode
  align?: React.CSSProperties['textAlign']
  width?: string
  // Optional styling or classnames for the header cell
  headerStyle?: React.CSSProperties
  headerClassName?: string
  // Optional styling or classnames for the body cell
  cellStyle?: React.CSSProperties
  cellClassName?: string
}

/**
 * TableProps<TData> describes the props that the DataTable component will accept.
 */
export interface TableProps<TData> {
  columns: ColumnDefinition<TData>[]
  data: TData[]

  /**
   * If provided, this function should return a unique string/number
   * to use as the React "key" for each row. This is better for performance
   * and helps React identify which rows have changed.
   */
  getRowKey?: (row: TData) => string | number

  /**
   * If you want to apply a conditional or static class name to each table row.
   * Can be either a string or a function returning a string.
   */
  rowClassName?: string | ((row: TData, rowIndex: number) => string)

  /**
   * If you want to apply inline styles to each table row.
   * Can be a static object or a function returning an object.
   */
  rowStyle?: React.CSSProperties | ((row: TData, rowIndex: number) => React.CSSProperties)

  /**
   * A fallback displayed whenever a cell's value is undefined (or if
   * the column key doesn't exist on the row). Defaults to an empty string.
   */
  placeholderForUndefined?: React.ReactNode
}
