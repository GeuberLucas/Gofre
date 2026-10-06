"use client";

import { ApiResult } from "../DTO/ApiResult";
import { IProfile } from "../DTO/profile";
import { ApiClient } from "../httpClient";

export async function GetProfile(): Promise<ApiResult> {
  try {
    const res = await ApiClient.request<IProfile>("profile", {
      method: "GET",
    });

    localStorage.setItem("saldo_incial", res.data.initialBalance.toString());
    return { success: res.success };
  } catch (error) {
    return { success: false, error: error.message };
  }
}
