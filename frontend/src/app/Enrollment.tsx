import './globals.scss'
export default function RootEnrollment({ children }: { children: React.ReactNode }) {
    return (
        <html lang="en">
            <body className="root-enrollment">{children}</body>
        </html>
    )
}
