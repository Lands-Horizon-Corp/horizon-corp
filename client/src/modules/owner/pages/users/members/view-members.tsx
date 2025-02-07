import { useState } from 'react'

import MembersTable from '@/components/tables/members-table'
import PageContainer from '@/components/containers/page-container'
import EnsureOwnerCompany from '@/modules/owner/components/ensure-company'
import { MemberCreateFormModal } from '@/components/forms/member-forms/member-create-form'

import { ICompanyResource } from '@/server'

const OwnerViewMembersPage = () => {
    const [createModal, setCreateModal] = useState(false)
    const [company, setCompany] = useState<undefined | ICompanyResource>()

    return (
        <PageContainer>
            <MemberCreateFormModal
                open={createModal}
                onOpenChange={setCreateModal}
            />
            <EnsureOwnerCompany disabled onSuccess={setCompany}>
                <MembersTable
                    toolbarProps={{
                        createActionProps: {
                            onClick: () => setCreateModal(true),
                        },
                    }}
                    defaultFilter={{
                        companyId: {
                            mode: 'equal',
                            displayText: '',
                            dataType: 'number',
                            value: company?.id,
                        },
                    }}
                    className="min-h-[90vh] w-full"
                />
            </EnsureOwnerCompany>
        </PageContainer>
    )
}

export default OwnerViewMembersPage
