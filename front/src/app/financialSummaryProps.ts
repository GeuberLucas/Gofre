export interface FinancialSummaryProps {
  expectedAmount: number;
  actualAmount: number;
  pendingAmount: number;
  currentBalance: number;
  invoiceAmount?: number;
  monthlyAmount?: number;
  variableAmount?: number;
}
