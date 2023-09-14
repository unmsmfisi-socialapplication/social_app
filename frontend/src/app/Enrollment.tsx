import { Image } from '@mui/icons-material';
import './globals.scss';

export default function RootEnrollment({
    children,
  }: {
    children: React.ReactNode
  }) {
    return (
      <html lang="en">
        <body style={{margin: "0px", padding: "0px"}}>
            <img
                src='/images/background.png'
            ></img>
        </body>
      </html>
    )
  }