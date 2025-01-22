import {
    Dialog,
    DialogContent,
    DialogTrigger,
} from '@/components/ui/dialog'

import { ReactNode } from 'react'

interface InsertTableProps {
    table: ReactNode,
    trigger?: ReactNode,
}

export const InsertTable =({
    trigger,
    table
}: InsertTableProps) => {
    

    return (
        <Dialog>
            <DialogTrigger asChild>
              {trigger}
            </DialogTrigger>
            <DialogContent className="max-w-fit pt-10">
                {table}
            </DialogContent>
        </Dialog>
    )
}

export default InsertTable
