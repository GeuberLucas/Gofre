"use client";

import { ColumnDef } from "@tanstack/react-table";
import { Expense } from "./model/expense";
import DropDownActions from "@/components/actions-dropdown";
import { TransactionType } from "@/enums/TypeTransactions";

export const columns: ColumnDef<Expense>[] = [
  {
    accessorKey: "description",
    header: "Descrição",
  },
  {
    accessorKey: "target",
    header: "Destino",
  },
  {
    accessorKey: "category",
    header: "Categoria",
  },
  {
    accessorKey: "type",
    header: "Tipo",
  },
  {
    accessorKey: "paymentMethod",
    header: "Pagamento",
  },
  {
    accessorKey: "paymentDate",
    header: "Data",
    cell: ({ row }) => {
      const dateValue = row.getValue("paymentDate");
      if (!dateValue) return "-";

      const date = new Date(dateValue as string);
      return date.toLocaleDateString("pt-BR");
    },
  },
  {
    accessorKey: "amount",
    header: "Valor",
    cell: ({ row }) => {
      const amount = Number.parseFloat(row.getValue("amount"));
      return new Intl.NumberFormat("pt-BR", {
        style: "currency",
        currency: "BRL",
      }).format(amount);
    },
  },
  {
    accessorKey: "isPaid",
    header: "Pago?",
    cell: ({ row }) => {
      const isPaid = row.getValue("isPaid");
      return isPaid ? "Sim" : "Não";
    },
  },
  {
    id: "Actions",
    header: "Ações",
    cell: ({ row }) => {
      const expense = row.original;

      return (
        <DropDownActions
          idTransaction={expense.id}
          transactionType={TransactionType.Expense}
          executedTransaction={expense.isPaid}
        />
      );
    },
  },
];
