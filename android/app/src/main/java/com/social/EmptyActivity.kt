package com.social

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import androidx.fragment.app.Fragment
import com.social.databinding.ActivityEmptyBinding
import com.social.presentation.friends.FriendsFragment
import com.social.presentation.home.HomeFragment
import com.social.presentation.profile.UserProfileFragment
import com.social.presentation.search.SearchFragment
import com.social.presentation.work.WorkFragment
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class EmptyActivity : AppCompatActivity() {
    private lateinit var binding: ActivityEmptyBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = ActivityEmptyBinding.inflate(layoutInflater)
        setContentView(binding.root)
        navigation()
    }

    private fun navigation() {
        replaceFragment(HomeFragment())
        binding.bottomNavbar.setOnItemSelectedListener {
            when (it.itemId) {
                R.id.navbar_icon_home -> replaceFragment(HomeFragment())
                R.id.navbar_icon_search -> replaceFragment(SearchFragment())
                R.id.navbar_icon_business -> replaceFragment(WorkFragment())
                R.id.navbar_icon_friends -> replaceFragment(FriendsFragment())
                R.id.navbar_icon_user -> replaceFragment(UserProfileFragment())
                else -> {
                }
            }
            true
        }
    }

    private fun replaceFragment(fragment: Fragment) {
        val fragmentTransaction = supportFragmentManager.beginTransaction()
        fragmentTransaction.replace(R.id.navHostFragment, fragment)
        fragmentTransaction.commit()
    }
}
