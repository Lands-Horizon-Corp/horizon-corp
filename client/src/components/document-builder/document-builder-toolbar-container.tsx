import React from 'react'
import { ReportBuilderToolbar } from './report-builder-toolbar'
import { Editor } from '@tiptap/react'
import { useSidebar } from '../ui/sidebar'
import { cn } from '@/lib'

interface TableToolBarProps extends React.ComponentProps<'div'> {
    editor: Editor | null
    handleGeneratePdf?: () => void
}

const DocumentBuilderTools = ({
    editor,
    handleGeneratePdf,
}: TableToolBarProps) => {
    const { open } = useSidebar()

    return (
        <div
            className={cn(
                ' z-10 flex h-14 w-full  items-center justify-center bg-transparent',
                // 'fixed left-[60%] top-[15%] -translate-x-1/2 -translate-y-1/2'
            )}
        >
            <div className="h-full w-fit rounded-lg bg-white px-5 shadow-md dark:bg-secondary">
                <ReportBuilderToolbar
                    handleGeneratePdf={handleGeneratePdf}
                    className="grow"
                    editor={editor}
                />
            </div>
            {open && <div className="w-[16rem] flex-none"></div>}
        </div>
    )
}

export default DocumentBuilderTools
