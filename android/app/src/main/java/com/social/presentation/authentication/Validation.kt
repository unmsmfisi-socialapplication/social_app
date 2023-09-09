package com.social.presentation.authentication

import android.content.Context
import android.content.res.ColorStateList
import android.util.Patterns
import androidx.core.content.ContextCompat
import com.google.android.material.textfield.TextInputLayout
import com.social.R
object Validation {
    fun validateAndHighlightInput(context: Context, inputLayout: TextInputLayout, text: String, isValid: Boolean, errorColorResId: Int) {
        if (!isValid) {
            changeTextInputLayoutColor(context, inputLayout, errorColorResId)
        } else {
            resetTextInputLayoutColor(context, inputLayout)
        }
    }

    private fun changeTextInputLayoutColor(context: Context, textInputLayout: TextInputLayout, colorResId: Int) {
        val color = ContextCompat.getColor(context, colorResId)
        textInputLayout.defaultHintTextColor = ColorStateList.valueOf(color)
    }

    private fun resetTextInputLayoutColor(context: Context, textInputLayout: TextInputLayout) {
        val color = ContextCompat.getColor(context, R.color.black)
        textInputLayout.boxStrokeColor = color
        textInputLayout.defaultHintTextColor = ColorStateList.valueOf(color)
    }

    fun isEmailValid(email: String): Boolean {
        return Patterns.EMAIL_ADDRESS.matcher(email).matches()
    }

    fun isPasswordValid(password: String): Boolean {
        return password.isNotEmpty()
    }
}