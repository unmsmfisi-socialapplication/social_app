package com.social

import android.content.Intent
import android.os.Bundle
import android.os.CountDownTimer
import androidx.appcompat.app.AppCompatActivity

class MainActivity : AppCompatActivity() {
    private val delay: Long = 2000

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        splashScreen()
    }

    private fun splashScreen() {
        object : CountDownTimer(delay, delay) {
            override fun onTick(millisUntilFinished: Long) {
            }

            override fun onFinish() {
                if (isSessionActive()) {
                    startActivity(
                        Intent(
                            this@MainActivity,
                            EmptyActivity::class.java,
                        ),
                    )
                } else {
                    startActivity(
                        Intent(
                            this@MainActivity,
                            OnboardingActivity::class.java,
                        ),
                    )
                }
                finish()
            }
        }.start()
    }

    private fun isSessionActive(): Boolean {
        val preference = getSharedPreferences("login_saved", MODE_PRIVATE)
        return preference.getBoolean("check", false)
    }
}
