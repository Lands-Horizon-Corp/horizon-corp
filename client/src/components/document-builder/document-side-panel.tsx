import { Editor } from '@tiptap/react'
import React from 'react'
import ReactDOMServer from 'react-dom/server'
import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
    SidebarMenuSub,
    SidebarMenuSubItem,
} from '@/components/ui/sidebar'

import { Table2Icon } from 'lucide-react'
import {
    Collapsible,
    CollapsibleContent,
    CollapsibleTrigger,
} from '../ui/collapsible'
import {
    FaUsers,
    FaCreditCard,
    FaDollarSign,
    FaBuilding,
    FaExchangeAlt,
} from 'react-icons/fa'
import { MdAdd } from 'react-icons/md'
import InsertTable from './document-insert-table'
// import companiesTableColumns from '../tables/companies-table/columns'
import Test from './test'
const headers = ['Id', 'Name', 'Bday', 'Age', 'Gender']

const Person: any[] = []

for (let i = 1; i <= 50; i++) {
    const randomYear = Math.floor(Math.random() * (2005 - 1970 + 1)) + 1970
    const randomMonth = Math.floor(Math.random() * 12) + 1
    const randomDay = Math.floor(Math.random() * 28) + 1
    const bday = `${randomYear}-${String(randomMonth).padStart(2, '0')}-${String(randomDay).padStart(2, '0')}`
    const age = new Date().getFullYear() - randomYear
    const gender = Math.random() > 0.5 ? 'Male' : 'Female'
    const name = `Person ${i}`

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

interface SidebarSubItem {
    label: string
    onClick?: () => void
    icon?: React.ReactNode
}

interface SidebarMenuItem {
    label: string
    subItems: SidebarSubItem[]
    icon?: React.ReactNode
}

interface SidebarMenuProps {
    menuItems: SidebarMenuItem[]
    editor: Editor | null
}
const SidebarMenuContainer: React.FC<SidebarMenuProps> = ({
    menuItems,
    editor,
}) => {

    const insertTable = (tableName: string) => () => {
        console.log('tableName', tableName)
        if (!tableName) return
        if (tableName === 'Accounts') {
            console.log(ReactDOMServer.renderToStaticMarkup(<Test />))
            editor?.commands.insertContent(ReactDOMServer.renderToStaticMarkup(<Test />))
        }
    }       

    return (
        <div>
            {menuItems.map((item, index) => (
                <SidebarMenu key={index}>
                    <Collapsible>
                        <CollapsibleTrigger asChild>
                            <SidebarMenuButton>
                                {item.icon && (
                                    <span className="mr-2">{item.icon}</span>
                                )}
                                {item.label}
                            </SidebarMenuButton>
                        </CollapsibleTrigger>
                        <CollapsibleContent>
                            {item.subItems.map((subItem, subIndex) => (
                                <SidebarMenuSub key={subIndex}>
                                    <SidebarMenuSubItem onClick={insertTable('Accounts')} className="flex cursor-pointer items-center rounded-sm px-2 py-2 text-xs hover:bg-secondary">
                                        {subItem.icon && (
                                            <span className="mr-2 flex size-4 items-center">
                                                {subItem.icon}
                                            </span>
                                        )}
                                        <span className="grow">
                                            {' '}
                                            {subItem.label}
                                        </span>

                                        {/* <InsertTable
                                            content={<Test />}
                                            onClick={insertTable(subItem.label)}
                                            trigger={<MdAdd size={18} />}
                                        /> */}
                                    </SidebarMenuSubItem>
                                </SidebarMenuSub>
                            ))}
                        </CollapsibleContent>
                    </Collapsible>
                </SidebarMenu>
            ))}
        </div>
    )
}

const menuItems: SidebarMenuItem[] = [
    {
        label: 'Companies Management',
        icon: <FaUsers />,
        subItems: [
            {
                label: 'Companies',
                icon: <FaDollarSign />,
            },
        ],
    },
    {
        label: 'Members Management',
        icon: <FaDollarSign />,
        subItems: [
            {
                label: 'Members',
                icon: <FaExchangeAlt />,
                onClick: () => console.log('Transactions clicked'),
            },
        ],
    },
    {
        label: 'Loans',
        icon: <FaExchangeAlt />,
        subItems: [
            {
                label: 'Loan Payments',
                icon: <FaDollarSign />,
                onClick: () => console.log('Loan Payments clicked'),
            },
            {
                label: 'Late Fees',
                icon: <FaDollarSign />,
                onClick: () => console.log('Late Fees clicked'),
            },
            {
                label: 'Collateral',
                icon: <FaBuilding />,
                onClick: () => console.log('Collateral clicked'),
            },
        ],
    },
    {
        label: 'Cards',
        icon: <FaCreditCard />,
        subItems: [
            {
                label: 'Card Transactions',
                icon: <FaExchangeAlt />,
                onClick: () => console.log('Card Transactions clicked'),
            },
            {
                label: 'Rewards',
                icon: <FaDollarSign />,
                onClick: () => console.log('Rewards clicked'),
            },
            {
                label: 'Limit Adjustments',
                icon: <FaDollarSign />,
                onClick: () => console.log('Limit Adjustments clicked'),
            },
        ],
    },
    {
        label: 'Branches',
        icon: <FaBuilding />,
        subItems: [
            {
                label: 'Employees',
                icon: <FaUsers />,
                onClick: () => console.log('Employees clicked'),
            },
            {
                label: 'ATMs',
                icon: <FaDollarSign />,
                onClick: () => console.log('ATMs clicked'),
            },
            {
                label: 'Services',
                icon: <FaBuilding />,
                onClick: () => console.log('Services clicked'),
            },
        ],
    },
]

interface DocumenetSidePanelProps {
    editor: Editor | null
}

const DocumenetSidePanel = ({ editor }: DocumenetSidePanelProps) => {
    return (
        <Sidebar side="right" className="">
            <SidebarGroup>
                <SidebarGroupLabel>
                    {' '}
                    <Table2Icon className="mr-2" size={18} /> Tables
                </SidebarGroupLabel>
                <SidebarGroupContent>
                    <SidebarContent className="">
                        <SidebarMenuContainer
                            editor={editor}
                            menuItems={menuItems}
                        />
                    </SidebarContent>
                </SidebarGroupContent>
            </SidebarGroup>
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
