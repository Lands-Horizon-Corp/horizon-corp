export interface TUsersReportFormat1 {
    date: string;
    orNumber: string; // O.R. No.
    pbNumber: string; // P.B. No.
    memberName: string; // Member's Name
    cashOnHand: number; // Cash on Hand (DR)
    savings: number; // Savings (CR)
    timeDeposit: number; // Time Dep. (CR)
    shareCapital: number; // ShareCap (CR)
    regularLoan: number; // Reg. Loan (CR)
    salaryLoan: number; // Sal. Loan (CR)
    insuranceCash: number; // Ins. Cash (CR)
    pcl: number | string; // P C L (CR)
    interestOnLoans: number; // Int. On Loans (CR)
    finesPenalties: number; // Fines/Pen. (CR)
    sundries: number; // Sundries (CR)
}

const dummyData: TUsersReportFormat1[] = [
    {
      date: "01/22/25",
      orNumber: "00000001",
      pbNumber: "0003842",
      memberName: "ABAAG ISAGANI",
      cashOnHand: 500.00,
      savings: 500.00,
      timeDeposit: 0,
      shareCapital: 0,
      regularLoan: 0,
      salaryLoan: 0,
      insuranceCash: 0,
      pcl: 0,
      interestOnLoans: 0,
      finesPenalties: 0,
      sundries: 80.00,
    },
    {
      date: "01/22/25",
      orNumber: "00000002",
      pbNumber: "0003842",
      memberName: "ABAAG ISAGANI",
      cashOnHand: 80.00,
      savings: 0,
      timeDeposit: 0,
      shareCapital: 0,
      regularLoan: 0,
      salaryLoan: 0,
      insuranceCash: 0,
      pcl: 0,
      interestOnLoans: 0,
      finesPenalties: 0,
      sundries: 80.00,
    },
    {
      date: "01/22/25",
      orNumber: "00000003",
      pbNumber: "0003842",
      memberName: "ABAAG ISAGANI",
      cashOnHand: 2000.00,
      savings: 2000.00,
      timeDeposit: 0,
      shareCapital: 0,
      regularLoan: 0,
      salaryLoan: 0,
      insuranceCash: 0,
      pcl: 0,
      interestOnLoans: 0,
      finesPenalties: 0,
      sundries: 0,
    },
    {
      date: "01/22/25",
      orNumber: "00000004",
      pbNumber: "0003842",
      memberName: "ABAAG ISAGANI",
      cashOnHand: 4100.00,
      savings: 3000.00,
      timeDeposit: 200.00,
      shareCapital: 0,
      regularLoan: 0,
      salaryLoan: 0,
      insuranceCash: 0,
      pcl: 0,
      interestOnLoans: 0,
      finesPenalties: 0,
      sundries: 700.00,
    },
    {
      date: "01/22/25",
      orNumber: "00000005",
      pbNumber: "0005745",
      memberName: "YEPES ROJAN LAID",
      cashOnHand: 200.00,
      savings: 200.00,
      timeDeposit: 0,
      shareCapital: 0,
      regularLoan: 0,
      salaryLoan: 0,
      insuranceCash: 0,
      pcl: 0,
      interestOnLoans: 0,
      finesPenalties: 0,
      sundries: 200.00,
    },
    {
      date: "01/22/25",
      orNumber: "00012225",
      pbNumber: "0003842",
      memberName: "ABAAG ISAGANI",
      cashOnHand: 100.00,
      savings: 100.00,
      timeDeposit: 0,
      shareCapital: 0,
      regularLoan: 0,
      salaryLoan: 0,
      insuranceCash: 0,
      pcl: 0,
      interestOnLoans: 0,
      finesPenalties: 0,
      sundries: 0,
    },
    {
      date: "01/22/25",
      orNumber: "00012225",
      pbNumber: "0005745",
      memberName: "YEPES ROJAN LAID",
      cashOnHand: 1000.00,
      savings: 1000.00,
      timeDeposit: 0,
      shareCapital: 0,
      regularLoan: 0,
      salaryLoan: 0,
      insuranceCash: 0,
      pcl: 0,
      interestOnLoans: 0,
      finesPenalties: 0,
      sundries: 0,
    },
  ];
  