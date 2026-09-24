import {
  Field,
  FieldLabel,
  FieldError,
  FieldGroup,
} from "@/components/ui/field";
import {
  InputGroup,
  InputGroupInput,
  InputGroupAddon,
  InputGroupButton,
} from "@/components/ui/input-group";
import { zodResolver } from "@hookform/resolvers/zod";
import { EyeOffIcon } from "lucide-react";
import { useState } from "react";
import { Controller, useForm } from "react-hook-form";
import z from "zod";
import { DoResetPass } from "../services/auth-service";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { useSearchParams } from "next/navigation";

const formSchema = z
  .object({
    password: z.string(),
    confirmPass: z.string(),
  })
  .refine((data) => data.password === data.confirmPass, {
    error: "As senhas não são iguais",
    path: ["confirmPass"],
  });
export default function ResetPassword() {
  const params = useSearchParams();
  const resetCode = params.get("code");
  const [showPassword, setShowPassword] = useState(false);
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      password: "",
      confirmPass: "",
    },
  });

  async function onSubmit(data: z.infer<typeof formSchema>) {
    DoResetPass(data.password, resetCode);
  }
  return (
    <div className="mx-auto flex w-full max-w-md flex-1 flex-col justify-center py-6">
      <header className="mb-4 text-center">
        <p className="mb-2 text-[11px] font-bold uppercase text-investment-accent">
          Nova senha
        </p>
        <h1 className="font-display text-3xl font-bold text-foreground">
          Redefina sua senha
        </h1>
        <p className="mt-2 text-sm text-muted-foreground">
          Escolha uma senha forte e diferente das anteriores.
        </p>
      </header>
      <main>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FieldGroup>
            <div className="flex gap-3">
              <Controller
                name="password"
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
            <div className="flex gap-3">
              <Controller
                name="confirmPass"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor={field.name}>
                      Confirmar Senha
                    </FieldLabel>
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
        <div className="flex gap-3 my-4">
          <Button
            type="button"
            onClick={form.handleSubmit(onSubmit)}
            className="h-12 w-full rounded-full"
          >
            Salvar senha nova
          </Button>
        </div>
        <div className="w-full text-center mt-2">
          <Link
            href="/session?view=login"
            className="w-full my-2 text-action-primary underline decoration-action-primary underline-offset-4 hover:opacity-80 transition-colors"
          >
            Entre Agora Mesmo
          </Link>
        </div>
      </main>
    </div>
  );
}
