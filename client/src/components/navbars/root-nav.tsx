import { Link, ReactNode } from '@tanstack/react-router'

import EcoopLogo from '@/components/ecoop-logo'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {
    iconHref?: string
    hideIcon?: boolean
    midContents?: ReactNode
    leftContents?: ReactNode
    rightContents?: ReactNode
    midGroupClassName?: string
    leftGroupClassName?: string
    rightGroupClassName?: string
}

const RootNav = ({
    className,
    midContents,
    leftContents,
    rightContents,
    iconHref = '/',
    hideIcon = false,
    midGroupClassName,
    leftGroupClassName,
    rightGroupClassName,
}: Props) => {
    return (
        <nav
            className={cn(
                'z-10 flex min-h-[68px] items-center justify-between gap-x-2 px-4 lg:px-16',
                className
            )}
        >
            <div
                className={cn(
                    'flex w-fit items-center justify-center',
                    leftGroupClassName
                )}
            >
                <Link to={iconHref}>
                    <EcoopLogo
                        className={cn('size-[46px]', hideIcon && 'hidden')}
                    />
                </Link>
                {leftContents}
            </div>
            <div
                className={cn(
                    'flex h-full items-center justify-center space-x-[14px] text-[15px]',
                    midGroupClassName
                )}
            >
                {midContents}
            </div>
            <div
                className={cn(
                    'flex h-full items-center justify-center gap-x-4',
                    rightGroupClassName
                )}
            >
                {rightContents}
            </div>
        </nav>
    )
}

export default RootNav
