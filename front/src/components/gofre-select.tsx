import {
  ControllerFieldState,
  ControllerRenderProps,
  FieldPath,
  FieldValues,
} from "react-hook-form";
import {
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
  Select,
} from "@/components/ui/select";
import { enumToFormattedOptions } from "@/lib/enum-to-option";

export interface SelectOption {
  label: string;
  value: string | number;
}
interface GofreSelectProps<
  TFieldValues extends FieldValues,
  TName extends FieldPath<TFieldValues>,
  E extends Record<string, string | number> = undefined,
> {
  enumObj?: E;
  options?: SelectOption[];
  field: ControllerRenderProps<TFieldValues, TName>;
  fieldState: ControllerFieldState;
}

/**
 * Componente de seleção customizado para o ecossistema Gofre.
 * * Este componente integra-se ao `react-hook-form` e garante que o valor
 * retornado para o formulário seja sempre do tipo **number**.
 * * @example
 * <GofreSelect enumObj={MeusStatus} field={field} fieldState={fieldState} />
 * * @returns Um componente de Select que emite valores numéricos via `field.onChange`.
 */
export function GofreSelect<
  TFieldValues extends FieldValues,
  TName extends FieldPath<TFieldValues>,
  E extends Record<string, string | number> = undefined,
>({
  enumObj,
  options,
  field,
  fieldState,
}: Readonly<GofreSelectProps<TFieldValues, TName, E>>) {
  let items: SelectOption[] = [];

  if (options) {
    items = options;
  } else if (enumObj) {
    items = enumToFormattedOptions(enumObj).map((item) => ({
      label: item.texto,
      value: item.valor,
    }));
  }
  return (
    <Select
      name={field.name}
      value={
        field.value !== undefined && field.value !== null
          ? field.value.toString()
          : ""
      }
      onValueChange={(value) => field.onChange(Number(value))}
    >
      <SelectTrigger className="w-full" aria-invalid={fieldState.invalid}>
        <SelectValue placeholder="Selecione" />
      </SelectTrigger>
      <SelectContent>
        {items.map((type) => {
          return (
            <SelectItem value={type.value.toString()} key={type.value}>
              {type.label}
            </SelectItem>
          );
        })}
      </SelectContent>
    </Select>
  );
}
