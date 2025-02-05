import { BiBook, BiMoney, BiListUl, BiCoinStack, BiBarChart, BiWallet, BiReceipt, BiCreditCard, BiUser, BiStore } from 'react-icons/bi';
import { FaUsers, FaFileInvoice, FaBalanceScale, FaChartPie, FaRegMoneyBillAlt, FaHandHoldingUsd, FaShoppingCart } from 'react-icons/fa';

export interface TReportRoute {
    path: string
    title: string
    description: string
    icon: React.ElementType
}

export const reportRoutes: TReportRoute[] = [
    { path: 'daily-coll-detail', title: 'Daily Collection Detail', description: 'Detailed report of daily collections.', icon: BiMoney },
    { path: 'daily-coll-summary', title: 'Daily Collection Summary', description: 'Summary report of daily collections.', icon: BiBarChart },
    { path: 'daily-withdrawal', title: 'Daily Withdrawal', description: 'Report on daily withdrawals.', icon: BiWallet },
    { path: 'member-listing', title: 'Member Listing', description: 'List of registered members.', icon: FaUsers },
    { path: 'number-tag', title: 'Number Tag', description: 'Report on assigned number tags.', icon: BiListUl },
    { path: 'loan-rel-tabulated', title: 'Loan Related Tabulated', description: 'Tabulated report on loan relations.', icon: BiCoinStack },
    { path: 'loan-rel-summary', title: 'Loan Related Summary', description: 'Summary of loan-related transactions.', icon: BiBook },
    { path: 'deposit-balances', title: 'Deposit Balances', description: 'Current deposit balance reports.', icon: FaBalanceScale },
    { path: 'statement-of-account', title: 'Statement of Account', description: 'Detailed account statements.', icon: FaFileInvoice },
    { path: 'subscription-fee', title: 'Subscription Fee', description: 'Report on collected subscription fees.', icon: BiMoney },
    { path: 'loan-rel-detail', title: 'Loan Related Detail', description: 'Detailed report on loan-related transactions.', icon: BiBook },
    { path: 'time-deposit', title: 'Time Deposit', description: 'Details on time deposit accounts.', icon: BiCoinStack },
    { path: 'td-balance', title: 'TD Balance', description: 'Balance details for TD accounts.', icon: FaRegMoneyBillAlt },
    { path: 'td-bal-ytd', title: 'TD Balance / YTD', description: 'Year-to-date report of TD balances.', icon: FaChartPie },
    { path: 'td-accrued', title: 'TD Accrued', description: 'Accrued TD details.', icon: FaHandHoldingUsd },
    { path: 'cash-disbursement', title: 'Cash Disbursement', description: 'Reports on cash disbursements.', icon: BiWallet },
    { path: 'journal-voucher', title: 'Journal Voucher', description: 'Voucher transactions report.', icon: BiReceipt },
    { path: 'account-hold-out', title: 'Account Hold Out', description: 'Accounts on hold report.', icon: FaFileInvoice },
    { path: 'account-balance', title: 'Account Balance', description: 'Balance details for various accounts.', icon: FaBalanceScale },
    { path: 'share-cap-withdrawal', title: 'Share Capital Withdrawal', description: 'Withdrawal reports for share capital.', icon: FaHandHoldingUsd },
    { path: 'cash-receipt-journal', title: 'Cash Receipt Journal', description: 'Journal report of cash receipts.', icon: BiReceipt },
    { path: 'blotter', title: 'Blotter', description: 'Summary of financial transactions.', icon: BiBook },
    { path: 'loan-balances', title: 'Loan Balances', description: 'Current loan balance reports.', icon: BiCreditCard },
    { path: 'loan-receivable', title: 'Loan Receivable', description: 'Details of loans receivable.', icon: BiCreditCard },
    { path: 'earned-unearned', title: 'Earned / Unearned', description: 'Breakdown of earned and unearned funds.', icon: FaBalanceScale },
    { path: 'adjustment', title: 'Adjustment', description: 'Financial adjustments report.', icon: BiBook },
    { path: 'direct-adjustment', title: 'Direct Adjustment', description: 'Directly adjusted financial data.', icon: BiBook },
    { path: 'portfolio-at-risk', title: 'Portfolio at Risk', description: 'Risk assessment of financial portfolio.', icon: FaChartPie },
    { path: 'loan-coll-detail', title: 'Loan Collection Detail', description: 'Detailed loan collection reports.', icon: BiCreditCard },
    { path: 'comaker', title: 'Comaker', description: 'Report on comaker agreements.', icon: FaUsers },
    { path: 'ledger', title: 'Ledger', description: 'General ledger report.', icon: BiBook },
    { path: 'other-funds', title: 'Other Funds', description: 'Details on miscellaneous funds.', icon: BiCoinStack },
    { path: 'interest-share-cap', title: 'Interest Share Capital', description: 'Report on interest-bearing share capital.', icon: FaHandHoldingUsd },
    { path: 'loan-coll-summary', title: 'Loan Collection Summary', description: 'Summary of loan collection activities.', icon: BiCreditCard },
    { path: 'past-due-on-installment', title: 'Past Due on Installment', description: 'Overdue installment payments report.', icon: BiWallet },
    { path: 'clpp', title: 'CLPP', description: 'CLPP-related financial report.', icon: BiBook },
    { path: 'grocery-loan', title: 'Grocery Loan', description: 'Report on grocery loans.', icon: FaShoppingCart },
    { path: 'rebates', title: 'Rebates', description: 'Details of rebates issued.', icon: FaHandHoldingUsd },
    { path: 'proof-of-purchase', title: 'Proof of Purchase', description: 'Proof of purchase reports.', icon: BiStore },
    { path: 'teller-monitor', title: 'Teller Monitor', description: 'Teller monitoring report.', icon: BiUser },
];
