import { TEntityId } from "./common";

export type Cheque = {
  id: TEntityId;
  bank: string;
  check_date: Date;
  media_id: TEntityId;
  description: string;
};
