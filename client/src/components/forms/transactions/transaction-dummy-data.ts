import {
    IAccountingLedgerResource,
    TransactionSource,
} from '@/server/types/accounts/accounting-ledger'
import { IAccountsPaginatedResource } from '@/server/types/accounts/accounts'
import { IMemberPaginatedResource } from '@/server'

export const TAccountingAccounts = [
    'Deposit',
    'Loan',
    'A/R-Ledger',
    'A/R-Aging',
    'Fines',
    'Interest',
    'SVF-Ledger',
    'W-Off',
    'A/P-Ledger',
    'Other',
] as const

export const TEarnedUnearnedInterest = [
    'By Formula',
    'By Formula + Actual Pay',
    'By Advance Int. + Actual Pay',
    'None',
] as const

export const TOtherAccountInformation = [
    'None',
    'Jewelry',
    'Field',
    'Grocery',
    'Track Loan Ded',
    'CIB/CIK Acct.',
    'COH Acct.',
] as const

const generateRandomId = (length = 12) => {
    const chars =
        'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
    return Array.from(
        { length },
        () => chars[Math.floor(Math.random() * chars.length)]
    ).join('')
}

export const generateDummyAccounts = (
    count: number
): IAccountsPaginatedResource => {
    const accounts = Array.from({ length: count }, () => ({
        id: generateRandomId(),
        companyId: generateRandomId(),
        accountCode: `ACCT-${Math.floor(Math.random() * 9000) + 1000}`,
        description: `Account ${Math.floor(Math.random() * 100)}`,
        altDescription: `Alternative ${Math.floor(Math.random() * 100)}`,
        type: TAccountingAccounts[
            Math.floor(Math.random() * TAccountingAccounts.length)
        ],
        maxAmount: parseFloat(
            (Math.random() * (10000 - 1000) + 1000).toFixed(2)
        ),
        minAmount: parseFloat((Math.random() * (999 - 100) + 100).toFixed(2)),
        computationType: generateRandomId(),
        headerRow: Math.floor(Math.random() * 10) + 1,
        centerRow: Math.floor(Math.random() * 10) + 1,
        totalRow: Math.floor(Math.random() * 10) + 1,
        print: Math.random() < 0.5,
        addOn: Math.random() < 0.5,
        allowRebate: Math.random() < 0.5,
        taxable: Math.random() < 0.5,
        finesAmort: parseFloat((Math.random() * (500 - 100) + 100).toFixed(2)),
        finesMaturity: Math.floor(Math.random() * 36) + 1,
        interestStandard: parseFloat((Math.random() * (15 - 1) + 1).toFixed(2)),
        interestSecured: parseFloat((Math.random() * (15 - 1) + 1).toFixed(2)),
        schemeNo: Math.floor(Math.random() * 10) + 1,
        altCode: Math.floor(Math.random() * 900) + 100,
        glCode: Math.floor(Math.random() * 9000) + 1000,
        finesGpAmort: parseFloat(
            (Math.random() * (500 - 100) + 100).toFixed(2)
        ),
        addtlGp: `Additional-${Math.floor(Math.random() * 50)}`,
        noGracePeriodDaily: Math.random() < 0.5,
        finesGpMaturity: Math.floor(Math.random() * 36) + 1,
        earnedUnearnedInterest:
            TEarnedUnearnedInterest[
                Math.floor(Math.random() * TEarnedUnearnedInterest.length)
            ],
        otherInformationOfAnAccount:
            TOtherAccountInformation[
                Math.floor(Math.random() * TOtherAccountInformation.length)
            ],
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
        deletedAt: null,
    }))
    return {
        data: accounts,
        pageIndex: 1,
        totalPage: 1,
        pageSize: count,
        totalSize: count,
        pages: [],
    }
}

export const SampleAccountsData = generateDummyAccounts(100)

