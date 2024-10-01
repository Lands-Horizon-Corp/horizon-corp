import { useState } from 'react'
import { useLocation, useRouter } from '@tanstack/react-router'

import { GoChevronDown } from 'react-icons/go'

import { Button } from '@/components/ui/button'
import { useSidebarContext } from '@/components/sidebar'

import { cn } from '@/lib/utils'
import type { ISidebarItem } from '@/types/component/sidebar/index'

interface Props extends ISidebarItem {}

const SidebarItem = ({
    Icon,
    url,
    text,
    baseUrl,
    subItems,
    isSub = false,
    collapsed = false,
}: Props) => {
    const router = useRouter()
    const pathname = useLocation({
        select: (location) => location.pathname,
    })

    const expand = useSidebarContext()

    const [collapse, setcollapse] = useState(collapsed)

    const handleClick = (url: string) => {
        router.navigate({ to: url })
    }

    const handlecollapse = () => setcollapse((prev) => !prev)

    const routeMatched = subItems ? false : pathname.startsWith(url || '')

    return (
        <>
            {subItems && subItems.length > 0 ? (
                <>
                    <div
                        onClick={handlecollapse}
                        className={cn(
                            'relative flex cursor-pointer items-center justify-between gap-x-3 rounded-lg border border-transparent px-3 py-2 duration-300 ease-in-out hover:bg-secondary/80 hover:text-foreground',
                            !expand && 'w-fit p-1'
                        )}
                    >
                        <span className="flex items-center gap-x-2">
                            {Icon && <Icon className="size-4" />}
                            {expand && text}
                        </span>
                        <Button
                            size="icon"
                            variant="ghost"
                            className={cn(
                                'size-fit p-1 hover:bg-transparent',
                                !expand && 'absolute -right-2'
                            )}
                        >
                            <GoChevronDown
                                className={cn(
                                    'size-4 duration-300 ease-in-out',
                                    collapse && '-rotate-180 text-primary',
                                    !expand && 'size-2'
                                )}
                            />
                        </Button>
                    </div>
                    {collapse && (
                        <div
                            className={cn(
                                'ml-3.5 flex flex-col gap-y-1 border-l pl-2',
                                !expand && 'pl-1'
                            )}
                        >
                            {subItems.map(
                                ({ url, isSub = true, ...subItem }, key) => (
                                    <SidebarItem
                                        key={key}
                                        {...subItem}
                                        isSub={isSub}
                                        url={`${baseUrl ?? ''}${url ?? ''}`}
                                    />
                                )
                            )}
                        </div>
                    )}
                </>
            ) : (
                <div
                    onClick={() => (url ? handleClick(url) : undefined)}
                    className={cn(
                        'group relative flex cursor-pointer items-center justify-between gap-x-2 rounded-lg px-3 py-2 text-foreground/80 duration-300 ease-in-out hover:bg-secondary/80 hover:text-foreground',
                        routeMatched && 'bg-secondary/70 dark:bg-secondary/40',
                        !expand && 'w-fit p-1'
                    )}
                >
                    {(isSub || routeMatched) && (
                        <div
                            className={cn(
                                'absolute left-[-10.5px] size-1 rounded-full bg-secondary duration-300 ease-in-out',
                                routeMatched && 'bg-primary',
                                !expand && 'left-[-6.8px]'
                            )}
                        />
                    )}
                    <span className="flex items-center gap-x-2">
                        {Icon && (
                            <Icon
                                className={cn(
                                    'size-4 text-foreground/40 duration-500 group-hover:text-foreground',
                                    routeMatched && 'text-foreground'
                                )}
                            />
                        )}
                        {expand && text}
                    </span>
                </div>
            )}
        </>
    )
}

export default SidebarItem
