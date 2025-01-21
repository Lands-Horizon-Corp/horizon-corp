import PageContainer from "@/components/containers/page-container"
import EnsureOwnerCompany from "@/modules/owner/components/ensure-company"

const OwnerViewEmployeesPage = () => {
    return (
        <PageContainer>
            <EnsureOwnerCompany></EnsureOwnerCompany>
        </PageContainer>
    )
}

export default OwnerViewEmployeesPage
