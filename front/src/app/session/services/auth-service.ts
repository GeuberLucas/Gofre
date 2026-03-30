"use server";

import { ApiClient } from "@/lib/httpClient";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export async function DoLogin(user, pass) {
  const obj = {
    username: user,
    password: pass,
  };

  const res = await ApiClient.request<{ token: string }>("auth/login", {
    method: "POST",
    body: JSON.stringify(obj),
  });

  if (!res.success) {
    const errorBody = res.data;
    console.error({ status: res.statusCode, msg: errorBody });
    return;
  }
  const cookieStore = await cookies();
  cookieStore.set("session", res.data.token, {
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    sameSite: "lax",
    path: "/",
    maxAge: 3600,
  });
  redirect("/expense");
}
