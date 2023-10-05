package com.social

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import com.social.databinding.ActivityEmptyBinding

class EmptyActivity : AppCompatActivity() {
    private lateinit var binding: ActivityEmptyBinding
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = ActivityEmptyBinding.inflate(layoutInflater)
        setContentView(binding.root)

    }
}