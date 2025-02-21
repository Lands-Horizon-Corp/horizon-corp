import { useState } from 'react'

import MembersTable from '@/components/tables/members-table'
import PageContainer from '@/components/containers/page-container'
import EnsureOwnerCompany from '@/modules/owner/components/ensure-company'
import { MemberCreateUpdateFormModal } from '@/components/forms/member-forms/member-create-update-form'
import MembersTableOwnerAction from '@/components/tables/members-table/row-actions/members-table-owner-action'

import { ICompanyResource } from '@/server'

const OwnerViewMembersPage = () => {
    const [createModal, setCreateModal] = useState(false)
    const [company, setCompany] = useState<ICompanyResource | undefined>({
        id: '7f76efd0-940a-42f9-afa9-8644453e20aa',
        name: 'Example Tech Co.',
        description:
            'A leading technology company specializing in software solutions.',
        address: '123 Example Street, Tech City, TC 12345',
        longitude: 123.456,
        latitude: -78.91,
        contactNumber: '+1-800-555-1234',
        isAdminVerified: true,
        owner: {
            id: '24d740f2-f57f-4579-a9a3-03d6ddec44b5',
            username: 'John Doe',
            firstName: 'John',
            lastName: 'doe',
            fullName: 'haha',
            email: 'john.doe@example.com',
            contactNumber: '+1-800-555-5678',
            accountType: '',
            birthDate: '2025-01-15T12:00:00Z',
            status: 'Pending',
            isEmailVerified: false,
            isContactVerified: false,
            isSkipVerification: false,
            createdAt: '',
            updatedAt: '',
        },
        media: {
            id: '9482cd20-df59-49a2-b6ce-bf88bed349e3',
            fileName: 'company-logo.png',
            fileSize: 204800, // 200 KB
            fileType: 'image/png',
            storageKey: 'media/company-logo.png',
            url: 'https://example.com/media/company-logo.png',
            bucketName: 'example-bucket',
            createdAt: '2025-01-01T12:00:00Z',
            updatedAt: '2025-01-15T12:00:00Z',
            downloadURL:
                'https://example.com/media/company-logo.png?download=true',
        },
        branches: [
            {
                id: 'f13a44b8-d627-47db-8584-614eb016f6d6',
                name: 'Tech Branch HQ',
                address: '456 Innovation Avenue, Tech City, TC 12345',
                longitude: 124.123,
                latitude: -79.456,
                email: 'hq@exampletech.com',
                contactNumber: '+1-800-555-7890',
                isAdminVerified: true,
                media: {
                    id: '88350bb7-29de-4e99-bcf0-2f668bcfae92',
                    fileName: 'branch-image.jpg',
                    fileSize: 102400, // 100 KB
                    fileType: 'image/jpeg',
                    storageKey: 'media/branch-image.jpg',
                    url: 'https://example.com/media/branch-image.jpg',
                    bucketName: 'example-bucket',
                    createdAt: '2025-01-01T12:00:00Z',
                    updatedAt: '2025-01-15T12:00:00Z',
                    downloadURL:
                        'https://example.com/media/branch-image.jpg?download=true',
                },
                createdAt: '2025-01-10t08:00:00z',
                updatedAt: '2025-01-15t08:00:00z',
            },
        ],
        createdAt: '2025-01-01T12:00:00Z',
        updatedAt: '2025-01-15T12:00:00Z',
    })

    return (
        <PageContainer>
            <MemberCreateUpdateFormModal
                open={createModal}
                onOpenChange={setCreateModal}
                formProps={{
                    defaultValues: {
                        mode: 'create',
                        companyId: company?.id,
                    },
                    onSuccess: () => {},
                }}
            />
            <EnsureOwnerCompany disabled onSuccess={setCompany}>
                <MembersTable
                    actionComponent={(props) => (
                        <MembersTableOwnerAction {...props} />
                    )}
                    toolbarProps={{
                        createActionProps: {
                            onClick: () => setCreateModal(true),
                        },
                    }}
                    defaultFilter={{
                        companyId: {
                            mode: 'equal',
                            displayText: '',
                            dataType: 'number',
                            value: company?.id,
                            isStaticFilter: true,
                        },
                    }}
                    className="max-h-[90vh] min-h-[90vh] w-full"
                />
            </EnsureOwnerCompany>
        </PageContainer>
    )
}

export default OwnerViewMembersPage
