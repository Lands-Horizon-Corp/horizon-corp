export type ReportConfig = {
    [key: string]: string | number | boolean; // Dynamic fields
  };
  
  export type ReportRoute = {
    path: string;
    title: string;
    description: string;
    icon: React.ElementType;
    configFields: { name: string; label: string; type: "text" | "number" | "checkbox" }[];
  };