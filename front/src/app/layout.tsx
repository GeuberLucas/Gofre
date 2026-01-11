import { Navbar } from "@/components/Navbar";
import { Geist, Poppins, Roboto, Roboto_Mono } from "next/font/google";
import { HeroUIProvider } from "@heroui/react";
import "./globals.css";

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
    <html lang="en">
      <body
        className={`${robotoFont.variable} ${poppinsFont.variable} antialiased `}
      >
        <main className="z-1 min-h-screen min-w-screen bg-black text-white pt-10">
          {children}
        </main>
        <Navbar />
      </body>
    </html>
  );
}
