import { Button } from '@/components/ui/button'
import TimeInBar from '@/components/time-in-bar'
import { SidebarIcon } from '@/components/icons'
import RootNav from '@/components/navbars/root-nav'
import { ThemeToggleMenu } from '@/components/theme-toggle'
import MobileSidebar from '@/components/sidebar/mobile-sidebar'
import { employeeSidebarItems } from '@/modules/employee/components/employee-sidebar'
import { UserBase } from '@/types'

const EmployeeNavbar = () => {
    return (
        <RootNav
            hideIcon
            iconHref="/employee"
            leftContents={
                <>
                    <MobileSidebar
                        items={employeeSidebarItems}
                        triggerComponent={
                            <Button
                                size="icon"
                                variant="secondary"
                                className="size-fit rounded-lg p-2 sm:hidden"
                            >
                                <SidebarIcon className="size-4" />
                            </Button>
                        }
                    />
                </>
            }
            rightContents={
                <>
                    <TimeInBar currentUser={{} as UserBase} />
                    <ThemeToggleMenu />
                </>
            }
        />
    )
}

export default EmployeeNavbar
