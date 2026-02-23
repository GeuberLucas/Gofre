"use client";
import { Revenue } from "./revenue";
import { DataTable } from "./data-table";
import { columns } from "./columns";
import { InfoComponent } from "./info-component";
import { FinancialSummaryProps } from "../financialSummaryProps";
import { useState } from "react";
import Detail from "./detail-dialog";
import DetailRevenue from "./detail-dialog";

const props: FinancialSummaryProps = {
  expectedAmount: 0,
  actualAmount: 0,
  pendingAmount: 0,
  currentBalance: 0,
};
const rows: Revenue[] = [
  {
    id: 0,
    description: "salario",
    origin: "empresa x",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: true,
  },
  {
    id: 0,
    description: "salario",
    origin: "fulano",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: false,
  },
  {
    id: 0,
    description: "salario",
    origin: "empresa x",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: true,
  },
  {
    id: 0,
    description: "salario",
    origin: "fulano",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: false,
  },
  {
    id: 0,
    description: "salario",
    origin: "empresa x",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: true,
  },
  {
    id: 0,
    description: "salario",
    origin: "fulano",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: false,
  },
  {
    id: 0,
    description: "salario",
    origin: "empresa x",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: true,
  },
  {
    id: 0,
    description: "salario",
    origin: "fulano",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: false,
  },
  {
    id: 0,
    description: "salario",
    origin: "empresa x",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: true,
  },
  {
    id: 0,
    description: "salario",
    origin: "fulano",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: false,
  },
  {
    id: 0,
    description: "salario",
    origin: "empresa x",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: true,
  },
  {
    id: 0,
    description: "salario",
    origin: "fulano",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: false,
  },
  {
    id: 0,
    description: "salario",
    origin: "empresa x",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: true,
  },
  {
    id: 0,
    description: "salario",
    origin: "fulano",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: false,
  },
  {
    id: 0,
    description: "salario",
    origin: "empresa x",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: true,
  },
  {
    id: 0,
    description: "salario",
    origin: "fulano",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: false,
  },
  {
    id: 0,
    description: "salario",
    origin: "empresa x",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: true,
  },
  {
    id: 0,
    description: "salario",
    origin: "fulano",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: false,
  },
  {
    id: 0,
    description: "salario",
    origin: "empresa x",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: true,
  },
  {
    id: 0,
    description: "salario",
    origin: "fulano",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: false,
  },
  {
    id: 0,
    description: "salario",
    origin: "empresa x",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: true,
  },
  {
    id: 0,
    description: "salario",
    origin: "fulano",
    type: "Trabalho",
    date: new Date(),
    amount: 3766.04,
    recieved: false,
  },
];
export default function Revenues() {
  const [isOpenDialog, setIsOpenDialog] = useState(false);
  const handleCloseDialog = () => setIsOpenDialog(false);
  const handleOpenDialog = () => setIsOpenDialog(true);

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
          <DataTable columns={columns} data={rows} />
        </div>
      </div>
      <aside className="w-96 hidden xl:flex flex-col border-l bg-muted/10 h-full p-6 overflow-y-auto">
        <div className="sticky top-0">
          <h2 className="text-sm font-semibold text-muted-foreground uppercase tracking-wider mb-6">
            Resumo
          </h2>

          <InfoComponent
            expectedAmount={props.expectedAmount}
            pendingAmount={props.pendingAmount}
            actualAmount={props.actualAmount}
            currentBalance={props.currentBalance}
          />
        </div>
      </aside>

      <DetailRevenue onClose={handleCloseDialog} open={isOpenDialog} id={0} />
    </div>
  );
}
