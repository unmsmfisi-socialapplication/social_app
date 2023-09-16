package com.social

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.view.View
import android.widget.Button
import android.widget.EditText
import android.widget.ImageView
import android.widget.TextView

class InsertBibliographyUser : AppCompatActivity() {
    lateinit var btnBack_name: Button
    lateinit var btnAccept_bibliography : Button
    lateinit var txtNameUser: TextView
    lateinit var userPhoto_Final: ImageView
    lateinit var inputBibliography: EditText

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_insert_bibliography_user)

        btnBack_name = findViewById(R.id.btnBack_name)
        btnAccept_bibliography = findViewById(R.id.btnAccept_bibliography)
        txtNameUser = findViewById(R.id.txtNameUser)
        inputBibliography = findViewById(R.id.inputbibliography)
        userPhoto_Final = findViewById(R.id.userPhoto_Final)

        // Regresar a la View  para subir foto de perfil
        btnBack_name.setOnClickListener {
            // Crear un Intent para volver a InsertNameUser
            val intent = Intent(this, InsertNameUser::class.java)
            startActivity(intent)
            finish() // Cierra la actividad actual para que no quede en segundo plano
        }

        // Obtener la URI de la imagen y el nombre pasada desde InsertNameUser
        val selectedImageUriString = intent.getStringExtra("selectedImageUri")
        val name = intent.getStringExtra("name")

        // Si se pasó una imagen, establecerla en el ImageView
        if (!selectedImageUriString.isNullOrEmpty()) {
            val selectedImageUri = android.net.Uri.parse(selectedImageUriString)
            userPhoto_Final.setImageURI(selectedImageUri)
        }

        // Si se pasó un nombre, establecerlo en el TextView
        if (!name.isNullOrEmpty()) {
            // Se divide el nombre en palabras usando espacio como separador
            val separatorName = name.split(" ")

            // Inicializa una lista para almacenar las palabras capitalizadas
            val formattedListName = mutableListOf<String>()

            // Convierte la primera letra a mayúscula y el resto a minúscula
            // Capitaliza cada palabra y agrégala a la lista
            for (Name in separatorName) {
                if (Name.isNotEmpty()) {
                    val palabraCapitalizada = Name.substring(0, 1).toUpperCase() + Name.substring(1).toLowerCase()
                    formattedListName.add(palabraCapitalizada)
                }
            }

            // Capitalizamos en una sola cadena con espacios
            val textNameFormatted = formattedListName.joinToString(" ")

            // Establece el texto formateado en el TextView
            txtNameUser.text = textNameFormatted
        }

        inputBibliography.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(s: CharSequence?, start: Int, count: Int, after: Int) {
                // No es necesario implementar este método
            }

            override fun onTextChanged(s: CharSequence?, start: Int, before: Int, count: Int) {
                if (!s.isNullOrEmpty()) {
                    btnAccept_bibliography.visibility = View.VISIBLE
                } else {
                    btnAccept_bibliography.visibility = View.INVISIBLE
                }
            }

            override fun afterTextChanged(s: Editable?) {
                // No es necesario implementar este método
            }
        })
    }
}