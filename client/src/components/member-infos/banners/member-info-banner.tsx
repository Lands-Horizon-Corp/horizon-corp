import { useState } from 'react'

import UserAvatar from '@/components/user-avatar'
import { Badge } from '@/components/ui/badge'
import CopyTextButton from '@/components/copy-text-button'
import { BadgeCheckFillIcon, BadgeQuestionFillIcon } from '@/components/icons'
import { MemberProfileCreateUpdateFormModal } from '@/components/forms/member-forms/member-application-form/member-profile-create-update-form'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { IMemberProfileResource } from '@/server'

interface Props extends IBaseCompNoChild {
    memberProfile: IMemberProfileResource
}

const MemberInfoBanner = ({ className, memberProfile }: Props) => {
    const [editProfile, setEditProfile] = useState(false)

    return (
        <div className={cn('flex justify-between', className)}>
            <MemberProfileCreateUpdateFormModal
                open={editProfile}
                onOpenChange={setEditProfile}
                formProps={{
                    profileId: memberProfile.id,
                    defaultValues: memberProfile,
                }}
            />
            <div className="flex gap-x-6">
                <UserAvatar
                    src={memberProfile.media?.downloadURL ?? ''}
                    className="size-32 !rounded-xl"
                    fallbackClassName="!rounded-xl"
                />
                <div className="space-y-1">
                    <p className="text-xl">{`${memberProfile.member?.fullName ?? 'no name'}`}</p>
                    <p className="text-sm text-muted-foreground/80">
                        {`${memberProfile.member?.email ?? 'no email'}`}
                        {memberProfile.member && (
                            <>
                                {memberProfile.member.isEmailVerified ? (
                                    <BadgeCheckFillIcon className="ml-1 inline text-primary" />
                                ) : (
                                    <BadgeQuestionFillIcon className="ml-1 inline text-amber-400" />
                                )}
                            </>
                        )}
                    </p>
                    <p className="text-sm text-muted-foreground/80">
                        {`${memberProfile.member?.contactNumber ?? 'no email'}`}
                        {memberProfile.member && (
                            <>
                                {memberProfile.member.isContactVerified ? (
                                    <BadgeCheckFillIcon className="ml-1 inline text-primary" />
                                ) : (
                                    <BadgeQuestionFillIcon className="ml-1 inline text-amber-400" />
                                )}
                            </>
                        )}
                    </p>
                    <p className="!mt-4 truncate whitespace-nowrap text-muted-foreground/90">
                        {`${memberProfile.passbookNumber ?? 'PB Missing'}`}
                        {memberProfile.passbookNumber && (
                            <CopyTextButton
                                className="ml-2"
                                textContent={memberProfile.passbookNumber}
                            />
                        )}
                    </p>
                </div>
            </div>
            <div className="flex flex-col items-end space-y-1.5">
                <span>
                    <Badge variant="secondary" className="gap-1.5">
                        {memberProfile.memberType?.name}
                    </Badge>
                </span>
                <p className="!mt-4 truncate whitespace-nowrap text-xs text-muted-foreground/60">
                    <span>Profile ID: </span>
                    {`${memberProfile.id ?? 'no id'}`}
                    {memberProfile.id && (
                        <CopyTextButton
                            className="ml-1"
                            textContent={memberProfile.id}
                        />
                    )}
                </p>
                <p className="truncate whitespace-nowrap text-xs text-muted-foreground/60">
                    <span>Member ID: </span>
                    {`${memberProfile.member?.id ?? 'no id'}`}
                    {memberProfile.member?.id && (
                        <CopyTextButton
                            className="ml-1"
                            textContent={memberProfile.member?.id}
                        />
                    )}
                </p>
            </div>
        </div>
    )
}

export default MemberInfoBanner
