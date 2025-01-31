import FootstepTable from '@/components/tables/footsteps-table'
import PageContainer from '@/components/containers/page-container'
import EnsureOwnerCompany from '@/modules/owner/components/ensure-company'

const OwnerEmployeeFootstepsPage = () => {
    return (
        <PageContainer>
            <EnsureOwnerCompany disabled>
                <FootstepTable
                    className="min-h-[90vh] w-full"
                    mode="team"
                    defaultFilter={{
                        accountType: {
                            mode: 'equal',
                            value: 'Employee',
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

export default OwnerEmployeeFootstepsPage
