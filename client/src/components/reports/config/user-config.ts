
// config/tellerReportConfig.ts
import { TellerTransactionRecord } from '../types/user-teller';
import { createElement } from 'react';
import { ColumnDefinition } from '../types/table';

export const FinancialRecords: TellerTransactionRecord[] = [
    {
        id: 1,
        date: '01/22/25',
        orNumber: '00000001',
        pbNumber: '0003842',
        memberName: 'ABAAG ISAGANI',
        cashOnHand: 500.0,
        savings: 500.0,
        timeDeposit: 0,
        shareCapital: 0,
        regularLoan: 0,
        salaryLoan: 0,
        insuranceCash: 0,
        pcl: 0,
        interestOnLoans: 0,
        finesPenalties: 0,
        sundries: 0,
        description: 'PASSBOOK',
        descriptionAmount: 80.0,
    },
    {
        id: 2,
        date: '01/22/25',
        orNumber: '00000004',
        pbNumber: '0003842',
        memberName: 'ABAAG ISAGANI',
        cashOnHand: 4100.0,
        savings: 3000.0,
        timeDeposit: 200.0,
        shareCapital: 0,
        regularLoan: 0,
        salaryLoan: 0,
        insuranceCash: 0,
        pcl: 0,
        interestOnLoans: 0,
        finesPenalties: 0,
        sundries: 0,
        description: 'MEMBERSHIP FEES',
        descriptionAmount: 200.0,
    },
];

// define columns as ColumnDefinition<FinancialRecord> so TS can help:
export const columns: ColumnDefinition<TellerTransactionRecord>[] = [
    { key: 'date', header: 'Date', cellStyle: { fontWeight: 'bold', fontStyle:'italic', fontSize:'2px' } },
    { key: 'orNumber', header: 'O.R. No.' },
    { key: 'pbNumber', header: 'P.B. No.' },
    { key: 'memberName', header: "Member's Name" },
    { key: 'cashOnHand', header: 'Cash on Hand' },
    { key: 'savings', header: 'Savings' },
    { key: 'timeDeposit', header: 'Time Dep.' },
    { key: 'shareCapital', header: 'ShareCap' },
    { key: 'regularLoan', header: 'Reg. Loan' },
    { key: 'salaryLoan', header: 'Sal. Loan' },
    { key: 'insuranceCash', header: 'Ins. Cash' },
    { key: 'pcl', header: 'PCL' },
    { key: 'interestOnLoans', header: 'Int. On Loans' },
    { key: 'finesPenalties', header: 'Fines/Pen.' },
    { key: 'sundries', header: 'Sundries' },
    {
        key: 'description',
        header: 'Description',
        renderCell: (rowData: TellerTransactionRecord) => {
            return rowData.description ? createElement('p', null, rowData.description) : "";
        },
    }, {
        key: 'descriptionAmount',
        header: 'Description Amount',
        renderCell: (row) =>
            row.descriptionAmount !== undefined ? row.descriptionAmount.toFixed(2) : '',
    },
];

// We can store the rest of the report config here, or just keep it in the UI layer
export const tellerReportConfig = {
    title: 'Teller Monitoring Report',
    subtitle: 'January 2025',
    filters: {
        Region: 'North America',
        'Date Range': '01/01/2025 - 01/31/2025',
    },
    footer: { totals: { label: 'Grand Total', value: 3100 } },
};
