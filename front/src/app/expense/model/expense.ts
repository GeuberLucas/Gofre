import { CategoryExpenseEnum } from "../enums/category-expense-enum";
import { PaymentMethodEnum } from "../enums/payment-method-enum";
import { TypeExpenseEnum } from "../enums/type-expense-enum";

export interface Expense {
  id: number;
  userId: number;
  description: string;
  target: string;
  category: CategoryExpenseEnum;
  type: TypeExpenseEnum;
  paymentMethod: PaymentMethodEnum;
  paymentDate: Date;
  isPaid: boolean;
  amount: number;
}
