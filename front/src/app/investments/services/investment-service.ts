"use server";

import { ApiClient } from "@/lib/httpClient";
import { Portfolio } from "../model/portfolio";

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
