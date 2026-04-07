import { IncomeType } from "../enums/income-type";

export interface Revenue {
  id: number;
  description: string;
  origin: string;
  type: IncomeType;
  receiveDate: Date;
  amount: number;
  recieved: boolean;
}
