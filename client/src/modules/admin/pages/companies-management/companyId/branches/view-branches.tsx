import { useParams } from '@tanstack/react-router'
import BranchesTable from '@/components/tables/branches-table'

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
                className="min-h-[90vh] w-full"
            />
        </div>
    )
}

export default CompanyBranchesPage
