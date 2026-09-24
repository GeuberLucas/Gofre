"use client";

import { useSearchParams } from "next/navigation";
import Login from "./_components/Login";
import Register from "./_components/Register";
import ResetPassword from "./_components/ResetPassword";
import ForgotPassword from "./_components/ForgotPassword";
import { UserRound } from "lucide-react";
import { ThemeToggle } from "@/components/button-theme-toggle";
import Image from "next/image";

export default function Session() {
  const params = useSearchParams();
  const appVersion = process.env.APP_VERSION;
  const view = params.get("view") || "login";
  const formView = () => {
    switch (view) {
      case "register":
        return <Register />;
      case "forgotpassword":
        return <ForgotPassword />;
      case "resetpassword":
        return <ResetPassword />;
      default:
        return <Login />;
    }
  };
  return (
    <section className="flex w-full flex-col items-center justify-center bg-site-bg dark:text-white lg:w-1/2">
      <div className="flex w-full items-center justify-between md:justify-end p-3">
        <div className="flex items-center gap-3 px-2 lg:hidden">
          <div className="flex aspect-square size-16 shrink-0 items-center justify-center rounded-lg bg-sidebar-primary/10 p-1">
            <Image
              src="/gofre-icon.png"
              width={500}
              height={500}
              alt="Gofre Logo"
              className="object-contain"
            />
          </div>

          <div className="flex flex-col min-w-0">
            <div className="flex items-center gap-2">
              <span className="font-display font-bold text-sm md:text-base truncate leading-tight">
                Gofre
              </span>
              <span className="rounded-md bg-muted px-1.5 py-0.5 font-number text-[10px] text-muted-foreground">
                v{appVersion}
              </span>
            </div>
            <span className="font-display truncate text-xs text-muted-foreground">
              Gestão Financeira <br /> Do caos ao patrimônio.
            </span>
          </div>
        </div>
        <ThemeToggle />
      </div>

      <div className="mx-auto flex w-full max-w-md flex-1 flex-col justify-center py-10">
        <header className="mb-4 text-center">
          <span className="mx-auto mb-5 flex size-12 items-center justify-center rounded-full bg-badge">
            <UserRound size={50} className="text-muted-foreground" />
          </span>
        </header>
        <div className="px-3 md:px-0">{formView()}</div>
      </div>
    </section>
  );
}
