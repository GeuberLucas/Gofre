import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  CircleCheckBig,
  CircleOff,
  MoreHorizontal,
  PenLine,
  Trash,
} from "lucide-react";
import { useState } from "react";

import { TransactionType } from "@/enums/TypeTransactions";
import DetailRevenue from "@/app/revenue/_components/detail-dialog";
import DetailExpense from "@/app/expense/_components/detail-dialog";
import DetailInvestment from "@/app/investments/_components/detail-dialog";
import { deleteRevenue } from "@/app/revenue/services/revenue-service";
import { deleteExpense } from "@/app/expense/services/expense-service";
import { deleteInvestment } from "@/app/investments/services/investment-service";

//TODO:IMPLEMENT LOGIC ALTER STATUS
//TODO:ALTER DELETE LOGIC
const deleteTransaction = (type: TransactionType, idTransaction: number) => {
  switch (type) {
    case TransactionType.Revenue:
      deleteRevenue(idTransaction);
      break;
    case TransactionType.Expense:
      deleteExpense(idTransaction);
      break;
    case TransactionType.Investment:
      deleteInvestment(idTransaction);
      break;
  }
};
const getIconByBool = (executedTransaction) => {
  if (executedTransaction) {
    return <CircleCheckBig className="text-emerald-600" />;
  }

  return <CircleOff className="text-red-500" />;
};
const textStatus = (type: TransactionType): string => {
  switch (type) {
    case TransactionType.Revenue:
      return "Recebimento";
    case TransactionType.Expense:
      return "Pagamento";
    case TransactionType.Investment:
      return "Aporte";
  }
};
const DetailDialog = (
  type: TransactionType,
  handleCloseDialog: () => void,
  isOpenDialog: boolean,
  idTransaction: number,
  handleSucces?: () => void,
) => {
  switch (type) {
    case TransactionType.Revenue:
      return (
        <DetailRevenue
          onSuccess={handleSucces}
          onClose={handleCloseDialog}
          open={isOpenDialog}
          id={idTransaction}
        />
      );
    case TransactionType.Expense:
      return (
        <DetailExpense
          onClose={handleCloseDialog}
          open={isOpenDialog}
          id={idTransaction}
        />
      );
    case TransactionType.Investment:
      return (
        <DetailInvestment
          onClose={handleCloseDialog}
          open={isOpenDialog}
          id={idTransaction}
        />
      );
  }
};
const DropDownActions = ({
  idTransaction,
  transactionType,
  executedTransaction,
}) => {
  const [isOpenDialog, setIsOpenDialog] = useState(false);
  const handleCloseDialog = () => setIsOpenDialog(false);
  const handleOpenDialog = () => setIsOpenDialog(true);

  return (
    <div>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="ghost" className="h-8 w-8 p-0">
            <span className="sr-only">Abrir menu</span>
            <MoreHorizontal className="h-4 w-4" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuItem onClick={handleOpenDialog}>
            <PenLine />
            <span className="font-display">Editar</span>
          </DropdownMenuItem>
          <DropdownMenuItem
            onClick={() => deleteTransaction(transactionType, idTransaction)}
          >
            <Trash />
            <span className="font-display">Remover</span>
          </DropdownMenuItem>
          <DropdownMenuItem>
            {getIconByBool(executedTransaction)}
            <span className="font-display">
              Alterar {textStatus(transactionType)}
            </span>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
      {DetailDialog(
        transactionType,
        handleCloseDialog,
        isOpenDialog,
        idTransaction,
      )}
    </div>
  );
};

export default DropDownActions;
