package com.social.presentation.authentication.confirm

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.findNavController
import com.social.R
import com.social.databinding.FragmentConfirmationEmailBinding
import com.social.utils.CodeGenerator.generateCode
import com.social.utils.Toast.showMessage
import com.social.utils.Validation
import com.social.utils.Validation.setupValidation

class ConfirmationEmailFragment : Fragment(R.layout.fragment_confirmation_email) {
    private lateinit var binding: FragmentConfirmationEmailBinding
    private lateinit var email: String
    private lateinit var codeArg: String
    private lateinit var bundle: Bundle

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentConfirmationEmailBinding.bind(view)

        action()
    }

    private fun action() {
        binding.buttonContinue.setOnClickListener {
            continueButton()
        }
        setupEmailValidation()
    }

    private fun setupEmailValidation() {
        setupValidation(binding.inputEmailRecover, binding.errorEmail) { email ->
            Validation.isEmailValid(email)
        }
    }

    private fun continueButton() {
        email = binding.inputEmailRecover.text.toString()
        if (Validation.isEmailValid(email)) {
            sendCode()
        } else {
            showMessage(requireContext(), "Campos vacios")
        }
    }

    private fun sendCode() {
        codeArg = generateCode()
        showMessage(requireContext(), "Codigo: $codeArg")
        bundle =
            Bundle().apply {
                putString("codeConfirm", codeArg)
            }
        findNavController()
            .navigate(R.id.action_confirmationEmailFragment_to_codeConfirmationFragment, bundle)
    }
}
