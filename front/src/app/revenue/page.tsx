import { div, th } from "framer-motion/client";
import React from "react";

interface Revenue {
  id: number;
  description: string;
  origin: string;
  type: string;
  date: Date;
  amount: number;
  recieved: boolean;
}

const initId = 1;
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
const columns = [
  {
    key: "description",
    label: "Descrição",
  },
  {
    key: "origin",
    label: "Origem",
  },
  {
    key: "type",
    label: "Tipo",
  },
  {
    key: "date",
    label: "Data",
  },
  {
    key: "amount",
    label: "Valor",
  },
  {
    key: "recieved",
    label: "Recebido?",
  },
];
function GetValue(field: keyof Revenue, revenue: Revenue) {
  const value = revenue[field];
  if (value instanceof Date) {
    return <span className="font-display ">{value.toLocaleDateString()}</span>;
  }
  if (typeof value == "number") {
    return <span className="font-number ">{TransformInCurrency(value)}</span>;
  }
  if (typeof value == "boolean") {
    return <span className="font-display">{value ? "Sim" : "Nâo"}</span>;
  }

  return <span className="font-display">{value}</span>;
}
function TransformInCurrency(value: number) {
  return value.toLocaleString("pt-BR", {
    style: "currency",
    currency: "BRL",
  });
}
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
  return (
    <div className="h-screen w-full flex flex-col p-4">
      <div className="h-1/4 w-full flex flex-row align-middle">
        {/* header information */}
        <div className="w-full inline-grid grid-cols-12 gap-3">
          <div className="col-span-4 flex justify-center align-middle">
            <button className="lg:text-2xl h-fit md:text-sm py-3 px-1 rounded-2xl w-50 bg-revenue text-white font-display">
              Nova entrada
            </button>
          </div>
          <div className="col-span-2 flex flex-col text-center">
            <span className="font-display text-sm">PREVISTO</span>
            <span className="text-2xl font-bold font-number">00,00</span>
          </div>
          <div className="col-span-2 flex flex-col text-center align-middle">
            <span className="font-display text-sm">REALIZADO</span>
            <span className="text-2xl font-bold font-number">00,00</span>
          </div>
          <div className="col-span-2 flex flex-col text-center align-middle">
            <span className="font-display text-sm">PENDENTE</span>
            <span className="text-2xl font-bold font-number">00,00</span>
          </div>
          <div className="col-span-2 flex flex-col text-center align-middle">
            <span className="font-display text-sm">SALDO</span>
            <span className="text-2xl font-bold font-number">00,00</span>
          </div>
        </div>
      </div>
      <div className="h-3/4 w-full px-4">
        {/* table data */}
        {/* table header */}
        <div className="w-full bg-revenue text-white flex flex-row">
          {columns.map((column) => {
            return (
              <div key={column.key} className="flex-1 p-2 text-center">
                <span className="font-display text-xl text-center font-medium">
                  {column.label}
                </span>
              </div>
            );
          })}
          <div className="flex-1 p-2">
            <span className="font-display text-xl text-center">Ações</span>
          </div>
        </div>
        {/* {/* Table data */}
        <div className="w-full text-white overflow-auto h-full">
          {rows.map((row, index) => {
            row.id = initId + index;
            return (
              <div
                key={row.id}
                className="mt-2 even:bg-grey-dark odd:bg-grey-light flex flex-row"
              >
                {columns.map((column) => {
                  return (
                    <div
                      key={column.key}
                      className="flex-1 text-center text-xl "
                    >
                      {GetValue(column.key as keyof Revenue, row)}
                    </div>
                  );
                })}
                <div key="Actions" className="text-center p-2 flex-1 ">
                  <button className="p-2 mr-2 rounded-2xl bg-action-primary">
                    Editar
                  </button>
                  {ToggleState(row.recieved)}
                  <button className="p-2 mr-2 rounded-2xl bg-action-danger">
                    Remover
                  </button>
                </div>
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
}
