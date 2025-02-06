import { forwardRef } from 'react'

import { IconType } from 'react-icons/lib'
import { GoChevronDown } from 'react-icons/go'

import { Button } from '@/components/ui/button'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component'

interface Props extends IBaseCompNoChild {
    text: string
    isSub?: boolean
    Icon?: IconType
    className?: string
    isActive?: boolean
    isCollapsed?: boolean
    isExpanded: boolean
    onClick?: () => void
    onCollapse?: () => void
}

const SidebarItemContent = forwardRef<HTMLDivElement, Props>(
    (
        {
            Icon,
            text,
            className,
            isExpanded,
            isCollapsed,
            isSub = false,
            isActive = false,
            onCollapse,
            onClick,
            ...other
        },
        ref
    ) => {
        return (
            <div
                ref={ref}
                {...other}
                onClick={onCollapse ?? onClick}
                className={cn(
                    'group/navself relative flex cursor-pointer items-center justify-between gap-x-3 rounded-lg border border-transparent px-2 py-2 text-sm font-light text-foreground/80 duration-300 ease-in-out hover:bg-secondary/85 hover:text-foreground',
                    isActive && 'text-foreground',
                    !onCollapse &&
                        isActive &&
                        'bg-secondary/70 dark:bg-secondary/40',
                    !isExpanded && 'w-fit p-1',
                    className
                )}
            >
                {(isSub || isActive) && (
                    <div
                        className={cn(
                            'absolute left-[-11px] h-1/2 w-1.5 rounded-full bg-transparent delay-100 duration-300 ease-out group-hover/navself:bg-primary',
                            isActive && 'size-1.5 bg-primary',
                            isSub && 'left-[-12px]',
                            !isExpanded && 'left-[-7.8px]'
                        )}
                    />
                )}
                <span className="flex items-center gap-x-2">
                    {Icon && (!isSub || !isExpanded) && (
                        <Icon
                            className={cn(
                                'size-6 text-foreground/40 duration-500 group-hover/navself:text-foreground',
                                isActive && 'text-foreground'
                            )}
                        />
                    )}
                    {isExpanded && text}
                </span>
                {onCollapse && (
                    <Button
                        size="icon"
                        variant="ghost"
                        className={cn(
                            'size-fit p-[3px] hover:bg-transparent',
                            !isExpanded &&
                                'absolute -right-2.5 delay-200 group-hover:bg-secondary'
                        )}
                    >
                        <GoChevronDown
                            className={cn(
                                'size-3 duration-300 ease-in-out',
                                isCollapsed && '-rotate-180 text-primary',
                                !isExpanded && 'size-2'
                            )}
                        />
                    </Button>
                )}
            </div>
        )
    }
)

SidebarItemContent.displayName = 'SidebarItemContent'

export default SidebarItemContent
