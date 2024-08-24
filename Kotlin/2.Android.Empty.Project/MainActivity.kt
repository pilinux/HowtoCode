package com.example.greetingcard

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.greetingcard.ui.theme.GreetingCardTheme

class MainActivity : ComponentActivity() {
    /*
     * In Kotlin programs, the main() function is the entry point/starting point of execution.
     * In Android apps, the onCreate() function fills that role.
     * It calls other functions to build the user interface.
     * */
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        /*
         * The setContent() function within the onCreate() function is used to define your layout
         * through composable functions.
         * All functions marked with the @Composable annotation can be called from the setContent()
         * function or from other Composable functions.
         * The annotation tells the Kotlin compiler that this function is used by Jetpack Compose
         * to generate the UI.
         * */
        setContent {
            GreetingCardTheme {
                Scaffold(modifier = Modifier.fillMaxSize()) { innerPadding ->
                    Greeting(
                            name = "Android",
                            // A Modifier is used to augment or decorate your composable.
                            modifier = Modifier.padding(innerPadding)
                    )
                }
            }
        }
    }
}

/*
 * @Composable function names are capitalized.
 * @Composable functions can't return anything.
 * */
@Composable
fun Greeting(name: String, modifier: Modifier = Modifier) {
    Surface(color = Color.Cyan, modifier = Modifier.fillMaxSize()) {
        Text(text = "Hi, I am $name!", modifier = modifier.padding(24.dp))
    }
}

/*
 * The GreetingPreview() function is a cool feature that lets you see what your composable looks like
 * without having to build your entire app.
 * To enable a preview of a composable, annotate with @Composable and @Preview.
 * */
@Preview(showBackground = true)
@Composable
fun GreetingPreview() {
    GreetingCardTheme { Greeting("Android") }
}

/*
 * If pairing over Wi-Fi fails, do it manually:
 * $ export PATH=~/Library/Android/sdk/platform-tools:$PATH
 * $ adb pair ip:port
 * */
