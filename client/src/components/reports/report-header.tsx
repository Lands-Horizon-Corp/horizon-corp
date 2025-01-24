// components/ReportHeader.tsx
import React from 'react';

interface ReportHeaderProps {
  metadata?: Record<string, string>;
}

const ReportHeader: React.FC<ReportHeaderProps> = ({ metadata }) => {
  if (!metadata) return null;

  return (
    <div className="report-header">
      {Object.entries(metadata).map(([key, value]) => (
        <div key={key} className="metadata">
          <strong>{key}:</strong> {value}
        </div>
      ))}
    </div>
  );
};

export default ReportHeader;
