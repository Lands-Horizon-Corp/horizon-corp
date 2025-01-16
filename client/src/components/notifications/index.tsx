import { useMemo, useState } from 'react'
import { NotificationsResource } from '@/horizon-corp/types'
import { cn } from '@/lib'

import { Button } from '../ui/button'
import NotificationsItems from './notifications-items'
import BumpOnClick from '../bump-element'

import { CheckIcon, OptionsIcon, SettingsIcon } from '../icons'

import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import { BellIcon } from 'lucide-react'
import { Role } from '@/types'
// import { useNotifications } from '@/hooks/api-hooks/use-notifications'

enum NotificationGroup {
    all = 'all',
    unRead = 'unRead',
}

type NotificationGroupType = NotificationGroup

const NotificationsTabs = ({
    selectedTab,
    onSelectTab,
}: {
    selectedTab: NotificationGroupType
    onSelectTab: (tab: NotificationGroupType) => void
}) => (
    <div className="flex space-x-2">
        {Object.values(NotificationGroup).map((tab) => (
            <BumpOnClick key={tab}>
                <Button
                    variant="outline"
                    role="tab"
                    aria-selected={selectedTab === tab}
                    className={cn(
                        'rounded-full border-0 bg-transparent font-bold',
                        selectedTab === tab
                            ? 'bg-primary/20 text-primary hover:bg-primary/30 dark:bg-[#424247]'
                            : 'hover:bg-secondary'
                    )}
                    onClick={() => onSelectTab(tab)}
                >
                    {tab === NotificationGroup.all ? 'All' : 'Unread'}
                </Button>
            </BumpOnClick>
        ))}
    </div>
)

const NotificationsPopover = ({
    onMarkAllAsRead,
}: {
    onMarkAllAsRead: () => void
}) => (
    <Popover>
        <PopoverTrigger>
            <BumpOnClick duration={200}>
                <div
                    className={`cursor-pointer rounded-full p-2 text-[#222225] hover:bg-secondary dark:text-accent-foreground dark:hover:bg-[#303033]`}
                >
                    <OptionsIcon />
                </div>
            </BumpOnClick>
        </PopoverTrigger>
        <PopoverContent
            align="end"
            className="border-0 p-0 shadow-lg dark:bg-[#27272b]"
        >
            <div className="flex flex-col">
                <Button
                    variant="outline"
                    onClick={onMarkAllAsRead}
                    className={cn(
                        'w-full justify-start space-x-2 rounded-b-none border-0 bg-transparent text-xs dark:hover:bg-[#35353a]'
                    )}
                >
                    <CheckIcon size={19} />
                    <span>Mark all as read</span>
                </Button>
                <Button
                    variant="outline"
                    className={cn(
                        'w-full justify-start space-x-2 rounded-t-none border-0 bg-transparent text-xs dark:hover:bg-[#35353a]'
                    )}
                >
                    <SettingsIcon size={19} />
                    <span>Notification Settings</span>
                </Button>
            </div>
        </PopoverContent>
    </Popover>
)

interface NotificationProps {
    userId: number
    role: Role
    // fetchNotifications: (userId: number, role: Role) => Promise<NotificationsResource[]>;
}

