import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'
import { CameraFillIcon } from '@/components/icons'
import ActionTooltip from '@/components/action-tooltip'

import { UserData } from '@/horizon-corp/types'
import { useState } from 'react'
import ProfileUpload from './profile-upload'

const AccountProfileBanner = ({ currentUser }: { currentUser: UserData }) => {
    const [uploadState, toggleProfileUpload] = useState(false)

    return (
        <div className="overflow-clip rounded-2xl bg-secondary shadow-md">
            <Dialog open={uploadState} onOpenChange={toggleProfileUpload}>
                <DialogContent className="!rounded-2xl">
                    <DialogHeader>
                        <DialogTitle>Change Profile Picture</DialogTitle>
                        <DialogDescription>
                            Change your profile photo
                        </DialogDescription>
                    </DialogHeader>
                    <ProfileUpload
                        currentUserPhotoUrl={currentUser.media?.downloadURL ?? null}
                    />
                </DialogContent>
            </Dialog>

            <div className="h-[180px] bg-[url('/profile-cover.png')] bg-cover bg-center" />
            <div className="relative p-8">
                <div className="group absolute -top-16 left-8 size-fit">
                    <div className="relative size-fit">
                        <UserAvatar
                            className="size-28 border-4 border-popover shadow-sm"
                            src={
                                currentUser.media?.downloadURL ??
                                'https://avatars.githubusercontent.com/u/48374007?s=80&v=4'
                            }
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
                <p className="mt-9 text-2xl font-medium">
                    {currentUser.username}
                </p>
                <span className="text-sm text-foreground/80">
                    {currentUser.firstName} {currentUser.lastName}
                </span>
            </div>
        </div>
    )
}

export default AccountProfileBanner