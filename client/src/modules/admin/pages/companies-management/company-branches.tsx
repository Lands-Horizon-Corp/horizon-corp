import z from 'zod'
// import { useParams } from '@tanstack/react-router'

export const CompanyIdPathSchema = z.object({
    companyId: z.number({
        required_error: 'Company ID is required',
        invalid_type_error: 'Invalid Company ID',
    }),
})

const CompanyBranchesPage = () => {
    // const pathParams = useParams({ from: '/admin/companies-management/view-companies/$companyId/view-branches' })

    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            X
        </div>
    )
}

export default CompanyBranchesPage
