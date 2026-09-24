"use client";

import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";

import { Button } from "@/components/ui/button";
import { Controller, useForm } from "react-hook-form";
import z from "zod";
import { DoSendForgotPassword } from "../services/auth-service";
import Link from "next/link";

const formSchema = z.object({
  userName: z.string(),
});

export default function ForgotPassword() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      userName: "",
    },
  });

  async function onSubmit(data: z.infer<typeof formSchema>) {
    DoSendForgotPassword(data.userName);
  }
  return (
    <div className="mx-auto flex w-full max-w-md flex-1 flex-col justify-center py-6">
      <header className="mb-3 text-center">
        <p className="mb-2 text-[11px] font-bold uppercase text-investment-accent">
          Recuperação
        </p>
        <h1 className="font-display text-3xl font-bold text-foreground">
          Recupere seu acesso
        </h1>
        <p className="mt-2 text-sm text-muted-foreground">
          Informe seu e-mail ou usuário para receber as instruções.
        </p>
      </header>
      <main>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FieldGroup className="mt-2">
            <div className="flex gap-3">
              <Controller
                name="userName"
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
          </FieldGroup>
        </form>
        <div className="flex gap-3 my-4">
          <Button
            type="button"
            onClick={form.handleSubmit(onSubmit)}
            className="h-12 w-full rounded-full"
          >
            Enviar
          </Button>
        </div>
        <div className="w-full text-center mt-2">
          <Link
            href="/session?view=Login"
            className="w-full my-2 text-action-primary underline decoration-action-primary underline-offset-4 hover:opacity-80 transition-colors"
          >
            Voltar para o Login
          </Link>
        </div>
      </main>
    </div>
  );
}
