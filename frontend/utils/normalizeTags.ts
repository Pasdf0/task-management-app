export const normalizeTags = (value: string): string[] => {
  const cleanedTags = value
    .split(',')
    .map((item) => item.trim())
    .filter((item) => item.length > 0)

  return [...new Set(cleanedTags)]
}
