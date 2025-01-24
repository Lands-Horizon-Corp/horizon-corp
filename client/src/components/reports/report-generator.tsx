// components/ReportGenerator.tsx
import ReportWrapper from './report-wrapper';
import ReportHeader from './report-header';
import ReportFooter from './report-footer';
import DataTable from './reports-data-table';
import { TableProps } from './types';

interface ReportGeneratorConfig<TData> {
  title: string;
  subtitle?: string;
  filters?: Record<string, string>;
  footer?: { totals: { label: string; value: number } };
  columns: TableProps<TData>['columns'];
  data: TData[];
}

interface ReportGeneratorProps<TData> {
  reportConfig: ReportGeneratorConfig<TData>;
}

function ReportGenerator<TData>(props: ReportGeneratorProps<TData>) {
  const { reportConfig } = props;
  const { title, subtitle, filters, columns, data, footer } = reportConfig;

  return (
    <ReportWrapper
      title={title}
      subtitle={subtitle}
      header={filters ? <ReportHeader metadata={filters} /> : null}
      table={<DataTable columns={columns} data={data} />}
      footer={footer ? <ReportFooter totals={footer.totals} /> : null}
    />
  );
}

export default ReportGenerator;
