import { FinancialSummaryProps } from "../financialSummaryProps";

const formatCurrency = (value: number) =>
  Intl.NumberFormat("pt-BR", {
    style: "currency",
    currency: "BRL",
  }).format(value);

export function InfoComponent(summaryInfo: Readonly<FinancialSummaryProps>) {
  return (
    <div className="h-full flex flex-col justify-center">
      <div className="col-span-2 flex flex-col text-start my-2">
        <span className="font-display text-sm md:text-3xl">PREVISTO</span>
        <span className="text-2xl md:text-4xl font-bold font-number text-blue-400">
          {formatCurrency(summaryInfo.expectedAmount)}
        </span>
      </div>

      <div className="col-span-2 flex flex-col text-start my-2">
        <span className="font-display text-sm md:text-3xl">REALIZADO</span>
        <span className="text-2xl md:text-4xl font-bold font-number text-emerald-400">
          {formatCurrency(summaryInfo.actualAmount)}
        </span>
      </div>

      <div className="col-span-2 flex flex-col text-start my-2">
        <span className="font-display text-sm md:text-3xl">PENDENTE</span>
        <span className="text-2xl md:text-4xl font-bold font-number text-amber-400">
          {formatCurrency(summaryInfo.pendingAmount)}
        </span>
      </div>

      <div className="col-span-2 flex flex-col text-start my-2">
        <span className="font-display text-sm md:text-3xl">SALDO ATUAL</span>
        <span
          className={`text-2xl md:text-4xl font-bold font-number ${summaryInfo.currentBalance < 0 ? "text-red-500" : "text-emerald-600"}`}
        >
          {formatCurrency(summaryInfo.currentBalance)}
        </span>
      </div>
    </div>
  );
}
