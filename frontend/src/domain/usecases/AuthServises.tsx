/* eslint-disable prettier/prettier */
import AuthRepository from '../repositories/AuthRepository';

export default class AuthServices {

    static async authRequest(request : any) {
        const { data , error } = await AuthRepository.authRequest(request);
        return { data , error }
    }
}