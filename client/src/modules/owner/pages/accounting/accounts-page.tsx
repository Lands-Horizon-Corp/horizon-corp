import PageContainer from '@/components/containers/page-container'
import AccountsTable from '@/components/tables/accounts-table'

const AccountsPage = () => {
    return (
        <PageContainer>
                <AccountsTable
                    className="min-h-[90vh] w-full"
                    actionComponent={({ row }) => <></>}
                />
        </PageContainer>
    )
}

export default AccountsPage
