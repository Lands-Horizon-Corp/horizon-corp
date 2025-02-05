import { Editor } from '@tiptap/react'
import React, {
    //  useEffect, useRef 
    } from 'react'
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
// import { useTellerTransactionStore } from '@/store/report-store'
// import ReportGenerator from '../reports/report-generator'
// import { TellerTransactionRecord } from '../reports/types'
// import { getCoreRowModel, useReactTable } from '@tanstack/react-table'
// import TestReport from './test'
// import Test from './test'
import UserTeller from '@/modules/owner/pages/reports/user-teller/user-teller'


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
    menuItems,editor}) => {


    const insertContent = ( ) =>{
        if(editor){
            editor?.commands.insertContent(ReactDOMServer.renderToStaticMarkup(<UserTeller />))
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
                                    <SidebarMenuSubItem onClick={insertContent}
                                        className="flex cursor-pointer items-center rounded-sm px-2 py-2 text-xs hover:bg-secondary"
                                    >
                                        {subItem.icon && (
                                            <span className="mr-2 flex size-4 items-center">
                                                {subItem.icon}
                                            </span>
                                        )}
                                        <span className="grow">
                                            {' '}
                                            {subItem.label}
                                        </span>
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

export default DocumenetSidePanel
