'use client'
import RootLayout from './layout'
import './globals.scss'
export default function RootEnrollment({ children }: { children: React.ReactNode }) {
    return (
        <RootLayout>
            <body className="root-enrollment">{children}</body>
        </RootLayout>
    )
}
