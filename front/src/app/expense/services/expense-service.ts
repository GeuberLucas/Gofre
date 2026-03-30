"use server";

import { ApiClient } from "@/lib/httpClient";
import { Expense } from "../model/expense";

const baseUrl = `transaction/expense`;

function buildUrl(id?: number) {
  return id && id > 0 ? `${baseUrl}/${id}` : baseUrl;
}

export async function getExpense(id?: number): Promise<Expense[] | Expense> {
  return ApiClient.request<Expense[] | Expense>(buildUrl(id), {
    method: "GET",
  }).then((res) => {
    if (res == undefined) {
      return;
    }

    const expenses: Expense[] | Expense = res.data;
    return expenses;
  });
}

export async function sendExpense(expense: Expense) {
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
