package com.social

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import android.widget.Button
import android.widget.ImageView
import android.content.Intent
import android.widget.TextView
import android.content.SharedPreferences
import android.preference.PreferenceManager

class UploadPhotoUser : AppCompatActivity() {
    lateinit var btnUploadPhoto: Button
    lateinit var userPhoto: ImageView
    lateinit var txtSkip: TextView

    val SELECT_IMAGE_REQUEST = 1001

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_upload_photo_user)

        btnUploadPhoto = findViewById(R.id.btnUploadPhoto)
        userPhoto = findViewById(R.id.userPhoto)
        txtSkip = findViewById(R.id.txtSkip)

        btnUploadPhoto.setOnClickListener {
            val intent = Intent(Intent.ACTION_GET_CONTENT)
            intent.type = "image/*" // Filtrar por tipo de archivo (imágenes en este caso)
            startActivityForResult(intent, SELECT_IMAGE_REQUEST)
        }

        // Omitir paso de subir foto de perfil
        txtSkip.setOnClickListener {
            // Crear un Intent para ir a InsertNameUser
            val intent = Intent(this, InsertNameUser::class.java)
            startActivity(intent)
            finish() // Cierra la actividad actual para que no quede en segundo plano
        }
    }

    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
        super.onActivityResult(requestCode, resultCode, data)

        if (requestCode == SELECT_IMAGE_REQUEST && resultCode == RESULT_OK) {
            val selectedImageUri = data?.data
            if (selectedImageUri != null) {
                // Cargar la imagen seleccionada
                userPhoto.setImageURI(selectedImageUri)

                // Guardar la URI en SharedPreferences
                val sharedPreferences: SharedPreferences =
                    PreferenceManager.getDefaultSharedPreferences(this)
                val editor = sharedPreferences.edit()
                editor.putString("selectedImageUri", selectedImageUri.toString())
                editor.apply()

                // Pasar a la segunda actividad
                val intent = Intent(this, InsertNameUser::class.java)
                startActivity(intent)
            } else {
                // No se seleccionó ninguna imagen
                Log.i("aris", "No se seleccionó ninguna imagen")
            }
        }
    }
}