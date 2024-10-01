import { createContext, useContext, useState } from 'react'

import { GoChevronLeft } from 'react-icons/go'

import { Button } from '@/components/ui/button'
import SidebarItem from '@/components/sidebar/sidebar-item'

import { cn } from '@/lib/utils'
import { IBaseComp } from '@/types/component/base'
import { ISidebarItem } from '@/types/component/sidebar'

interface Props extends IBaseComp {
    items: ISidebarItem[]
    enableCollapse?: boolean
}

const ExpandContext = createContext<boolean | null>(true)

export const useSidebarContext = () => {
    const sidebarContext = useContext(ExpandContext)

    if (sidebarContext === null)
        throw new Error('This hook must only be used within the sidebar')

    return sidebarContext
}

/*

TODO:
  Tooltip on !expand
  space-y- on !expand
  mobile view implement sheet form shadcn

  Other routes & Pages

  on Bottom current user + sign out, help, dev mode

  add Batch Floating fixed bottom right Ecoop Beta 0.0.1
*/

const Sidebar = ({ className, items, enableCollapse = false }: Props) => {
    const [expand, setExpand] = useState(true)

    return (
        <div
            className={cn(
                'flex h-screen min-w-[260px] flex-col gap-y-4 border-r bg-background py-4 duration-300 ease-in-out',
                className,
                !expand && 'min-w-fit'
            )}
        >
            <div
                className={cn(
                    'relative flex items-center justify-between px-5 py-3',
                    expand && 'flex-col'
                )}
            >
                <div className="relative z-0 flex w-fit items-center justify-center">
                    <img
                        src="/e-coop-logo-1.png"
                        className={cn('z-10 size-24', !expand && 'size-6')}
                    />
                    <img
                        src="/e-coop-logo-1.png"
                        className={cn(
                            'absolute inset-0 left-1/2 top-1/2 z-0 size-28 -translate-x-1/2 -translate-y-1/2 blur-lg selection:bg-none',
                            !expand && 'size-6'
                        )}
                    />
                </div>
                {enableCollapse && (
                    <Button
                        size="icon"
                        variant="secondary"
                        onClick={() => setExpand((val) => !val)}
                        className={cn(
                            'absolute -right-2 z-10 size-fit p-1',
                            expand && '-right-3'
                        )}
                    >
                        <GoChevronLeft
                            className={cn(
                                'size-2 rotate-180 cursor-pointer duration-300 ease-in-out',
                                expand && 'size-4 rotate-0'
                            )}
                        />
                    </Button>
                )}
            </div>
            <ExpandContext.Provider value={expand}>
                <div className="relative h-full max-h-full flex-1 overflow-y-hidden">
                    <div className="pointer-events-none absolute left-0 top-0 z-10 h-5 w-full bg-gradient-to-b from-background to-transparent" />
                    <div className="ecoop-scroll z-0 flex max-h-full flex-col gap-y-1 overflow-y-scroll px-4 py-4">
                        {items.map((prop, key) => (
                            <SidebarItem key={key} {...prop} />
                        ))}
                    </div>
                    <div className="pointer-events-none absolute bottom-0 left-0 z-10 h-5 w-full bg-gradient-to-t from-background to-transparent" />
                </div>
            </ExpandContext.Provider>
        </div>
    )
}

export default Sidebar
