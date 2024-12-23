import MembersTable from "@/components/tables/members-table"

const AdminViewMembersPage = () => {
    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            <MembersTable className="min-h-[90vh] w-full" />
        </div>
    )
}

export default AdminViewMembersPage
