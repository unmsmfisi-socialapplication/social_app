package com.social.presentation.onboarding

import android.os.Bundle
import android.text.Html
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.LinearLayout
import android.widget.TextView
import androidx.constraintlayout.widget.ConstraintLayout
import androidx.navigation.fragment.findNavController
import androidx.viewpager.widget.ViewPager
import com.social.R

class OnboardingFragment : Fragment() {
    private lateinit var mSLideViewPager: ViewPager
    private lateinit var mDotLayout: LinearLayout
    private lateinit var headerLogo: LinearLayout
    private lateinit var dots: Array<TextView>
    private lateinit var viewPagerAdapter: ViewPagerAdapter
    private lateinit var btnRegister: Button
    private lateinit var textLogin: TextView

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val view = inflater.inflate(R.layout.fragment_onboarding, container, false)

        initializeViews(view)
        action()

        viewPagerAdapter = ViewPagerAdapter(requireContext())
        mSLideViewPager.adapter = viewPagerAdapter

        setUpIndicator(0)
        mSLideViewPager.addOnPageChangeListener(viewListener)

        return view
    }

    private fun action() {
        btnRegister.setOnClickListener {
            findNavController().navigate(R.id.action_onboardingFragment_to_registerFragment)
        }
        textLogin.setOnClickListener {
            findNavController().navigate(R.id.action_onboardingFragment_to_loginFragment)
        }
    }

    private fun initializeViews(view: View) {
        mSLideViewPager = view.findViewById(R.id.sliderViewPager)
        mDotLayout = view.findViewById(R.id.indicator_layout)
        headerLogo = view.findViewById(R.id.header_logo)
        btnRegister = view.findViewById(R.id.button_register)
        textLogin = view.findViewById(R.id.text_login)
    }

    private fun setUpIndicator(position: Int) {
        dots = Array(4) { TextView(requireContext()) }
        mDotLayout.removeAllViews()

        for (i in dots.indices) {
            dots[i] = TextView(requireContext())
            dots[i].text = Html.fromHtml("&#8226")
            dots[i].textSize = 35f
            dots[i].setTextColor(requireContext().getColor(R.color.color03))
            mDotLayout.addView(dots[i])
        }
        dots[position]?.setTextColor(requireContext().getColor(R.color.color01))

        // Configurar las restricciones para indicator_layout
        configureIndicatorLayout(position > 0)

        // Gestionar la visibilidad de elementos
        setElementVisibility(position > 0)
    }

    private fun configureIndicatorLayout(shouldModify: Boolean) {
        val params = mDotLayout.layoutParams as ConstraintLayout.LayoutParams
        if (shouldModify) {
            params.topToTop = ConstraintLayout.LayoutParams.PARENT_ID
        } else {
            params.topToTop = ConstraintLayout.LayoutParams.UNSET
        }
        mDotLayout.layoutParams = params
    }

    private fun setElementVisibility(shouldShow: Boolean) {
        if (shouldShow) {
            btnRegister.visibility = View.VISIBLE
            textLogin.visibility = View.VISIBLE
            headerLogo.visibility = View.INVISIBLE
        } else {
            btnRegister.visibility = View.INVISIBLE
            textLogin.visibility = View.INVISIBLE
            headerLogo.visibility = View.VISIBLE
        }
    }

    private val viewListener: ViewPager.OnPageChangeListener = object :
        ViewPager.OnPageChangeListener {
        override fun onPageScrolled(
            position: Int,
            positionOffset: Float,
            positionOffsetPixels: Int
        ) {}

        override fun onPageSelected(position: Int) {
            setUpIndicator(position)
        }

        override fun onPageScrollStateChanged(state: Int) {}
    }

    private fun getItem(i: Int): Int {
        return mSLideViewPager.currentItem + i
    }
}