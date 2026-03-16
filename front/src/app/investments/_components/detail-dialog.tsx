"use client";

import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import {
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectGroup,
  SelectItem,
  Select,
} from "@/components/ui/select";
import { Controller, useForm } from "react-hook-form";
import { Field, FieldGroup, FieldLabel } from "@/components/ui/field";
import { Calendar } from "@/components/ui/calendar";
import { Switch } from "@/components/ui/switch";
import { useEffect, useState } from "react";
import { NumericFormat } from "react-number-format";
import { Portfolio } from "../model/portfolio";
import { getPortfolio, sendPortfolio } from "../services/investment-service";
interface DetailProps {
  open: boolean;
  onClose: () => void;
  id: number;
}

const formSchema = z.object({
  description: z.string().optional(),
  broker: z.string().optional(),
  deposit_date: z.date().optional(),
  amount: z.number().optional(),
  is_done: z.boolean().optional(),
  asset_id: z.string().optional(),
});

//TODO:IMPLEMENTS SERVER ACTION GET ASSETS TYPE

export default function DetailInvestment(props: Readonly<DetailProps>) {
  const action = props.id > 0 ? "Editar" : "Adicionar";
  const [month, setMonth] = useState<Date | undefined>(undefined);
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      description: "",
      broker: "",
      deposit_date: new Date(),
      amount: 0,
      is_done: false,
      asset_id: "0",
    },
  });

  useEffect(() => {
    if (props.id > 0) {
      getPortfolio(props.id).then((initialValues: Portfolio) => {
        form.reset(
          {
            description: initialValues.description,
            broker: initialValues.broker,
            deposit_date: initialValues.deposit_date,
            amount: initialValues.amount,
            is_done: initialValues.is_done,
            asset_id: initialValues.asset_id.toString(),
          },
          { keepDirty: true, keepDirtyValues: true },
        );
        setMonth(new Date(initialValues.deposit_date));
      });
    }
  }, [props.id, form]);

  function onSubmit(data: z.infer<typeof formSchema>) {
    if (!form.formState.isDirty) {
      props.onClose();
      return;
    }
    const investment: Portfolio = {
      id: props.id,
      user_id: 0,
      description: data.description,
      broker: data.broker,
      deposit_date: data.deposit_date,
      amount: data.amount,
      is_done: data.is_done,
      asset_id: Number(data.asset_id),
    };

    sendPortfolio(investment).then();
  }
  return (
    <Dialog open={props.open} onOpenChange={props.onClose}>
      <DialogContent>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <DialogHeader>
            <DialogTitle>{action} Aporte de Investimento</DialogTitle>
          </DialogHeader>
          <FieldGroup className="p-3">
            <div className="flex gap-3">
              <Controller
                name="description"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor={field.name}>Decrição</FieldLabel>
                    <Input placeholder="decrição" {...field} id={field.name} />
                  </Field>
                )}
              />
              <Controller
                name="broker"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor={field.name}>
                      Corretora/Banco
                    </FieldLabel>
                    <Input
                      placeholder="Corretora/Banco"
                      {...field}
                      id={field.name}
                    />
                  </Field>
                )}
              />
            </div>
            <div className="flex gap-3">
              <Controller
                name="deposit_date"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor={field.name}>Data</FieldLabel>
                    <Calendar
                      mode="single"
                      selected={field.value || undefined}
                      onSelect={field.onChange}
                      month={month}
                      onMonthChange={setMonth}
                      className="rounded-lg border"
                      captionLayout="dropdown"
                    />
                  </Field>
                )}
              />
              <div className="flex flex-col gap-2">
                <Controller
                  name="asset_id"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid} className="p-2">
                      <FieldLabel htmlFor={field.name}>Ativo</FieldLabel>
                      <Select
                        name={field.name}
                        value={field.value}
                        onValueChange={field.onChange}
                      >
                        <SelectTrigger
                          className="w-full"
                          aria-invalid={fieldState.invalid}
                        >
                          <SelectValue placeholder="Selecione" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="0">Trabalho</SelectItem>
                          <SelectItem value="Extra">Extra</SelectItem>
                          <SelectItem value="Investimento">
                            Investimento
                          </SelectItem>
                          <SelectItem value="Aposentadoria">
                            Aposentadoria
                          </SelectItem>
                          <SelectItem value="Resgate">Resgate</SelectItem>
                          <SelectItem value="Outros">Outros</SelectItem>
                        </SelectContent>
                      </Select>
                    </Field>
                  )}
                />

                <Controller
                  name="amount"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid} className="p-2">
                      <FieldLabel htmlFor={field.name}>Valor Total</FieldLabel>
                      <NumericFormat
                        id={field.name}
                        value={field.value}
                        thousandSeparator="."
                        decimalSeparator=","
                        decimalScale={2}
                        fixedDecimalScale
                        prefix="R$ "
                        className="w-full"
                        customInput={Input}
                        placeholder="Valor Total"
                        onValueChange={(values) => {
                          field.onChange(values.floatValue);
                        }}
                      />
                    </Field>
                  )}
                />
                <Controller
                  name="is_done"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field
                      orientation="horizontal"
                      data-invalid={fieldState.invalid}
                      className="p-2"
                    >
                      <FieldLabel htmlFor={field.name}>Feito ?</FieldLabel>
                      <Switch
                        name={field.name}
                        checked={field.value}
                        onCheckedChange={field.onChange}
                        aria-invalid={fieldState.invalid}
                      />
                    </Field>
                  )}
                />
              </div>
            </div>
          </FieldGroup>
          <DialogFooter>
            <DialogClose asChild>
              <Button variant="outline" type="button">
                Cancelar
              </Button>
            </DialogClose>
            <Button type="submit">Salvar</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
