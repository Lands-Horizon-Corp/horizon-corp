import CompaniesTable from '@/components/tables/companies-table'
import logger from '@/helpers/loggers/logger'

const AdminViewCompaniesPage = () => {
    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            <CompaniesTable
                className="min-h-[90vh] w-full"
                onSelectedData={(data) => logger.log(data)}
            />
        </div>
    )
}

export default AdminViewCompaniesPage
