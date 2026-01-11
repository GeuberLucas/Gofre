import React from "react";
const rows = [
  {
    id: 1,
    description: "",
    origin: "",
    type: "",
    date: new Date(),
    amount: 0,
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
    label: "Recebido",
  },
];
export default function Revenues() {
  return (
    <div className="w-full flex justify-center">
      {/* header information */}
      <div className="inline-grid grid-cols-5 gap-1">
        <div className="p-4 mr-4">
          <button className="lg:text-2xl md:text-sm p-2 rounded-2xl w-fit bg-[#2b6051ff] text-white font-display">
            Nova entrada
          </button>
        </div>
        <div className="flex flex-col text-end">
          <span className="font-display text-sm">PREVISTO</span>
          <span className="text-2xl font-bold font-number">00,00</span>
        </div>
        <div className="flex flex-col text-end">
          <span className="font-display text-sm">REALIZADO</span>
          <span className="text-2xl font-bold font-number">00,00</span>
        </div>
        <div className="flex flex-col text-end">
          <span className="font-display text-sm">PENDENTE</span>
          <span className="text-2xl font-bold font-number">00,00</span>
        </div>
        <div className="flex flex-col text-end">
          <span className="font-display text-sm">SALDO</span>
          <span className="text-2xl font-bold font-number">00,00</span>
        </div>
      </div>
      {/* table data */}
      <div></div>
    </div>
  );
}
