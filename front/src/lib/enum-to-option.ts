export function enumToFormattedOptions<
  E extends Record<string, string | number>,
>(enumObj: E) {
  return Object.keys(enumObj)
    .filter((key) => Number.isNaN(Number(key)))
    .map((key) => ({
      texto: key,
      valor: enumObj[key as keyof E],
    }));
}
