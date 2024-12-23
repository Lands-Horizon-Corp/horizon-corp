import { useParams } from '@tanstack/react-router'
import { useSuspenseQuery } from '@tanstack/react-query'

import CompanyBasicInfoForm from '@/modules/admin/components/forms/company-edit-basic-info-form'
import { companyLoader } from '@/hooks/api-hooks/use-company'

const CompanyEditPage = () => {
    const { companyId } = useParams({
        from: '/admin/companies-management/$companyId/edit',
    })

    const { data: company } = useSuspenseQuery(companyLoader(companyId))

    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            <div className="w-full max-w-5xl space-y-4 p-4">
                <h4 className="font-medium">Edit Company</h4>
                <CompanyBasicInfoForm
                    companyId={companyId}
                    defaultValues={company}
                />
            </div>
        </div>
    )
}

export default CompanyEditPage
