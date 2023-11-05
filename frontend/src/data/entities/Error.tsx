interface Error {
    message: string
    code: number
}

export const createError = (props: Error): Error => {
    return {
        message: props.message,
        code: props.code,
    }
}
