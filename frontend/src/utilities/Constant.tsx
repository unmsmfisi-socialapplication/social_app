export const nameRegex = /^[a-zA-ZÀ-ÿ\s'-]{1,150}$/
export const emailRegex = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/
export const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/
export const usernameRegex = /^(?=.+[a-zA-Z])(?=.*\d)[a-zA-Z\d_-]{6,40}$/

export const apiSattus = {
    SUCCES: 'SUCCESS',
    LOADING: 'LOADING',
    FAILED: 'FAILED',
    IDLE: 'IDLE',
}
