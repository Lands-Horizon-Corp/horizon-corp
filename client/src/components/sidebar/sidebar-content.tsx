import { Link } from '@tanstack/react-router'

import { GoChevronLeft } from 'react-icons/go'

import EcoopLogo from '@/components/ecoop-logo'
import { Button } from '@/components/ui/button'
import SidebarItem from '@/components/sidebar/sidebar-item'
import SidebarUserBar from '@/components/sidebar/sidebar-user-bar'

import { cn } from '@/lib/utils'
import { IBaseComp } from '@/types/component'
import { useSidebarContext } from './sidebar-context'
import type { TSidebarItem } from '@/components/sidebar/sidebar-types'



export interface ISidebarProps extends IBaseComp {
    items: TSidebarItem[]
    enableCollapse?: boolean
    logoRedirectUrl?: string
    enableFocusBlur?: boolean
    defaultExpanded?: boolean
}
const SidebarContent = ({
    items,
    className,
    logoRedirectUrl = '/',
    enableCollapse = false,
    enableFocusBlur = false,
}: ISidebarProps) => {

    const { isExpanded, toggleExpanded, isMobileSidebarVisible } =
        useSidebarContext()


    return (
        <div
            className={cn(
                'flex h-screen min-w-[260px] flex-col gap-y-2 bg-background pt-4 transition-none duration-300 ease-in-out',
                className,
                !isExpanded && 'min-w-fit',
                enableCollapse && 'border-r'
            )}
        >
            <div
                className={cn(
                    'relative flex items-center justify-between px-5 py-3',
                    isExpanded && 'flex-col'
                )}
            >
                <div className="relative z-0 flex w-fit items-center justify-center">
                    <Link to={logoRedirectUrl}>
                        <EcoopLogo
                            className={cn('size-10', isExpanded && 'size-24')}
                        />
                    </Link>
                </div>
                {enableCollapse && !isMobileSidebarVisible && (
                    <Button
                        size="icon"
                        variant="secondary"
                        onClick={() => toggleExpanded(!isExpanded)}
                        className={cn(
                            'absolute -right-2 z-50 size-fit p-1',
                            isExpanded && '-right-3'
                        )}
                    >
                        <GoChevronLeft
                            className={cn(
                                'size-3 rotate-180 cursor-pointer duration-300 ease-in-out',
                                isExpanded && 'size-4 rotate-0'
                            )}
                        />
                    </Button>
                )}
            </div>
            <div className="relative max-h-full flex-1 overflow-y-hidden">
                <div className="pointer-events-none absolute left-0 top-0 z-10 h-5 w-full bg-gradient-to-b from-background to-transparent" />
                <div
                    className={cn(
                        'ecoop-scroll z-0 flex max-h-full flex-col gap-y-2 overflow-y-scroll px-4 py-4',
                        isExpanded && 'gap-y-1',
                        enableFocusBlur && 'group'
                    )}
                >
                    {items.map((prop, key) => (
                        <SidebarItem key={key} {...prop} />
                    ))}
                </div>
                <div className="pointer-events-none absolute bottom-0 left-0 z-10 h-5 w-full bg-gradient-to-t from-background to-transparent" />
            </div>
            <div className="px-4 pb-4">
                <SidebarUserBar isExpanded={isExpanded} />
            </div>
        </div>
    )
}

export default SidebarContent
