"use client";
import { useEffect, useState } from "react";
import { DoLogin } from "../services/auth-service";
import {
  FieldGroup,
  Field,
  FieldLabel,
  FieldError,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { Controller, useForm } from "react-hook-form";
import z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  InputGroup,
  InputGroupAddon,
  InputGroupButton,
  InputGroupInput,
} from "@/components/ui/input-group";
import { EyeOffIcon } from "lucide-react";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

const formSchema = z.object({
  email: z.string(),
  pass: z.string(),
});

export default function Login() {
  const router = useRouter();
  const [showPassword, setShowPassword] = useState(false);
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      pass: "",
    },
  });

  async function onSubmit(data: z.infer<typeof formSchema>) {
    const result = await DoLogin(data.email, data.pass);
    if (!result?.success) {
      toast.error(result?.error || "Erro ao tentar fazer login.");
      return;
    }
    toast.success("Login efetuado com sucesso!", {});

    router.push("/");
  }
  return (
    <div className="mx-auto flex w-full max-w-md flex-1 flex-col justify-center py-8">
      <header className="mb-6 text-center">
        <p className="mb-2 text-[11px] font-bold uppercase text-investment-accent">
          Boas Vindas
        </p>
        <h1 className="font-display text-3xl font-bold text-foreground">
          Que bom te ver de novo!
        </h1>
        <p className="mt-2 text-sm text-muted-foreground">
          Informe seus dados para entrar.
        </p>
      </header>
      <main>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FieldGroup className="mt-2">
            <div className="flex gap-3">
              <Controller
                name="email"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor={field.name}>Usuario</FieldLabel>
                    <Input
                      className="border-input-border border"
                      placeholder="usuario"
                      {...field}
                      id={field.name}
                    />
                    <FieldError>{fieldState.error?.message}</FieldError>
                  </Field>
                )}
              />
            </div>
            <div className="flex gap-3">
              <Controller
                name="pass"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor={field.name}>Senha</FieldLabel>
                    <InputGroup className="border-input-border border">
                      <InputGroupInput
                        placeholder="senha"
                        type={showPassword ? "text" : "password"}
                        {...field}
                        id={field.name}
                      />
                      <InputGroupAddon align="inline-end">
                        <InputGroupButton
                          onClick={() => setShowPassword(!showPassword)}
                          size="icon-xs"
                        >
                          <EyeOffIcon />
                        </InputGroupButton>
                      </InputGroupAddon>
                    </InputGroup>

                    <FieldError>{fieldState.error?.message}</FieldError>
                  </Field>
                )}
              />
            </div>
          </FieldGroup>
        </form>
        <div className="w-full flex gap-3 mt-2 text-right">
          <hr />
          <Link
            href="/session?view=forgotpassword"
            className="w-full my-2 text-action-primary underline decoration-action-primary underline-offset-4 hover:opacity-80 transition-colors"
          >
            Esqueceu a Senha?
          </Link>
        </div>
        <div className="flex gap-3 mt-2">
          <Button
            type="button"
            onClick={form.handleSubmit(onSubmit)}
            className="h-12 w-full rounded-full"
          >
            Entrar
          </Button>
        </div>
        <div className="w-full text-center mt-2">
          <p>Você ainda não possui uma Conta?</p>
          <Link
            href="/session?view=register"
            className="w-full my-2 text-action-primary underline decoration-action-primary underline-offset-4 hover:opacity-80 transition-colors"
          >
            Registre-se Agora Mesmo
          </Link>
        </div>
      </main>
    </div>
  );
}
