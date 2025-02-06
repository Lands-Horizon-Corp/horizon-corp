import PageContainer from '@/components/containers/page-container'
import MemberTypeTable from '@/components/tables/member-type-table'
import EnsureOwnerCompany from '@/modules/owner/components/ensure-company'

const OwnerViewMembersPage = () => {
    return (
        <PageContainer>
            <EnsureOwnerCompany disabled>
                <MemberTypeTable />
            </EnsureOwnerCompany>
        </PageContainer>
    )
}

export default OwnerViewMembersPage
