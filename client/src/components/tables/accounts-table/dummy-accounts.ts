import { IAccountResource } from "@/server/types/accounts/accounts";

export const TAccountingAccounts = [
    'Deposit', 'Loan', 'A/R-Ledger', 'A/R-Aging', 'Fines', 'Interest',
    'SVF-Ledger', 'W-Off', 'A/P-Ledger', 'Other'
] as const;

export const TEarnedUnearnedInterest = [
    'By Formula', 'By Formula + Actual Pay', 'By Advance Int. + Actual Pay', 'None'
] as const;

export const TOtherAccountInformation = [
    'None', 'Jewelry', 'Field', 'Grocery', 'Track Loan Ded', 'CIB/CIK Acct.', 'COH Acct.'
] as const;

const generateRandomId = (length = 12) => {
    const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    return Array.from({ length }, () => chars[Math.floor(Math.random() * chars.length)]).join('');
};

export const generateDummyAccounts = (count: number): IAccountResource[] => {
    return Array.from({ length: count }, () => ({
        id: generateRandomId(),
        companyId: generateRandomId(),
        accountCode: `ACCT-${Math.floor(Math.random() * 9000) + 1000}`,
        description: `Account ${Math.floor(Math.random() * 100)}`,
        altDescription: `Alternative ${Math.floor(Math.random() * 100)}`,
        type: TAccountingAccounts[Math.floor(Math.random() * TAccountingAccounts.length)],
        maxAmount: parseFloat((Math.random() * (10000 - 1000) + 1000).toFixed(2)),
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
        finesGpAmort: parseFloat((Math.random() * (500 - 100) + 100).toFixed(2)),
        addtlGp: `Additional-${Math.floor(Math.random() * 50)}`,
        noGracePeriodDaily: Math.random() < 0.5,
        finesGpMaturity: Math.floor(Math.random() * 36) + 1,
        earnedUnearnedInterest: TEarnedUnearnedInterest[Math.floor(Math.random() * TEarnedUnearnedInterest.length)],
        otherInformationOfAnAccount: TOtherAccountInformation[Math.floor(Math.random() * TOtherAccountInformation.length)],
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
        deletedAt: null
    }));
};

export const DummyAccountsData = generateDummyAccounts(100);
