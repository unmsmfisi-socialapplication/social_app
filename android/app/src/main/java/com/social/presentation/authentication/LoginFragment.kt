package com.social.presentation.authentication

import android.os.Bundle
import android.util.Log
import android.view.View
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.lifecycle.lifecycleScope
import androidx.navigation.fragment.findNavController
import com.social.R
import com.social.databinding.FragmentLoginBinding
import com.social.utils.Toast.showMessage
import com.social.utils.Validation
import com.social.utils.Validation.setupValidation
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.flow.collectLatest

@AndroidEntryPoint
class LoginFragment : Fragment(R.layout.fragment_login) {
    private lateinit var binding: FragmentLoginBinding
    private lateinit var globalView: View
    private val viewModel: LoginViewModel by viewModels()

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentLoginBinding.bind(view)
        globalView = view

        lifecycleScope.launchWhenCreated {
            viewModel.evenFlow.collectLatest { event ->
                when (event) {
                    is LoginViewModel.UILoginEvent.GetData -> {
                        val userData = viewModel.state.value!!.dataLogin?.get(0)
                        Log.i("dato_usuario", userData.toString())
                    }

                    is LoginViewModel.UILoginEvent.ShowMessage -> {
                        showMessage(requireContext(), event.message)
                    }
                }
            }
        }

        // setupEmailValidation()
        // setupPasswordValidation()
        action()
    }

    private fun action() {
        binding.textRegister.setOnClickListener {
            findNavController()
                .navigate(R.id.action_loginFragment_to_registerFragment)
        }
        binding.textForgotPassword.setOnClickListener {
            findNavController()
                .navigate(R.id.action_loginFragment_to_confirmationEmailFragment)
        }
        binding.buttonGoogle.setOnClickListener {
            googleLoginButton()
        }
        binding.buttonMastodon.setOnClickListener {
            mastodonLoginButton()
        }
        binding.buttonLogin.setOnClickListener {
            loginButton()
        }
    }

    private fun googleLoginButton() {
        showMessage(requireContext(), "Pendiente")
    }

    private fun mastodonLoginButton() {
        showMessage(requireContext(), "Pendiente")
    }

    private fun setupEmailValidation() {
        setupValidation(binding.inputEmail, binding.errorEmail) { email ->
            Validation.isEmailValid(email)
        }
    }

    private fun setupPasswordValidation() {
        setupValidation(binding.inputPassword, binding.errorPassword) { password ->
            Validation.isPasswordValid(password)
        }
    }

    private fun loginButton() {
        val email = binding.inputEmail.text.toString()
        val password = binding.inputPassword.text.toString()
        /*if (Validation.isEmailValid(email) && Validation.isPasswordValid(password)) {
            performLogin(email, password)
        } else {
            showMessage(requireContext(), "Campos vacíos")
        }*/
        viewModel.getData(LoginEvent.EnterUser(value = email))
        viewModel.getData(LoginEvent.EnterPassword(value = password))
        viewModel.getData(LoginEvent.SearchUser)
        // viewModel.getData(LoginEvent.SearchUser)
    }

    private fun performLogin(
        email: String,
        password: String,
    ) {
        showMessage(requireContext(), "$email - $password")
    }
}
