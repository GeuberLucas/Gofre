"use server";

import { ApiClient } from "@/lib/httpClient";
import { Portfolio } from "../model/portfolio";
import { AssetClass } from "../model/asset-class";

const defaultHeaders = {
  "Content-Type": "application/json",
  Authorization: `Bearer ${process.env.TOKEN}`,
};

const baseUrl = `investments`;

function buildUrl(id?: number) {
  return id && id > 0 ? `${baseUrl}/${id}` : baseUrl;
}

export async function getPortfolio(
  id?: number,
): Promise<Portfolio[] | Portfolio> {
  const res = await ApiClient.request<Portfolio[] | Portfolio>(buildUrl(id), {
    method: "GET",
  });
  if (!res.success) {
    const errorBody = res.data;
    console.error({ status: res.statusCode, msg: errorBody });
    return;
  }

  const expenses: Portfolio[] | Portfolio = res.data;
  return expenses;
}

export async function sendPortfolio(expense: Portfolio) {
  const id = expense.id;
  const url = buildUrl(id);
  const method = id && id > 0 ? "PUT" : "POST";
  const json = JSON.stringify(expense);
  const res = await ApiClient.request<{ token: string }>(url, {
    method: method,
    body: json,
  });
  if (!res.success) {
    const errorBody = res.data;
    console.error({ status: res.statusCode, msg: errorBody });
    return;
  }
  return res.success;
}

export async function getAssetClasses() {
  //const url = "investments/assets";
  // const res = await ApiClient.request<AssetClass[]>(url, {
  //   method: "GET",
  // });
  // if (!res.success) {
  //   const errorBody = res.data;
  //   console.error({ status: res.statusCode, msg: errorBody });
  //   return;
  // }
  return [
    { id: 1, name: "Títulos privados" },
    { id: 2, name: "Títulos públicos" },
    { id: 3, name: "Ações" },
    { id: 4, name: "ETFs" },
    { id: 5, name: "FIIs" },
    { id: 6, name: "Fundos" },
    { id: 7, name: "Commodities" },
    { id: 8, name: "Derivativos" },
    { id: 9, name: "Criptomoeda" },
    { id: 10, name: "Exterior" },
    { id: 11, name: "Poupança" },
    { id: 12, name: "Outros" },
  ];
}

export async function deleteInvestment(id: number) {
  const res = await ApiClient.request(buildUrl(id), {
    method: "DELETE",
  });
  if (!res.success) {
    const errorBody = res.data;
    console.error({ status: res.statusCode, msg: errorBody });
    return;
  }
  return res.success;
}
export async function updateIsDone(id: number, isDone: boolean) {
  console.log(isDone)
  const url = `${buildUrl(id)}/update-status`;
  const json = JSON.stringify({ isDone: isDone });
  const res = await ApiClient.request(url, {
    method: "PATCH",
    body: json,
  });
  if (!res.success) {
    const errorBody = res.data;
    console.error({ status: res.statusCode, msg: errorBody });
    return;
  }
  return res.success;
}
