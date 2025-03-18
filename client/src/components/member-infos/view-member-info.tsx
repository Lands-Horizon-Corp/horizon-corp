import { ReactNode } from 'react'
import { IconType } from 'react-icons/lib'

import { UserIcon, BankIcon, UserCogIcon } from '../icons'
import Modal, { IModalProps } from '../modals/modal'
import { ScrollArea, ScrollBar } from '../ui/scroll-area'
import MemberInfoBanner from './banners/member-info-banner'
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
    Component: (props?: IBaseCompNoChild) => ReactNode
}[] = [
    {
        value: 'general-infos',
        title: 'General/Membership',
        Icon: UserIcon,
        Component: () => (
            <div className="flex min-h-[90%] flex-1 flex-col gap-y-4 rounded-xl bg-background p-4">
                <p className="text-sm">Member Center History of this member</p>
            </div>
        ),
    },
    {
        value: 'personal-infos',
        title: 'Personal Info',
        Icon: UserCogIcon,
        Component: () => (
            <div className="flex min-h-[90%] flex-1 flex-col gap-y-4 rounded-xl bg-background p-4">
                <p className="text-sm">
                    Member Classification history of this member
                </p>
            </div>
        ),
    },
    {
        value: 'member-type-history',
        title: 'Member Type',
        Icon: UserCogIcon,
        Component: () => (
            <div className="flex min-h-[90%] flex-1 flex-col gap-y-4 rounded-xl bg-background p-4">
                <p className="text-sm">Member Type history of this member</p>
            </div>
        ),
    },
    {
        value: 'member-accounts',
        title: 'Member Accounts',
        Icon: BankIcon,
        Component: () => (
            <div className="flex min-h-[90%] flex-1 flex-col gap-y-4 rounded-xl bg-background p-4">
                <p className="text-sm">Member Gender history for this member</p>
            </div>
        ),
    },
]

