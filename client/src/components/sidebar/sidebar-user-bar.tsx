import {
    DropdownMenu,
    DropdownMenuItem,
    DropdownMenuGroup,
    DropdownMenuContent,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import {
    DevIcon,
    LogoutIcon,
    QuestionIcon,
    ChevronRightIcon,
} from '@/components/icons'
import UserAvatar from '@/components/user-avatar'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {
    isExpanded: boolean
}

const SidebarUserBar = ({ className, isExpanded }: Props) => {
    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <div
                    className={cn(
                        'group flex cursor-pointer items-center justify-between gap-x-2 rounded-lg bg-none duration-100 ease-out hover:bg-none',
                        className,
                        isExpanded &&
                            'bg-secondary/20 px-3 py-2 hover:bg-secondary/70'
                    )}
                >
                    <span className="inline-flex items-center gap-x-2">
                        <UserAvatar
                            src="https://github.com/shadcn.png"
                            fallback="A"
                            className={cn(
                                'size-9 rounded-[4rem] duration-150 ease-in-out',
                                !isExpanded && 'hover:rounded-[6px]'
                            )}
                        />
                        {isExpanded && <p>Amada</p>}
                    </span>
                    {isExpanded && (
                        <ChevronRightIcon className="opacity-30 delay-200 duration-150 ease-in group-hover:opacity-100" />
                    )}
                </div>
            </DropdownMenuTrigger>
            <DropdownMenuContent
                side="right"
                align="end"
                className="rounded-xl border-secondary/30 !bg-secondary"
            >
                <DropdownMenuGroup>
                    <DropdownMenuItem className="group cursor-pointer rounded-lg text-foreground/70 hover:text-foreground">
                        <QuestionIcon className="mr-2 size-4 text-foreground/60 duration-150 ease-in-out group-hover:text-foreground" />
                        <span>Help</span>
                    </DropdownMenuItem>
                    <DropdownMenuItem className="group cursor-pointer rounded-lg text-foreground/70 hover:text-foreground">
                        <DevIcon className="mr-2 size-4 text-foreground/60 duration-150 ease-in-out group-hover:text-foreground" />
                        <span>Dev Mode</span>
                    </DropdownMenuItem>
                    <DropdownMenuItem className="group cursor-pointer rounded-lg text-foreground/70 hover:text-foreground">
                        <LogoutIcon className="mr-2 size-4 text-foreground/60 duration-150 ease-in-out group-hover:text-foreground" />
                        <span>Sign Out</span>
                    </DropdownMenuItem>
                </DropdownMenuGroup>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

export default SidebarUserBar
