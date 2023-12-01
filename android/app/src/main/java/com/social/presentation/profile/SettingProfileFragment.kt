package com.social.presentation.profile

import android.content.Context
import android.content.Intent
import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import com.social.OnboardingActivity
import com.social.R
import com.social.databinding.FragmentSettingProfileBinding

class SettingProfileFragment : Fragment(R.layout.fragment_setting_profile) {
    private lateinit var binding: FragmentSettingProfileBinding

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentSettingProfileBinding.bind(view)

        action()
    }

    private fun action() {
        binding.tvLogout.setOnClickListener {
            logout()
        }
    }

    private fun logout() {
        val preference =
            requireContext().getSharedPreferences("login_saved", Context.MODE_PRIVATE)
        val editor = preference.edit()
        editor.remove("psw")
        editor.remove("check")
        editor.apply()
        startActivity(
            Intent(
                requireContext(),
                OnboardingActivity::class.java,
            ),
        )
    }
}
