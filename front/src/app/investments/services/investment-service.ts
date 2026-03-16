"use server";

import { Portfolio } from "../model/portfolio";

const defaultHeaders = {
  "Content-Type": "application/json",
  Authorization: `Bearer ${process.env.TOKEN}`,
};

const baseUrl = `${process.env.API_URL}investments/`;

function buildUrl(id?: number) {
  return id && id > 0 ? `${baseUrl}/${id}` : baseUrl;
}

export async function getPortfolio(
  id?: number,
): Promise<Portfolio[] | Portfolio> {
  const url = buildUrl(id);
  const res = await fetch(url, { headers: defaultHeaders });
  if (!res.ok) {
    const errorBody = await res.text();
    console.error({ status: res.status, msg: errorBody });
    return;
  }
  const expenses: Portfolio[] | Portfolio = await res.json();
  return expenses;
}

export async function sendPortfolio(expense: Portfolio) {
  const id = expense.id;
  const url = buildUrl(id);
  const method = id && id > 0 ? "PUT" : "POST";
  const json = JSON.stringify(expense);
  console.log(json);
  const res = await fetch(url, {
    headers: defaultHeaders,
    method: method,
    body: json,
  });
  if (!res.ok) {
    const errorBody = await res.json();
    console.error({ status: res.status, msg: errorBody });
    return;
  }

  return res;
}