const Notification = ({}: NotificationProps) => {
    const [selectedTab, setSelectedTab] = useState<NotificationGroupType>(
        NotificationGroup.all
    )

    // const { data } = useNotifications({ userId, role })

    const [notifications, setNotifications] = useState<NotificationsResource[]>(
        [
            {
                id: 101,
                type: 'error',
                message: 'Database connection failed on Server #3.',
                title: 'Critical System Alert',
                createdAt: '2025-01-12T09:45:00Z',
                actionUrl: '/admin/system/logs/101',
                read: false,
                priority: 'high',
                userId: null,
                category: 'System',
            },
            {
                id: 102,
                type: 'warning',
                message: 'There are 5 pending user account approvals.',
                title: 'Pending Approvals',
                createdAt: '2025-01-12T08:30:00Z',
                actionUrl: '/admin/users/pending-approvals',
                read: false,
                priority: 'medium',
                userId: null,
                category: 'User Management',
            },
            {
                id: 103,
                type: 'success',
                message: 'Backup process completed successfully.',
                title: 'Backup Complete',
                createdAt: '2025-01-12T07:15:00Z',
                actionUrl: '/admin/system/backup-logs',
                read: true,
                priority: 'low',
                userId: null,
                category: 'System',
            },
            {
                id: 104,
                type: 'info',
                message: 'The weekly analytics report is ready for review.',
                title: 'Analytics Report',
                createdAt: '2025-01-11T22:00:00Z',
                actionUrl: '/admin/reports/weekly',
                read: false,
                priority: 'medium',
                userId: null,
                category: 'Analytics',
            },
            {
                id: 105,
                type: 'error',
                message: 'Payment processing failed for Invoice #5678.',
                title: 'Payment Failure',
                createdAt: '2025-01-12T05:00:00Z',
                actionUrl: '/admin/payments/failure/5678',
                read: false,
                priority: 'high',
                userId: 42,
                category: 'Finance',
            },
            {
                id: 106,
                type: 'success',
                message: "New user 'John Doe' has registered.",
                title: 'New Registration',
                createdAt: '2025-01-12T06:00:00Z',
                actionUrl: '/admin/users/42',
                read: true,
                priority: 'low',
                userId: 42,
                category: 'User Management',
            },
            {
                id: 107,
                type: 'info',
                message:
                    'Version 2.5.0 of the application has been deployed successfully.',
                title: 'Deployment Success',
                createdAt: '2025-01-12T03:30:00Z',
                actionUrl: '/admin/deployments/2.5.0',
                read: true,
                priority: 'low',
                userId: null,
                category: 'System',
            },
            {
                id: 108,
                type: 'warning',
                message: 'The API usage limit is at 85%.',
                title: 'API Usage Warning',
                createdAt: '2025-01-11T21:15:00Z',
                actionUrl: '/admin/api/usage',
                read: false,
                priority: 'medium',
                userId: null,
                category: 'System',
            },
        ]
    )

    const [readNotifications, setReadNotifications] = useState<Set<number>>(
        new Set(
            notifications.filter((item) => item.read).map((item) => item.id)
        )
    )

    const handleDeleteNotification = (id: number) => {
        setNotifications(notifications.filter((item) => item.id !== id))
    }

    const handleMarkItemAsRead = (id: number) => {
        setReadNotifications((prev) => new Set([...prev, id]))
    }

    const handleMarkAllAsRead = () => {
        setReadNotifications(new Set(notifications.map((item) => item.id)))
    }

    const filteredNotifications = useMemo(() => {
        return notifications
            .map((item) => ({
                ...item,
                read: readNotifications.has(item.id),
            }))
            .filter(
                (item) => selectedTab === NotificationGroup.all || !item.read
            )
    }, [notifications, readNotifications, selectedTab])

    return (
        <div className="mx-auto w-[100%] space-y-4 rounded-lg bg-[#fcfcfc] p-4 shadow-md dark:border dark:bg-[#222225] xl:w-[80%]">
            <div className="flex items-center justify-between">
                <div className="flex flex-col">
                    <h1 className="text-xl font-bold">Notifications</h1>
                    <p className="text-sm text-muted-foreground">
                        You have {filteredNotifications.length} notifications.
                    </p>
                </div>
                <NotificationsPopover onMarkAllAsRead={handleMarkAllAsRead} />
            </div>
            <NotificationsTabs
                selectedTab={selectedTab}
                onSelectTab={setSelectedTab}
            />

            <div className="">
                {filteredNotifications.length > 0 ? (
                    filteredNotifications.map((notif) => (
                        <div key={notif.id}>
                            <NotificationsItems
                                handleDeleteNotification={
                                    handleDeleteNotification
                                }
                                handleMarkItemAsRead={handleMarkItemAsRead}
                                data={notif}
                            />
                        </div>
                    ))
                ) : (
                    <div className="flex h-32 items-center justify-center space-x-2">
                        <BellIcon size={30} />
                        <p>No notifications found.</p>
                    </div>
                )}
            </div>
        </div>
    )
}

export default Notification
