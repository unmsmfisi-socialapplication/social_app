package com.social.presentation.authentication

import android.app.Application
import android.util.Log
import androidx.compose.runtime.State
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.social.domain.model.LoginBody
import com.social.domain.use_case.ValidateUser
import com.social.utils.Resource
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableSharedFlow
import kotlinx.coroutines.flow.asSharedFlow
import kotlinx.coroutines.flow.launchIn
import kotlinx.coroutines.flow.onEach
import javax.inject.Inject

@HiltViewModel
class LoginViewModel @Inject constructor(
    private val authenticationUseCase: AuthenticationUseCase
) : ViewModel() {
    private val _state = MutableLiveData<LoginDataState>()
    val state: LiveData<LoginDataState> get() = _state
    private val _username = mutableStateOf(TextState())
    val username: State<TextState> = _username
    private val _password = mutableStateOf(TextState())
    val password: State<TextState> = _password
    private val _eventFlow = MutableSharedFlow<UILoginEvent>()
    val evenFlow = _eventFlow.asSharedFlow()

    fun getData(event: LoginEvent) {
        when (event) {
            is LoginEvent.EnterUser -> {
                _username.value = username.value.copy(text = event.value)
            }

            is LoginEvent.EnterPassword -> {
                _password.value = password.value.copy(text = event.value)
            }

            is LoginEvent.SearchUser -> {
                authenticationUseCase.validateUser(
                    LoginBody(
                        username = username.value.text,
                        password = password.value.text
                    )
                ).onEach { user ->
                    Log.i("Autenticacion: ","user: $username and password: $password")
                    when (user){
                        is Resource.Loading->{
                            _state.value = LoginDataState(isLoading = true)
                        }
                        is Resource.Error->{
                            _state.value = LoginDataState(error = user.message)
                            val status = _state.value?.error
                            val status2 = user.data?.status
                            val status3 = user.message
                            Log.i("Monchito", status.toString()?:"a")
                            Log.i("Monchito", status2.toString()?:"b")
                            Log.i("Monchito", status3.toString()?:"c")
                            _state.value = LoginDataState(error = status!!)
                            /*if(user.message!!.contains("404")){
                                _eventFlow.emit(
                                    UILoginEvent.ShowMessage(user.message?:"Error")
                                )
                            }else if(user.message.contains("401")){
                                _eventFlow.emit(
                                    UILoginEvent.ShowMessage(user.message?:"Error")
                                )
                            }else{
                                _state.value = LoginDataState(error = user.message ?: "Error")
                            }*/
                            _eventFlow.emit(
                                UILoginEvent.ShowMessage(status?:"Error")
                            )

                        }
                        is Resource.Success->{
                            /*_state.value = LoginDataState(dataLogin = user.data!!)
                            if(_state.value!!.dataLogin.isNotEmpty()){
                                _eventFlow.emit(UILoginEvent.GetData)
                            }else{
                                UILoginEvent.ShowMessage(user.message?:"Error de datos usuario")
                            }*/
                            val status = _state.value?.dataLogin
                            val status2 = user.data?.status
                            val status3 = user.message
                            Log.i("Monchito", status.toString()?:"a")
                            Log.i("Monchito", status2.toString()?:"b")
                            Log.i("Monchito", status3.toString()?:"c")
                            _state.value = LoginDataState(dataLogin = status)
                            _eventFlow.emit(UILoginEvent.ShowMessage(status?:"Bien"))
                            //_eventFlow.emit(UILoginEvent.ShowMessage(user.message?:"Authentication successful"))

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