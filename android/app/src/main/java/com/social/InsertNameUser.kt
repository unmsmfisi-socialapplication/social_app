package com.social

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.ImageView
import android.widget.ImageButton
import android.content.Intent
import android.widget.Button
import android.widget.EditText
import android.view.View
import android.text.Editable
import android.text.TextWatcher

class InsertNameUser : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_insert_name_user)

        val profileImageView = findViewById<ImageView>(R.id.userFoto_name)
        val ibtnAtras = findViewById<ImageButton>(R.id.ibtnAtras)
        val inputName = findViewById<EditText>(R.id.inputName)
        val btnAceptar = findViewById<Button>(R.id.btnAceptar)

        // Obtener la imagen pasada desde MainActivity
        val selectedImageUriString = intent.getStringExtra("selectedImageUri")

        // Si se pasó una imagen, establecerla en el ImageView
        if (!selectedImageUriString.isNullOrEmpty()) {
            val selectedImageUri = android.net.Uri.parse(selectedImageUriString)
            profileImageView.setImageURI(selectedImageUri)
        }

        // Volver
        ibtnAtras.setOnClickListener {
            // Crear un Intent para volver a MainActivity
            val intent = Intent(this, MainActivity::class.java)
            startActivity(intent)
            finish() // Cierra la actividad actual para que no quede en segundo plano
        }

        //
        inputName.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(s: CharSequence?, start: Int, count: Int, after: Int) {
                // No es necesario implementar este método
            }

            override fun onTextChanged(s: CharSequence?, start: Int, before: Int, count: Int) {
                // Comprueba si el texto del EditText no está vacío y muestra el botón si es cierto
                if (!s.isNullOrEmpty()) {
                    btnAceptar.visibility = View.VISIBLE
                } else {
                    btnAceptar.visibility = View.INVISIBLE
                }
            }

            override fun afterTextChanged(s: Editable?) {
                // No es necesario implementar este método
            }
        })
    }
}