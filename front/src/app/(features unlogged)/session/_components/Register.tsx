import { Button } from "@/components/ui/button";
import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import {
  InputGroup,
  InputGroupAddon,
  InputGroupButton,
  InputGroupInput,
} from "@/components/ui/input-group";
import { zodResolver } from "@hookform/resolvers/zod";
import { EyeOffIcon } from "lucide-react";
import Link from "next/link";
import { useState } from "react";
import { Controller, useForm } from "react-hook-form";
import z from "zod";
import { DoRegister } from "../services/auth-service";
import { IRegister } from "../model/register";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

const formSchema = z
  .object({
    username: z.string("Informe um nome de usuário"),
    complete_name: z.string().optional(),
    cellphone: z.string().optional(),
    email: z.email("E-mail inválido"),
    password: z
      .string("A senha é obrigatória")
      .min(8, "A senha deve ter no mínimo 8 caracteres.")
      .regex(/[a-z]/, "A senha deve conter pelo menos uma letra minúscula.")
      .regex(/[A-Z]/, "A senha deve conter pelo menos uma letra maiúscula.")
      .regex(/[/d]/, "A senha deve conter pelo menos um número."),
    confirmPass: z.string("Por gentileza confirme a sua senha"),
  })
  .refine((data) => data.password === data.confirmPass, {
    error: "As senhas não são iguais",
    path: ["confirmPass"],
    when() {
      return true;
    },
  });
export default function Register() {
  const router = useRouter();
  const [showPassword, setShowPassword] = useState(false);
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: undefined,
      cellphone: undefined,
      email: undefined,
      password: undefined,
      confirmPass: undefined,
    },
  });

  async function onSubmit(data: z.infer<typeof formSchema>) {
    const obj: IRegister = {
      username: data.username,
      complete_name: data.complete_name,
      cellphone: data.cellphone,
      email: data.email,
      password: data.password,
    };

    const result = await DoRegister(obj);
    if (!result?.success) {
      toast.error(result?.error || "Erro ao tentar fazer login.");

      return;
    }
    toast.success("Login efetuado com sucesso!", {});

    router.push("/register");
  }
  return (
    <div className="mx-auto flex w-full max-w-md flex-1 flex-col justify-center py-6">
      <header className="mb-4 text-center">
        <p className="mb-2 text-[11px] font-bold uppercase text-investment-accent">
          Primeiro passo
        </p>
        <h1 className="font-display text-3xl font-bold text-foreground">
          Crie sua conta
        </h1>
        <p className="mt-2 text-sm text-muted-foreground">
          Comece hoje a colocar sua vida financeira em ordem.
        </p>
      </header>
      <main>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FieldGroup className="mt-2">
            <div className="flex gap-3">
              <Controller
                name="username"
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
                name="email"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor={field.name}>Email</FieldLabel>
                    <Input
                      className="border-input-border border"
                      placeholder="exemple@example.com"
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

          <div className="flex gap-3 my-4">
            <Button type="submit" className="h-12 w-full rounded-full">
              Registrar
            </Button>
          </div>
        </form>
        <div className="w-full text-center mt-2">
          <p>Você já possui uma Conta?</p>
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
