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
  BookUserIcon,
  CheckIcon,
  CreditCardIcon,
  LoaderCircleIcon,
  LockIcon,
} from "lucide-react";
import Image from "next/image";
import { useState } from "react";

const steps = [
  {
    title: "User Details",
    icon: <BookUserIcon className="size-4" />,
  },
  {
    title: "Payment Info",
    icon: <CreditCardIcon className="size-4" />,
  },
  {
    title: "Auth OTP",
    icon: <LockIcon className="size-4" />,
  },
];

export default function Onboarding() {
  const appVersion = process.env.APP_VERSION;
  const [currentStep, setCurrentStep] = useState(2);

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
            <Stepper
              value={currentStep}
              onValueChange={setCurrentStep}
              indicators={{
                completed: <CheckIcon className="size-3.5" />,
                loading: <LoaderCircleIcon className="size-3.5 animate-spin" />,
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
                      <StepperIndicator className="size-8 border-2 data-[state=inactive]:border-border data-[state=inactive]:bg-transparent data-[state=inactive]:text-muted-foreground data-[state=completed]:bg-success data-[state=completed]:text-white">
                        {step.icon}
                      </StepperIndicator>
                      <div className="flex flex-col items-start gap-1">
                        <div className="text-[10px] font-semibold uppercase text-muted-foreground">
                          Step {index + 1}
                        </div>
                        <StepperTitle className="text-start text-base font-semibold group-data-[state=inactive]/step:text-muted-foreground">
                          {step.title}
                        </StepperTitle>
                        <div>
                          <Badge className="hidden group-data-[state=active]/step:inline-flex">
                            In Progress
                          </Badge>
                          <Badge className="hidden group-data-[state=completed]/step:inline-flex">
                            Completed
                          </Badge>
                          <Badge className="hidden text-muted-foreground group-data-[state=inactive]/step:inline-flex">
                            Pending
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
                {steps.map((step, index) => (
                  <StepperContent
                    key={index}
                    value={index + 1}
                    className="flex items-center justify-center"
                  >
                    {step.title} content
                  </StepperContent>
                ))}
              </StepperPanel>
            </Stepper>
          </CardContent>
          <CardFooter className="flex w-full gap-3 p-0 pt-6">
            <Button
              type="button"
              className="h-12 flex-1 rounded-full"
              onClick={() => setCurrentStep((prev) => prev - 1)}
              disabled={currentStep === 1}
            >
              Voltar
            </Button>
            <Button
              type="button"
              className="h-12 flex-1 rounded-full"
              onClick={() => setCurrentStep((prev) => prev + 1)}
              disabled={currentStep === steps.length}
            >
              Próximo
            </Button>
          </CardFooter>
        </Card>
      </main>
    </section>
  );
}
