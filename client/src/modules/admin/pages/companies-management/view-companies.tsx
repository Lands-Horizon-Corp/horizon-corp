import CompaniesTable from '@/components/tables/companies-table'

const AdminViewCompaniesPage = () => {
    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            <CompaniesTable onSelectData={(data)=> console.log('data', data) } className="min-h-[90vh] w-full" />
        </div>
    )
}

export default AdminViewCompaniesPage
