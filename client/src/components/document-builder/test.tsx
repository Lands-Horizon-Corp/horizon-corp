import React from 'react';
import {
  tellerReportConfig,
  FinancialRecords,
  columns,
} from '@/components/reports/config/user-config';
import ReportGenerator from '../reports/report-generator';
import { TellerTransactionRecord } from '@/components/reports/types';

const fullReportConfig = {
  ...tellerReportConfig,
  columns,
  data: FinancialRecords,
};

const TestReport: React.FC = () => {
  return <ReportGenerator<TellerTransactionRecord> reportConfig={fullReportConfig} />;
};

export default TestReport;