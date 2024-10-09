import {
    DropdownMenu,
    DropdownMenuSub,
    DropdownMenuItem,
    DropdownMenuGroup,
    DropdownMenuPortal,
    DropdownMenuContent,
    DropdownMenuTrigger,
    DropdownMenuSubTrigger,
    DropdownMenuSubContent,
} from '@/components/ui/dropdown-menu'
import {
    DevIcon,
    SunIcon,
    MoonIcon,
    LogoutIcon,
    SunMoonIcon,
    QuestionIcon,
    ChevronRightIcon,
} from '@/components/icons'
import UserAvatar from '@/components/user-avatar'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'
import { useTheme } from '@/components/providers/theme-provider'

interface Props extends IBaseCompNoChild {
    isExpanded: boolean
}

const SidebarUserBar = ({ className, isExpanded }: Props) => {
    const { setTheme, resolvedTheme } = useTheme()

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
                            src="https://avatars.githubusercontent.com/u/48374007?s=80&v=4"
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
            <DropdownMenuContent side="right" align="end">
                <DropdownMenuGroup>
                    <DropdownMenuItem>
                        <QuestionIcon className="mr-2 size-4 duration-150 ease-in-out" />
                        <span>Help</span>
                    </DropdownMenuItem>
                    <DropdownMenuSub>
                        <DropdownMenuSubTrigger>
                            {resolvedTheme === 'light' && (
                                <SunIcon className="mr-2 size-4" />
                            )}
                            {resolvedTheme === 'dark' && (
                                <MoonIcon className="mr-2 size-4" />
                            )}
                            <span>Theme</span>
                        </DropdownMenuSubTrigger>
                        <DropdownMenuPortal>
                            <DropdownMenuSubContent>
                                <DropdownMenuItem
                                    onClick={() => setTheme('light')}
                                >
                                    <SunIcon className="mr-2 size-4" />
                                    <span>Light</span>
                                </DropdownMenuItem>
                                <DropdownMenuItem
                                    onClick={() => setTheme('dark')}
                                >
                                    <MoonIcon className="mr-2 size-4" />
                                    <span>Dark</span>
                                </DropdownMenuItem>
                                <DropdownMenuItem
                                    onClick={() => setTheme('system')}
                                >
                                    <SunMoonIcon className="mr-2 size-4" />
                                    <span>System</span>
                                </DropdownMenuItem>
                            </DropdownMenuSubContent>
                        </DropdownMenuPortal>
                    </DropdownMenuSub>
                    <DropdownMenuItem>
                        <DevIcon className="mr-2 size-4 duration-150 ease-in-out" />
                        <span>Dev Mode</span>
                    </DropdownMenuItem>
                    <DropdownMenuItem>
                        <LogoutIcon className="mr-2 size-4 duration-150 ease-in-out" />
                        <span>Sign Out</span>
                    </DropdownMenuItem>
                </DropdownMenuGroup>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

export default SidebarUserBar
