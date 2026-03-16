"use server";

import { redirect } from "next/navigation";

const baseUrl = `${process.env.API_URL}auth`;
const defaultHeaders = {
  "Content-Type": "application/json",
};
function buildUrl(controller: string) {
  return `${baseUrl}/${controller}`;
}
export async function DoLogin(user, pass) {
  const url = buildUrl("login");
  const obj = {
    username: user,
    password: pass,
  };

  const res = await fetch(url, {
    headers: defaultHeaders,
    method: "POST",
    body: JSON.stringify(obj),
  });
  if (!res.ok) {
    const errorBody = await res.json();
    console.error({ status: res.status, msg: errorBody });
    return;
  }
  const result = await res.json();
  process.env.TOKEN = result.token;
  redirect("/expense");
}
