import FootstepTable from '@/components/tables/footsteps-table'
import PageContainer from '@/components/containers/page-container'
import EnsureOwnerCompany from '@/modules/owner/components/ensure-company'

const OwnerMembersActivityPage = () => {
    return (
        <PageContainer>
            <EnsureOwnerCompany disabled>
                <FootstepTable
                    className="min-h-[90vh] w-full"
                    mode="team"
                    defaultFilter={{
                        accountType: {
                            mode: 'equal',
                            value: 'Member',
                            dataType: 'text',
                            isStaticFilter: true,
                            displayText: 'Account Type',
                        },
                    }}
                />
            </EnsureOwnerCompany>
        </PageContainer>
    )
}

export default OwnerMembersActivityPage
