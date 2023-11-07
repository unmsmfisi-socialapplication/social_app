
package com.social.domain.usecase

import com.social.data.source.remote.dto.aRegister
import com.social.data.source.remote.dto.aRegisterError
import com.social.domain.SocialAppRepository
import com.social.domain.model.InvalidRegisterException
import com.social.domain.model.RegisterBody
import com.social.domain.model.RegisterData
import com.social.utils.Resource
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import javax.inject.Inject

class RegisterNewUser
    @Inject
    constructor(
        private val socialAppRepository: SocialAppRepository,
    ) {
        operator fun invoke(registerBody: RegisterBody): Flow<Resource<List<RegisterData>>> {
            return flow {
                try {
                    if (registerBody.User_name.isBlank()) {
                        throw InvalidRegisterException("Ingrese usuario")
                    }
                    if (registerBody.Email.isBlank()) {
                        throw InvalidRegisterException("Ingrese correo")
                    }
                    if (registerBody.Password.isBlank()) {
                        throw InvalidRegisterException("Ingrese contraseña")
                    } else {
                        if (!registerBody.Password.matches("^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]).+\$".toRegex()) ||
                            registerBody.Password.length < 8
                        ) {
                            throw InvalidRegisterException("La contraseña debe ser segura")
                        }
                    }
                    emit(Resource.Loading())
                    val register = socialAppRepository.registerNewUser(registerBody).aRegister()
                    if (register.user_name.isNotEmpty()) {
                        emit(Resource.Success("Registrado", listOf(register)))
                    } else {
                        val registerError = socialAppRepository.registerNewUserError(registerBody).aRegisterError()
                        emit(Resource.Error(registerError.response))
                    }
                } catch (e: Exception) {
                    emit(Resource.Error(e.message.toString()))
                }
            }
        }
    }
