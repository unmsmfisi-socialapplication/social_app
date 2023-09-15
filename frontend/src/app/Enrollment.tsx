import './globals.scss';

export default function RootEnrollment({
    children,
  }: {
    children: React.ReactNode
  }) {
    const divStyle = {
      margin: '0px',
      padding: '0px',
      background: `url('/images/background.jpg')`, 
      backgroundSize: 'cover', 
    };

    
    return (
      <html lang="en">
        <body style={divStyle}>
          {children}
        </body>
      </html>
    )
  }