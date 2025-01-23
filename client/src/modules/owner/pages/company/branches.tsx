import { useState } from 'react'

import BranchesTable from '@/components/tables/branches-table'
import PageContainer from '@/components/containers/page-container'

import EnsureOwnerCompany from '../../components/ensure-company'

import { ICompanyResource } from '@/server'

const OwnerCompanyBranchesPage = () => {
    const [company, setCompany] = useState<ICompanyResource | undefined>(
        undefined
    )

    return (
        <PageContainer>
            <EnsureOwnerCompany onSuccess={setCompany}>
                {company && (
                    <BranchesTable
                        toolbarProps={{
                            createActionProps: {
                                onClick: () => {},
                            },
                        }}
                        className="min-h-[90vh] w-full"
                        defaultFilter={{
                            companyId: {
                                value: company.id,
                                mode: 'equal',
                                displayText: '',
                                isStaticFilter: true,
                                dataType: 'number',
                            },
                        }}
                    />
                )}
            </EnsureOwnerCompany>
        </PageContainer>
    )
}

export default OwnerCompanyBranchesPage
