"use server";

import { Expense } from "../model/expense";

const defaultHeaders = {
  "Content-Type": "application/json",
  Authorization: `Bearer ${process.env.TOKEN}`,
};

const baseUrl = `${process.env.API_URL}transaction/expense`;

function buildUrl(id?: number) {
  return id && id > 0 ? `${baseUrl}/${id}` : baseUrl;
}

export async function getExpense(id?: number): Promise<Expense[] | Expense> {
  const url = buildUrl(id);
  const res = await fetch(url, { headers: defaultHeaders });
  if (!res.ok) {
    const errorBody = await res.text();
    console.error({ status: res.status, msg: errorBody });
    return;
  }
  const expenses: Expense[] | Expense = await res.json();
  return expenses;
}

export async function sendExpense(expense: Expense) {
  const id = expense.id;
  const url = buildUrl(id);
  const method = id && id > 0 ? "POST" : "PUT";
  const json = JSON.stringify(expense);
  const res = await fetch(url, {
    headers: defaultHeaders,
    method: method,
    body: json,
  });
  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }

  return res;
}
