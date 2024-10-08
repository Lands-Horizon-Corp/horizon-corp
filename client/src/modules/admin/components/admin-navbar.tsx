import { Button } from '@/components/ui/button'
import { SidebarIcon } from '@/components/icons'
import RootNav from '@/components/navbars/root-nav'
import TimeInBar from '@/components/time-in-out/time-in-bar'
import MobileSidebar from '@/components/sidebar/mobile-sidebar'

import { UserBase } from '@/types'
import { adminSidebarItems } from './admin-sidebar'

const AdminNavbar = () => {
    return (
        <RootNav
            hideIcon
            iconHref="/employee"
            className="lg:px-4"
            leftContents={
                <>
                    <MobileSidebar
                        items={adminSidebarItems}
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
                </>
            }
        />
    )
}

export default AdminNavbar
