
export interface TellerTransactionRecord {
    id: number;
    date: string;
    orNumber: string;
    pbNumber: string;
    memberName: string;
    cashOnHand: number;
    savings: number;
    timeDeposit: number;
    shareCapital: number;
    regularLoan: number;
    salaryLoan: number;
    insuranceCash: number;
    pcl: number;
    interestOnLoans: number;
    finesPenalties: number;
    sundries: number;
    description?: string;
    descriptionAmount?: number;
  }
  