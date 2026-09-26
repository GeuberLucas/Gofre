"use client";

import { ThemeToggle } from "@/components/button-theme-toggle";
import {
  Stepper,
  StepperContent,
  StepperIndicator,
  StepperItem,
  StepperNav,
  StepperPanel,
  StepperSeparator,
  StepperTitle,
  StepperTrigger,
} from "@/components/reui/stepper";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import {
  CheckIcon,
  Cog,
  LoaderCircleIcon,
  Receipt,
  ShieldUser,
  UserStar,
} from "lucide-react";
import Image from "next/image";
import { useState } from "react";
import z from "zod";
import { IProfile } from "./model/profile";
import { toast } from "sonner";
import router from "next/router";
import { DoUpdateProfile } from "./services/onboarding-service";
import { Controller, useForm } from "react-hook-form";
import { Field, FieldError, FieldLabel } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { NumericFormat } from "react-number-format";
import Link from "next/link";

const steps = [
  {
    title: "Dados de Acesso",
    icon: <ShieldUser className="size-4" />,
  },
  {
    title: "Perfil do usuário",
    icon: <UserStar className="size-4" />,
  },
  {
    title: "Balanço inicial",
    icon: <Receipt className="size-4" />,
  },
  // {
  //   title: "Upload do extrato",
  //   icon: <Cog className="size-4" />,
  // },
];

