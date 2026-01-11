import { Navbar } from "@/components/Navbar";
import { Poppins, Roboto_Mono } from "next/font/google";
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
        className={`${robotoFont.variable} ${poppinsFont.variable} antialiased`}
      >
        <main className="z-1 h-screen w-screen max-w-screen dark:text-white bg-site-bg">
          {children}
        </main>
        {/* <Navbar /> */}
      </body>
    </html>
  );
}
