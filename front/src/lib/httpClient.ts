import { redirect } from "next/navigation";

interface RequestOptions extends RequestInit {
  headers?: Record<string, string>;
}

interface ApiResponse<T = unknown> {
  data: T;
  statusCode: number;
  timestamp: string;
  success: boolean;
}
const baseUrl = process.env.API_URL;

function buildUrl(endpoint: string) {
  return new URL(endpoint, baseUrl).toString();
}
function UnauthorizedResponse() {
  redirect("/session/login");
}
function getCookieToken(): string | null {
  if (typeof document === "undefined") return null;

  const cookies = document.cookie.split(";");

  for (const cookie of cookies) {
    const [name, value] = cookie.trim().split("=");
    if (name === "session") return value;
  }
  return null;
}

export class ApiClient {
  static async request<T = unknown>(
    endpoint: string,
    options: RequestOptions = {},
  ): Promise<ApiResponse<T>> {
    const url = buildUrl(endpoint);
    const headers: Record<string, string> = {
      ...options.headers,
    };
    if (["POST", "PUT", "PATCH"].includes(options.method || "")) {
      headers["Content-Type"] = "application/json";
    }

    let token: string | null = null;

    if (globalThis.window === undefined) {
      const { cookies } = await import("next/headers");
      const cookieStore = await cookies();
      token = cookieStore.get("session")?.value || null;
    } else {
      token = getCookieToken();
    }
    if (token) {
      headers.Authorization = `Bearer ${token}`;
    }
    const config = {
      ...options,
      headers,
    };

    try {
      const response = await fetch(url, config);
      if (response.status === 401) {
        UnauthorizedResponse();
        return;
      }
      if (
        response.status === 204 ||
        response.headers.get("content-length") === "0"
      ) {
        return {
          success: response.ok,
          data: { success: true } as unknown as T,
          statusCode: response.status,
          timestamp: new Date().toISOString(),
        };
      }

      const data = await ProcessData<T>(response);
      if (!response.ok) {
        throw new Error(
          (data as { message?: string }).message || "An error occurred",
        );
      }
      return {
        success: response.ok,
        data: data as T,
        statusCode: response.status,
        timestamp: new Date().toISOString(),
      };
    } catch (error) {
      console.error("API request error:", error);
      throw error;
    }
  }
}
async function ProcessData<T>(
  response: Response,
): Promise<T | { message?: string }> {
  if (!response.ok) {
    try {
      const errorData = await response.clone().json();

      return {
        message:
          errorData.erro || errorData.message || "Erro desconhecido na API",
      };
    } catch (parseError) {
      console.error(`Falha ao ler o erro como JSON: ${parseError}`);

      const text = await response.text();
      return { message: text || `Erro HTTP: ${response.status}` };
    }
  }

  try {
    const data = await response.clone().json();
    return data as T;
  } catch (parseError) {
    console.error(`Falha ao ler o erro como JSON: ${parseError}`);
    return { success: true } as unknown as T;
  }
}
