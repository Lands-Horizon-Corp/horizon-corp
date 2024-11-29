import CompaniesTable from '@/components/tables/companies-table'

const AdminViewCompaniesPage = () => {
    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 sm:px-8">
            <CompaniesTable className="w-full" />
        </div>
    )
}

export default AdminViewCompaniesPage
