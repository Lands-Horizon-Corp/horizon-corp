import { useParams } from '@tanstack/react-router'
import { useSuspenseQuery } from '@tanstack/react-query'

import PageContainer from '@/components/containers/page-container'
import MemberApplicationForm from '@/components/forms/member-forms/member-application-form/member-application-form'

import { memberLoader } from '@/hooks/api-hooks/user-member'

const MemberApplication = () => {
    const { memberId } = useParams({
        from: '/owner/users/members/$memberId/member-application',
    })

    const { data: member } = useSuspenseQuery(memberLoader(memberId))

    return (
        <PageContainer>
            <p className="mb-4 self-start text-lg">Member Application</p>
            <MemberApplicationForm
                defaultValues={{
                    memberID: member.id,
                }}
                className="max-w-7xl"
                hiddenFields={['memberID']}
            />
        </PageContainer>
    )
}

export default MemberApplication
