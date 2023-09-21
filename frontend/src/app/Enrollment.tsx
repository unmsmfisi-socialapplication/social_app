import "./globals.scss";
export default function RootEnrollment({
  children,
}: {
  children: React.ReactNode;
}) {
  const divStyle = {
    padding: "90px 100px 120px 10px",
    gap: "25px",
    background: `url('/images/enrollment.png')`,
    backgroundSize: "cover",
    backgroundRepeat: "no-repeat",
    minHeight: '80vh', 
    display: "flex", 
    alignItems: "center", 
  };
  return (
    <html lang="en">
      <body style={divStyle}>
        {children}
      </body>
    </html>
  );
}
