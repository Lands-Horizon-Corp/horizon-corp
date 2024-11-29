import CompaniesTable from '@/components/tables/companies-table'

const AdminViewCompaniesPage = () => {
    return (
        <div className="flex w-full max-w-full flex-col pb-6 items-center px-4 sm:px-8">
            <CompaniesTable className="w-full min-h-96" />
        </div>
    )
}

export default AdminViewCompaniesPage
