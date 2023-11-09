package com.social.presentation.onboarding

import android.os.Bundle
import android.view.View
import android.widget.TextView
import androidx.constraintlayout.widget.ConstraintLayout
import androidx.core.text.HtmlCompat
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.findNavController
import androidx.viewpager.widget.ViewPager
import com.social.R
import com.social.databinding.FragmentOnboardingBinding

class OnboardingFragment : Fragment(R.layout.fragment_onboarding) {
    private lateinit var binding: FragmentOnboardingBinding
    private lateinit var globalView: View
    private lateinit var dots: Array<TextView>

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentOnboardingBinding.bind(view)
        globalView = view

        initializeViews()
        setUpIndicator(0)
        binding.sliderViewPager.addOnPageChangeListener(viewListener)

        binding.textLogin.setOnClickListener {
            findNavController().navigate(R.id.action_onboardingFragment_to_loginFragment)
        }
        binding.buttonRegister.setOnClickListener {
            findNavController().navigate(R.id.action_onboardingFragment_to_registerFragment)
        }
    }

    private fun initializeViews() {
        binding.sliderViewPager.adapter = ViewPagerAdapter(requireContext())
    }

    private fun setUpIndicator(position: Int) {
        dots = Array(4) { TextView(requireContext()) }
        binding.indicatorLayout.removeAllViews()
        for (i in dots.indices) {
            dots[i] = TextView(requireContext())
            dots[i].text = HtmlCompat.fromHtml("&#8226;", HtmlCompat.FROM_HTML_MODE_LEGACY)
            dots[i].textSize = 35f
            dots[i].setTextColor(requireContext().getColor(R.color.color03))
            binding.indicatorLayout.addView(dots[i])
        }
        dots[position].setTextColor(requireContext().getColor(R.color.color01))
        configureIndicatorLayout(position > 0)
        setElementVisibility(position > 0)
    }

    private fun configureIndicatorLayout(shouldModify: Boolean) {
        val params = binding.indicatorLayout.layoutParams as ConstraintLayout.LayoutParams
        if (shouldModify) {
            params.topToTop = ConstraintLayout.LayoutParams.PARENT_ID
        } else {
            params.topToTop = ConstraintLayout.LayoutParams.UNSET
        }
        binding.indicatorLayout.layoutParams = params
    }

    private fun setElementVisibility(shouldShow: Boolean) {
        if (shouldShow) {
            binding.buttonRegister.visibility = View.VISIBLE
            binding.textLogin.visibility = View.VISIBLE
            binding.headerLogo.visibility = View.INVISIBLE
        } else {
            binding.buttonRegister.visibility = View.INVISIBLE
            binding.textLogin.visibility = View.INVISIBLE
            binding.headerLogo.visibility = View.VISIBLE
        }
    }

    private val viewListener: ViewPager.OnPageChangeListener =
        object : ViewPager.OnPageChangeListener {
            override fun onPageScrolled(
                position: Int,
                positionOffset: Float,
                positionOffsetPixels: Int,
            ) {
            }

            override fun onPageSelected(position: Int) {
                setUpIndicator(position)
            }

            override fun onPageScrollStateChanged(state: Int) {}
        }
}
