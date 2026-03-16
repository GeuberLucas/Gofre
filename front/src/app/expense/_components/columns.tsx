"use client";

import { ColumnDef } from "@tanstack/react-table";
import DropDownActions from "@/components/actions-dropdown";
import { TransactionType } from "@/enums/TypeTransactions";
import { enumToFormattedOptions } from "@/app/enum-to-option";
import { CategoryExpenseEnum } from "../enums/category-expense-enum";
import { PaymentMethodEnum } from "../enums/payment-method-enum";
import { TypeExpenseEnum } from "../enums/type-expense-enum";
import { Expense } from "../model/expense";

const categorys = enumToFormattedOptions(CategoryExpenseEnum);
const typeExpense = enumToFormattedOptions(TypeExpenseEnum);
const paymentMethods = enumToFormattedOptions(PaymentMethodEnum);
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
    cell: ({ row }) => {
      const dateValue = row.getValue("category");
      const category = categorys.findLast((key) => key.valor == dateValue);
      return category.texto;
    },
  },
  {
    accessorKey: "type",
    header: "Tipo",
    cell: ({ row }) => {
      const dateValue = row.getValue("type");
      const type = typeExpense.findLast((key) => key.valor == dateValue);
      return type.texto;
    },
  },
  {
    accessorKey: "paymentMethod",
    header: "Pagamento",
    cell: ({ row }) => {
      const dateValue = row.getValue("paymentMethod");
      const paymentMethod = paymentMethods.findLast(
        (key) => key.valor == dateValue,
      );
      return paymentMethod.texto;
    },
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
