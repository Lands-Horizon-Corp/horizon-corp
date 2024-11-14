import { toast } from 'sonner'
import { Table } from '@tanstack/react-table'

import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuGroup,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'
import { ExportIcon, CsvIcon } from '@/components/icons'
import { useCallback, useEffect } from 'react'

type TExportFileType = 'xlsx' | 'csv'

interface Props<TData> {
    table: Table<TData>
}

const DataTableExportButton = <TData,>({ table }: Props<TData>) => {
    const selectedData = table
        .getSelectedRowModel()
        .flatRows.map(({ original }) => original)

    const exportData = useCallback(
        (fileType: TExportFileType) => {
            try {
                if (selectedData.length === 0)
                    return toast.warning('Please select data to export.')

                // TODO Export Data Logic

                // ignore this for now, I'll remove it later
                console.log('Exporting these...', selectedData)

                toast.success(
                    `Kunwari Exported lol ${fileType} with ${selectedData.length} rows`
                )
            } catch (e) {
                console.error(e)
                toast.error('Failed to export')
            }
        },
        [selectedData]
    )

    useEffect(() => {
        const exportShortcut = (event: KeyboardEvent) => {
            if (event.metaKey && event.key === 'x') {
                exportData('xlsx')
                event.preventDefault()
            }

            if (event.metaKey && event.key === 'c') {
                exportData('csv')
                event.preventDefault()
            }
        }

        window.addEventListener('keydown', exportShortcut)
        return () => {
            window.removeEventListener('keydown', exportShortcut)
        }
    }, [exportData])

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button
                    size="sm"
                    variant={'secondary'}
                    className="gap-x-1 rounded-md"
                >
                    <ExportIcon className="size-4" />
                    Export
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="min-w-44">
                {/* <DropdownMenuLabel>Export Format</DropdownMenuLabel> */}
                {/* <DropdownMenuSeparator /> */}
                <DropdownMenuGroup>
                    {selectedData.length > 0 && (
                        <DropdownMenuItem onClick={() => exportData('csv')}>
                            <CsvIcon className="mr-2 size-4 text-emerald-600" />
                            <span>Export Selected</span>
                        </DropdownMenuItem>
                    )}
                    <DropdownMenuItem onClick={() => exportData('csv')}>
                        <CsvIcon className="mr-2 size-4 text-emerald-600" />
                        <span>Export All Filtered</span> 
                        {/* No ID */}
                    </DropdownMenuItem>
                    <DropdownMenuItem onClick={() => exportData('csv')}>
                        <CsvIcon className="mr-2 size-4 text-emerald-600" />
                        <span>Export All</span>
                        {/* No ID, no filters */}
                    </DropdownMenuItem>
                    <DropdownMenuItem onClick={() => exportData('csv')}>
                        <CsvIcon className="mr-2 size-4 text-emerald-600" />
                        <span>Export Current Page</span>
                        {/* With ID's */}
                    </DropdownMenuItem>
                </DropdownMenuGroup>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

export default DataTableExportButton
