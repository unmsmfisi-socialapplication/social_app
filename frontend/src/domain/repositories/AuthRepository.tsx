import { createError } from '@/data/entities/Error'
import { doPost } from '../../data/api/apiService'

export default class AuthRepository {
    static async authRequest(userData: any) {
        const { ...props } = userData
        let data = {} as any
        let error = null

        try {
            console.log('request', props)
            const dataResponse = await doPost(`${process.env.NEXT_PUBLIC_API_URL_LOGIN}`, props)

            if (dataResponse && dataResponse?.status === 200) {
                data = dataResponse?.data
                return { data, error }
            } else {
                error = new Error('Respuesta inesperada del servidor')
            }
        } catch (errorRequest: any) {
            let { response } = errorRequest
            error = createError({
                message: response?.data.response,
                code: response?.status,
            })
            //TODO: STANDARD ERROR LOG
        }
        return { data, error }
    }
    static async registerRequest(userData: any) {
        const { ...props } = userData

        let data = {} as any
        let error = null

        try {
            console.log('request', props)
            const dataResponse = await doPost(`${process.env.NEXT_PUBLIC_API_URL_REGISTER}`)

            if (dataResponse && dataResponse?.status === 200) {
                data = dataResponse?.data
                return { data, error }
            } else {
                error = new Error('Respuesta inesperada del servidor')
            }
        } catch (errorRequest: any) {
            let { response } = errorRequest
            error = createError({
                message: response?.data.response,
                code: response?.status,
            })
            //TODO: STANDARD ERROR LOG
        }

        return { data, error }
    }
}
