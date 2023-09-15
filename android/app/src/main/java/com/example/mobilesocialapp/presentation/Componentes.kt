package com.example.mobilesocialapp.presentation

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.text.ClickableText
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.font.Font
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import com.example.mobilesocialapp.R


@Composable
fun EspaciadorHorizontal(tamaño: Double) {
    Spacer(
        modifier = Modifier
            .height(tamaño.dp)
            .width(0.dp)
    )
}


@Composable
fun EspaciadorVertical(tamaño: Double) {
    Spacer(
        modifier = Modifier
            .width(tamaño.dp)
            .width(0.dp)
    )
}
@Composable
fun Label(texto: String, alineacion: TextAlign?, tamaño: Int) {
    val poppins = FontFamily(
        Font(R.font.poppinsregular, FontWeight.Normal)
    )

    Text(
        text = texto,
        modifier = Modifier.fillMaxWidth(),
        fontFamily = poppins,
        textAlign = alineacion,
        fontSize = tamaño.sp
    )
}
