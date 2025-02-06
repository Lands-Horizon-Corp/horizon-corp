import { useState } from 'react'

import PageContainer from '@/components/containers/page-container'
import MemberTypeTable from '@/components/tables/member-type-table'
import EnsureOwnerCompany from '@/modules/owner/components/ensure-company'
import { MemberTypeCreateFormModal } from '@/components/forms/member-forms/member-type-create-form'

const OwnerViewMembersPage = () => {
    const [modalState, setModalState] = useState(false)

    return (
        <PageContainer>
            <EnsureOwnerCompany disabled>
                <MemberTypeCreateFormModal
                    open={modalState}
                    onOpenChange={setModalState}
                />
                <MemberTypeTable
                    toolbarProps={{
                        createActionProps: {
                            onClick: () => setModalState(true),
                        },
                    }}
                    className="min-h-[90vh] w-full"
                />
            </EnsureOwnerCompany>
        </PageContainer>
    )
}

export default OwnerViewMembersPage
