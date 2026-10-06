"use server";

import { ApiResult } from "@/lib/DTO/ApiResult";
import { IProfile } from "../../../../lib/DTO/profile";
import { ApiClient } from "@/lib/httpClient";

export async function DoUpdateProfile(profile: IProfile): Promise<ApiResult> {
  try {
    const res = await ApiClient.request<{ token: string }>("profile", {
      method: "PATCH",
      body: JSON.stringify(profile),
    });

    return { success: res.success };
  } catch (error) {
    return { success: false, error: error.message };
  }
}