export const sampleLedgerData: IAccountingLedgerResource[] = [
    {
        id: 'ledger-001',
        description: 'Loan Repayment',
        notes: 'Monthly loan repayment for January',
        company_id: 'comp-001',
        member_profile_id: 'member-002',
        account_id: 'acct-002',
        employee_id: 'emp-002',
        or_number: 'OR-789012',
        transaction_source: TransactionSource.payment,
        debit: 300.0,
        credit: 0.0,
        balance: 4700.0,
        transaction_date: new Date('2024-02-20'),
        entry_date: new Date('2024-02-21'),
        line_number: 2,
        transaction_type_id: 'trans-002',
        created_by: 'admin-002',
        created_at: new Date('2024-02-21'),
        createdAt: new Date('2024-02-21').toISOString(),
        updatedAt: new Date('2024-02-21').toISOString(),
        summary: {
            credit: 1000,
            debit: 2000,
            balance: 5400,
        },
    },
    {
        id: 'ledger-002',
        description: 'Service Income',
        notes: 'Consulting services for February',
        company_id: 'comp-001',
        member_profile_id: 'member-003',
        account_id: 'acct-003',
        employee_id: 'emp-001',
        or_number: 'OR-789013',
        transaction_source: TransactionSource.withdrawal,
        debit: 0.0,
        credit: 1500.0,
        balance: 6200.0,
        transaction_date: new Date('2024-02-22'),
        entry_date: new Date('2024-02-23'),
        line_number: 3,
        transaction_type_id: 'trans-003',
        created_by: 'admin-003',
        created_at: new Date('2024-02-23'),
        createdAt: new Date('2024-02-23').toISOString(),
        updatedAt: new Date('2024-02-23').toISOString(),
        summary: {
            credit: 2500,
            debit: 2000,
            balance: 6200,
        },
    },
    {
        id: 'ledger-003',
        description: 'Equipment Purchase',
        notes: 'Laptop purchase for new employee',
        company_id: 'comp-002',
        member_profile_id: 'member-001',
        account_id: 'acct-001',
        employee_id: 'emp-003',
        or_number: 'OR-789014',
        transaction_source: TransactionSource.payment,
        debit: 1200.0,
        credit: 0.0,
        balance: 5000.0,
        transaction_date: new Date('2024-02-24'),
        entry_date: new Date('2024-02-25'),
        line_number: 4,
        transaction_type_id: 'trans-004',
        created_by: 'admin-001',
        created_at: new Date('2024-02-25'),
        createdAt: new Date('2024-02-25').toISOString(),
        updatedAt: new Date('2024-02-25').toISOString(),
        summary: {
            credit: 2500,
            debit: 3200,
            balance: 5000,
        },
    },
    {
        id: 'ledger-004',
        description: 'Utilities Payment',
        notes: 'Electricity bill for February',
        company_id: 'comp-003',
        member_profile_id: 'member-004',
        account_id: 'acct-004',
        employee_id: 'emp-004',
        or_number: 'OR-789015',
        transaction_source: TransactionSource.payment,
        debit: 400.0,
        credit: 0.0,
        balance: 4600.0,
        transaction_date: new Date('2024-02-26'),
        entry_date: new Date('2024-02-27'),
        line_number: 5,
        transaction_type_id: 'trans-005',
        created_by: 'admin-004',
        created_at: new Date('2024-02-27'),
        createdAt: new Date('2024-02-27').toISOString(),
        updatedAt: new Date('2024-02-27').toISOString(),
        summary: {
            credit: 2500,
            debit: 3600,
            balance: 4600,
        },
    },
    {
        id: 'ledger-005',
        description: 'Interest Income',
        notes: 'Savings account interest for February',
        company_id: 'comp-001',
        member_profile_id: 'member-005',
        account_id: 'acct-005',
        employee_id: 'emp-005',
        or_number: 'OR-789016',
        transaction_source: TransactionSource.deposit,
        debit: 0.0,
        credit: 200.0,
        balance: 4800.0,
        transaction_date: new Date('2024-02-28'),
        entry_date: new Date('2024-02-29'),
        line_number: 6,
        transaction_type_id: 'trans-006',
        created_by: 'admin-005',
        created_at: new Date('2024-02-29'),
        createdAt: new Date('2024-02-29').toISOString(),
        updatedAt: new Date('2024-02-29').toISOString(),
        summary: {
            credit: 2700,
            debit: 3600,
            balance: 4800,
        },
    },
]

