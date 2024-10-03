import { Link } from '@tanstack/react-router'
import { createContext, useContext, useState } from 'react'

import { GoChevronLeft } from 'react-icons/go'

import { Button } from '@/components/ui/button'
import SidebarItem from '@/components/sidebar/sidebar-item'
import SidebarUserBar from '@/components/sidebar/sidebar-user-bar'

import { cn } from '@/lib/utils'
import { IBaseComp } from '@/types/component/base'
import type { TSidebarItem } from '@/types/component/sidebar'
import { useTheme } from '@/components/providers/theme-provider'

interface Props extends IBaseComp {
    items: TSidebarItem[]
    defaultExpanded?: boolean
    enableCollapse?: boolean
    enableFocusBlur?: boolean
}

const ExpandContext = createContext<boolean | null>(true)

export const useSidebarExpandContext = () => {
    const sidebarExpandContext = useContext(ExpandContext)

    if (sidebarExpandContext === null)
        throw new Error('This hook must only be used within the sidebar')

    return sidebarExpandContext
}

const Sidebar = ({
    items,
    className,
    defaultExpanded = true,
    enableCollapse = false,
    enableFocusBlur = false,
}: Props) => {
    const { resolvedTheme } = useTheme()
    const [isExpanded, setIsExpanded] = useState(defaultExpanded)

    return (
        <div
            className={cn(
                'flex h-screen min-w-[260px] flex-col gap-y-2 bg-background pt-4 duration-300 ease-in-out',
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
                    <Link to="/member">
                        <img
                            src={
                                resolvedTheme === 'light'
                                    ? '/e-coop-logo-1.png'
                                    : '/e-coop-logo-white.png'
                            }
                            className={cn(
                                'z-10 size-7',
                                isExpanded && 'size-24'
                            )}
                        />
                    </Link>
                    <img
                        src="/e-coop-logo-1.png"
                        className={cn(
                            'pointer-events-none absolute inset-0 left-1/2 top-1/2 z-0 size-8 -translate-x-1/2 -translate-y-[40%] -rotate-90 blur-sm selection:bg-none',
                            isExpanded && 'size-28 blur-lg'
                        )}
                    />
                </div>
                {enableCollapse && (
                    <Button
                        size="icon"
                        variant="secondary"
                        onClick={() => setIsExpanded((val) => !val)}
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
            <ExpandContext.Provider value={isExpanded}>
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
            </ExpandContext.Provider>
            <div className="px-4 pb-4">
                <SidebarUserBar isExpanded={isExpanded} />
            </div>
        </div>
    )
}

export default Sidebar
