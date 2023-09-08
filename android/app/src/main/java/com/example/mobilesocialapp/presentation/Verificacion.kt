package com.example.mobilesocialapp.presentation

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.mobilesocialapp.R
import com.example.mobilesocialapp.presentation.ui.theme.MobileSocialAPPTheme

class Verificacion : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            BodyValidate()
        }
    }
}

@Preview(showBackground = true)
@Composable
fun BodyValidate() {
    Box(
        modifier = Modifier
            .fillMaxSize()
            .background(Color.White), contentAlignment = Alignment.TopCenter
    ) {
        Column(
            modifier = Modifier
                .padding(top = 40.dp)
                .verticalScroll(rememberScrollState())
            , horizontalAlignment = Alignment.CenterHorizontally
        ) {
            EspaciadorHorizontal(tamaño = 30.0)
            Column(modifier = Modifier.padding(horizontal = 20.dp)) {
                Label("Verficiacion de cuenta", TextAlign.Center, 24)
            }
            Column {
                Column(modifier = Modifier.fillMaxSize()) {
                    Column(modifier = Modifier.padding(top = 20.dp, start = 20.dp, end = 20.dp)) {
                        Label("Ingrese los siguientes datos", TextAlign.Left, 14)
                    }
                    Column(modifier = Modifier
                        .padding(top = 20.dp, start = 20.dp, end = 20.dp)
                        ) {
                        Image(
                            painter = painterResource(id = R.drawable.imgemail),
                            contentDescription = "Logo pepsico",
                            modifier = Modifier
                                .size(238.dp, 180.86.dp),
                            alignment = Alignment.Center
                        )
                    }
                    Column(modifier = Modifier.padding(top = 20.dp, start = 20.dp, end = 20.dp),
                        horizontalAlignment = Alignment.CenterHorizontally) {
                        Label("¡Listo, se ha enviado un enlace de verificación a su correo electrónico!", TextAlign.Left, 14)
                    }
                    Column(modifier = Modifier.padding(top = 20.dp, start = 20.dp, end = 20.dp),
                        horizontalAlignment = Alignment.CenterHorizontally) {
                        Label("Puede minimizar la pantalla para realizar la verificación", TextAlign.Left, 14)
                    }
                    Column(modifier = Modifier.padding(top = 40.dp, end = 20.dp, start = 20.dp)) {
                        Button(
                            onClick = {
                            }
                        ) {
                            Label(texto = "Reenviar enlace", alineacion = TextAlign.Center, 22)
                        }
                    }
                }
            }
        }
    }
}