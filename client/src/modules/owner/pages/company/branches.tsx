import { useState } from 'react'

import { BranchCreateFormModal } from '@/components/forms'
import BranchesTable from '@/components/tables/branches-table'
import PageContainer from '@/components/containers/page-container'

import EnsureOwnerCompany from '../../components/ensure-company'

import { ICompanyResource } from '@/server'

const OwnerCompanyBranchesPage = () => {
    const [modal, setModal] = useState(false)
    const [company, setCompany] = useState<ICompanyResource | undefined>(
        undefined
    )

    return (
        <PageContainer>
            <EnsureOwnerCompany onSuccess={setCompany}>
                {company && (
                    <>
                        <BranchCreateFormModal
                            open={modal}
                            onOpenChange={setModal}
                            formProps={{
                                disabledFields: ['companyId'],
                                defaultValues: {
                                    companyId: company.id,
                                    ownerId: company.owner?.id,
                                    longitude: company.longitude,
                                    latitude: company.latitude,
                                },
                            }}
                        />
                        <BranchesTable
                            toolbarProps={{
                                createActionProps: {
                                    onClick: () => setModal((val) => !val),
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
                    </>
                )}
            </EnsureOwnerCompany>
        </PageContainer>
    )
}

export default OwnerCompanyBranchesPage
