package com.social.presentation.authentication

import androidx.compose.runtime.State
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.social.domain.model.LoginBody
import com.social.utils.Resource
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableSharedFlow
import kotlinx.coroutines.flow.asSharedFlow
import kotlinx.coroutines.flow.launchIn
import kotlinx.coroutines.flow.onEach
import javax.inject.Inject

@HiltViewModel
class LoginViewModel
    @Inject
    constructor(
        private val authenticationUseCase: AuthenticationUseCase,
    ) : ViewModel() {
        private val _state = MutableLiveData<LoginDataState>()
        val state: LiveData<LoginDataState> get() = _state
        private val usernameState = mutableStateOf(TextState())
        val username: State<TextState> = usernameState
        private val passwordState = mutableStateOf(TextState())
        val password: State<TextState> = passwordState
        private val eventFlowM = MutableSharedFlow<UILoginEvent>()
        val evenFlow = eventFlowM.asSharedFlow()

        fun getData(event: LoginEvent) {
            when (event) {
                is LoginEvent.EnterUser -> {
                    usernameState.value = username.value.copy(text = event.value)
                }

                is LoginEvent.EnterPassword -> {
                    passwordState.value = password.value.copy(text = event.value)
                }

                is LoginEvent.SearchUser -> {
                    authenticationUseCase.validateUser(
                        LoginBody(
                            username = username.value.text,
                            password = password.value.text,
                        ),
                    ).onEach { user ->
                        when (user) {
                            is Resource.Loading -> {
                                _state.value = LoginDataState(isLoading = true)
                            }

                            is Resource.Error -> {
                                _state.value = LoginDataState(error = user.message ?: "Error")
                                if (user.message!!.contains("HTTP 404")) {
                                    eventFlowM.emit(
                                        UILoginEvent.ShowMessage("User not found"),
                                    )
                                } else if (user.message.contains("HTTP 401")) {
                                    eventFlowM.emit(
                                        UILoginEvent.ShowMessage("Invalid credentials"),
                                    )
                                } else {
                                    eventFlowM.emit(
                                        UILoginEvent.ShowMessage(user.message ?: "Error"),
                                    )
                                }
                            }

                            is Resource.Success -> {
                                _state.value = LoginDataState(error = user.message ?: "Autenticado")
                                eventFlowM.emit(UILoginEvent.ShowMessage(user.message.toString()))
                            }
                        }
                    }.launchIn(viewModelScope)
                }
            }
        }

        sealed class UILoginEvent {
            object GetData : UILoginEvent()

            data class ShowMessage(val message: String) : UILoginEvent()
        }
    }
