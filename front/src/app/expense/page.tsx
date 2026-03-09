"use client";
import { useEffect, useState } from "react";
import DetailExpense from "./_components/detail-dialog";
import { DataTable } from "./_components/data-table";
import { InfoComponent } from "./_components/info-component";
import { columns } from "./_components/columns";
import { getExpense } from "./services/expense-service";
import { Expense } from "./model/expense";
import { PaymentMethodEnum } from "./enums/payment-method-enum";
import { TypeExpenseEnum } from "./enums/type-expense-enum";

function getSummary(expenses: Expense[]) {
  const today = new Date();
  const actualMonth = today.getMonth();
  const actualYear = today.getFullYear();
  const expensesThisMonth = expenses.filter((expense) => {
    const dateObj = new Date(expense.paymentDate);

    if (Number.isNaN(dateObj.getTime())) return false;

    return (
      dateObj.getMonth() === actualMonth && dateObj.getFullYear() === actualYear
    );
  });
  const notCredit = expensesThisMonth.filter(
    (exp) => exp.paymentMethod.toLocaleUpperCase() != PaymentMethodEnum.CREDITO,
  );
  const expectedAmount = notCredit.reduce((accumulator, expense) => {
    return accumulator + expense.amount;
  }, 0);
  const actualAmount = notCredit
    .filter((exp) => exp.isPaid)
    .reduce((accumulator, expense) => {
      return accumulator + expense.amount;
    }, 0);
  const pendingAmount = notCredit
    .filter((exp) => !exp.isPaid)
    .reduce((accumulator, expense) => {
      return accumulator + expense.amount;
    }, 0);
  const invoiceAmount = expensesThisMonth
    .filter((exp) => exp.type.toUpperCase() == TypeExpenseEnum.FATURA)
    .reduce((accumulator, expense) => {
      return accumulator + expense.amount;
    }, 0);
  const variableAmount = expensesThisMonth
    .filter((exp) => exp.type.toUpperCase() == TypeExpenseEnum.VARIAVEL)
    .reduce((accumulator, expense) => {
      return accumulator + expense.amount;
    }, 0);
  const monthlyAmount = expensesThisMonth
    .filter((exp) => exp.type.toUpperCase() == TypeExpenseEnum.MENSAL)
    .reduce((accumulator, expense) => {
      return accumulator + expense.amount;
    }, 0);
  return {
    expectedAmount: expectedAmount,
    actualAmount: actualAmount,
    pendingAmount: pendingAmount,
    currentBalance: 0,
    invoiceAmount: invoiceAmount,
    variableAmount: variableAmount,
    monthlyAmount: monthlyAmount,
  };
}

export default function Revenues() {
  const [isOpenDialog, setIsOpenDialog] = useState(false);
  const [data, setData] = useState([]);
  const [financialSummary, setfinancialSummary] = useState({
    expectedAmount: 0,
    actualAmount: 0,
    pendingAmount: 0,
    currentBalance: 0,
    invoiceAmount: 0,
    variableAmount: 0,
    monthlyAmount: 0,
  });
  const handleCloseDialog = () => setIsOpenDialog(false);
  const handleOpenDialog = () => setIsOpenDialog(true);

  //fetch data
  useEffect(() => {
    getExpense().then((data: Expense[]) => {
      setData(data);
      setfinancialSummary(getSummary(data));
    });
  }, []);

  return (
    <div className="flex h-screen w-full overflow-hidden">
      <div className="flex-1 flex flex-col overflow-y-auto p-4 md:p-6 gap-6">
        <div className="w-full flex justify-start">
          <button
            className="bg-expense hover:bg-expense/90 text-white font-display py-2 px-6 rounded-xl shadow-sm transition-colors text-sm md:text-base"
            onClick={handleOpenDialog}
          >
            Adicionar saida
          </button>
        </div>

        <div className="w-full rounded-xl border bg-card shadow-sm">
          <DataTable columns={columns} data={data} />
        </div>
      </div>
      <aside className="w-96 hidden xl:flex flex-col border-l bg-muted/10 h-full p-6 overflow-y-auto">
        <div className="sticky top-0">
          <h2 className="text-sm font-semibold text-muted-foreground uppercase tracking-wider mb-4">
            Resumo
          </h2>

          <InfoComponent
            expectedAmount={financialSummary.expectedAmount}
            pendingAmount={financialSummary.pendingAmount}
            actualAmount={financialSummary.actualAmount}
            currentBalance={financialSummary.currentBalance}
            invoiceAmount={financialSummary.invoiceAmount}
            monthlyAmount={financialSummary.monthlyAmount}
            variableAmount={financialSummary.variableAmount}
          />
        </div>
      </aside>
      <DetailExpense onClose={handleCloseDialog} open={isOpenDialog} id={0} />
    </div>
  );
}
