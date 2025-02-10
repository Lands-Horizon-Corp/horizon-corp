export const abbreviateUUID = (uuid: string, abbrevLength = 7) => {
    const cleanedUUID = uuid.replace(/-/g, '')
    return cleanedUUID.substring(0, abbrevLength)
}
