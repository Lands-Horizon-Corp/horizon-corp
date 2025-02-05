import DocumentBuilder from '@/components/document-builder';
// import { ColumnDef } from '@tanstack/react-table';

export type Payment = {
  id: string;
  amount: number;
  status: "pending" | "processing" | "success" | "failed";
  email: string;
};

// const payments: Payment[] = [
//     { id: '1', amount: 100, status: 'success', email: 'test1@example.com' },
//     { id: '2', amount: 200, status: 'failed', email: 'test2@example.com' },
// ]

// const columns: ColumnDef<Payment>[] = [
//     {
//         accessorKey: 'id',
//         header: 'ID',
//     },
//     {
//         accessorKey: 'email',
//         header: 'Email',
//     },
//     {
//         accessorKey: 'amount',
//         header: 'Amount',
//         cell: ({ row }) => `$${row.getValue('amount')}`,
//     },
//     {
//         accessorKey: 'status',
//         header: 'Status',
//     },
// ]

export type Person = {
  id: string;
  name: string;
  age: number;
  occupation: string;
  email: string;
};

// const persons: Person[] = [
//   { id: "1", name: "John Doe", age: 30, occupation: "Engineer", email: "john.doe@example.com" },
//   { id: "2", name: "Jane Smith", age: 25, occupation: "Designer", email: "jane.smith@example.com" },
// ];

// // Column definitions
// const personColumns: ColumnDef<Person>[] = [
//   {
//     accessorKey: "id",
//     header: "ID",
//   },
//   {
//     accessorKey: "name",
//     header: "Name",
//   },
//   {
//     accessorKey: "age",
//     header: "Age",
//     cell: ({ row }) => <div>{row.getValue("age")} years old</div>,
//   },
//   {
//     accessorKey: "occupation",
//     header: "Occupation",
//   },
//   {
//     accessorKey: "email",
//     header: "Email",
//     cell: ({ row }) => <a href={`mailto:${row.getValue("email")}`}>{row.getValue("email")}</a>,
//   },
// ];

const document = () => {
    return (
        <div className="h-screen w-screen">
            <DocumentBuilder />
            {/* <InsertTable data={payments} columns={columns} /> */}
            {/* <InsertTable data={persons} columns={personColumns} />; */}
        </div>
    )
}

export default document
