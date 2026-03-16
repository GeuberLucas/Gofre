"use client";
import { Revenue } from "./model/revenue";
import { DataTable } from "./_components/data-table";
import { columns } from "./_components/columns";
import { InfoComponent } from "./_components/info-component";
import { useEffect, useState } from "react";
import DetailRevenue from "./_components/detail-dialog";
import { getRevenue } from "./services/revenue-service";

function getSummary(expenses: Revenue[]) {
  const today = new Date();
  const actualMonth = today.getMonth();
  const actualYear = today.getFullYear();
  const investmentThisMonth = expenses.filter((expense) => {
    const dateObj = new Date(expense.date);
    if (Number.isNaN(dateObj.getTime())) return false;
    return (
      dateObj.getMonth() === actualMonth && dateObj.getFullYear() === actualYear
    );
  });

  const expectedAmount = investmentThisMonth.reduce((accumulator, expense) => {
    return accumulator + expense.amount;
  }, 0);
  const actualAmount = investmentThisMonth
    .filter((exp) => exp.recieved)
    .reduce((accumulator, expense) => {
      return accumulator + expense.amount;
    }, 0);
  const pendingAmount = investmentThisMonth
    .filter((exp) => !exp.recieved)
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
export default function Revenues() {
  const [isOpenDialog, setIsOpenDialog] = useState(false);
  const handleCloseDialog = () => setIsOpenDialog(false);
  const handleOpenDialog = () => setIsOpenDialog(true);
  const [data, setData] = useState([]);
  const [financialSummary, setfinancialSummary] = useState({
    expectedAmount: 0,
    actualAmount: 0,
    pendingAmount: 0,
    currentBalance: 0,
  });
  useEffect(() => {
    getRevenue().then((data: Revenue[]) => {
      setData(data);
      setfinancialSummary(getSummary(data));
    });
  }, []);
  return (
    <div className="flex h-screen w-full overflow-hidden">
      <div className="flex-1 flex flex-col overflow-y-auto p-4 md:p-6 gap-6">
        <div className="w-full flex justify-start">
          <button
            className="bg-revenue hover:bg-revenue/90 text-white font-display py-2 px-6 rounded-xl shadow-sm transition-colors text-sm md:text-base"
            onClick={handleOpenDialog}
          >
            Adicionar entrada
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

      <DetailRevenue onClose={handleCloseDialog} open={isOpenDialog} id={0} />
    </div>
  );
}
