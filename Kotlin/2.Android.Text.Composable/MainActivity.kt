package com.example.happybirthday

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import com.example.happybirthday.ui.theme.HappyBirthdayTheme

/*
 * Composable function naming conventions:
 * https://github.com/androidx/androidx/blob/androidx-main/compose/docs/compose-api-guidelines.md#naming-unit-composable-functions-as-entities
 *
 * The compose function that returns nothing and bears the @Composable annotation MUST be named using Pascal case.
 * MUST be a noun: DoneButton().
 * NOT a verb or verb phrase: DrawTextField().
 * NOT a nouned preposition: TextFieldWithLink().
 * NOT an adjective: Bright().
 * NOT an adverb: Outside().
 * Nouns MAY be prefixed by descriptive adjectives: RoundIcon().
 *
 * Three basic standard layout elements in Compose are Column, Row, and Box.
 * */
class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContent {
            HappyBirthdayTheme {
                Scaffold(modifier = Modifier.fillMaxSize()) { innerPadding ->
                    Surface(
                            modifier = Modifier.fillMaxSize(),
                            color = MaterialTheme.colorScheme.background
                    ) {
                        GreetingText(
                                message = "Happy Birthday Android 14!",
                                from = "From John",
                                modifier = Modifier.padding(innerPadding)
                        )
                    }
                }
            }
        }
    }
}

@Composable
fun GreetingText(message: String, modifier: Modifier = Modifier, from: String = "") {
    Column(verticalArrangement = Arrangement.Center, modifier = modifier.padding(8.dp)) {
        Text(
                text = message,
                fontSize = 100.sp,
                lineHeight = 116.sp,
                textAlign = TextAlign.Center,
                modifier = modifier
        )

        Text(
                text = from,
                fontSize = 36.sp,
                modifier = Modifier.padding(16.dp).align(alignment = Alignment.End)
        )
    }
}

/*
 * The code you added to the BirthdayCardPreview() function with the @Preview annotation
 * is only for previewing in the Design pane in Android Studio.
 * These changes aren't reflected in the app.
 * */
@Preview(showBackground = true)
@Composable
fun BirthdayCardPreview() {
    HappyBirthdayTheme { GreetingText(message = "Happy Birthday Android!", from = "From John") }
}
