
export const toBase64 = (val: any) => {
    const stringified = JSON.stringify(val)
    const base64String = btoa(stringified)
    return base64String
}
