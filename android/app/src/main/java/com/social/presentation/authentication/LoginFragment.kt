package com.social.presentation.authentication

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.findNavController
import com.social.R
import com.social.databinding.FragmentLoginBinding
import com.social.utils.Toast.showMessage
import com.social.utils.Validation
import com.social.utils.Validation.setupValidation


class LoginFragment : Fragment(R.layout.fragment_login) {
    private lateinit var binding: FragmentLoginBinding
    private lateinit var globalView: View

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentLoginBinding.bind(view)
        globalView = view

        setupEmailValidation()
        setupPasswordValidation()
        action()
    }
    private fun action() {
        binding.textRegister.setOnClickListener {
            findNavController().navigate(R.id.action_loginFragment_to_registerFragment)
        }

        binding.textForgotPassword.setOnClickListener {
            recoverPassword()
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

    private fun recoverPassword(){
        showMessage(requireContext(), "Pendiente")
    }

    private fun googleLoginButton() {
        showMessage(requireContext(), "Pendiente")
    }

    private fun mastodonLoginButton(){
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
        if (Validation.isEmailValid(email) && Validation.isPasswordValid(password)) {
            performLogin(email, password)
        } else {
            showMessage(requireContext(), "Campos vacios")
        }
    }

    private fun performLogin(email: String, password: String) {
        showMessage(requireContext(), "$email - $password")
    }

}