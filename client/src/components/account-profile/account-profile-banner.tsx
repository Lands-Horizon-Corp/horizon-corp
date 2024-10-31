import { useState } from 'react'
import { format } from 'date-fns'

import {
    Dialog,
    DialogTitle,
    DialogHeader,
    DialogContent,
    DialogDescription,
} from '@/components/ui/dialog'
import ProfileUpload from './profile-upload'
import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'
import ActionTooltip from '@/components/action-tooltip'
import {
    EmailIcon,
    CalendarIcon,
    DotMediumIcon,
    CameraFillIcon,
    PhoneOutlineIcon,
    CalendarCheckIcon,
} from '@/components/icons'

import { UserData } from '@/horizon-corp/types'

const AccountProfileBanner = ({
    currentUser,
    updateUserData,
}: {
    currentUser: UserData
    updateUserData: (newUserData: UserData) => void
}) => {
    const [uploadState, toggleProfileUpload] = useState(false)

    return (
        <div className="overflow-clip rounded-2xl bg-secondary shadow-md">
            <Dialog open={uploadState} onOpenChange={toggleProfileUpload}>
                <DialogContent
                    closeButtonClassName="sm:hidden"
                    className="!rounded-2xl"
                >
                    <DialogHeader>
                        <DialogTitle>Change Profile Picture</DialogTitle>
                        <DialogDescription>
                            Change your profile photo
                        </DialogDescription>
                    </DialogHeader>
                    <ProfileUpload
                        currentUser={currentUser}
                        onUploadComplete={(newUserData) => {
                            updateUserData(newUserData)
                            toggleProfileUpload(false)
                        }}
                    />
                </DialogContent>
            </Dialog>

            <div className="h-[180px] bg-[url('/profile-cover.png')] bg-cover bg-center" />
            <div className="relative p-3 sm:p-5">
                <div className="group absolute -top-24 left-8 size-fit">
                    <div className="relative size-fit">
                        <UserAvatar
                            className="size-28 border-4 border-popover shadow-sm"
                            src={currentUser.media?.downloadURL ?? ''}
                            fallback={currentUser.username.charAt(0) ?? '-'}
                        />
                        <ActionTooltip
                            tooltipContent="Change Profile Photo"
                            align="center"
                            side="right"
                        >
                            <Button
                                variant="secondary"
                                onClick={() => toggleProfileUpload(true)}
                                className="absolute bottom-2 right-2 size-fit rounded-full border border-transparent p-1 hover:border-foreground/20"
                            >
                                <CameraFillIcon className="size-4 opacity-50 duration-300 ease-in-out group-hover:opacity-80" />
                            </Button>
                        </ActionTooltip>
                    </div>
                </div>
                <p className="text-2xl font-medium">{currentUser.username}</p>
                <span className="text-sm text-foreground/80">
                    {currentUser.firstName} {currentUser.lastName}
                </span>
                <div className="flex flex-row gap-y-1 mt-2 sm:items-center flex-wrap gap-x-2 text-xs text-foreground/60">
                    <span>
                        <EmailIcon className="inline" /> {currentUser.email}
                    </span>
                    <DotMediumIcon className="hidden lg:block" />
                    <span>
                        <PhoneOutlineIcon className="inline" />{' '}
                        {currentUser.contactNumber}
                    </span>
                    <DotMediumIcon className="hidden lg:block" />
                    <span>
                        <CalendarIcon className="inline" /> Joined Since{' '}
                        {format(currentUser.createdAt, 'MMM dd yyyy')}
                    </span>
                    <DotMediumIcon className="hidden lg:block" />
                    <span>
                        <CalendarCheckIcon className="inline" /> Update at{' '}
                        {format(currentUser.updatedAt, 'MMM dd yyyy')}
                    </span>
                </div>
            </div>
        </div>
    )
}

export default AccountProfileBanner
