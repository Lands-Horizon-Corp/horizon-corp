import React from 'react'
import { ReportBuilderToolbar } from './report-builder-toolbar'
import { Editor } from '@tiptap/react'
import { useSidebar } from '../ui/sidebar'

interface TableToolBarProps extends React.ComponentProps<'div'> {
    editor: Editor | null
}

const DocumentBuilderTools = ({ editor }: TableToolBarProps) => {
    const { open } = useSidebar()

    return (
        <div className="fixed z-20 flex h-14 w-full items-center justify-center border bg-white dark:bg-black">
            <ReportBuilderToolbar className="grow" editor={editor} />
            {open && <div className="w-[16rem] flex-none"></div>}
        </div>
    )
}

export default DocumentBuilderTools
