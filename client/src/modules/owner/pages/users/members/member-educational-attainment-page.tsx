import { useState } from 'react'

import PageContainer from '@/components/containers/page-container'
import MemberEducationalAttainmentTable from '@/components/tables/member-educational-attainment-table'
import { MemberEducationalAttainmentCreateUpdateFormModal } from '@/components/forms/member-forms/member-educational-attainment-create-update-form'

const OwnerMemberEducationalAttainmentPage = () => {
    const [createModal, setCreateModal] = useState(false)

    return (
        <PageContainer>
            <MemberEducationalAttainmentCreateUpdateFormModal
                open={createModal}
                onOpenChange={setCreateModal}
            />
            <MemberEducationalAttainmentTable
                className="min-h-[90vh] w-full"
                toolbarProps={{
                    createActionProps: { onClick: () => setCreateModal(true) },
                }}
            />
        </PageContainer>
    )
}

export default OwnerMemberEducationalAttainmentPage
