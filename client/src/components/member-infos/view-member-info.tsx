import { ReactNode } from 'react'
import { IconType } from 'react-icons/lib'

import {
    UserIcon,
    BankIcon,
    UserCogIcon,
    UserTagIcon,
    CreditCardIcon,
} from '../icons'
import Modal, { IModalProps } from '../modals/modal'
import MemberPersonalInfo from './member-personal-info'
import { ScrollArea, ScrollBar } from '../ui/scroll-area'
import MemberInfoBanner from './banners/member-info-banner'
import MemberMembershipInfo from './member-general-membership-info'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '../ui/tabs'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { IMemberProfileResource, TEntityId } from '@/server'
import { useMemberProfile } from '@/hooks/api-hooks/member/use-member-profile'

interface MemberOverallInfoProps {
    memberProfileId: TEntityId
    defaultMemberProfile?: IMemberProfileResource
}

const memberInfoTabs: {
    value: string
    title: string
    Icon?: IconType
    Component: (
        props: IBaseCompNoChild & {
            profileId: TEntityId
            defaultData?: IMemberProfileResource
        }
    ) => ReactNode
}[] = [
    {
        value: 'general-infos',
        title: 'General/Membership',
        Icon: UserTagIcon,
        Component: (props) => <MemberMembershipInfo {...props} />,
    },
    {
        value: 'personal-infos',
        title: 'Personal Info',
        Icon: UserIcon,
        Component: (props) => <MemberPersonalInfo {...props} />,
    },
    {
        value: 'government-benefits',
        title: 'Government Benefits',
        Icon: UserCogIcon,
        Component: () => (
            <div className="flex min-h-[90%] flex-1 flex-col gap-y-4 rounded-xl bg-background p-4">
                <p className="text-sm">Government Benefits</p>
            </div>
        ),
    },
    {
        value: 'financial',
        title: 'Financial',
        Icon: CreditCardIcon,
        Component: () => (
            <div className="flex min-h-[90%] flex-1 flex-col gap-y-4 rounded-xl bg-background p-4">
                <p className="text-sm">Financial Info</p>
            </div>
        ),
    },
    {
        value: 'accounts',
        title: 'Accounts Info',
        Icon: BankIcon,
        Component: () => (
            <div className="flex min-h-[90%] flex-1 flex-col gap-y-4 rounded-xl bg-background p-4">
                <p className="text-sm">Accounts</p>
            </div>
        ),
    },
]

const MemberOverallInfo = ({ memberProfileId }: MemberOverallInfoProps) => {
    const { data: memberProfile } = useMemberProfile({
        profileId: memberProfileId,
    })

    return (
        <div className="min-h-[80vh] min-w-[80vw] space-y-4 pt-4">
            {memberProfile && (
                <>
                    <MemberInfoBanner memberProfile={memberProfile} />
                </>
            )}
            <Tabs defaultValue="general-infos" className="flex-1 flex-col">
                <ScrollArea>
                    <TabsList className="mb-3 h-auto min-w-full justify-start gap-2 rounded-none border-b bg-transparent px-0 py-1 text-foreground">
                        {memberInfoTabs.map((tab) => (
                            <TabsTrigger
                                key={tab.value}
                                value={tab.value}
                                className="relative after:absolute after:inset-x-0 after:bottom-0 after:-mb-1 after:h-0.5 after:duration-300 after:ease-in-out hover:bg-accent hover:text-foreground data-[state=active]:bg-transparent data-[state=active]:shadow-none data-[state=active]:after:bg-primary data-[state=active]:hover:bg-accent"
                            >
                                {tab.Icon && (
                                    <tab.Icon
                                        className="-ms-0.5 me-1.5 opacity-60"
                                        size={16}
                                        aria-hidden="true"
                                    />
                                )}
                                {tab.title}
                            </TabsTrigger>
                        ))}
                    </TabsList>
                    <ScrollBar orientation="horizontal" />
                </ScrollArea>
                {memberInfoTabs.map((tab) => (
                    <TabsContent value={tab.value} key={tab.value} asChild>
                        {tab.Component({ profileId: memberProfileId })}
                    </TabsContent>
                ))}
            </Tabs>
        </div>
    )
}

export default MemberOverallInfo

export const MemberOverallInfoModal = ({
    overallInfoProps,
    className,
    ...props
}: IModalProps & { overallInfoProps: MemberOverallInfoProps }) => {
    return (
        <Modal
            {...props}
            titleClassName="hidden"
            descriptionClassName="hidden"
            className={cn('!max-w-[90vw]', className)}
        >
            <MemberOverallInfo {...overallInfoProps} />
        </Modal>
    )
}
