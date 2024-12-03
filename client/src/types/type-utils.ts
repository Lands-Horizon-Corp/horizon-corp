export type DeepKeys<T, Prefix extends string = ''> = T extends object
    ? {
          [K in keyof T]: K extends string
              ? T[K] extends object
                  ? `${Prefix}${K}` | DeepKeys<T[K], `${Prefix}${K}.`>
                  : `${Prefix}${K}`
              : never
      }[keyof T]
    : never

export type KeysOfType<T, ValueType> = {
    [K in keyof T]: T[K] extends ValueType ? K : never
}[keyof T]
