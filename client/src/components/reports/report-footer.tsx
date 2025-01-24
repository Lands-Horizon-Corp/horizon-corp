// components/ReportFooter.tsx
import React from 'react';

interface ReportFooterProps {
  totals: { label: string; value: number };
}

const ReportFooter: React.FC<ReportFooterProps> = ({ totals }) => {
  return (
    <div className="report-footer">
      <strong>{totals.label}:</strong> {totals.value.toLocaleString()}
    </div>
  );
};

export default ReportFooter;
