import FeedBackTable from '@/components/tables/feedback-table'
import { useState } from 'react'

const AdminFeedbacksTable = () => {

    const [_, setCreateModal] = useState(false)

    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            <FeedBackTable
                className="min-h-[90vh] w-full"
                actionComponent={() => {}}
                toolbarProps={{
                    createActionProps: {
                        onClick: () => setCreateModal(true),
                    },
                }}
            />
        </div>
    )
}

export default AdminFeedbacksTable
