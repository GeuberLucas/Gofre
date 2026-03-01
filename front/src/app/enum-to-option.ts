function formatLabel(label: string) {
  return label
    .toLowerCase()
    .replaceAll("_", " ")
    .replaceAll(/\b\w/g, (l) => l.toUpperCase());
}

export function enumToFormattedOptions<
  E extends Record<string, string | number>,
>(enumObj: E) {
  return Object.keys(enumObj)
    .filter((key) => Number.isNaN(Number(key)))
    .map((key) => ({
      texto: formatLabel(key),
      valor: enumObj[key as keyof E],
    }));
}
