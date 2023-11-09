package com.social.presentation.authentication

import android.os.Bundle
import android.util.Log
import android.view.View
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.lifecycle.lifecycleScope
import androidx.navigation.fragment.findNavController
import com.social.R
import com.social.databinding.FragmentRegisterBinding
import com.social.utils.CodeGenerator
import com.social.utils.Toast.showMessage
import com.social.utils.Validation
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.flow.collectLatest
import kotlinx.coroutines.launch

@AndroidEntryPoint
class RegisterFragment : Fragment(R.layout.fragment_register) {
    private lateinit var binding: FragmentRegisterBinding
    private lateinit var globalView: View
    private val viewModel: RegisterUserViewModel by viewModels()

    private lateinit var codeArg: String
    private lateinit var bundle: Bundle

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentRegisterBinding.bind(view)
        globalView = view

        viewLifecycleOwner.lifecycleScope.launch {
            viewModel.evenFlow.collectLatest { event ->
                when (event) {
                    is RegisterUserViewModel.UIRegisterUserEvent.GetData -> {
                        val userRegisterData = viewModel.state.value!!.dataRegister[0]
                        Log.i("registrado", userRegisterData.email)
                    }

                    is RegisterUserViewModel.UIRegisterUserEvent.ShowMessage -> {
                        showMessage(requireContext(), event.message)
                    }
                }
            }
        }

        action()
        setup()
    }

    private fun action() {
        binding.textLogIn.setOnClickListener {
            findNavController().navigate(R.id.action_registerFragment_to_loginFragment)
        }
        binding.btnRegister.setOnClickListener {
            registerButton()
        }
        binding.buttonGoogle.setOnClickListener {
            googleLoginButton()
        }
        binding.buttonMastodon.setOnClickListener {
            mastodonLoginButton()
        }
    }

    private fun setup() {
        setupUserValidation()
        setupPasswordValidation()
        setupEmailValidation()
        setupCoincidence()
    }

    private fun googleLoginButton() {
        showMessage(requireContext(), "Proximamente")
    }

    private fun mastodonLoginButton() {
        showMessage(requireContext(), "Proximamente")
    }

    private fun registerButton() {
        val email = binding.txtEmailRegister.text.toString()
        val user = binding.txtUserRegister.text.toString()
        val password = binding.txtPasswordRegister.text.toString()
        val passwordConfirm = binding.txtConfirmPasswordRegister.text.toString()

        if (Validation.isPasswordValid(password) && Validation.isUserValid(user) && Validation.isEmailValid(email)
        ) {
            viewModel.insertData(RegisterEvent.EnterUsername(value = user))
            viewModel.insertData(RegisterEvent.EnterEmail(value = email))
            viewModel.insertData(RegisterEvent.EnterPassword(value = passwordConfirm))
            viewModel.insertData(RegisterEvent.RegisterUser)
            // navigateToPin()
        }
    }

    private fun setupUserValidation() {
        Validation.setupValidation(
            binding.txtUserRegister,
            binding.errorUser,
        ) { user ->
            Validation.isUserValid(user)
        }
    }

    private fun setupPasswordValidation() {
        Validation.setupValidation(
            binding.txtPasswordRegister,
            binding.errorPassword,
        ) { password ->
            Validation.isPasswordValid(password)
        }
    }

    private fun setupEmailValidation() {
        Validation.setupValidation(
            binding.txtEmailRegister,
            binding.errorEmail,
        ) { email ->
            Validation.isEmailValid(email)
        }
    }

    private fun setupCoincidence() {
        Validation.setupCoincidence(
            binding.txtPasswordRegister,
            binding.txtConfirmPasswordRegister,
            binding.errorCoincidence,
        )
    }

    private fun navigateToPin() {
        sendCode()
    }

    private fun sendCode() {
        codeArg = CodeGenerator.generateCode()
        showMessage(requireContext(), "Codigo: $codeArg")
        bundle =
            Bundle().apply {
                putString("codeConfirm", codeArg)
            }
        findNavController()
            .navigate(R.id.action_registerFragment_to_codeConfirmationFragment, bundle)
    }
}
