
export const randomChoose = <T,> (data: Array<T>) => {
    return data[~~(Math.random() * data.length)]
}
