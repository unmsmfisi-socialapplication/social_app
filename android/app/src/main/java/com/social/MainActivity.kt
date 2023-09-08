package com.social

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import android.widget.Button
import android.widget.ImageView
import android.content.Intent

class MainActivity : AppCompatActivity() {
    lateinit var btnSubirFoto: Button
    lateinit var userFoto: ImageView

    val SELECT_IMAGE_REQUEST = 1001

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        btnSubirFoto = findViewById(R.id.btnSubirFoto)
        userFoto = findViewById(R.id.userFoto)

        btnSubirFoto.setOnClickListener {
            val intent = Intent(Intent.ACTION_GET_CONTENT)
            intent.type = "image/*" // Filtrar por tipo de archivo (imágenes en este caso)
            startActivityForResult(intent, SELECT_IMAGE_REQUEST)
        }
    }

    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
        super.onActivityResult(requestCode, resultCode, data)

        if (requestCode == SELECT_IMAGE_REQUEST && resultCode == RESULT_OK) {
            val selectedImageUri = data?.data
            if (selectedImageUri != null) {
                // Cargar la imagen seleccionada
                userFoto.setImageURI(selectedImageUri)

                // Pasar la imagen a la segunda actividad
                val intent = Intent(this, InsertNameUser::class.java)
                intent.putExtra("selectedImageUri", selectedImageUri.toString())
                startActivity(intent)
            } else {
                // No se seleccionó ninguna imagen
                Log.i("aris", "No se seleccionó ninguna imagen")
            }
        }
    }
}