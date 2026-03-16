"use server";

import { Revenue } from "../model/revenue";

const defaultHeaders = {
  "Content-Type": "application/json",
  Authorization: `Bearer ${process.env.TOKEN}`,
};

const baseUrl = `${process.env.API_URL}transaction/revenue`;

function buildUrl(id?: number) {
  return id && id > 0 ? `${baseUrl}/${id}` : baseUrl;
}

export async function getRevenue(id?: number): Promise<Revenue[] | Revenue> {
  const url = buildUrl(id);
  const res = await fetch(url, { headers: defaultHeaders });
  if (!res.ok) {
    const errorBody = await res.text();
    console.error({ status: res.status, msg: errorBody });
    return;
  }
  const expenses: Revenue[] | Revenue = await res.json();
  return expenses;
}

export async function sendRevenue(expense: Revenue) {
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
