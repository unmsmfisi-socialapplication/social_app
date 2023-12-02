import { emailRegex, nameRegex, passwordRegex, phoneRegex, usernameRegex } from './Constant'

export const validateName = (inputString: string): boolean => {
    return nameRegex.test(inputString)
}

export const validateEmail = (inputString: string): boolean => {
    return emailRegex.test(inputString)
}

export const validatePassword = (inputString: string): boolean => {
    return passwordRegex.test(inputString)
}

export const validateUsername = (inputString: string): boolean => {
    return usernameRegex.test(inputString)
}
export const validatePhone = (inputString: string): boolean => {
    return phoneRegex.test(inputString)
}
