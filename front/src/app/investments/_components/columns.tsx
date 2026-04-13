"use client";

import { ColumnDef } from "@tanstack/react-table";

import { TransactionType } from "@/enums/TypeTransactions";
import DropDownActions from "@/components/actions-dropdown";
import { Portfolio } from "../model/portfolio";
import { AssetClass } from "../model/asset-class";

export const getColumns = (
  assetsClasses: AssetClass[],
): ColumnDef<Portfolio>[] => [
  {
    accessorKey: "description",
    header: "Descrição",
  },
  {
    accessorKey: "broker",
    header: "Corretora/Banco",
  },
  {
    accessorKey: "asset_id",
    header: "Ativo",
    cell: ({ row }) => {
      const asset = assetsClasses.find(
        (ass) => ass.id == row.getValue("asset_id"),
      );
      return asset?.name;
    },
  },
  {
    accessorKey: "deposit_date",
    header: "Data",
    cell: ({ row }) => {
      const date = new Date(row.getValue("deposit_date"));
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
    accessorKey: "is_done",
    header: "Feito ?",
    cell: ({ row }) => {
      const recieved = Boolean(row.getValue("is_done"));

      return recieved ? "Sim" : "Não";
    },
  },
  {
    id: "Actions",
    header: "Ações",
    cell: ({ row }) => {
      const investment = row.original;

      return (
        <DropDownActions
          idTransaction={investment.id}
          transactionType={TransactionType.Investment}
          executedTransaction={investment.is_done}
        />
      );
    },
  },
];
