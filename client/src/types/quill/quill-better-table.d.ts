declare module 'quill-better-table' {
    import { Quill } from 'quill';
  
    interface BetterTableOptions {
      operationMenu?: {
        items?: {
          insertColumnLeft?: { text: string };
          insertColumnRight?: { text: string };
          deleteColumn?: { text: string };
          insertRowUp?: { text: string };
          insertRowDown?: { text: string };
          deleteRow?: { text: string };
          deleteTable?: { text: string };
        };
      };
    }
  
    export default class QuillBetterTable {
      constructor(quill: Quill, options: BetterTableOptions);
    }
  }