"use client";
import { DataTable } from "./data-table";
import { columns } from "./columns";
import { InfoComponent } from "./info-component";
import { Expense } from "./expense";
import { FinancialSummaryProps } from "../financialSummaryProps";
import Detail from "./detail-dialog";
import { useState } from "react";
import DetailExpense from "./detail-dialog";

const props: FinancialSummaryProps = {
  expectedAmount: 0,
  actualAmount: 0,
  pendingAmount: 0,
  currentBalance: 0,
  invoiceAmount: 0,
  variableAmount: 0,
  monthlyAmount: 0,
};
const rows: Expense[] = [
  {
    id: 1,
    userId: 101,
    description: "Aluguel Apartamento",
    target: "Imobiliária Silva",
    category: "Moradia",
    type: "Fixo",
    paymentMethod: "Boleto",
    paymentDate: new Date("2026-01-01T00:00:00Z"),
    isPaid: false,
    amount: 1500,
  },
  {
    id: 2,
    userId: 101,
    description: "Supermercado Mensal",
    target: "Carrefour",
    category: "Alimentação",
    type: "Variável",
    paymentMethod: "Cartão de Crédito",
    paymentDate: new Date("2026-01-25T00:00:00Z"), // Já foi pago
    isPaid: true,
    amount: 850.4,
  },
  {
    id: 3,
    userId: 101,
    description: "Conta de Luz",
    target: "Cemig",
    category: "Utilidades",
    type: "Fixo",
    paymentMethod: "Pix",
    paymentDate: new Date("2026-02-01T00:00:00Z"), // Vence Hoje (considerando 01/02/2026))
    isPaid: false,
    amount: 125.3,
  },
  {
    id: 4,
    userId: 101,
    description: "Assinatura Netflix",
    target: "Netflix",
    category: "Lazer",
    type: "Fixo",
    paymentMethod: "Cartão de Crédito",
    paymentDate: new Date("2026-02-15T00:00:00Z"), // Futur)o
    isPaid: false,
    amount: 55.9,
  },
  {
    id: 5,
    userId: 101,
    description: "Manutenção Carro",
    target: "Oficina do João",
    category: "Transporte",
    type: "Eventual",
    paymentMethod: "Débito",
    paymentDate: new Date("2026-01-20T00:00:00Z"), // Atrasado, mas está pag)o
    isPaid: true,
    amount: 450.0,
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
            className="bg-expense hover:bg-expense/90 text-white font-display py-2 px-6 rounded-xl shadow-sm transition-colors text-sm md:text-base"
            onClick={handleOpenDialog}
          >
            Adicionar saida
          </button>
        </div>

        <div className="w-full rounded-xl border bg-card shadow-sm">
          <DataTable columns={columns} data={rows} />
        </div>
      </div>
      <aside className="w-96 hidden xl:flex flex-col border-l bg-muted/10 h-full p-6 overflow-y-auto">
        <div className="sticky top-0">
          <h2 className="text-sm font-semibold text-muted-foreground uppercase tracking-wider mb-4">
            Resumo
          </h2>

          <InfoComponent
            expectedAmount={props.expectedAmount}
            pendingAmount={props.pendingAmount}
            actualAmount={props.actualAmount}
            currentBalance={props.currentBalance}
            invoiceAmount={props.invoiceAmount}
            monthlyAmount={props.monthlyAmount}
            variableAmount={props.variableAmount}
          />
        </div>
      </aside>
      <DetailExpense onClose={handleCloseDialog} open={isOpenDialog} id={0} />
    </div>
  );
}
