package com.example.mobilesocialapp.presentation

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.twotone.Face
import androidx.compose.material.icons.twotone.Info
import androidx.compose.material3.Button
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.text.input.VisualTransformation
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.mobilesocialapp.ui.theme.MobileSocialAPPTheme

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            BodyRegister()
        }
    }
}


@OptIn(ExperimentalMaterial3Api::class)
@Preview(showBackground = true, showSystemUi = true)
@Composable
fun BodyRegister() {
    var name by remember { mutableStateOf("") }
    var lastname by remember { mutableStateOf("") }
    var user by remember { mutableStateOf("") }
    var email by remember { mutableStateOf("") }
    var date by remember { mutableStateOf("") }
    var genrer by remember{ mutableStateOf("") }
    var psw1 by remember{ mutableStateOf("") }
    var psw2 by remember{ mutableStateOf("") }
    var isPasswordVisible by remember {
        mutableStateOf(false)
    }

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
                Label("Registro", TextAlign.Center, 30)
            }
            Column(modifier = Modifier.fillMaxSize()) {
                Column(modifier = Modifier.padding(top = 20.dp, start = 20.dp, end = 20.dp)) {
                    Label("Ingrese los siguientes datos", TextAlign.Left, 14)
                }
                Column(modifier = Modifier.padding(top = 20.dp, start = 20.dp, end = 20.dp)) {
                    OutlinedTextField(
                        value = name,
                        modifier = Modifier.fillMaxWidth(),
                        onValueChange = { name = it },
                        label = { Text("Nombres") },
                        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.NumberPassword)
                    )
                }
                Column(modifier = Modifier.padding(top = 20.dp, start = 20.dp, end = 20.dp)) {
                    OutlinedTextField(
                        value = lastname,
                        modifier = Modifier.fillMaxWidth(),
                        onValueChange = { lastname = it },
                        label = { Text("Apellidos") },
                        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.NumberPassword)
                    )
                }
                Column(modifier = Modifier.padding(top = 20.dp, start = 20.dp, end = 20.dp)) {
                    OutlinedTextField(
                        value = user,
                        modifier = Modifier.fillMaxWidth(),
                        onValueChange = { user = it },
                        label = { Text("Nombre de usuario") },
                        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.NumberPassword)
                    )
                }

                Column(modifier = Modifier.padding(top = 20.dp, start = 20.dp, end = 20.dp)) {
                    OutlinedTextField(
                        value = email,
                        modifier = Modifier.fillMaxWidth(),
                        onValueChange = { email = it },
                        label = { Text("Correo electrónico") },
                        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.NumberPassword)
                    )
                }
                Row(modifier = Modifier.padding(top = 20.dp, start = 20.dp, end = 20.dp), horizontalArrangement = Arrangement.SpaceBetween){
                    Column(modifier = Modifier.weight(2f)) {
                        OutlinedTextField(
                            value = date,
                            modifier = Modifier.fillMaxWidth(),
                            onValueChange = { date = it },
                            label = { Text("Fecha de nacimiento") },
                            keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.NumberPassword)
                        )
                    }
                    EspaciadorVertical(tamaño = 10.0)
                    Column(modifier = Modifier.weight(1f)) {
                        OutlinedTextField(
                            value = genrer,
                            modifier = Modifier.fillMaxWidth(),
                            onValueChange = { genrer = it },
                            label = { Text("Sexo") },
                            keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.NumberPassword)
                        )
                    }
                }

                Column(modifier = Modifier.padding(top = 20.dp, start = 20.dp, end = 20.dp)) {
                    OutlinedTextField(
                        value = psw1,
                        onValueChange = { psw1 = it },
                        modifier = Modifier.fillMaxWidth(),
                        label = { Text("Contraseña") },
                        //visualTransformation = PasswordVisualTransformation(),
                        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Password),
                        visualTransformation = if (isPasswordVisible) VisualTransformation.None else PasswordVisualTransformation(),
                        trailingIcon = {
                            IconButton(
                                onClick = { isPasswordVisible = !isPasswordVisible },
                                modifier = Modifier.padding(horizontal = 8.dp)
                            ) {
                                val eyeIcon = if (isPasswordVisible) {
                                    Icons.TwoTone.Face
                                } else {
                                    Icons.TwoTone.Info
                                }
                                Icon(

                                    imageVector = eyeIcon,
                                    contentDescription = "Toggle Password Visibility"
                                )
                            }
                        }
                    )
                }

                Column(modifier = Modifier.padding(top = 20.dp, start = 20.dp, end = 20.dp)) {
                    OutlinedTextField(
                        value = psw2,
                        onValueChange = { psw2 = it },
                        modifier = Modifier.fillMaxWidth(),
                        label = { Text("Confirmar contraseña") },
                        //visualTransformation = PasswordVisualTransformation(),
                        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Password),
                        visualTransformation = if (isPasswordVisible) VisualTransformation.None else PasswordVisualTransformation(),
                        trailingIcon = {
                            IconButton(
                                onClick = { isPasswordVisible = !isPasswordVisible },
                                modifier = Modifier.padding(horizontal = 8.dp)
                            ) {
                                val eyeIcon = if (isPasswordVisible) {
                                    Icons.TwoTone.Face
                                } else {
                                    Icons.TwoTone.Info
                                }
                                Icon(

                                    imageVector = eyeIcon,
                                    contentDescription = "Toggle Password Visibility"
                                )
                            }
                        }
                    )
                }



                Column(modifier = Modifier.padding(top = 40.dp, end = 20.dp, start = 20.dp)) {
                    Button(
                        onClick = {
                        }
                    ) {
                        Label(texto = "Registrarse", alineacion = TextAlign.Center, 22)
                    }
                }
                Column(modifier = Modifier.padding(10.dp)) {
                    //LinkText("¿Olvidaste tu contraseña?")
                }
            }
        }
    }

    }