package com.social.utils

import android.content.Context
import android.widget.Toast

object Toast {
    fun showMessage(context:Context ,message: String) {
        Toast.makeText(context, message, Toast.LENGTH_SHORT).show()
    }
}