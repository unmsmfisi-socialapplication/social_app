import "./globals.scss";

export default function RootEnrollment({
  children,
}: {
  children: React.ReactNode;
}) {
  const divStyle = {
    margin: "0px",
    padding: "0px",
    background: `url('/images/enrollment.png')`,
    backgroundSize: "cover",
    backgroundRepeat: "no-repeat",
    width: "100vh",
  };

  return (
    <html lang="en">
      <body style={divStyle}>
        <div > {children}</div>
      </body>
    </html>
  );
}
