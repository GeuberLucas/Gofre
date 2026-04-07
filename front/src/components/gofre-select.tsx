import { ControllerFieldState, ControllerRenderProps } from "react-hook-form";
import {
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
  Select,
} from "@/components/ui/select";
import { enumToFormattedOptions } from "@/lib/enum-to-option";

interface GofreSelectProps<E extends Record<string, string | number>> {
  enumObj: E;
  field: ControllerRenderProps;
  fieldState: ControllerFieldState;
}

export function GofreSelect<E extends Record<string, string | number>>({
  enumObj,
  field,
  fieldState,
}: Readonly<GofreSelectProps<E>>) {
  const items = enumToFormattedOptions(enumObj);
  return (
    <Select
      name={field.name}
      value={field.value.toString()}
      onValueChange={(value) => field.onChange(Number(value))}
    >
      <SelectTrigger className="w-full" aria-invalid={fieldState.invalid}>
        <SelectValue placeholder="Selecione" />
      </SelectTrigger>
      <SelectContent>
        {items.map((type) => {
          return (
            <SelectItem value={type.valor.toString()} key={type.valor}>
              {type.texto}
            </SelectItem>
          );
        })}
      </SelectContent>
    </Select>
  );
}
