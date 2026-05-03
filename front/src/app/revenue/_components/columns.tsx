"use client";

import { ColumnDef } from "@tanstack/react-table";
import { Revenue } from "../model/revenue";
import DropDownActions from "../../../components/actions-dropdown";
import { TransactionType } from "@/enums/TypeTransactions";

export const columns: ColumnDef<Revenue>[] = [
  {
    accessorKey: "description",
    header: "Descrição",
  },
  {
    accessorKey: "origin",
    header: "Origem",
  },
  {
    accessorKey: "type",
    header: "Tipo",
  },
  {
    accessorKey: "receiveDate",
    header: "Data",
    cell: ({ row }) => {
      const date = new Date(row.getValue("receiveDate"));
      return date.toLocaleDateString("pt-BR");
    },
  },
  {
    accessorKey: "amount",
    header: "Valor",
    cell: ({ row }) => {
      const amount = Number.parseFloat(row.getValue("amount"));
      const formatted = new Intl.NumberFormat("pt-BR", {
        style: "currency",
        currency: "BRL",
      }).format(amount);

      return formatted;
    },
  },
  {
    accessorKey: "isRecieved",
    header: "Recebido ?",
    cell: ({ row }) => {
      const recieved = Boolean(row.getValue("isRecieved"));
      return recieved ? "Sim" : "Não";
    },
  },
  {
    id: "Actions",
    header: "Ações",
    cell: ({ row }) => {
      const revenue = row.original;

      return (
        <DropDownActions
          idTransaction={revenue.id}
          transactionType={TransactionType.Revenue}
          executedTransaction={revenue.isRecieved}
        />
      );
    },
  },
];
