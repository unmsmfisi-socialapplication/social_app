package com.social.presentation.authentication

import android.os.Bundle
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import com.social.R
import com.social.databinding.ActivityLoginBinding

class LoginActivity : AppCompatActivity() {
    private lateinit var binding: ActivityLoginBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = ActivityLoginBinding.inflate(layoutInflater)
        setContentView(binding.root)

        binding.buttonContinueGoogle.setOnClickListener {
            Toast.makeText(this@LoginActivity, "Pendiente", Toast.LENGTH_SHORT).show()
        }

        binding.buttonCreateAccount.setOnClickListener {
            Toast.makeText(this@LoginActivity, "Pendiente", Toast.LENGTH_SHORT).show()
        }

        binding.buttonLogin.setOnClickListener {
            val isEmailValid = Validation.isEmailValid(binding.inputEmail.text.toString())
            val isPasswordValid = Validation.isPasswordValid(binding.inputPassword.text.toString())

            Validation.validateAndHighlightInput(
                this,
                binding.layoutEmail,
                binding.inputEmail.text.toString(),
                isEmailValid,
                R.color.red
            )

            Validation.validateAndHighlightInput(
                this,
                binding.layoutPassword,
                binding.inputPassword.text.toString(),
                isPasswordValid,
                R.color.red
            )

            if (isEmailValid && isPasswordValid) {
                performLogin(binding.inputEmail.text.toString(), binding.inputPassword.text.toString())
            } else {
                Toast.makeText(this@LoginActivity, "Correo o contraseña no válidos", Toast.LENGTH_SHORT).show()
            }
        }
    }

    private fun performLogin(email: String, password: String) {
        // Implement your login logic here
    }
}