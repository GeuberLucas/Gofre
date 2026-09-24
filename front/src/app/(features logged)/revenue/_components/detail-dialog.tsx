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
import { Revenue } from "../model/revenue";
import { Input } from "@/components/ui/input";
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
import { getRevenue, sendRevenue } from "../services/revenue-service";
import { IncomeType } from "../enums/income-type";
import { GofreSelect } from "@/components/gofre-select";
import { NumericFormat } from "react-number-format";
interface DetailProps {
  open: boolean;
  onClose: () => void;
  onSuccess: () => void;
  id: number;
}

const formSchema = z.object({
  description: z.string().optional(),
  origin: z.string("Informe de onde você recebeu"),
  type: z.enum(IncomeType, "Selecione um tipo de entrada válida"),
  receiveDate: z.date("A data de recebimento é obrigatória"),
  amount: z
    .number("O valor é obrigatório")
    .min(0.01, "O valor deve ser maior que zero"),
  recieved: z.boolean().optional(),
});

export default function DetailRevenue(props: Readonly<DetailProps>) {
  const action = props.id > 0 ? "Editar" : "Adicionar";
  const [month, setMonth] = useState<Date | undefined>(undefined);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      description: "",
      origin: "",
      type: IncomeType.Outros,
      receiveDate: undefined,
      amount: 0,
      recieved: false,
    },
  });
  useEffect(() => {
    if (props.id > 0) {
      getRevenue(props.id).then((data: Revenue) => {
        const parsedDate = data.receiveDate
          ? new Date(data.receiveDate)
          : undefined;
        form.reset({
          description: data.description,
          amount: data.amount,
          recieved: data.isRecieved,
          origin: data.origin,
          type: data.type,
          receiveDate: parsedDate,
        });

        setMonth(new Date(parsedDate));
      });
    }
  }, [props.id, form, props.open]);

  function onSubmit(data: z.infer<typeof formSchema>) {
    if (!form.formState.isDirty && props.id > 0) {
      props.onClose();
      return;
    }
    const expense: Revenue = {
      id: props.id,
      description: data.description,
      amount: data.amount,
      isRecieved: data.recieved,
      origin: data.origin,
      type: data.type,
      receiveDate: data.receiveDate,
    };

    sendRevenue(expense).then((sucess) => {
      if (sucess) {
        form.reset();
        setTimeout(() => props.onSuccess(), 100);
      }
    });
  }
  return (
    <Dialog open={props.open} onOpenChange={props.onClose}>
      <DialogContent>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <DialogHeader>
            <DialogTitle>{action} Entrada</DialogTitle>
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
                name="origin"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor={field.name}>origem</FieldLabel>
                    <Input placeholder="origem" {...field} id={field.name} />
                    <FieldError>{fieldState.error?.message}</FieldError>
                  </Field>
                )}
              />
            </div>
            <div className="flex gap-3">
              <Controller
                name="receiveDate"
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
                    <FieldError>{fieldState.error?.message}</FieldError>
                  </Field>
                )}
              />
              <div className="flex flex-col gap-2">
                <Controller
                  name="type"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid} className="p-2">
                      <FieldLabel htmlFor={field.name}>Tipo</FieldLabel>
                      <GofreSelect
                        enumObj={IncomeType}
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
                  name="recieved"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field
                      orientation="horizontal"
                      data-invalid={fieldState.invalid}
                      className="p-2"
                    >
                      <FieldLabel htmlFor={field.name}>Recebido ?</FieldLabel>
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
