import { Operator } from './operator';

export interface Filter {
  field: string;
  operator: Operator;
  value: any;
}
