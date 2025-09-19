import Notification from '@/components/notifications'

const AdminNotificationsPageTable = () => {
    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 pt-10 sm:px-8">
            <Notification userId={101} />
        </div>
    )
}

export default AdminNotificationsPageTable
