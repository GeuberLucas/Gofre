export interface Expense {
  id: number;
  userId: number;
  description: string;
  target: string;
  category: string;
  type: string;
  paymentMethod: string;
  paymentDate: Date;
  isPaid: boolean;
  amount: number;
}
