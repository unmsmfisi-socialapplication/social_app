package com.social

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.ImageView
import android.content.Intent
import android.content.SharedPreferences
import android.preference.PreferenceManager
import android.widget.Button
import android.widget.EditText
import android.view.View
import android.text.Editable
import android.text.TextWatcher

class InsertNameUser : AppCompatActivity() {
    lateinit var userPhoto_select: ImageView
    lateinit var btnBack_uploadP: Button
    lateinit var btnAccept_name: Button
    lateinit var inputName: EditText

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_insert_name_user)

        userPhoto_select = findViewById(R.id.userPhoto_select)
        btnBack_uploadP = findViewById(R.id.btnBack_uploadP)
        inputName = findViewById(R.id.inputName)
        btnAccept_name = findViewById(R.id.btnAccept_name)

        // Obtener la imagen pasada desde UploadPhotoUser en SharedPreferences
        val sharedPreferences: SharedPreferences =
            PreferenceManager.getDefaultSharedPreferences(this)
        val selectedImageUriString = sharedPreferences.getString("selectedImageUri", "")

        // Si se pasó una imagen, establecerla en el ImageView
        if (!selectedImageUriString.isNullOrEmpty()) {
            val selectedImageUri = android.net.Uri.parse(selectedImageUriString)
            userPhoto_select.setImageURI(selectedImageUri)
        }

        // Regresar a la View para subir foto de perfil
        btnBack_uploadP.setOnClickListener {
            // Crear un Intent para volver a UploadPhotoUser
            val intent = Intent(this, UploadPhotoUser::class.java)
            startActivity(intent)
            finish() // Cierra la actividad actual para que no quede en segundo plano
        }

        inputName.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(s: CharSequence?, start: Int, count: Int, after: Int) {
                // No es necesario implementar este método
            }

            override fun onTextChanged(s: CharSequence?, start: Int, before: Int, count: Int) {
                if (!s.isNullOrEmpty()) {
                    btnAccept_name.visibility = View.VISIBLE
                } else {
                    btnAccept_name.visibility = View.INVISIBLE
                }
            }

            override fun afterTextChanged(s: Editable?) {
                // No es necesario implementar este método
            }
        })

        btnAccept_name.setOnClickListener {
            // Obtener el texto ingresado en el EditText
            val name = inputName.text.toString()

            // Obtener la imagen pasada desde UploadPhotoUser en SharedPreferences
            val sharedPreferences: SharedPreferences =
                PreferenceManager.getDefaultSharedPreferences(this)
            val selectedImageUriString = sharedPreferences.getString("selectedImageUri", "")

            // Si se pasó una imagen, establecerla en el ImageView de InsertBibliographyUser
            if (!selectedImageUriString.isNullOrEmpty()) {
                val selectedImageUri = android.net.Uri.parse(selectedImageUriString)

                // Crear un Intent para ir a InsertBibliographyUser y pasar la imagen y el nombre
                val intent = Intent(this, InsertBibliographyUser::class.java)
                intent.putExtra("selectedImageUri", selectedImageUri.toString())
                intent.putExtra("name", name)
                startActivity(intent)
            }
        }

    }
}
