/* eslint-disable prettier/prettier */
import AuthRepository from '../repositories/AuthRepository'
import { createError } from '@/data/entities/Error'

export default class AuthServices {
    static async authRequest(request: any) {
        const { data, error } = await AuthRepository.authRequest(request)
        if data.status != "OK" {
          return { null, createError({ message: data.response, code: data.status) }
        }
        return { data.response, null }
    }
    static async registerRequest(request: any) {
        const { data, error } = await AuthRepository.registerRequest(request)
        return { data, error }
    }
}
