package com.social.presentation.authentication

import androidx.compose.runtime.State
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.social.domain.model.RegisterBody
import com.social.utils.Resource
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableSharedFlow
import kotlinx.coroutines.flow.asSharedFlow
import kotlinx.coroutines.flow.launchIn
import kotlinx.coroutines.flow.onEach
import javax.inject.Inject

@HiltViewModel
class RegisterUserViewModel
    @Inject
    constructor(
        private val authenticationUseCase: AuthenticationUseCase,
    ) : ViewModel() {
        private val _state = MutableLiveData<RegisterUserState>()
        val state: LiveData<RegisterUserState> get() = _state
        private val eMailState = mutableStateOf(TextState())
        val eMail: State<TextState> = eMailState
        private val usernameState = mutableStateOf(TextState())
        val username: State<TextState> = usernameState
        private val passwordState = mutableStateOf(TextState())
        val psw: State<TextState> = passwordState
        private val eventFlowM = MutableSharedFlow<UIRegisterUserEvent>()
        val evenFlow = eventFlowM.asSharedFlow()

        fun insertData(event: RegisterEvent) {
            when (event) {
                is RegisterEvent.EnterEmail -> {
                    eMailState.value = eMail.value.copy(text = event.value)
                }
                is RegisterEvent.EnterUsername -> {
                    usernameState.value = username.value.copy(text = event.value)
                }
                is RegisterEvent.EnterPassword -> {
                    passwordState.value = psw.value.copy(text = event.value)
                }
                is RegisterEvent.RegisterUser -> {
                    authenticationUseCase.registerNewUser(
                        RegisterBody(
                            Phone = "",
                            Email = eMail.value.text,
                            User_name = username.value.text,
                            Password = psw.value.text,
                        ),
                    ).onEach { register ->
                        when (register) {
                            is Resource.Loading -> {
                                _state.value = RegisterUserState(isLoading = true)
                            }
                            is Resource.Error -> {
                                _state.value = RegisterUserState(error = register.message ?: "Error")
                                if (register.message!!.contains("HTTP 400")) {
                                    eventFlowM.emit(
                                        UIRegisterUserEvent.ShowMessage("Email already in use"),
                                    )
                                } else {
                                    eventFlowM.emit(
                                        UIRegisterUserEvent.ShowMessage(register.message),
                                    )
                                }
                            }
                            is Resource.Success -> {
                                _state.value = RegisterUserState(dataRegister = register.data!!)
                                eventFlowM.emit(UIRegisterUserEvent.GetData)
                                eventFlowM.emit(UIRegisterUserEvent.ShowMessage(register.message.toString()))
                            }
                        }
                    }.launchIn(viewModelScope)
                }
            }
        }

        sealed class UIRegisterUserEvent {
            object GetData : UIRegisterUserEvent()

            data class ShowMessage(val message: String) : UIRegisterUserEvent()
        }
    }
