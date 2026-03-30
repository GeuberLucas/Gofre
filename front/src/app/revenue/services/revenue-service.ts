"use server";

import { ApiClient } from "@/lib/httpClient";
import { Revenue } from "../model/revenue";

const baseUrl = `transaction/revenue`;

function buildUrl(id?: number) {
  return id && id > 0 ? `${baseUrl}/${id}` : baseUrl;
}

export async function getRevenue(id?: number): Promise<Revenue[] | Revenue> {
  const res = await ApiClient.request<Revenue[] | Revenue>(buildUrl(id), {
    method: "GET",
  });
  if (!res.success) {
    const errorBody = res.data;
    console.error({ status: res.statusCode, msg: errorBody });
    return;
  }
  const expenses: Revenue[] | Revenue = res.data;
  return expenses;
}

export async function sendRevenue(expense: Revenue) {
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
