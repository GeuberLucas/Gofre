export interface IRevenue {
  id: number;
  description: string;
  origin: string;
  type: string;
  date: Date;
  amount: number;
  recieved: boolean;
}

export class Revenue implements IRevenue {
  id: number = 0;
  description: string = "";
  origin: string = "";
  type: string = "";
  date: Date = new Date();
  amount: number = 0.0;
  recieved: boolean = false;
}
