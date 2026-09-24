"use server";

import { ApiClient, ApiResponse } from "@/lib/httpClient";
import { cookies } from "next/headers";
import { IRegister } from "../model/register";

export async function DoLogin(user, pass) {
  try {
    const obj = {
      username: user,
      password: pass,
    };

    const res = await ApiClient.request<{ token: string }>("auth/login", {
      method: "POST",
      body: JSON.stringify(obj),
    });

    return { success: await SetCookies(res) };
  } catch (error) {
    return { success: false, error: error.message };
  }
}

export async function DoSendForgotPassword(user: string) {
  const res = await ApiClient.request<{ token: string }>(
    "auth/fogort-passord",
    {
      method: "POST",
      body: JSON.stringify({ email: user }),
    },
  );

  if (!res.success) {
    const errorBody = res.data;
    console.error({ status: res.statusCode, msg: errorBody });
    return;
  }
  return res.success;
}
export async function DoRegister(registerObj: IRegister) {
  try {
    const res = await ApiClient.request<{ token: string }>("auth/register", {
      method: "POST",
      body: JSON.stringify(registerObj),
    });

    return { success: await SetCookies(res) };
  } catch (error) {
    return { success: false, error: error.message };
  }
}
export async function DoResetPass(pass: string, code: string) {
  const res = await ApiClient.request<{ token: string }>(
    "auth/reset-password/" + code,
    {
      method: "POST",
      body: JSON.stringify({ new_password: pass }),
    },
  );

  if (!res.success) {
    const errorBody = res.data;
    console.error({ status: res.statusCode, msg: errorBody });
    return;
  }
  return res.success;
}

async function SetCookies(
  res: ApiResponse<{
    token: string;
  }>,
) {
  const setCookieHeader = res.headers?.get("set-cookie");
  if (setCookieHeader) {
    const [nameValuePair] = setCookieHeader.split(";");
    const [_, cookieValue] = nameValuePair.split("=");

    const cookieStore = await cookies();
    cookieStore.set("session", cookieValue, {
      httpOnly: true,
      secure: process.env.NODE_ENV === "production",
      sameSite: "lax",
      path: "/",
      maxAge: 3600,
    });
    return true;
  } else {
    console.warn("A API não enviou o cabeçalho Set-Cookie");
    return false;
  }
}
