import { useParams } from '@tanstack/react-router'

import BranchesTable from '@/components/tables/branches-table'
import BranchesTableAdminAction from '@/components/tables/branches-table/row-actions/branches-table-admin-action'

const CompanyBranchesPage = () => {
    const { companyId } = useParams({
        from: '/admin/companies-management/$companyId/branch',
    })

    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            <BranchesTable
                defaultFilter={{
                    companyId: {
                        mode: 'equal',
                        value: companyId,
                        dataType: 'number',
                        isStaticFilter: true,
                        displayText: 'Company',
                    },
                }}
                actionComponent={({ row }) => (
                    <BranchesTableAdminAction row={row} />
                )}
                className="min-h-[90vh] w-full"
            />
        </div>
    )
}

export default CompanyBranchesPage
