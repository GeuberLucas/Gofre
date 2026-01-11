import Link from "next/link";

export function Navbar() {
  return (
    <div className="z-2 w-full flex justify-center fixed bottom-4">
      <nav className="flex backdrop-blur-2xl rounded-2xl w-fit bg-[#251351] text-white">
        <Link href={"/"} className="p-2 my-2">
          Home
        </Link>
        <Link href={"/"} className="p-2 my-2">
          Entradas
        </Link>
        <Link href={"/"} className="p-2 my-2">
          Saidas
        </Link>
        <Link href={"/"} className="p-2 my-2">
          Aporte
        </Link>
        <Link href={"/"} className="p-2 my-2">
          Mensal
        </Link>
        <Link href={"/"} className="p-2 my-2">
          Previsao
        </Link>
        <Link href={"/"} className="p-2 my-2">
          Patrimonio
        </Link>
        <Link href={"/"} className="p-2 my-2">
          Simulador
        </Link>
      </nav>
    </div>
  );
}
