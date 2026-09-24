import { ThemeToggle } from "@/components/button-theme-toggle";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import Image from "next/image";

export default function Onboarding() {
  const appVersion = process.env.APP_VERSION;

  return (
    <section className="flex w-full flex-col items-center justify-center bg-site-bg dark:text-white lg:w-1/2">
      <header className="flex w-full items-center justify-between md:justify-end p-3">
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
      </header>
      <main>
        <Card>
          <CardContent>
            <p>Card Content</p>
          </CardContent>
          <CardFooter>
            <Button type="button" className="h-12 w-full rounded-full">
              Voltar
            </Button>
            <Button type="button" className="h-12 w-full rounded-full">
              Proximo
            </Button>
          </CardFooter>
        </Card>
      </main>
    </section>
  );
}
