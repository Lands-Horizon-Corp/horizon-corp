import PageContainer from '@/components/containers/page-container'
import MemberTypeTable from '@/components/tables/member-type-table'
import EnsureOwnerCompany from '@/modules/owner/components/ensure-company'

const OwnerViewMembersPage = () => {
    return (
        <PageContainer>
            <EnsureOwnerCompany disabled>
                <MemberTypeTable
                    toolbarProps={{
                        createActionProps: {
                            onClick: () => {},
                        },
                    }}
                    className="min-h-[90vh] w-full"
                />
            </EnsureOwnerCompany>
        </PageContainer>
    )
}

export default OwnerViewMembersPage
