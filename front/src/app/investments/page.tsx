"use client";
import { DataTable } from "./data-table";
import { columns } from "./columns";
import { InfoComponent } from "./info-component";
import { FinancialSummaryProps } from "../financialSummaryProps";
import { Portfolio } from "./portfolio";
import { useState } from "react";
import Detail from "./detail-dialog";
import DetailInvestment from "./detail-dialog";

const props: FinancialSummaryProps = {
  expectedAmount: 0,
  actualAmount: 0,
  pendingAmount: 0,
  currentBalance: 0,
};
const rows: Portfolio[] = [
  {
    id: 1,
    user_id: 101,
    asset_id: 5001,
    deposit_date: new Date("2023-11-15T14:30:00Z"),
    broker: "XP Investimentos",
    amount: 5000.0,
    is_done: true,
    description: "Aporte inicial em FIIs",
  },
  {
    id: 2,
    user_id: 101,
    asset_id: 3005,
    deposit_date: new Date("2024-01-10T09:15:00Z"),
    broker: "Binance",
    amount: 1250.5,
    is_done: true,
    description: "Compra fracionada de Bitcoin",
  },
  {
    id: 3,
    user_id: 102,
    asset_id: 1022,
    deposit_date: new Date("2024-02-01T10:00:00Z"),
    broker: "NuInvest",
    amount: 300.0,
    is_done: false, // Transação pendente
    description: "Reinvestimento de dividendos",
  },
  {
    id: 4,
    user_id: 103,
    asset_id: 4004,
    deposit_date: new Date("2024-02-28T16:45:00Z"),
    broker: "BTG Pactual",
    amount: 15000.0,
    is_done: true,
    description: "Tesouro Direto IPCA+ 2035",
  },
  {
    id: 5,
    user_id: 101,
    asset_id: 5001,
    deposit_date: new Date("2024-03-05T11:20:00Z"),
    broker: "Rico",
    amount: 75.9,
    is_done: true,
    description: "Sobra de caixa",
  },
];

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
  const handleCloseDialog = () => setIsOpenDialog(false);
  const handleOpenDialog = () => setIsOpenDialog(true);

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
      <DetailInvestment
        onClose={handleCloseDialog}
        open={isOpenDialog}
        id={0}
      />
    </div>
  );
}
