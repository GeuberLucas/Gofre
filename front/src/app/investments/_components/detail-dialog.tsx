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
import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field";
import { Calendar } from "@/components/ui/calendar";
import { Switch } from "@/components/ui/switch";
import { useEffect, useState } from "react";
import { NumericFormat } from "react-number-format";
import { Portfolio } from "../model/portfolio";
import {
  getAssetClasses,
  getPortfolio,
  sendPortfolio,
} from "../services/investment-service";
import { AssetClass } from "../model/asset-class";
import { GofreSelect } from "@/components/gofre-select";
interface DetailProps {
  open: boolean;
  onClose: () => void;
  id: number;
}

const formSchema = z.object({
  description: z.string().optional(),
  broker: z.string().optional(),
  deposit_date: z.date(),
  amount: z.number().optional(),
  is_done: z.boolean().optional(),
  asset_id: z.coerce.number(),
});

export default function DetailInvestment(props: Readonly<DetailProps>) {
  const [assetClasses, setAssetClasses] = useState<AssetClass[]>([]);
  const action = props.id > 0 ? "Editar" : "Adicionar";
  const [month, setMonth] = useState<Date | undefined>(undefined);
  const form = useForm({
    resolver: zodResolver(formSchema),
    defaultValues: {
      description: "",
      broker: "",
      deposit_date: undefined,
      amount: 0,
      is_done: false,
      asset_id: 0,
    },
  });

  useEffect(() => {
    getAssetClasses().then((assetsResult) => {
      setAssetClasses(assetsResult);
    });
    if (props.id > 0) {
      getPortfolio(props.id).then((initialValues: Portfolio) => {
        form.reset({
          description: initialValues.description,
          broker: initialValues.broker,
          deposit_date: initialValues.deposit_date,
          amount: initialValues.amount,
          is_done: initialValues.is_done,
          asset_id: initialValues.asset_id,
        });
        setMonth(new Date(initialValues.deposit_date));
      });
    }
  }, [props.id, props.open, form]);

  function onSubmit(data: z.infer<typeof formSchema>) {
    //  if (Object.keys(form.formState.dirtyFields).length === 0) {
    //   props.onClose();
    //   return;
    // }
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
    props.onClose();
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
                    <FieldError>{fieldState.error?.message}</FieldError>
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
                    <FieldError>{fieldState.error?.message}</FieldError>
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
                      selected={field.value}
                      onSelect={field.onChange}
                      month={month}
                      onMonthChange={setMonth}
                      className="rounded-lg border"
                      captionLayout="dropdown"
                    />
                    <FieldError>{fieldState.error?.message}</FieldError>
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
                      <GofreSelect
                        options={assetClasses.map((item) => ({
                          label: item.name,
                          value: item.id,
                        }))}
                        field={field}
                        fieldState={fieldState}
                      />
                      <FieldError>{fieldState.error?.message}</FieldError>
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
                      <FieldError>{fieldState.error?.message}</FieldError>
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
                      <FieldError>{fieldState.error?.message}</FieldError>
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
