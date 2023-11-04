package com.social

import android.app.Application
import dagger.hilt.android.HiltAndroidApp

/*TODO: THIS CLASS (THAT EXTENDS FROM APPLICATION()) HAS BEEN CREATED TO USE DAGGER HILT
   AND IS ABLE TO USE DEPENDENCY INJECTION IN ALL THE PROJECT.
   IF THIS FILE IS NOT ADDED IT WILL CAUSE AN
   'java.lang.IllegalStateException' EXCEPTION
* */
@HiltAndroidApp
class Social : Application()
