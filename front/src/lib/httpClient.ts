import { redirect } from "next/navigation";
import { toast } from "sonner";
interface RequestOptions extends RequestInit {
  headers?: Record<string, string>;
}

export interface ApiResponse<T = unknown> {
  data: T;
  statusCode: number;
  timestamp: string;
  success: boolean;
  headers: Headers;
}
const baseUrl = process.env.API_URL;

function buildUrl(endpoint: string) {
  return new URL(endpoint, baseUrl).toString();
}
function UnauthorizedResponse() {
  redirect("/session");
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
  static async request<T = unknown | undefined>(
    endpoint: string,
    options: RequestOptions = {},
  ): Promise<ApiResponse<T>> {
    const url = buildUrl(endpoint);
    console.log(`Request for ${url}`);
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
      headers.Cookie = `jwt-token=${token}`;
    }
    const config = {
      ...options,
      headers,
    };

    try {
      const response = await fetch(url, config);

      if (response.status === 401) {
        UnauthorizedResponse();
        throw new Error("Não autorizado");
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
          headers: response.headers,
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
        headers: response.headers,
      };
    } catch (error) {
      let isConnectionRefused = false;

      if (error instanceof TypeError && error.cause instanceof AggregateError) {
        error.cause.errors.forEach((err) => {
          if (
            err &&
            typeof err === "object" &&
            "code" in err &&
            err.code === "ECONNREFUSED"
          ) {
            isConnectionRefused = true;
          }
        });
      } else if (error instanceof AggregateError) {
        error.errors.forEach((err) => {
          if (
            err &&
            typeof err === "object" &&
            "code" in err &&
            err.code === "ECONNREFUSED"
          )
            isConnectionRefused = true;
        });
      } else if (
        error &&
        typeof error === "object" &&
        "code" in error &&
        error.code === "ECONNREFUSED"
      ) {
        isConnectionRefused = true;
      }

      if (isConnectionRefused) {
        throw new Error(
          "Serviço temporariamente indisponível. Tente novamente mais tarde.",
        );
      }

      if (error instanceof Error) {
        throw error;
      }

      console.error("API request error:", error);
      throw new Error("Ocorreu um erro inesperado.");
    }
  }
}
async function ProcessData<T>(
  response: Response,
): Promise<T | { message?: string }> {
  if (!response.ok) {
    try {
      const errorData = await response.clone().json();
      const errorMessage =
        errorData.erro || errorData.message || "Erro desconhecido na API";
      toast(errorMessage);
      return {
        message: errorMessage,
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