export const sampleMemberPaginatedResource: IMemberPaginatedResource = {
    data: [
        {
            id: 'member-001',
            accountType: 'Member',
            username: 'johndoe01',
            fullName: 'John Doe',
            firstName: 'John',
            lastName: 'Doe',
            birthDate: '1990-01-01',
            email: 'johndoe01@example.com',
            contactNumber: '1234567890',
            permanentAddress: '123 Main St, Cityville',
            isEmailVerified: true,
            isContactVerified: true,
            isSkipVerification: false,
            status: 'Verified',
            media: {
                id: '7f76efd0-940a-42f9-afa9-8644453e20aa',
                fileName: 'sample-image.png',
                fileSize: 5242880, // 5 MB in bytes
                fileType: 'image/jpeg',
                storageKey: 'media/resources/sample-image.jpg',
                url: 'https://images.pexels.com/photos/415829/pexels-photo-415829.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1', // This is viewable directly in the browser
                bucketName: 'example-media-bucket',
                createdAt: '2024-10-29T08:45:00Z',
                updatedAt: '2024-10-29T10:20:00Z',
                downloadURL:
                    'https://images.pexels.com/photos/415829/pexels-photo-415829.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1',
            },
            gender: {
                id: 'gender-001',
                name: 'Male',
                createdAt: new Date().toISOString(),
                updatedAt: new Date().toISOString(),
                description: '',
            },
            footsteps: [],
            memberProfile: {
                id: 'profile-001',
                notes: 'Regular member',
                description: 'Interested in investments',
                contactNumber: '1234567890',
                civilStatus: 'Single',
                status: 'Verified',
                isClosed: false,
                isMutualFundMember: true,
                isMicroFinanceMember: false,
                occupation: {
                    id: 'occupation-001',
                    name: 'Engineer',
                    createdAt: new Date().toISOString(),
                    updatedAt: new Date().toISOString(),
                    description: '',
                },
                createdAt: new Date().toISOString(),
                updatedAt: new Date().toISOString(),
            },
            createdAt: new Date().toISOString(),
            updatedAt: new Date().toISOString(),
        },
        {
            id: 'member-002',
            accountType: 'Member',
            username: 'janedoe02',
            fullName: 'Jane Doe',
            firstName: 'Jane',
            lastName: 'Doe',
            birthDate: '1985-05-15',
            email: 'janedoe02@example.com',
            contactNumber: '0987654321',
            permanentAddress: '456 Park Ave, Townsville',
            isEmailVerified: false,
            isContactVerified: true,
            isSkipVerification: false,
            status: 'Pending',
            media: {
                id: '324d740f2-f57f-4579-a9a3-03d6ddec44b5',
                fileName: 'sample-image2.png',
                fileSize: 7340032, // 7 MB in bytes
                fileType: 'audio/mpeg',
                storageKey: 'media/resources/music-track.mp3',
                url: 'https://images.pexels.com/photos/3763188/pexels-photo-3763188.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1', // Direct URL
                bucketName: 'example-media-bucket',
                createdAt: '2024-10-28T09:30:00Z',
                updatedAt: '2024-10-28T11:00:00Z',
                downloadURL:
                    'https://images.pexels.com/photos/3763188/pexels-photo-3763188.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1',
            },
            gender: {
                id: 'gender-002',
                name: 'Female',
                createdAt: new Date().toISOString(),
                updatedAt: new Date().toISOString(),
                description: '',
            },
            footsteps: [],
            memberProfile: {
                id: 'profile-002',
                notes: 'New member',
                description: 'Looking for savings plans',
                contactNumber: '0987654321',
                civilStatus: 'Married',
                status: 'Pending',
                isClosed: false,
                isMutualFundMember: false,
                isMicroFinanceMember: true,
                occupation: {
                    id: 'occupation-002',
                    name: 'Teacher',
                    createdAt: new Date().toISOString(),
                    updatedAt: new Date().toISOString(),
                    description: '',
                },
                createdAt: new Date().toISOString(),
                updatedAt: new Date().toISOString(),
            },
            createdAt: new Date().toISOString(),
            updatedAt: new Date().toISOString(),
        },
    ],
    totalSize: 2,
    totalPage: 1,
    pageIndex: 1,
    pageSize: 10,
    pages: [],
}
