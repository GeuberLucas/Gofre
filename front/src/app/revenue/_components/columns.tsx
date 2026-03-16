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
    accessorKey: "date",
    header: "Data",
    cell: ({ row }) => {
      const unixDate = Date.parse(row.getValue("date"));
      const date = new Date(unixDate);
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
    accessorKey: "recieved",
    header: "Recebido ?",
    cell: ({ row }) => {
      const recieved = Boolean(row.getValue("recieved"));

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
          executedTransaction={revenue.recieved}
        />
      );
    },
  },
];