const MemberOverallInfo = ({ memberProfileId }: MemberOverallInfoProps) => {
    const { data: memberProfile } = useMemberProfile({
        profileId: memberProfileId,
        initialData: {
            id: '123e4567-e89b-12d3-a456-426614174000',
            oldReferenceId: 'REF-001',
            passbookNumber: 'PB-20240317',
            notes: 'This is a test member profile.',
            description: 'A sample description for testing purposes.',
            contactNumber: '09123456789',
            tinNumber: '123-456-789',
            civilStatus: 'Single',
            sssNumber: '12-3456789-0',
            businessAddress: '123 Business St., City, Region',
            businessContact: '09234567890',
            status: 'Verified',
            isClosed: false,
            pagibigNumber: '9876543210',
            philhealthNumber: 'PH-123456789',
            isMutualFundMember: true,
            isMicroFinanceMember: false,

            createdAt: '2025-03-17T12:00:00Z',
            updatedAt: '2025-03-17T12:00:00Z',

            memberAddresses: [
                {
                    id: '1d0f8c5c-41be-4adf-a98d-1ff3f3dcf3dc',
                    barangay: 'Balintad',
                    city: 'Baungon',
                    label: 'Home Address',
                    membersProfileId: '123e4567-e89b-12d3-a456-426614174000',
                    postalCode: '1233',
                    province: 'Bukidnon',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                },
            ],

            occupation: {
                id: '123e4567-e89b-12d3-a456-426614274928',
                name: 'Software Engineer',
                description: 'Develops software solutions',
                createdAt: '2025-03-17T12:00:00Z',
                updatedAt: '2025-03-17T12:00:00Z',
            },
            media: {
                id: '123e4567-e89b-12d3-a456-426614241526',
                fileName: 'profile.jpg',
                fileSize: 1024,
                fileType: 'image/jpeg',
                storageKey: 'uploads/profile.jpg',
                url: 'https://res.cloudinary.com/daqh9porl/image/upload/v1725602134/jerbee3_omq76s.jpg',
                bucketName: 'member-media',
                createdAt: '2025-03-17T12:00:00Z',
                updatedAt: '2025-03-17T12:00:00Z',
                downloadURL:
                    'https://res.cloudinary.com/daqh9porl/image/upload/v1725602134/jerbee3_omq76s.jpg',
            },
            member: {
                id: '123e4567-e89b-12d3-a456-426614847638',
                accountType: 'Member',
                username: 'johndoe',
                fullName: 'John Doe',
                isEmailVerified: true,
                isContactVerified: true,
                isSkipVerification: false,
                status: 'Verified',
                firstName: 'John',
                media: {
                    id: '123e4567-e89b-12d3-a456-426614241526',
                    fileName: 'profile.jpg',
                    fileSize: 1024,
                    fileType: 'image/jpeg',
                    storageKey: 'uploads/profile.jpg',
                    url: 'https://res.cloudinary.com/daqh9porl/image/upload/v1725602134/jerbee3_omq76s.jpg',
                    bucketName: 'member-media',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                    downloadURL: 'https://example.com/uploads/profile.jpg',
                },
                lastName: 'Doe',
                middleName: 'M',
                birthDate: '1990-01-01',
                email: 'johndoe@example.com',
                permanentAddress: '123 Main St, City',
                description: 'Test user',
                contactNumber: '09123456789',
                createdAt: '2025-03-17T12:00:00Z',
                updatedAt: '2025-03-17T12:00:00Z',
            },
            memberTypeId: '1d0f8c5c-41be-4adf-a98d-1ff3f3dcf3dc',
            memberType: {
                id: '1d0f8c5c-41be-4adf-a98d-1ff3f3dcf3dc',
                name: 'Regular Member',
                description: 'Entry-level membership',
                prefix: 'BASIC',
                createdAt: '2024-03-07T10:00:00Z',
                updatedAt: '2024-03-07T12:00:00Z',
                deletedAt: null,
            },
            memberClassification: {
                id: '123e4567-e89b-12d3-a456-426614174612',
                name: 'Gold',
                description: 'Premium classification',
                createdAt: '2025-03-17T12:00:00Z',
                updatedAt: '2025-03-17T12:00:00Z',
            },
            memberGender: {
                id: '123e4567-e89b-12d3-a456-426614174313',
                name: 'Male',

                description: 'Biological Male',
                createdAt: '2025-03-17T12:00:00Z',
                updatedAt: '2025-03-17T12:00:00Z',
            },
            branch: {
                id: '123e4567-e89b-12d3-a456-426614174882',
                name: 'Main Branch',
                email: 'branch@gmail.com',
                contactNumber: '0999999999',
                isAdminVerified: true,
                createdAt: '2025-03-17T12:00:00Z',
                updatedAt: '2025-03-17T12:00:00Z',
            },
            memberCenter: {
                id: '123e4567-e89b-12d3-a456-426614173066',
                name: 'Community Center',
                description: '',
                createdAt: '2025-03-17T12:00:00Z',
                updatedAt: '2025-03-17T12:00:00Z',
            },
            signatureMedia: {
                id: '123e4567-e89b-12d3-a456-426614172000',
                fileName: 'signature.png',
                fileSize: 500,
                fileType: 'image/png',
                storageKey: 'uploads/signature.png',
                url: 'https://example.com/uploads/signature.png',
                bucketName: 'member-media',
                createdAt: '2025-03-17T12:00:00Z',
                updatedAt: '2025-03-17T12:00:00Z',
                downloadURL: 'https://example.com/uploads/signature.png',
            },
            memberEducationalAttainmentId:
                '550e8400-e29b-41d4-a716-446655440000',
            memberEducationalAttainment: {
                id: '550e8400-e29b-41d4-a716-446655440000',
                name: 'High School Diploma',
                description: 'Completed secondary education',
                createdAt: '2024-03-12T10:00:00Z',
                updatedAt: '2024-03-12T12:00:00Z',
                deletedAt: null,
            },
            memberAssets: [
                {
                    id: '123e4567-e89b-12d3-a456-426614174078',
                    membersProfileId: '123e4567-e89b-12d3-a456-426614174000',
                    entryDate: '2025-03-17',
                    description: 'Real estate property',
                    name: 'House & Lot',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                },
                {
                    id: '123e4567-e89b-12d3-a456-426614174082',
                    membersProfileId: '123e4567-e89b-12d3-a456-426614174000',
                    entryDate: '2025-03-17',
                    description: 'Car ownership',
                    name: 'Toyota Hilux',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                },
            ],
            memberIncome: [
                {
                    id: '123e4567-e89b-12d3-a456-426614174032',
                    membersProfileId: '123e4567-e89b-12d3-a456-426614174000',
                    name: 'Salary',
                    amount: 50000,
                    date: '2025-03-17',
                    description: 'Monthly salary',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                },
                {
                    id: '123e4567-e89b-12d3-a456-426614174048',
                    membersProfileId: '123e4567-e89b-12d3-a456-426614174000',
                    name: 'Business Profit',
                    amount: 20000,
                    date: '2025-03-17',
                    description: 'Side business income',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                },
            ],
            memberExpenses: [
                {
                    id: '123e4567-e89b-12d3-a456-426614174041',
                    membersProfileId: '123e4567-e89b-12d3-a456-426614174000',
                    name: 'Rent',
                    amount: 10000,
                    date: '2025-03-17',
                    description: 'Monthly house rent',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                },
                {
                    id: '123e4567-e89b-12d3-a456-426614174042',
                    membersProfileId: '123e4567-e89b-12d3-a456-426614174000',
                    name: 'Utilities',
                    amount: 5000,
                    date: '2025-03-17',
                    description: 'Electricity & water bill',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                },
            ],
            memberContactNumberReferences: [
                {
                    id: '123e4567-e89b-12d3-a456-426614174033',
                    name: 'Jane Doe',
                    description: 'Spouse',
                    contactNumber: '09129876543',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                },
                {
                    id: '123e4567-e89b-12d3-a456-426614174222',
                    name: 'Mark Smith',
                    description: 'Brother',
                    contactNumber: '09123456788',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                },
            ],
            memberRelativeAccounts: [
                {
                    id: '123e4567-e89b-12d3-a456-426614174088',
                    membersProfileId: '123e4567-e89b-12d3-a456-426614174000',
                    relativeProfileMemberId: 'member-002',
                    familyRelationship: 'Brother',
                    description: 'Immediate family member',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                },
            ],
            memberRecruits: [
                {
                    id: '123e4567-e89b-12d3-a456-426614174001',
                    name: 'recruit 1',
                    membersProfileId: '123e4567-e89b-12d3-a456-426614174000',
                    membersProfileRecruitedId: 'member-003',
                    dateRecruited: '2025-02-01',
                    description: 'Recruited new member',
                    createdAt: '2025-03-17T12:00:00Z',
                    updatedAt: '2025-03-17T12:00:00Z',
                },
            ],
        },
    })

    return (
        <div className="min-h-[80vh] min-w-[80vw] space-y-4 p-4">
            {memberProfile && (
                <>
                    <MemberInfoBanner memberProfile={memberProfile} />
                </>
            )}
            <Tabs
                defaultValue="member-center-history"
                className="flex-1 flex-col"
            >
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
                        {tab.Component()}
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
