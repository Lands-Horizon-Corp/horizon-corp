import {
    Dialog,
    DialogContent,
    DialogTrigger,
} from '@/components/ui/dialog'
import {
    ColumnDef,
} from '@tanstack/react-table'

import { ReactNode, useState } from 'react'

interface InsertTableProps<TData> {
    data?: TData[]
    columns?: ColumnDef<TData>[]
    filterPlaceholder?: string,
    trigger?: ReactNode,
    content?: ReactNode,
}

export const InsertTable = <TData,>({
    data,
    columns,
    filterPlaceholder = 'Filter...',
    trigger,
    content
}: InsertTableProps<TData>) => {

    return (
        <Dialog>
            <DialogTrigger asChild>
              {trigger}
            </DialogTrigger>
            <DialogContent className="max-w-fit">
                <div className='w-full'>
                 {content}         
                </div>
            </DialogContent>
        </Dialog>
    )
}

export default InsertTable
