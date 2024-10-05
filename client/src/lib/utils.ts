import { clsx, type ClassValue } from 'clsx'
import { twMerge } from 'tailwind-merge'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}


type SimpleTypes = boolean | number | string | null | undefined | symbol;

export default function getSimpleProperties(
  objectToParse: object,
  include?: boolean | object | Array<string> | boolean
): { [key: string]: SimpleTypes } {
  const simpleProperties: { [key: string]: SimpleTypes } = {};

  if (include === false) return simpleProperties;

  for (const [key, value] of Object.entries(objectToParse)) {
    const isSimpleType =
      value === null ||
      value === undefined ||
      typeof value === 'boolean' ||
      typeof value === 'number' ||
      typeof value === 'string' ||
      typeof value === 'symbol';

    if (isSimpleType || (Array.isArray(include) && include.includes(key))) {
      simpleProperties[key] = value;
    }
  }

  return simpleProperties;
}
