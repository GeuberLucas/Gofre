"use client";
import { Poppins, Roboto_Mono } from "next/font/google";
import "../globals.css";
import Image from "next/image";
import { ThemeProvider } from "@/components/theme-provider";
import { Toaster } from "sonner";

const poppinsFont = Poppins({
  subsets: ["latin"],
  weight: "400",
  variable: "--font-poppins",
});
const robotoFont = Roboto_Mono({
  subsets: ["latin"],
  weight: "400",
  variable: "--font-roboto",
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body
        className={`${robotoFont.variable} ${poppinsFont.variable} antialiased `}
        suppressHydrationWarning
      >
        <ThemeProvider>
          <main className="flex min-h-screen w-full dark:text-white">
            <div className="flex min-h-screen w-full">
              {/* Lado Esquerdo - Fixo (Branding Gofre) */}
              <div className="hidden w-1/2 flex-col items-center justify-center bg-[#0a0a0a] text-white lg:flex">
                {/* Adicione sua imagem do mascote e os textos "Do caos ao patrimônio" aqui */}
                <Image
                  src="/gofre-icon.png"
                  width={400}
                  height={400}
                  alt="Gofre Logo"
                  className="object-contain"
                />
                <div className="mt-10 text-center">
                  <p className="font-display text-2xl font-bold text-panel-foreground">
                    Do caos ao patrimônio.
                  </p>
                  <p className="mt-3 max-w-xs text-sm leading-6 text-panel-muted">
                    Entenda o que passou, organize o agora e planeje o que vem
                    pela frente.
                  </p>
                </div>
                <div className="absolute bottom-8 grid w-full max-w-md grid-cols-3 border-t border-panel-foreground/15 px-8 pt-5">
                  <div className="text-center">
                    <p className="text-[10px] font-bold uppercase text-panel-muted">
                      Passado
                    </p>
                    <p className="mt-1 text-sm font-semibold text-panel-foreground">
                      Movimentações
                    </p>
                  </div>
                  <div className="text-center">
                    <p className="text-[10px] font-bold uppercase text-panel-muted">
                      Presente
                    </p>
                    <p className="mt-1 text-sm font-semibold text-panel-foreground">
                      Organização
                    </p>
                  </div>
                  <div className="text-center">
                    <p className="text-[10px] font-bold uppercase text-panel-muted">
                      Futuro
                    </p>
                    <p className="mt-1 text-sm font-semibold text-panel-foreground">
                      Patrimônio
                    </p>
                  </div>
                </div>
              </div>

              {/* Lado Direito - Dinâmico (Área do Formulário) */}

              {children}
            </div>
          </main>
          <Toaster />
        </ThemeProvider>
      </body>
    </html>
  );
}
