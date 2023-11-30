package com.social.utils

import androidx.fragment.app.Fragment
import androidx.fragment.app.FragmentManager
import androidx.fragment.app.FragmentTransaction
import com.social.R

object FragmentUtils {
    fun replaceFragment(
        fragmentManager: FragmentManager,
        fragment: Fragment,
    ) {
        val transaction: FragmentTransaction = fragmentManager.beginTransaction()
        transaction.replace(R.id.navHostFragment, fragment)
        transaction.addToBackStack(null)
        transaction.commit()
    }
}
