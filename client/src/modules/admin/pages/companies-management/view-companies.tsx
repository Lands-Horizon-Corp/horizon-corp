import { useState } from 'react'
import { useQueryClient } from '@tanstack/react-query'

import { CompanyCreateFormModal } from '@/components/forms'
import CompaniesTable from '@/components/tables/companies-table'
import CompaniesTableAdminAction from '@/components/tables/companies-table/row-actions/companies-table-admin-action'

const AdminViewCompaniesPage = () => {
    const queryClient = useQueryClient();
    const [createModal, setCreateModal] = useState(false)

    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            <CompanyCreateFormModal
                open={createModal}
                onOpenChange={(newState) => setCreateModal(newState)}
                formProps={{
                    onSuccess : () => queryClient.invalidateQueries({
                        queryKey : ['company', 'resource-query']
                    })
                }}
            />
            <CompaniesTable
                className="min-h-[90vh] w-full"
                actionComponent={({ row }) => (
                    <CompaniesTableAdminAction row={row} />
                )}
                toolbarProps={{
                    createActionProps: {
                        onClick: () => setCreateModal(true),
                    },
                }}
            />
        </div>
    )
}

export default AdminViewCompaniesPage
