import { useState } from 'react'
import { useRouter } from '@tanstack/react-router'

import { UserIcon } from '@/components/icons'
import { IMemberTableActionComponentProp } from '../columns'
import { DropdownMenuItem } from '@/components/ui/dropdown-menu'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'
import { MemberCreateUpdateFormModal } from '@/components/forms/member-forms/member-create-update-form'
import { MemberProfileCreateUpdateFormModal } from '@/components/forms/member-forms/member-application-form/member-profile-create-update-form'

import useConfirmModalStore from '@/store/confirm-modal-store'
import { useDeleteMember } from '@/hooks/api-hooks/member/use-member'

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
    const router = useRouter()
    const [editModal, setEditModal] = useState(false)
    const [editAccountModal, setEditAccountModal] = useState(false)

    const { onOpen } = useConfirmModalStore()

    const { isPending: isDeletingMember, mutate: deleteMember } =
        useDeleteMember({
            onSuccess: onDeleteSuccess,
        })

    return (
        <>
            <div onClick={(e) => e.stopPropagation()}>
                <MemberProfileCreateUpdateFormModal
                    title="Update Member Profile"
                    description="Update member profile details"
                    open={editModal}
                    onOpenChange={setEditModal}
                    className="max-w-7xl"
                    formProps={{
                        defaultValues: member.memberProfile,
                        disabledFields: ['memberId'],
                    }}
                />
                <MemberCreateUpdateFormModal
                    title="Update Member Account"
                    description="Update member account details."
                    open={editAccountModal}
                    onOpenChange={setEditAccountModal}
                    formProps={{
                        defaultValues: {
                            ...member,
                            birthDate: new Date(member.birthDate),
                        },
                    }}
                />
            </div>
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
                    text: 'Edit Account',
                    isAllowed: true,
                    onClick: () => {
                        // router.navigate({
                        //     to: '/admin/companies-management/$companyId/edit',
                        //     params: { companyId: member.id },
                        // })
                    },
                }}
                // onView={{
                //     text: 'View',
                //     isAllowed: true,
                //     onClick: () => {
                //         // router.navigate({
                //         //     to: '/admin/companies-management/$companyId/view',
                //         //     params: { companyId: member.id },
                //         // })
                //     },
                // }}
                otherActions={
                    <>
                        {!member.memberProfile ? (
                            <DropdownMenuItem
                                onClick={() => {
                                    router.navigate({
                                        to: `/owner/users/members/$memberId/member-application`,
                                        params: { memberId: member.id },
                                    })
                                }}
                            >
                                <UserIcon className="mr-2" />
                                Setup Profile
                            </DropdownMenuItem>
                        ) : (
                            <DropdownMenuItem
                                onClick={() => setEditModal((val) => !val)}
                            >
                                <UserIcon className="mr-2" />
                                Edit Profile
                            </DropdownMenuItem>
                        )}
                    </>
                }
            />
        </>
    )
}

export default MembersTableOwnerAction
