import { userReportsTableColumns } from '@/components/reports/config/user-config';
import { TellerTransactionRecord } from '@/components/reports/types';
import { ColumnDef } from '@tanstack/react-table';
import { create } from 'zustand';

interface TableStore<T, C> {
  data: T[];
  config: C;
  columns: ColumnDef<T>[];
  setData: (data: T[]) => void;
  setConfig: (config: C) => void;
  setColumns: (columns: ColumnDef<T>[]) => void;
}

export const createStore = <T, C>(defaultData: T[], defaultColumns: ColumnDef<T>[], defaultConfig: C) =>
  create<TableStore<T, C>>((set) => ({
    data: defaultData,
    columns: defaultColumns,
    config: defaultConfig,
    setData: (data) => set({ data }),
    setColumns: (columns) => set({ columns }),
    setConfig: (config) => set({ config }),
  }));


export const useTellerTransactionStore = createStore<TellerTransactionRecord, Record<string, any>>([], userReportsTableColumns(), {});