const formSchema = z.object({
  completeName: z.string().optional(),
  cellphone: z.string().optional(),
  initialBalance: z.number().optional(),
});
async function onSubmit(data: z.infer<typeof formSchema>) {
  const obj: IProfile = {
    completeName: "",
    cellphone: "",
    initialBalance: 0,
  };

  const result = await DoUpdateProfile(obj);
  if (!result?.success) {
    toast.error(result?.error || "Erro ao tentar fazer login.");

    return;
  }
  toast.success("Registro efetuado com sucesso!", {});

  router.push("/");
}
export default function Onboarding() {
  const appVersion = process.env.APP_VERSION;
  const [currentStep, setCurrentStep] = useState(2);
  const setStep = (step: number) => {
    if (step == 1) {
      return;
    }

    setCurrentStep(step);
  };
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      completeName: undefined,
      cellphone: undefined,
      initialBalance: undefined,
    },
  });
  return (
    <section className="flex min-h-screen w-full flex-col items-center bg-site-bg dark:text-white">
      {/* Header no topo com mt-2 */}
      <header className="mt-2 flex w-full items-center justify-between p-4 md:justify-end">
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

          <div className="flex min-w-0 flex-col">
            <div className="flex items-center gap-2">
              <span className="truncate font-display text-sm font-bold leading-tight md:text-base">
                Gofre
              </span>
              <span className="rounded-md bg-muted px-1.5 py-0.5 font-number text-[10px] text-muted-foreground">
                v{appVersion}
              </span>
            </div>
            <span className="truncate font-display text-xs text-muted-foreground">
              Gestão Financeira <br /> Do caos ao patrimônio.
            </span>
          </div>
        </div>
        <ThemeToggle />
      </header>

      {/* Container principal centralizado */}
      <main className="flex w-full flex-1 items-center justify-center p-4">
        <Card className="w-full max-w-xl p-6">
          <CardContent className="p-0">
            <form>
              <Stepper
                value={currentStep}
                onValueChange={setStep}
                indicators={{
                  completed: <CheckIcon className="size-3.5" />,
                  loading: (
                    <LoaderCircleIcon className="size-3.5 animate-spin" />
                  ),
                }}
                className="w-full space-y-8"
              >
                <StepperNav className="gap-3">
                  {steps.map((step, index) => (
                    <StepperItem
                      key={index}
                      step={index + 1}
                      className="relative flex-1 items-start"
                    >
                      <StepperTrigger className="flex grow flex-col items-start justify-center gap-2.5">
                        <StepperIndicator
                          className="size-8 border-2 
                      data-[state=inactive]:border-border 
                      data-[state=inactive]:bg-transparent 
                      data-[state=inactive]:text-muted-foreground 
                      data-[state=completed]:bg-action-realized 
                      data-[state=active]:bg-action-primary 
                      data-[state=completed]:text-white data-[state=active]:text-white"
                        >
                          {step.icon}
                        </StepperIndicator>
                        <div className="flex flex-col items-start gap-1">
                          <StepperTitle
                            className="text-start text-base font-semibold 
                        group-data-[state=inactive]/step:text-muted-foreground"
                          >
                            {step.title}
                          </StepperTitle>
                          <div>
                            <Badge
                              className="hidden group-data-[state=active]/step:inline-flex 
                          group-data-[state=active]/step:bg-action-primary group-data-[state=active]/step:text-white"
                            >
                              Em Progresso
                            </Badge>
                            <Badge
                              className="hidden group-data-[state=completed]/step:inline-flex 
                          group-data-[state=completed]/step:bg-action-realized group-data-[state=completed]/step:text-white"
                            >
                              Completo
                            </Badge>
                            <Badge
                              className="hidden text-muted-foreground group-data-[state=inactive]/step:inline-flex 
                          group-data-[state=inactive]/step:bg-action-pending group-data-[state=inactive]/step:text-white"
                            >
                              Pendente
                            </Badge>
                          </div>
                        </div>
                      </StepperTrigger>
                      {steps.length > index + 1 && (
                        <StepperSeparator className="absolute inset-x-0 start-9 top-4 m-0 group-data-[state=completed]/step:bg-success group-data-[orientation=horizontal]/stepper-nav:w-[calc(100%-2rem)] group-data-[orientation=horizontal]/stepper-nav:flex-none" />
                      )}
                    </StepperItem>
                  ))}
                </StepperNav>
                <StepperPanel className="py-4 text-sm">
                  <StepperContent key="2" value={2} className="">
                    <header className="mb-4 text-center">
                      <h1 className="font-display text-3xl font-bold text-foreground">
                        Perfil do usuário
                      </h1>
                      <p className="mt-2 text-sm text-muted-foreground">
                        Como podemos te chamar e qual telefone usar nos avisos
                        importantes? Os dois são opcionais.
                      </p>
                    </header>
                    <section>
                      <div className="flex gap-3 my-2">
                        <Controller
                          name="cellphone"
                          control={form.control}
                          render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                              <FieldLabel htmlFor={field.name}>
                                Nome Completo
                              </FieldLabel>
                              <Input
                                className="border-input-border border"
                                placeholder="Nome Completo"
                                {...field}
                                id={field.name}
                              />
                              <FieldError>
                                {fieldState.error?.message}
                              </FieldError>
                            </Field>
                          )}
                        />
                      </div>
                      <div className="flex gap-3 my-2">
                        <Controller
                          name="completeName"
                          control={form.control}
                          render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                              <FieldLabel htmlFor={field.name}>
                                Celular
                              </FieldLabel>
                              <Input
                                className="border-input-border border"
                                placeholder="Celular"
                                {...field}
                                id={field.name}
                              />
                              <FieldError>
                                {fieldState.error?.message}
                              </FieldError>
                            </Field>
                          )}
                        />
                      </div>
                    </section>
                  </StepperContent>

                  <StepperContent key="3" value={3}>
                    <header className="mb-4 text-center">
                      <h1 className="font-display text-3xl font-bold text-foreground">
                        Balanço inicial
                      </h1>
                      <p className="mt-2 text-sm text-muted-foreground">
                        Quanto você tem hoje? Some contas, carteira e
                        investimentos. <br />
                        Dá para ajustar quando quiser.
                      </p>
                    </header>
                    <section>
                      <div className="flex gap-3">
                        <Controller
                          name="initialBalance"
                          control={form.control}
                          render={({ field, fieldState }) => (
                            <Field
                              data-invalid={fieldState.invalid}
                              className="p-2"
                            >
                              <FieldLabel htmlFor={field.name}>
                                Saldo inicial
                              </FieldLabel>
                              <NumericFormat
                                value={field.value}
                                thousandSeparator="."
                                decimalSeparator=","
                                decimalScale={2}
                                fixedDecimalScale
                                prefix="R$ "
                                className="w-full"
                                customInput={Input}
                                placeholder="Saldo inicial"
                                onValueChange={(values) => {
                                  field.onChange(values.floatValue);
                                }}
                              />
                              <FieldError>
                                {fieldState.error?.message}
                              </FieldError>
                            </Field>
                          )}
                        />
                      </div>
                    </section>
                  </StepperContent>
                </StepperPanel>
              </Stepper>
            </form>
          </CardContent>
          <CardFooter className="grid grid-cols-2 grid-rows-2 gap-4">
            <div className="flex gap-3">
              <Button
                type="button"
                className="h-12 flex-1 rounded-full"
                onClick={() => setCurrentStep((prev) => prev - 1)}
                disabled={currentStep === 1}
              >
                Voltar
              </Button>
            </div>
            <div className="flex gap-3">
              <Button
                type="button"
                className="h-12 flex-1 rounded-full"
                onClick={() => setCurrentStep((prev) => prev + 1)}
                hidden={currentStep === steps.length}
              >
                Próximo
              </Button>
              <Button
                hidden={currentStep != steps.length}
                type="submit"
                className="h-12 flex-1 rounded-full"
              >
                Continuar
              </Button>
            </div>
            <div className="col-span-2 text-center">
              <Link
                className="text-action-primary underline decoration-action-primary"
                href="/"
              >
                Preencher depois
              </Link>
            </div>
          </CardFooter>
        </Card>
      </main>
    </section>
  );
}
