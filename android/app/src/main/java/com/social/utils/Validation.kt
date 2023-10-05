package com.social.utils

import android.text.Editable
import android.text.TextWatcher
import android.util.Patterns
import android.view.View
import android.widget.EditText

object Validation {

    fun isEmailValid(email: String): Boolean {
        return Patterns.EMAIL_ADDRESS.matcher(email).matches()
    }

    fun isPasswordValid(password: String): Boolean {
        return password.isNotEmpty() && password.length >= 8
    }

    fun setupValidation(
        inputField: EditText,
        errorView: View,
        validationFunction: (String) -> Boolean
    ) {
        inputField.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(s: CharSequence?, start: Int, count: Int, after: Int) {}
            override fun onTextChanged(s: CharSequence?, start: Int, before: Int, count: Int) {
                if (s?.isNotEmpty() == true) {
                    if (validationFunction(s.toString())) {
                        errorView.visibility = View.GONE
                    } else {
                        errorView.visibility = View.VISIBLE
                    }
                }
            }

            override fun afterTextChanged(s: Editable?) {}
        })
    }
}