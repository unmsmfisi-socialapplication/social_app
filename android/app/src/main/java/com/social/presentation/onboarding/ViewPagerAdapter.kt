package com.social.presentation.onboarding

import android.content.Context
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.LinearLayout
import android.widget.TextView
import androidx.viewpager.widget.PagerAdapter
import com.social.R

class ViewPagerAdapter(private val context: Context) : PagerAdapter() {
    private val images =
        arrayOf(
            R.drawable.online_world,
            R.drawable.connected_world,
            R.drawable.remote_meeting,
            R.drawable.online_friends,
        )

    private val titles =
        arrayOf(
            R.string.none,
            R.string.onboarding_connected_world_title,
            R.string.onboarding_remote_meeting_title,
            R.string.onboarding_online_friends_title,
        )

    private val subtitles =
        arrayOf(
            R.string.none,
            R.string.onboarding_connected_world_subtitle,
            R.string.onboarding_remote_meeting_subtitle,
            R.string.onboarding_online_friends_subtitle,
        )

    override fun getCount(): Int {
        return titles.size
    }

    override fun isViewFromObject(
        view: View,
        `object`: Any,
    ): Boolean {
        return view == `object` as LinearLayout
    }

    override fun instantiateItem(
        container: ViewGroup,
        position: Int,
    ): Any {
        val layoutInflater = context.getSystemService(Context.LAYOUT_INFLATER_SERVICE) as LayoutInflater
        val view = layoutInflater.inflate(R.layout.slider_layout, container, false)

        val slideTitleImage = view.findViewById<ImageView>(R.id.titleImage)
        val slideHeading = view.findViewById<TextView>(R.id.titleText)
        val slideDescription = view.findViewById<TextView>(R.id.subtitleText)

        slideTitleImage.setImageResource(images[position])
        slideHeading.setText(titles[position])
        slideDescription.setText(subtitles[position])

        container.addView(view)

        return view
    }

    override fun destroyItem(
        container: ViewGroup,
        position: Int,
        `object`: Any,
    ) {
        container.removeView(`object` as LinearLayout)
    }
}
