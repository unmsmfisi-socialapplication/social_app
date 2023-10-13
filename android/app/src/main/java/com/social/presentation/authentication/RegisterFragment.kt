package com.social.presentation.authentication

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.findNavController
import com.social.R
import com.social.databinding.FragmentRegisterBinding
import com.social.utils.CodeGenerator
import com.social.utils.Toast

class RegisterFragment : Fragment(R.layout.fragment_register) {
    private lateinit var binding: FragmentRegisterBinding
    private lateinit var globalView: View

    private lateinit var codeArg: String
    private lateinit var bundle: Bundle

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentRegisterBinding.bind(view)
        globalView = view

        action()
    }

    private fun action() {
        binding.textLogIn.setOnClickListener {
            findNavController().navigate(R.id.action_registerFragment_to_loginFragment)
        }
        binding.btnRegister.setOnClickListener {
            navegateToPin()
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
