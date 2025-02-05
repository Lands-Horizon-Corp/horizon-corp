import DailyCollDetail from "@/modules/owner/pages/reports/menu/daily-coll-detail"
import DailyCollSummary from "@/modules/owner/pages/reports/menu/daily-coll-summary"

export interface ReportProps<T> {
    onSubmit: (config: T) => void
    onClose: () => void
  }
  
  export interface ReportConfig<T = any> {
    [key: string]: React.FC<ReportProps<T>>
  }
  
 const ReportConfigComponents: ReportConfig = {
      'daily-coll-detail': DailyCollDetail,
      'daily-coll-summary':DailyCollSummary,
      // Add more report-specific configurations here
  }

  export default ReportConfigComponents
  
  
