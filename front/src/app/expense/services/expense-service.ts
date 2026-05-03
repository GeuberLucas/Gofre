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
  const res = await ApiClient.request(url, {
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

export async function deleteExpense(id: number) {
  console.log("Executado delete");
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

export async function updateIsPaid(id: number, isPaid: boolean) {
  const url = `${buildUrl(id)}/update-status`;
  const json = JSON.stringify({ isPaid: isPaid });
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
