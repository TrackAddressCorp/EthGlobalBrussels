import "~/styles/globals.css";

import { Providers } from './providers';
import { type Metadata } from "next";

export const metadata: Metadata = {
  title: "Pepetition",
  icons: [{ rel: "icon", url: "/favicon.ico" }],
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode,
}) {
  return (
    <html lang='en'>
      <body>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
