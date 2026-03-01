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
import { Expense } from "../model/expense";
import { CategoryExpenseEnum } from "../enums/category-expense-enum";
import { TypeExpenseEnum } from "../enums/type-expense-enum";
import { PaymentMethodEnum } from "../enums/payment-method-enum";
import { enumToFormattedOptions } from "../../enum-to-option";
interface DetailProps {
  open: boolean;
  onClose: () => void;
  id: number;
}

const formSchema = z.object({
  description: z.string().optional(),
  target: z.string().optional(),
  category: z.string().optional(),
  type: z.string().optional(),
  paymentMethod: z.string().optional(),
  paymentDate: z.date().optional(),
  amount: z.number().optional(),
  isPaid: z.boolean().optional(),
});
function getInicialValueForm(id): Expense | null {
  if (id == 0) {
    return null;
  }
  //TODO:IMPLEMENTS GET DATA FOR EDIT
}

export default function DetailExpense(props: Readonly<DetailProps>) {
  const action = props.id > 0 ? "Editar" : "Adicionar";
  const initialValues = getInicialValueForm(props.id);
  const categorys = enumToFormattedOptions(CategoryExpenseEnum);
  const typeExpense = enumToFormattedOptions(TypeExpenseEnum);
  const paymentMethods = enumToFormattedOptions(PaymentMethodEnum);
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      description: "",
      target: "",
      category: "",
      type: "",
      paymentMethod: "",
      paymentDate: new Date(),
      amount: 0,
      isPaid: false,
    },
  });
  if (initialValues) {
    form.setValue("description", initialValues.description);
    form.setValue("target", initialValues.target);
    form.setValue("category", initialValues.category);
    form.setValue("type", initialValues.type);
    form.setValue("paymentMethod", initialValues.target);
    form.setValue("paymentDate", initialValues.paymentDate);
    form.setValue("amount", initialValues.amount);
    form.setValue("isPaid", initialValues.isPaid);
  }

  function onSubmit(data: z.infer<typeof formSchema>) {
    console.log(data);
  }
  return (
    <Dialog open={props.open} onOpenChange={props.onClose}>
      <DialogContent>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <DialogHeader>
            <DialogTitle>{action} Despesa</DialogTitle>
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
                name="target"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor={field.name}>Destino</FieldLabel>
                    <Input placeholder="destino" {...field} id={field.name} />
                  </Field>
                )}
              />
            </div>
            <div className="flex gap-3">
              <Controller
                name="category"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid} className="p-2">
                    <FieldLabel htmlFor={field.name}>Categoria</FieldLabel>
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
                        {categorys.map((category) => {
                          return (
                            <SelectItem
                              value={category.valor.toString()}
                              key={category.valor}
                            >
                              {category.texto}
                            </SelectItem>
                          );
                        })}
                      </SelectContent>
                    </Select>
                  </Field>
                )}
              />
              <Controller
                name="type"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid} className="p-2">
                    <FieldLabel htmlFor={field.name}>Tipo de Saída</FieldLabel>
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
                        {typeExpense.map((type) => {
                          return (
                            <SelectItem
                              value={type.valor.toString()}
                              key={type.valor}
                            >
                              {type.texto}
                            </SelectItem>
                          );
                        })}
                      </SelectContent>
                    </Select>
                  </Field>
                )}
              />
              <Controller
                name="paymentMethod"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid} className="p-2">
                    <FieldLabel htmlFor={field.name}>Pagamento</FieldLabel>
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
                        {paymentMethods.map((paymentMethod) => {
                          return (
                            <SelectItem
                              value={paymentMethod.valor.toString()}
                              key={paymentMethod.valor}
                            >
                              {paymentMethod.texto}
                            </SelectItem>
                          );
                        })}
                      </SelectContent>
                    </Select>
                  </Field>
                )}
              />
            </div>
            <div className="flex gap-3">
              <Controller
                name="paymentDate"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor={field.name}>Data</FieldLabel>
                    <Calendar
                      mode="single"
                      selected={field.value}
                      onSelect={field.onChange}
                      className="rounded-lg border"
                      captionLayout="dropdown"
                    />
                  </Field>
                )}
              />
              <div className="flex flex-col gap-2">
                <Controller
                  name="amount"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={fieldState.invalid} className="p-2">
                      <FieldLabel htmlFor={field.name}>Valor Total</FieldLabel>
                      <Input
                        placeholder="Valor Total"
                        {...field}
                        id={field.name}
                      />
                    </Field>
                  )}
                />
                <Controller
                  name="isPaid"
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
