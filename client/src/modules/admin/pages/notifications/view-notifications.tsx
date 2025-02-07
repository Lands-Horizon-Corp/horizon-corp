import Notification from '@/components/notifications'
import { Role } from '@/types'

const sampleRoles: Role = {
    id: '1',
    name: 'Admin',
    description: 'Administrator role with full permissions.',
    createdAt: new Date('2025-01-01T10:00:00Z'),
    updatedAt: new Date('2025-01-05T12:00:00Z'),
    deletedAt: undefined,
    permissions: [
        {
            id: '101',
            name: 'Manage Users',
            description: 'Allows managing user accounts.',
            read: true,
            readDescription: 'Can view user accounts.',
            update: true,
            updateDescription: 'Can update user accounts.',
            create: true,
            createDescription: 'Can create user accounts.',
            roleId: '1',
            createdAt: new Date('2025-01-01T10:30:00Z'),
            updatedAt: new Date('2025-01-05T12:30:00Z'),
            deletedAt: undefined,
        },
    ],
}

const AdminNotificationsPageTable = () => {
    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 pt-10 sm:px-8">
            <Notification userId={101} role={sampleRoles} />
        </div>
    )
}

export default AdminNotificationsPageTable
