import { useState } from 'react'

import PageContainer from '@/components/containers/page-container'
import MemberTypeTable from '@/components/tables/member-type-table'
import EnsureOwnerCompany from '@/modules/owner/components/ensure-company'
import { MemberTypeCreateFormModal } from '@/components/forms/member-forms/member-type-create-form'
import MemberTypeTableOwnerAction from '@/components/tables/member-type-table/row-actions/member-type-owner-action'

const OwnerViewMemberTypesPage = () => {
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
                    actionComponent={(prop) => (
                        <MemberTypeTableOwnerAction {...prop} />
                    )}
                    className="min-h-[90vh] w-full"
                />
            </EnsureOwnerCompany>
        </PageContainer>
    )
}

export default OwnerViewMemberTypesPage
