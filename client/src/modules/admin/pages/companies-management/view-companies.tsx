import CompaniesTable from '@/components/tables/companies-table'
import CompaniesTableAdminAction from '@/components/tables/companies-table/row-actions/companies-table-admin-action'

const AdminViewCompaniesPage = () => {
    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            <CompaniesTable
                className="min-h-[90vh] w-full"
                actionComponent={({ row }) => (
                    <CompaniesTableAdminAction row={row} />
                )}
                toolbarProps={{
                    createActionProps: {
                        onClick: () => {},
                    },
                }}
            />
        </div>
    )
}

export default AdminViewCompaniesPage
