package com.social.utilsTest

import androidx.fragment.app.Fragment
import androidx.fragment.app.FragmentManager
import androidx.fragment.app.FragmentTransaction
import com.social.R
import com.social.utils.FragmentUtils
import org.junit.Before
import org.junit.Test
import org.mockito.Mock
import org.mockito.Mockito
import org.mockito.Mockito.mock
import org.mockito.MockitoAnnotations

class FragmentUtilsTest {

    @Mock
    private lateinit var fragmentManager: FragmentManager

    @Mock
    private lateinit var fragmentTransaction: FragmentTransaction

    @Before
    fun setup() {
        MockitoAnnotations.openMocks(this)
        Mockito.`when`(fragmentManager.beginTransaction()).thenReturn(fragmentTransaction)
    }

    @Test
    fun testReplaceFragment() {
        val fragment: Fragment = mock()

        FragmentUtils.replaceFragment(
            fragmentManager,
            fragment,
        )

        Mockito.verify(fragmentTransaction).replace(
            R.id.navHostFragment,
            fragment,
        )
        Mockito.verify(fragmentTransaction).addToBackStack(null)
        Mockito.verify(fragmentTransaction).commit()
    }
}
