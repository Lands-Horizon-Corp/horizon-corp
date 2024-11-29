import { Editor } from '@tiptap/react'
import React from 'react'
import ReactDOMServer from 'react-dom/server'
import { Button } from '../ui/button'
import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarHeader,
} from '@/components/ui/sidebar'

const headers = ['Id', 'Name', 'Bday', 'Age', 'Gender']

const Person: any[] = []

for (let i = 1; i <= 5; i++) {
    const randomYear = Math.floor(Math.random() * (2005 - 1970 + 1)) + 1970 // Random year between 1970 and 2005
    const randomMonth = Math.floor(Math.random() * 12) + 1 // Random month between 1 and 12
    const randomDay = Math.floor(Math.random() * 28) + 1 // Random day between 1 and 28
    const bday = `${randomYear}-${String(randomMonth).padStart(2, '0')}-${String(randomDay).padStart(2, '0')}`
    const age = new Date().getFullYear() - randomYear // Calculate age based on the year
    const gender = Math.random() > 0.5 ? 'Male' : 'Female' // Random gender
    const name = `Person ${i}` // Generate a placeholder name

    Person.push({
        id: i,
        bday: bday,
        name: name,
        age: age,
        gender: gender,
    })
}

interface TableRowData {
    [key: string]: string | number | null | undefined
}

interface TableContentProps {
    headers: string[]
    data: TableRowData[]
}

const TableContent: React.FC<TableContentProps> = ({ headers, data }) => {
    return (
        <table>
            <thead>
                <tr>
                    {headers.map((header) => (
                        <th key={header}>{header}</th>
                    ))}
                </tr>
            </thead>
            <tbody>
                {data.map((row, rowIndex) => (
                    <tr key={rowIndex}>
                        {headers.map((header) => (
                            <td key={header}>
                                {row[header.toLowerCase()] || ''}
                            </td>
                        ))}
                    </tr>
                ))}
            </tbody>
        </table>
    )
}

interface DocumenetSidePanelProps {
    editor: Editor | null
}

const DocumenetSidePanel = ({ editor }: DocumenetSidePanelProps) => {
    const insertTable = () => {
        console.log(
            editor?.commands.insertContent(generateTableHTML(headers, Person))
        )
    }
    return (
        <Sidebar side='right' className='z-[999]'>
            <SidebarHeader />
            <SidebarContent className='flex border'>
                <div className="">
                    <Button
                        className="rounded-md px-4 py-2 text-white shadow-md"
                        onClick={insertTable}
                    >
                        insert table
                    </Button>
                </div>
            </SidebarContent>
            <SidebarFooter />
        </Sidebar>
    )
}

export const generateTableHTML = (
    headers: string[],
    data: TableRowData[]
): string => {
    return ReactDOMServer.renderToStaticMarkup(
        <TableContent headers={headers} data={data} />
    )
}

export default DocumenetSidePanel
