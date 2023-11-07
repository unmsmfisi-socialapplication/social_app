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
import com.social.utils.Toast
import com.social.utils.Toast.showMessage
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
                        val userRegisterData = viewModel.state.value!!.dataRegister?.get(0)
                        Log.i("registrado", userRegisterData!!.email)
                    }
                    is RegisterUserViewModel.UIRegisterUserEvent.ShowMessage -> {
                        showMessage(requireContext(), event.message)
                    }
                }
            }
        }

        action()
    }

    private fun action() {
        binding.textLogIn.setOnClickListener {
            findNavController().navigate(R.id.action_registerFragment_to_loginFragment)
        }
        binding.btnRegister.setOnClickListener {
            val email = binding.txtEmailRegister.text.toString()
            val username = binding.txtFullNameRegister.text.toString()
            val passwordOne = binding.txtPasswordRegister.text.toString()
            val passwordTwo = binding.txtConfirmPasswordRegister.text.toString()
            viewModel.insertData(RegisterEvent.EnterUsername(value = username))
            viewModel.insertData(RegisterEvent.EnterEmail(value = email))
            viewModel.insertData(RegisterEvent.EnterPassword(value = passwordTwo))
            viewModel.insertData(RegisterEvent.RegisterUser)
            // navegateToPin()
        }
        binding.buttonGoogle.setOnClickListener {
            googleLoginButton()
        }
        binding.buttonMastodon.setOnClickListener {
            mastodonLoginButton()
        }
    }

    private fun googleLoginButton() {
        Toast.showMessage(requireContext(), "Proximamente")
    }

    private fun mastodonLoginButton() {
        Toast.showMessage(requireContext(), "Proximamente")
    }

    private fun navegateToPin() {
        sendCode()
    }

    private fun sendCode() {
        codeArg = CodeGenerator.generateCode()
        Toast.showMessage(requireContext(), "Codigo: $codeArg")
        bundle =
            Bundle().apply {
                putString("codeConfirm", codeArg)
            }
        findNavController()
            .navigate(R.id.action_registerFragment_to_codeConfirmationFragment, bundle)
    }
}
