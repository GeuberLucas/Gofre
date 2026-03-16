"use client";
import { DataTable } from "./_components/data-table";
import { FinancialSummaryProps } from "../financialSummaryProps";
import { Portfolio } from "./model/portfolio";
import { useEffect, useState } from "react";
import { getPortfolio } from "./services/investment-service";
import { InfoComponent } from "./_components/info-component";
import DetailInvestment from "./_components/detail-dialog";
import { columns } from "../revenue/_components/columns";

function getSummary(expenses: Portfolio[]) {
  const today = new Date();
  const actualMonth = today.getMonth();
  const actualYear = today.getFullYear();
  const investmentThisMonth = expenses.filter((expense) => {
    const dateObj = new Date(expense.deposit_date);
    if (Number.isNaN(dateObj.getTime())) return false;
    return (
      dateObj.getMonth() === actualMonth && dateObj.getFullYear() === actualYear
    );
  });

  const expectedAmount = investmentThisMonth.reduce((accumulator, expense) => {
    return accumulator + expense.amount;
  }, 0);
  const actualAmount = investmentThisMonth
    .filter((exp) => exp.is_done)
    .reduce((accumulator, expense) => {
      return accumulator + expense.amount;
    }, 0);
  const pendingAmount = investmentThisMonth
    .filter((exp) => !exp.is_done)
    .reduce((accumulator, expense) => {
      return accumulator + expense.amount;
    }, 0);

  return {
    expectedAmount: expectedAmount,
    actualAmount: actualAmount,
    pendingAmount: pendingAmount,
    currentBalance: 0,
  };
}

const ToggleState = (rowState: boolean) => {
  let stateRevenue = "bg-action-pending";
  if (rowState) {
    stateRevenue = "bg-action-realized";
  }
  return (
    <button className={`p-2 mr-2 rounded-2xl ${stateRevenue}`}>
      Toggle Recebido
    </button>
  );
};
export default function Revenues() {
  const [isOpenDialog, setIsOpenDialog] = useState(false);
  const [data, setData] = useState([]);
  const [financialSummary, setfinancialSummary] = useState({
    expectedAmount: 0,
    actualAmount: 0,
    pendingAmount: 0,
    currentBalance: 0,
  });
  const handleCloseDialog = () => setIsOpenDialog(false);
  const handleOpenDialog = () => setIsOpenDialog(true);

  useEffect(() => {
    getPortfolio().then((data: Portfolio[]) => {
      setData(data);
      setfinancialSummary(getSummary(data));
    });
  }, []);
  return (
    <div className="flex h-screen w-full overflow-hidden">
      <div className="flex-1 flex flex-col overflow-y-auto p-4 md:p-6 gap-6">
        <div className="w-full flex justify-start">
          <button
            className="bg-investment hover:bg-investment/90 text-white font-display py-2 px-6 rounded-xl shadow-sm transition-colors text-sm md:text-base"
            onClick={handleOpenDialog}
          >
            Adicionar aporte
          </button>
        </div>

        <div className="w-full rounded-xl border bg-card shadow-sm">
          <DataTable columns={columns} data={data} />
        </div>
      </div>
      <aside className="w-96 hidden xl:flex flex-col border-l bg-muted/10 h-full p-6 overflow-y-auto">
        <div className="sticky top-0">
          <h2 className="text-sm font-semibold text-muted-foreground uppercase tracking-wider mb-6">
            Resumo
          </h2>

          <InfoComponent
            expectedAmount={financialSummary.expectedAmount}
            pendingAmount={financialSummary.pendingAmount}
            actualAmount={financialSummary.actualAmount}
            currentBalance={financialSummary.currentBalance}
          />
        </div>
      </aside>
      <DetailInvestment
        onClose={handleCloseDialog}
        open={isOpenDialog}
        id={0}
      />
    </div>
  );
}
