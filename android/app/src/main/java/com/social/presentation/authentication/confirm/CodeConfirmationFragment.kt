package com.social.presentation.authentication.confirm

import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.view.View
import android.widget.EditText
import androidx.fragment.app.Fragment
import com.social.R
import com.social.databinding.FragmentCodeConfirmationBinding
import com.social.utils.CodeGenerator.generateCode
import com.social.utils.Toast.showMessage

class CodeConfirmationFragment : Fragment(R.layout.fragment_code_confirmation) {
    private lateinit var binding: FragmentCodeConfirmationBinding
    private lateinit var codeArg: String

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentCodeConfirmationBinding.bind(view)

        action()
        codeArg = arguments?.getString("codeConfirm") ?: ""
    }

    private fun action() {
        numberField(arrayOf(binding.uno, binding.dos, binding.tres, binding.cuatro))
        binding.buttonSendCode.setOnClickListener {
            compareCodes(code(), codeArg)
        }
        binding.resendCode.setOnClickListener {
            resendCode()
        }
    }

    private fun numberField(fields: Array<EditText>) {
        for (i in 0 until fields.size - 1) {
            fields[i].addTextChangedListener(
                object : TextWatcher {
                    override fun beforeTextChanged(
                        s: CharSequence?,
                        start: Int,
                        count: Int,
                        after: Int,
                    ) {}

                    override fun onTextChanged(
                        s: CharSequence?,
                        start: Int,
                        before: Int,
                        count: Int,
                    ) {
                        if (s?.length == 1) {
                            fields[i + 1].requestFocus()
                        }
                    }

                    override fun afterTextChanged(s: Editable?) {}
                },
            )
        }
    }

    private fun resendCode() {
        codeArg = generateCode()
        showMessage(requireContext(), "Reenvio: $codeArg")
    }

    private fun code(): String {
        val code = StringBuilder()
        for (field in arrayOf(binding.uno, binding.dos, binding.tres, binding.cuatro)) {
            code.append(field.text)
        }
        return code.toString()
    }

    private fun compareCodes(
        inputCode: String,
        codeArg: String,
    ) {
        showMessage(requireContext(), "Codigo: $codeArg - Input: ${code()}")
        if (inputCode == codeArg) {
            showMessage(requireContext(), "Correcto")
        } else {
            showMessage(requireContext(), "Incorrecto")
        }
    }
}
