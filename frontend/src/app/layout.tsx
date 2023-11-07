'use client'
import React from 'react'
import './globals.scss'
import { Providers } from '@/redux/provider'

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <html lang="en">
            <body>
                <Providers>{children}</Providers>
            </body>
        </html>
    )
}
