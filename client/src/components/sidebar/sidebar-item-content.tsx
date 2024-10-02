import { forwardRef } from 'react'

import { IconType } from 'react-icons/lib'
import { GoChevronDown } from 'react-icons/go'

import { Button } from '@/components/ui/button'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {
    Icon?: IconType
    text: string
    className?: string

    isSub?: boolean
    expand: boolean
    active?: boolean
    collapse?: boolean

    onClick?: () => void
    onCollapse?: () => void
}

const SidebarItemContent = forwardRef<HTMLDivElement, Props>(
    (
        {
            Icon,
            text,
            expand,
            collapse,
            className,
            active = false,
            isSub = false,
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
                    'group relative flex cursor-pointer items-center justify-between gap-x-3 rounded-lg border border-transparent px-3 py-2 font-light duration-300 ease-in-out hover:bg-secondary/80 hover:text-foreground group-hover:blur-[1px] group-hover:hover:blur-none',
                    active && 'font-medium',
                    !onCollapse &&
                        active &&
                        'bg-secondary/70 dark:bg-secondary/40',
                    !expand && 'w-fit p-1',
                    className
                )}
            >
                {(isSub || active) && (
                    <div
                        className={cn(
                            'absolute left-[-11px] size-1.5 rounded-full bg-secondary duration-300 ease-in-out',
                            active && 'bg-primary',
                            isSub && 'left-[-12px]',
                            // isSub && !expand && "left-[-6.8]",
                            !expand && 'left-[-7.8px]'
                        )}
                    />
                )}
                <span className="flex items-center gap-x-2">
                    {Icon && (
                        <Icon
                            className={cn(
                                'size-6 text-foreground/40 duration-500 group-hover:text-foreground',
                                active && 'text-foreground'
                            )}
                        />
                    )}
                    {expand && text}
                </span>
                {onCollapse && (
                    <Button
                        size="icon"
                        variant="ghost"
                        className={cn(
                            'size-fit p-[3px] hover:bg-transparent',
                            !expand &&
                                'absolute -right-2.5 delay-200 group-hover:bg-secondary'
                        )}
                    >
                        <GoChevronDown
                            className={cn(
                                'size-3 duration-300 ease-in-out',
                                collapse && '-rotate-180 text-primary',
                                !expand && 'size-2'
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
