import PageContainer from '@/components/containers/page-container'
import MemberOccupationTable from '@/components/tables/member-occupation-table'

const OwnerOccupationPage = () => {
    return (
        <PageContainer>
            <MemberOccupationTable className="min-h-[90vh] w-full" />
        </PageContainer>
    )
}

export default OwnerOccupationPage
