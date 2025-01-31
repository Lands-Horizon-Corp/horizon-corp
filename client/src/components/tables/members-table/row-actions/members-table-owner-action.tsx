import RowActionsGroup from '@/components/data-table/data-table-row-actions'

import { IMemberTableActionComponentProp } from '../columns'

import useConfirmModalStore from '@/store/confirm-modal-store'
import { useDeleteMember } from '@/hooks/api-hooks/user-member'

interface IMembersTableOwnerActionProps
    extends IMemberTableActionComponentProp {
    onMemberUpdate?: () => void
    onDeleteSuccess?: () => void
}

const MembersTableOwnerAction = ({
    row,
    onDeleteSuccess,
}: IMembersTableOwnerActionProps) => {
    const member = row.original

    const { onOpen } = useConfirmModalStore()

    const { isPending: isDeletingMember, mutate: deleteMember } =
        useDeleteMember({
            onSuccess: onDeleteSuccess,
        })

    return (
        <RowActionsGroup
            onDelete={{
                text: 'Delete',
                isAllowed: !isDeletingMember,
                onClick: () => {
                    onOpen({
                        title: 'Delete Member',
                        description: 'Are you sure to delete this Member?',
                        onConfirm: () => deleteMember(member.id),
                    })
                },
            }}
            onEdit={{
                text: 'Edit',
                isAllowed: true,
                onClick: () => {
                    // router.navigate({
                    //     to: '/admin/companies-management/$companyId/edit',
                    //     params: { companyId: member.id },
                    // })
                },
            }}
            onView={{
                text: 'View',
                isAllowed: true,
                onClick: () => {
                    // router.navigate({
                    //     to: '/admin/companies-management/$companyId/view',
                    //     params: { companyId: member.id },
                    // })
                },
            }}
            otherActions={
                <>
                    {/* {!member. && (
                        <DropdownMenuItem
                            onClick={() => {
                                // onOpen({
                                //     title: 'Approve Company',
                                //     description:
                                //         'Are you sure you want to approve this company? Approval will enable them to begin their operations.',
                                //     onConfirm: () => approveCompany(company.id),
                                // })
                            }}
                        >
                            {isApprovingCompany ? (
                                <LoadingSpinner />
                            ) : (
                                <BadgeCheckIcon className="mr-2" />
                            )}{' '}
                            Approve
                        </DropdownMenuItem>
                    )} */}
                </>
            }
        />
    )
}

export default MembersTableOwnerAction
