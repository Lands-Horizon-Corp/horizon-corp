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
        <div className="fixed z-10 mt-2 flex h-14 w-full items-center justify-center dark:bg-black">
            <div className='px-5 w-fit bg-white dark:bg-secondary h-full shadow-md rounded-lg'>
                <ReportBuilderToolbar className="grow" editor={editor} />
            </div>
            {open && <div className="w-[16rem] flex-none"></div>}
        </div>
    )
}

export default DocumentBuilderTools
